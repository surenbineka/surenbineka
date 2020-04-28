package server

import (
	"errors"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	//"strconv"

	"github.com/jpillora/archive"
)

const fileNumberLimit = 1000

type fsNode struct {
	Name     string
	Size     int64
	Modified time.Time
	Children []*fsNode
}

func (s *Server) listFiles() *fsNode {
	rootDir := s.state.Config.DownloadDirectory
	root := &fsNode{}
	if info, err := os.Stat(rootDir); err == nil {
		if err := list(rootDir, info, root, new(int)); err != nil {
			log.Printf("File listing failed: %s", err)
		}
	}
	return root
}

func (s *Server) serveFiles(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/download/") {
		url := strings.TrimPrefix(r.URL.Path, "/download/")
		//dldir is absolute
		dldir := s.state.Config.DownloadDirectory
		file := filepath.Join(dldir, url)
		//only allow fetches/deletes inside the dl dir
		if !strings.HasPrefix(file, dldir) || dldir == file {
			http.Error(w, "Nice try\n"+dldir+"\n"+file, http.StatusBadRequest)
			return
		}
		info, err := os.Stat(file)
		if err != nil {
			http.Error(w, "File stat error: "+err.Error(), http.StatusBadRequest)
			return
		}
		switch r.Method {
		case "HEAD":
			if info.IsDir() {
				
			} else {
				/*
				f, err := os.Open(file)
				if err != nil {
					http.Error(w, "File open error: "+err.Error(), http.StatusBadRequest)
					return
				}
				defer f.Close()
				
				fileHeader := make([]byte, 512)
				_, err = f.Read(fileHeader)	// File offset is now len(fileHeader)
				if err != nil {
					http.Error(w, "File read error: "+err.Error(), http.StatusBadRequest)
					return
				}
				
				w.Header().Set("Content-Disposition", "attachment; filename=\""+info.Name()+"\"")				
				w.Header().Set("Content-Type", http.DetectContentType(fileHeader))				
				w.Header().Set("Accept-Ranges", "bytes")
				requestRange := r.Header.Get("range")
				if requestRange == "" {
					w.Header().Set("Content-Length", strconv.Itoa(int(info.Size())))
					return
				}
				requestRange = requestRange[6:]
				splitRange := strings.Split(requestRange, "-")
				if len(splitRange) != 2 {
					http.Error(w, "invalid values for header 'Range'", http.StatusBadRequest)
					return
				}
				begin, err :=strconv.ParseInt(splitRange[0], 10, 64)
				if err != nil {
					http.Error(w, "File read error: "+err.Error(), http.StatusBadRequest)
					return
				}
				end, err := strconv.ParseInt(splitRange[1], 10, 64)
				if err != nil {
					http.Error(w, "File read error: "+err.Error(), http.StatusBadRequest)
					return
				}
				if begin > info.Size() || end > info.Size() {
					http.Error(w, "File read error: range out of bounds for file", http.StatusBadRequest)
					return
				}

				if begin >= end {
					http.Error(w, "File read error: range begin cannot be bigger than range end", http.StatusBadRequest)
					return
				}
				w.Header().Set("Content-Length", strconv.FormatInt(end-begin+1, 10))
				w.Header().Set("Content-Range", fmt.Sprintf( "bytes %d-%d/%d", begin, end, info.Size()))
				w.WriteHeader(http.StatusPartialContent)
				return
				*/
			}
			http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		case "GET":
			if info.IsDir() {
				w.Header().Set("Content-Disposition", "attachment; filename=\""+info.Name()+".zip\"")
				w.Header().Set("Content-Type", "application/zip")
				w.WriteHeader(200)
				//write .zip archive directly into response
				a := archive.NewZipWriter(w)
				a.AddDir(file)
				a.Close()
			} else {
				w.Header().Set("Content-Disposition", "attachment; filename=\""+info.Name()+"\"")
				http.ServeFile(w, r, file)
				/*
				f, err := os.Open(file)
				if err != nil {
					http.Error(w, "File open error: "+err.Error(), http.StatusBadRequest)
					return
				}
				defer f.Close()
				
				fileHeader := make([]byte, 512)
				_, err = f.Read(fileHeader)	// File offset is now len(fileHeader)
				if err != nil {
					http.Error(w, "File read error: "+err.Error(), http.StatusBadRequest)
					return
				}
				*/
				
				/*
				w.Header().Set("Content-Type", http.DetectContentType(fileHeader))				
				w.Header().Set("Accept-Ranges", "bytes")
				requestRange := r.Header.Get("range")
				if requestRange == "" {
					w.Header().Set("Content-Length", strconv.Itoa(int(info.Size())))
					f.Seek(0, 0)
					io.Copy(w, f)
					return
				}
				requestRange = requestRange[6:]
				splitRange := strings.Split(requestRange, "-")
				if len(splitRange) != 2 {
					http.Error(w, "invalid values for header 'Range'", http.StatusBadRequest)
					return
				}
				begin, err :=strconv.ParseInt(splitRange[0], 10, 64)
				if err != nil {
					http.Error(w, "File read error: "+err.Error(), http.StatusBadRequest)
					return
				}
				end, err := strconv.ParseInt(splitRange[1], 10, 64)
				if err != nil {
					http.Error(w, "File read error: "+err.Error(), http.StatusBadRequest)
					return
				}
				if begin > info.Size() || end > info.Size() {
					http.Error(w, "File read error: range out of bounds for file", http.StatusBadRequest)
					return
				}

				if begin >= end {
					http.Error(w, "File read error: range begin cannot be bigger than range end", http.StatusBadRequest)
					return
				}
				w.Header().Set("Content-Length", strconv.FormatInt(end-begin+1, 10))
				w.Header().Set("Content-Range", fmt.Sprintf( "bytes %d-%d/%d", begin, end, info.Size()))
				w.WriteHeader(http.StatusPartialContent)
				f.Seek(begin, 0)
				io.CopyN(w, f, end-begin)
				return
				*/
			}
		case "DELETE":
			duration := time.Since(info.ModTime())
			if ( duration.Hours() >= 12 ) {
				err := os.RemoveAll(file)
				if err != nil {
					http.Error(w, "Delete failed: "+err.Error(), http.StatusInternalServerError)
				}
			}
		default:
			http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	s.static.ServeHTTP(w, r)
}

//custom directory walk

func list(path string, info os.FileInfo, node *fsNode, n *int) error {
	if (!info.IsDir() && !info.Mode().IsRegular()) || strings.HasPrefix(info.Name(), ".") {
		return errors.New("Non-regular file")
	}
	(*n)++
	if (*n) > fileNumberLimit {
		return errors.New("Over file limit") //limit number of files walked
	}
	node.Name = info.Name()
	node.Size = info.Size()
	node.Modified = info.ModTime()
	if !info.IsDir() {
		return nil
	}
	children, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("Failed to list files: %w", err)
	}
	node.Size = 0
	for _, i := range children {
		c := &fsNode{}
		p := filepath.Join(path, i.Name())
		if err := list(p, i, c, n); err != nil {
			continue
		}
		node.Size += c.Size
		node.Children = append(node.Children, c)
	}

	return nil
}

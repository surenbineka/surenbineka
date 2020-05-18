![screenshot](https://raw.githubusercontent.com/sashithacj/torrentfast.net/master/torrentfast-scr.png)

![Build Status](https://raw.githubusercontent.com/sashithacj/torrentfast.net/master/go-passing.svg) 

**TorrentFast.net** is a a self-hosted remote torrent client, backend written in Go (golang), frontend written in AngularJS. Started torrents remotely, download sets of files on the local disk of the server, which are then retrievable or streamable via HTTP.

This project is re-branded from [simple-torrent](https://github.com/boypt/simple-torrent) by `boypt`.

# Features

* UI changes to give the best experience
* Downoaded files can be deleted in frontend only after 12 hours left
* Downoaded files will be automatically deleted by the system after 2 days
* Porn torrent names will be hidden
* File Pause/Resume support (Using apache, thanks to [Navinda](https://github.com/ipmanlk))


# Install

## Source

**Requirement**
- Latest [Golang](https://golang.org/dl/) (Go 1.13+)

``` sh
$ git clone https://github.com/sashithacj/torrentfast.net.git
$ cd torrentfast.net
$ sudo bash run.sh
```

## Binary

See [the latest release](https://github.com/sashithacj/torrentfast.net/releases/latest) or use the oneline script to do a quick install on modern Linux.

```
bash <(wget -qO- https://raw.githubusercontent.com/sashithacj/torrentfast.net/master/scripts/quickinstall.sh)
```

The script install a systemd unit (under `scripts/cloud-torrent.service`) as service. Read further intructions: [Auth And Security](https://github.com/sashithacj/torrentfast.net/wiki/AuthSecurity)

## Docker [![Docker Pulls](https://img.shields.io/docker/pulls/sashithacj/cf24w6g66.svg)][dockerhub] [![Image Size](https://images.microbadger.com/badges/version/sashithacj/cf24w6g66.svg)][dockerhub]

[dockerhub]: https://hub.docker.com/r/sashithacj/cf24w6g66/

``` sh
$ docker run -d -p 80:80 --restart always -v /root/downloads:/downloads sashithacj/cf24w6g66:latest --port 80
```

# Usage

* See Wiki [Command line Options](https://github.com/sashithacj/torrentfast.net/wiki/Command-line-Options) 
* See Wiki [Config File](https://github.com/sashithacj/torrentfast.net/wiki/Config-File) 
* See Wiki [Behind WebServer (reverse proxying)](https://github.com/sashithacj/torrentfast.net/wiki/ReverseProxy) 

# Credits 
* Credits to @boypt for [Simple Torrent](https://github.com/boypt/simple-torrent) 
* Credits to @jpillora for [Cloud Torrent](https://github.com/jpillora/cloud-torrent) 
* Credits to @anacrolix for https://github.com/anacrolix/torrent 

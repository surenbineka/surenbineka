#!/bin/bash
docker stop cloud-torrent cloud-torrent-files 
docker rm cloud-torrent cloud-torrent-files
docker system prune -a
rm -r torrentfast.net
git clone https://github.com/sashithacj/torrentfast.net.git
cd torrentfast.net
docker-compose up

#!/usr/bin/env bash
docker ps -a | egrep 'user-management-front-end' | awk '{print $1}'| xargs docker rm -f
docker build -t user-management-front-end .
docker run  -itd  --name user-management-front-end -p 80:80 user-management-front-end
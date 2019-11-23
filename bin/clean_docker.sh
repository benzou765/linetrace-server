#!/bin/sh

# delete container
docker rm `docker ps -a -q`
# delete images
docker rmi `docker images -q`

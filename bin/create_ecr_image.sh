#!/bin/sh

SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR
cd ..
pwd
cd src
tar czvf docker-ecr/go/src.tar.gz *
cd ../docker-ecr/go/
docker build -t linetrace-server .

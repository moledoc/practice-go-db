#!/bin/sh

set -e
sudo docker build --tag rest .
# sudo docker run -it --rm \
sudo docker run -it --rm -d \
	--name rest \
	--network prac \
	-p 8080:8080 \
	rest;
ip=$(sudo docker inspect -f '{{.NetworkSettings.Networks.prac.IPAddress}}' rest)
printf "Listening and Serving HTTP on %s:8080\n" "$ip"


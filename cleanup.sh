#!/bin/sh

set -e
sudo docker stop pg > /dev/null
sudo docker rm pg > /dev/null
# sudo docker rmi postgres:latest
# sudo docker volume ls | awk '{print }' | tail -n +2 | xargs -I {} sudo docker volume rm "{}"
printf "DONE\n"

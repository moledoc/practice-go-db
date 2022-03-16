#!/bin/sh

set -e
rm -f rest/rest.go
if [ $(sudo docker container ls | grep pg | wc -l) -gt 0 ] || [ $(sudo docker ps -a | grep pg | wc -l) -gt 0 ];
then
	printf "Closing pg container\n"
	sudo docker stop pg > /dev/null
	sudo docker rm pg > /dev/null
fi
if [ $(sudo docker container ls | grep rest | wc -l) -gt 0 ] || [ $(sudo docker ps -a | grep rest | wc -l) -gt 0 ];
then
	printf "Closing rest container\n"
	sudo docker stop rest > /dev/null
fi
# sudo docker rmi postgres:latest
printf "Removing docker volumes\n"
sudo docker volume ls | awk '{print $2}' | tail -n +2 | xargs -I {} sudo docker volume rm "{}"
printf "Removing docker images\n"
sudo docker image ls | grep none | awk '{print $3}' | tail -n +1 | xargs -I {} sudo docker image rm "{}"
printf "Removing docker network prac\n"
sudo docker network rm prac > /dev/null
printf "DONE\n"

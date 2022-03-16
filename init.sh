#!/bin/sh

set -e
if [ $(sudo docker network ls | grep prac | wc -l) -lt 1 ];
then
	printf "Setting up docker network prac\n"
	sudo docker network create -d bridge prac
fi
printf "Creating and running docker container\n"
# When port 5432 is already in use, run the following command
# sudo ss -lptn 'sport = :5432'
# find pid and then run
# kill <pid>
sudo docker run --name pg -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=db -p 5432:5432 --network prac -v postgre:/data -d postgres:latest
sleep 3
printf "Init database ddl\n"
cd init
go run init.go
cd ..
sudo docker inspect -f '{{.NetworkSettings.Networks.prac.IPAddress}}' pg > .dbip
go run gen_rest.go


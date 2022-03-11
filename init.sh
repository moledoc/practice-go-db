#!/bin/sh

set -e
printf "Creating and running docker container\n"
sudo docker run --name pg -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=db -p 5432:5432 -v postgre:/data -d postgres:latest
sleep 3
printf "Init database ddl\n"
cd init
go run init.go
cd ..
printf "DONE\n"

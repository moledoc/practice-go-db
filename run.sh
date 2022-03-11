#!/bin/sh

set -e
# sudo docker run --rm --name pg -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=db -p 5432:5432 -v postgre:/data -d postgres:latest
sudo docker run --name pg -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=db -p 5432:5432 -v postgre:/data -d postgres:latest
sleep 3
go run main.go
# sudo docker stop pg > /dev/null
# sudo docker rmi postgres:latest
# sudo docker volume ls | awk '{print }' | tail -n +2 | xargs -I {} sudo docker volume rm "{}"

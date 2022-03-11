##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /sql
COPY sql/add_idnames.sql .

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

#COPY *.go ./
COPY rest/rest.go ./

RUN go build -o /rest ./rest.go

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /rest /rest
COPY --from=build /sql /sql

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/rest"]

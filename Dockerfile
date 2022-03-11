##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY sql/add_idnames.sql .
COPY go.mod .
COPY go.sum .
RUN go mod download

#COPY *.go ./
COPY rest ./

RUN go build -o /rest

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /rest /rest

EXPOSE 5432

USER nonroot:nonroot

ENTRYPOINT ["/rest"]

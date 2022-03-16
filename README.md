## Purpose

This repository was used to practice:

* connecting to a db from go;
* setting up docker containers;
* rest API.

There won't be very extensive readme about this practice exercise, source code serves that purpose (might add some comments where necessary).
Also, there are shell scripts to set everything up (and show the necessary commands etc).
To get this practice exercise up-and-running, see the 'Quick startup' section.

## Quick startup

**Dependencies**

* go is installed
* docker is installed

Run the provided runscript

```sh
./run.sh
```

Example API command:

```sh
curl "<ip>:8080/add?id=1&name=test1"
```

where ip is given in the terminal when running the `run.sh`.

To check, that API indeed inserts data to database, run

```sh
go run main.go
```

To make a steady stream of data (can also be checked with main.go), run

```sh
go run datastream/datastream.go
```

## Author

Meelis Utt

## TODO:

Create readme and document go etc programs.

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

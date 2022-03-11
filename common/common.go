package common

import (
	"github.com/moledoc/check"
	"io/ioutil"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "postgres"
	Dbname   = "db"
)

func SqlFile(infile string) string {
	contents, err := ioutil.ReadFile(infile)
	check.Err(err)
	return string(contents)
}

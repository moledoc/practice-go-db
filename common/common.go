package common

import (
	"fmt"
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

type Data struct {
	Id   int
	Name string
}

func (d Data) String() string {
	return fmt.Sprintf("%v,%v", d.Id, d.Name)
}

var Psqlconn string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)

func SqlFile(infile string) string {
	contents, err := ioutil.ReadFile(infile)
	check.Err(err)
	return string(contents)
}

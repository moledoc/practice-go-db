package main

import (
	c "github.com/moledoc/practice-go-db/common"
	"github.com/moledoc/templ"
	"io/ioutil"
	"strings"
)

func main() {
	dbip := c.SqlFile(".dbip")
	dbip = strings.Replace(dbip, "\n", "", -1)
	ioutil.WriteFile("rest/rest.go", []byte(templ.ParseStr(c.SqlFile("rest/rest.templ"), map[string]string{"dbip": dbip})), 0600)
}

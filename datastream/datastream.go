package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/moledoc/check"
	c "github.com/moledoc/practice-go-db/common"
	"github.com/moledoc/templ"
	"math/rand"
	"time"
)

func Gen(n int) []c.Data {
	data := make([]c.Data, n)
	for i := 0; i < n; i++ {
		rnd := rand.Intn(1000)
		data[i] = c.Data{Id: rnd, Name: fmt.Sprintf("test%v", rnd)}
	}
	return data
}

func DataStream(db *sql.DB) {
	sql := c.SqlFile("sql/add_idnames.sql")
	//reusable map for idname sql params
	idnameParams := make(map[string]string)
	for {
		// for i := 0; i < 3; i++ {
		d := Gen(1)[0]
		idnameParams["id"] = fmt.Sprintf("%v", d.Id)
		idnameParams["name"] = d.Name
		_, err := db.Exec(templ.ParseStr(sql, idnameParams))
		check.Err(err)
		fmt.Printf("Added: %v,%v\n", d.Id, d.Name)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	db, err := sql.Open("postgres", c.Psqlconn)
	defer db.Close()
	check.Err(err)
	check.Err(db.Ping())
	DataStream(db)
}

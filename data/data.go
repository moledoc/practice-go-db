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

type Data struct {
	Id   int
	Name string
}

func Gen(n int) []Data {
	data := make([]Data, n)
	for i := 0; i < n; i++ {
		rnd := rand.Intn(1000)
		data[i] = Data{Id: rnd, Name: fmt.Sprintf("test%v", rnd)}
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
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Dbname)
	db, err := sql.Open("postgres", psqlconn)
	defer db.Close()
	check.Err(err)
	err = db.Ping()
	check.Err(err)
	DataStream(db)
}

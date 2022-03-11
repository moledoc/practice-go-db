package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/moledoc/check"
	c "github.com/moledoc/practice-go-db/common"
	// "github.com/moledoc/practice-go-db/data"
	// "github.com/moledoc/templ"
	"math/rand"
	"time"
)

// func dataStream(db *sql.DB) {
// sql := c.SqlFile("sql/add_idnames.sql")
// //reusable map for idname sql params
// idnameParams := make(map[string]string)
//for {
// for i := 0; i < 3; i++ {
// d := data.Gen(1)[0]
// idnameParams["id"] = fmt.Sprintf("%v", d.Id)
// idnameParams["name"] = d.Name
// _, err := db.Exec(templ.ParseStr(sql, idnameParams))
// check.Err(err)
// time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
// }
// }
//
// func dataStreamCh(ch chan data.Data) {
//for {
// for i := 0; i < 3; i++ {
// ch <- data.Gen(1)[0]
// time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
// }
// close(ch)
// }

func printRows(db *sql.DB) {
	rows, err := db.Query(`select * from idnames`)
	defer rows.Close()
	check.Err(err)
	fmt.Println("_________________________________________________________________________")
	fmt.Printf("|%-4v|%-10v|%-27v|%-27v|\n", "id", "name", "created", "modified")
	fmt.Println("|----|----------|---------------------------|---------------------------|")
	for rows.Next() {
		var id int
		var name string
		var created string
		var modified string
		check.Err(rows.Scan(&id, &name, &created, &modified))
		fmt.Printf("|%-4v|%-10v|%-27v|%-27v|\n", id, name, created, modified)
	}
	fmt.Println("|____|__________|___________________________|___________________________|")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Dbname)
	db, err := sql.Open("postgres", psqlconn)
	defer db.Close()
	check.Err(err)
	err = db.Ping()
	check.Err(err)

	// sql := c.SqlFile("sql/add_idnames.sql")
	// ch := make(chan data.Data)
	// go dataStreamCh(ch)
	// //reusable map for idname sql params
	// idnameParams := make(map[string]string)
	// loop := true
	// for loop {
	// select {
	// case d, ok := <-ch:
	// if !ok {
	// loop = false
	// break
	// }
	// idnameParams["id"] = fmt.Sprintf("%v", d.Id)
	// idnameParams["name"] = d.Name
	// //fmt.Println(templ.ParseStr(sql, idnameParams))
	// _, err = db.Exec(templ.ParseStr(sql, idnameParams))
	// check.Err(err)
	// //fmt.Println(res)
	// }
	// }

	// go dataStream(db)
	// dataStream(db)
	printRows(db)

	// fmt.Println("DONE")

}

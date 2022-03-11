package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/moledoc/check"
	c "github.com/moledoc/practice-go-db/common"
)

func main() {
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Dbname)
	db, err := sql.Open("postgres", c.Psqlconn)
	defer db.Close()
	check.Err(err)
	check.Err(db.Ping())

	sql := c.SqlFile("ddl/idnames.sql")
	_, err = db.Exec(sql)
	check.Err(err)

	fmt.Println("DDL: done")

}

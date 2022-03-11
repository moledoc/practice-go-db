package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/moledoc/check"
	com "github.com/moledoc/practice-go-db/common"
	"github.com/moledoc/templ"
	"net/http"
	"strconv"
)

var sql_script string
var db *sql.DB
var err error

func addIdname(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if _, err := strconv.Atoi(id); !ok || err != nil {
		return
	}
	name, ok := c.GetQuery("name")
	if !ok {
		return
	}
	_, err := db.Exec(templ.ParseStr(sql_script, map[string]string{"id": id, "name": name}))
	check.Err(err)
	c.String(http.StatusOK, fmt.Sprintf("Added w/ API: %v,%v\n", id, name))
}

func main() {
	fmt.Print("Opening connection to database ")
	// db, err := sql.Open("postgres", com.Psqlconn)
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "${dbip}", com.Port, com.User, com.Password, com.Dbname))
	defer db.Close()
	check.Err(err)
	check.Err(db.Ping())
	fmt.Println("DONE")

	fmt.Println("setting up a http server")
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/add", addIdname)
	// sql_script = com.SqlFile("../sql/add_idnames.sql")
	sql_script = com.SqlFile("/sql/add_idnames.sql")
	router.Run(":8080")
}

package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const MYSQL_DSN = "smart_db:Zxcvbnm123@tcp(127.0.0.1:3306)/smart_db"

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", MYSQL_DSN)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("MYSQL CONNECTED")
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("/users", getUsers)
	r.Run("0.0.0.0:8182")
}
func getUsers(c *gin.Context) {
	rows, err := db.Query("SELECT id,email FROM users")
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	defer rows.Close()
	type User struct {
		ID    int
		Email string
	}
	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Email)
		users = append(users, u)
	}
	c.JSON(200, users)
}

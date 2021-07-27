package user

import (
	"github.com/gin-gonic/gin"
	"github.com/github-hewei/go-gin-demo/utils"
	"log"
	"net/http"
)

type User struct {
	Id   int    `sql:"id"`
	Name string `sql:"name"`
	Age  int    `sql:"age"`
}

func Lists(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}()

	db, err := utils.Db()
	if err != nil {
		panic(err)
	}

	defer func() {
		log.Println("db closed")
		_ = db.Close()
	}()

	rows, err := db.Query("select * from `user`")
	if err != nil {
		panic(err)
	}

	var list []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			panic(err)
		}
		list = append(list, u)
	}

	log.Println(list)

	c.HTML(http.StatusOK, "user_lists.html", gin.H{
		"title": "Users",
		"list": list,
	})
}

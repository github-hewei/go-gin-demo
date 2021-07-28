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
		"list":  list,
	})
}

func Create(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}()

	c.HTML(http.StatusOK, "user_create.html", gin.H{
		"title": "Create User",
	})
}

func Edit(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}()

	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	db, err := utils.Db()
	if err != nil {
		panic(err)
	}

	var u User
	err = db.QueryRow("select * from `user` where id = ?", id).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "user_create.html", gin.H{
		"title": "Edit User",
		"user":  u,
	})
}

func Save(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}()

	id := c.PostForm("id")

	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "姓名不能为空",
		})
	}

	age := c.PostForm("age")
	if age == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "年龄不能为空",
		})
	}

	db, err := utils.Db()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close()
	}()

	if id != "" {
		var u User
		err := db.QueryRow("select * from `user` where id = ?", id).Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			panic(err)
		}
		ret, err := db.Exec("update `user` set `name` = ?, `age` = ? where `id` = ?", name, age, id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "保存成功",
			"data": gin.H{
				"ret": ret,
			},
		})
	} else {
		ret, err := db.Exec("insert into `user` (`name`, `age`) values (?, ?)", name, age)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "保存成功",
			"data": gin.H{
				"ret": ret,
			},
		})
	}
}

func Delete(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}()

	id := c.PostForm("id")
	if id == "" {
		c.AsciiJSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}

	db, err := utils.Db()
	if err != nil {
		panic(err)
	}

	ret, err := db.Exec("delete from `user` where id = ?", id)
	if err != nil {
		panic(err)
	}

	c.AsciiJSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除成功",
		"data": gin.H{
			"ret": ret,
		},
	})
}

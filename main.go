package main

import (
	"2021itmo/DB"
	"database/sql"
	_ "database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type user struct {
	name     string
	lastname string
	login    string
	password string
}

func main() {
	r := gin.Default()

	//регистрация нового пользователя
	//вход пользователя в систему
	//кодирование параши

	//r.LoadHTMLGlob("tamplates/*")

	//регистрация пользователя
	r.POST("/2021itmo/tamplates/index.html", func(c *gin.Context) {
		name := c.PostForm("name")
		lastname := c.PostForm("lastname")
		login := c.PostForm("login")
		password := c.PostForm("password")
		DB.Connect().Exec("INSERT INTO user (name, lastname, login, password) values ($1, $2, $3, $4)", name, lastname, login, password)
		//c.JSON(200, gin.H{"new user":login})// тут вернуть страницу с пользователем
	})

	//вход пользователя в систему
	r.GET("/2021itmo/tamplates/index.html", func(c *gin.Context) {
		login := c.PostForm("login")
		password := c.PostForm("password")
		usr := DB.Connect().QueryRow("SELECT * FROM user WHERE login = $1 and password =$2", login, password)
		us := new(user)
		err := usr.Scan(&us.name, &us.lastname)
		if err != sql.ErrNoRows {
			msg := "Добро пожаловать, " + us.name + " " + us.lastname

			c.JSON(200, gin.H{
				"message": msg,
			})
			// возвращаем страницу пользователя
		}
		if err != nil {
			msg := "неправильное имя пользователя или пароль"
			c.JSON(404, gin.H{
				"massage": msg,
			})
		}
	})

	//бработчик кнопки "кодировать"
	/*r.POST("/", func(c *gin.Context){

	})*/

	r.Run()
}

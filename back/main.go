package main

import (
	"bufio"
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"log"
	"mysql_con"
	"os"
)

var AttachmentDir string
var AttachmentDatabase string

func main() {

	file, err := os.Open("config")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var config [3]string
	var i int = 0
	for scanner.Scan() {
		config[i] = scanner.Text()
		i++
	}
	AttachmentDir = config[0]
	mysql_con.MySQLServer = config[1]
	mysql_con.MySQLPasswd = config[2]

	fmt.Println("Global Settings:")
	fmt.Println("AttachmentDir: ", AttachmentDir)
	fmt.Println("MySQLServer: @", mysql_con.MySQLServer+":"+mysql_con.MySQLPasswd)

	crs := cors.New(cors.Options{ //crs相当于一个中间件，允许所有主机通过
		AllowedOrigins:   []string{"*"}, //
		AllowCredentials: true,
	})
	//---------------------------------------------
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	index := app.Party("/", crs) //所有请求先过crs中间件
	index.Get("/renderIndexMain", renderIndexMain)
	index.Get("/renderThreadsView", renderThreadsView)
	index.Get("/getTotalPosts", getTotalPosts)
	index.Get("/getForums", getForums)
	index.Post("/setNewPost", setNewPost)
	index.Get("/getPost", getPost)

	index.Get("/secret", secret)
	index.Get("/login", login)
	index.Get("/logout", logout)
	index.Get("/users", users)
	index.Get("/resetPosts", resetPosts)

	index.Get("/setNewPasswd", SetNewPasswd)
	index.Get("/getProfile", GetProfile)
	index.Get("/getUserProfile", GetUserProfile)
	index.Post("/setProfile", SetProfile)

	data := app.Party("/data", crs) //所有请求先过crs中间件
	data.Get("/getGender", getGender)
	data.Get("/getLocation", getLocation)
	data.Get("/getLevel", getLevel)

	_ = app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))

}

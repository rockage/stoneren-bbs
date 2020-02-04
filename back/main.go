package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)



func main() {

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
	index.Get("/getTotalThreads", getTotalThreads)
	index.Get("/getTotalPosts", getTotalPosts)
	index.Get("/getForums", getForums)
	index.Post("/setNewPost", setNewPost)
	index.Get("/login", login)
	index.Get("/setNewPasswd", setNewPasswd)
	index.Get("/getProfile", getProfile)
	index.Get("/getUserProfile", getUserProfile)
	index.Post("/setProfile", setProfile)
	index.Get("/resetPosts", resetPosts)

	data:=app.Party("/data", crs) //所有请求先过crs中间件
	data.Get("/getGender", getGender)
	data.Get("/getLocation", getLocation)
	data.Get("/getLevel", getLevel)

	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))



}

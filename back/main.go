package main

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	//index.Get("/getTotalThreads", getTotalThreads)
	index.Get("/getTotalPosts", getTotalPosts)
	index.Get("/getForums", getForums)
	index.Post("/setNewPost", setNewPost)
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



	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))

}

func test() {

	message := Read3()
	var match [][]string
	reg := regexp.MustCompile(`<img.+?>`)
	match = reg.FindAllStringSubmatch(message, -1)

	for i := 0; i < len(match); i++ {
		old_string := match[i][0]
		img_string := match[i][0]
		fmt.Println("i=", i, "value:", img_string)
		//分析每一条img标签，控制宽高的有style和原始html两种方法
		//根据正则，截取内容的序号如下：
		//1:style_width
		//2:html_width
		//3:style_height
		//4:html_height
		//<img.+?width:(\d{1,})px|width="(\d{1,})"|height:(\d{1,})px|height="(\d{1,})"
		reg_a := regexp.MustCompile(`<img.+?width\s{0,}:\s{0,}(\d{1,})px|width\s{0,}=\s{0,}"(\d{1,})"|height\s{0,}:\s{0,}(\d{1,})px|height\s{0,}=\s{0,}"(\d{1,})"`)
		match_a := reg_a.FindAllStringSubmatch(img_string, -1)
		if len(match_a) > 0 {
			for j := 0; j < len(match_a); j++ {
				style_width_string := match_a[j][1] //处理第1种情况：style控制width
				if style_width_string != "" {
					style_width_value, _ := strconv.Atoi(style_width_string)
					if style_width_value > 1024 {
						img_string = strings.Replace(img_string, style_width_string+"px", "100%", 1) //将之替换为100%以自适应屏幕，如小图则不做处理
					}
				}

				html_width_string := match_a[j][2] //处理第2种情况：html控制width
				if html_width_string != "" {
					html_width_value, _ := strconv.Atoi(html_width_string)
					if html_width_value > 1024 {
						img_string = strings.Replace(img_string, html_width_string, "100%", 1)
					}
				}

				style_height_string := match_a[j][3] //处理第3种情况：style_height
				if style_height_string != "" {
					style_height_value, _ := strconv.Atoi(style_height_string)
					if style_height_value > 768 {
						img_string = strings.Replace(img_string, style_height_string+"px", "100%", 1)
					}
				}

				html_height_string := match_a[j][4] //处理第4种情况：html_height
				if html_height_string != "" {
					html_height_value, _ := strconv.Atoi(html_height_string)
					if html_height_value > 768 {
						img_string = strings.Replace(img_string, html_height_string, "100%", 1)
					}
				}

			} //for j
			message = strings.Replace(message, old_string, img_string, -1)
		} // if
	} // for i

	reg = regexp.MustCompile(`<img.+?>`)
	match = reg.FindAllStringSubmatch(message, -1)

	fmt.Println()

	for i := 0; i < len(match); i++ {
		fmt.Println("i=", i, "after:", match[i][0])
	}

	//	reg_a:=regexp.MustCompile(`<img.+?width\s{0,}:\s{0,}(\d{1,})px`)
	//	match := reg_a.FindStringSubmatch(s)

}

func Read3() string {
	f, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}

func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}

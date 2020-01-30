package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"mysql_con"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//FillMain
func renderIndexMain(ctx iris.Context) {
	var sql string
	forumsId := ctx.FormValue("fid")
	page := ctx.FormValue("page")
	p1, _ := strconv.Atoi(page)
	t1 := p1*20 - 20

	startRec := strconv.Itoa(t1)
	stopRec := "20"

	if forumsId == "0" {
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread ORDER BY dateline DESC limit " + startRec + "," + stopRec
	} else {
		sql = "select tid,author,subject,dateline,lastpost,lastposter,views,replies from pre_forum_thread where fid = " + forumsId + " ORDER BY dateline DESC limit " + startRec + "," + stopRec
	}
	var ok bool
	var rst []map[string]string

	rst, ok = mysql_con.Query(sql)
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			_, _ = ctx.JSON(string(b)) //不返回错误代码，强行执行
		}
	}
}
func getTotalThreads(ctx iris.Context) {
	var sql string
	forumsId := ctx.FormValue("fid")
	if forumsId == "0" {
		sql = "explain select * from pre_forum_thread"
	} else {
		sql = "select  COUNT(1) as rows from pre_forum_thread where fid = " + forumsId
	}

	var rst []map[string]string
	rst, _ = mysql_con.Query(sql) //求出总行数
	_, _ = ctx.JSON(rst[0]["rows"])

}

//ThreadsView
func renderThreadsView(ctx iris.Context) {
	page := ctx.FormValue("page")
	tid := ctx.FormValue("tid")
	p1, _ := strconv.Atoi(page)
	t1 := p1*20 - 20
	startRec := strconv.Itoa(t1)
	stopRec := "20"

	var ok bool
	var rst []map[string]string
	rst, ok = mysql_con.Query("select pid,fid,tid,author,dateline,message from pre_forum_post where tid = " + tid + " ORDER BY dateline limit " + startRec + "," + stopRec)
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			_, _ = ctx.JSON(string(b)) //不返回错误代码，强行执行
		}
	}
}
func getTotalPosts(ctx iris.Context) {
	posts := ctx.FormValue("tid")
	var rst []map[string]string
	rst, _ = mysql_con.Query("select COUNT(1) as rows from pre_forum_post where tid = " + posts) //求出总行数
	_, _ = ctx.JSON(rst[0]["rows"])
}

//forumView
func getForums(ctx iris.Context) {
	var rst []map[string]string
	var ok bool
	rst, ok = mysql_con.Query("select fid,name,status,displayorder from pre_forum_forum where status = 1 ORDER BY displayorder")
	if ok {
		b, err := json.Marshal(rst)
		if err == nil {
			_, _ = ctx.JSON(string(b))
		}
	}
}

//New Post
func setNewPost(ctx iris.Context) {
	fid := ctx.FormValue("fid")
	tid := ctx.FormValue("tid")
	uid := ctx.FormValue("uid")
	threadsTitle := ctx.FormValue("threadsTitle")
	postContens := ctx.FormValue("postContens")
	fmt.Println("fid = " + fid)
	fmt.Println("tid = " + tid)
	fmt.Println("uid = " + uid)
	fmt.Println("threadsTitle = " + threadsTitle)
	// 105003
	postContens = strings.Replace(postContens, "'", "\\'", -1)
	postContens = strings.Replace(postContens, "</a>", "", -1)
	re, _ := regexp.Compile(`<a\s{1,}href(.+?)>`) //删除全部原生超链接
	postContens = re.ReplaceAllString(postContens, "")
	re = regexp.MustCompile(`<img src="data:.+?;base64,(.+?)">`)
	for _, match := range re.FindAllStringSubmatch(postContens, -1) {
		data, _ := base64.StdEncoding.DecodeString(match[1])
		fileName := saveAttachment(data)
		src := "<img src=\"" + fileName + "\">"
		postContens = strings.Replace(postContens, match[0], src, -1)
	}
	mysql_con.Exec("UPDATE pre_forum_post set message = '" + postContens + "' where pid = 105003")
}

func saveAttachment(data []byte) string {
	dir := "../front/static/attachment/" + time.Now().Format("2006-01") + "/"
	myUuid := uuid.NewV4()
	_, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, 0777)
	}
	fileName := dir + myUuid.String() + ".jpg"
	err = ioutil.WriteFile(fileName, data, 0666)
	fileName = strings.Replace(fileName, "../front", "", -1) //去掉../front，否则前端无法读取
	return fileName
}

// login
func login(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	var rst []map[string]string
	rst, _ = mysql_con.Query("select uid from pre_members where username = '" + username + "' and `password` = '" + password + "'")
	if rst != nil {
		b, err := json.Marshal(rst)
		if err == nil {
			_, _ = ctx.JSON(string(b))
		}
	} else {
		ctx.Text("not found")
	}
}

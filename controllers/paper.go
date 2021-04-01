package controllers

import (
	"Blog/solr"
	"Blog/models"
	// "encoding/json"
	"fmt"
	// "github.com/astaxie/beego/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var (
	sc *solr.Connection
)

// Operations about Blog
type BlogController struct {
	beego.Controller
}

// @Title CreateBlog
// @Description create blogs
// @Param	body		body 	models.Blog	true		"body for blog content"
// @Success 200 {int} models.Blog.Id
// @Failure 403 body is empty
// @router / [post]
func (p *BlogController) Post() {
	var blog models.Blog
    title := p.GetString("title")
    introduction := p.GetString("introduction")
	content := p.GetString("content")

	blog.Title = title
	id := models.AddBlog(blog)
	f := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{"id": id, "title": title,"introduction":introduction,"content":content},
		},
	}
	resp, err := sc.Update(f, true)

	if err != nil {
		fmt.Println("error =>", err)
	} else {
		fmt.Println("resp =>", resp)
	}
}

// @Title GetAll
// @Description get all Blogs
// @Success 200 {object} models.Blog
// @router / [get]
func (p *BlogController) GetAll() {
	blogs := models.GetAllBlogs()
	p.Data["json"] = blogs
	p.ServeJSON()
}


// @Title Get
// @Description get blog by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Blog
// @Failure 403 :id is empty
// @router /:id [get]
func (p *BlogController) Get() {
	id,err := p.GetInt(":id")
	blog, err := models.GetBlog(int64(id))
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{fmt.Sprintf("id:%d",id)},
		},
		Rows: 1,
	}
	res, err := sc.Select(&q)

	if err != nil {
		return
	}
	results := res.Results
	result := results.Get(0)
	

	if err != nil {
		p.Data["json"] = err.Error()
	} else {
		p.Data["json"] = map[string]interface{}{
			"title":result.Field("title"),
			"introduction":result.Field("introduction"),
			"content":result.Field("content"),
			"id":blog.Id,
			"create_time":blog.CreateTime,
		}
	}
	p.ServeJSON()
}

func init() {
	var err error
	sc, err = solr.Init("localhost", 8983,"blog")
	if err != nil{
		fmt.Print(err)
	}
}
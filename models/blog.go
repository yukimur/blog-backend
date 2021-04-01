package models

import (
	"time"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/astaxie/beego/logs"
)

var (
	BlogList map[string]*Blog
)

type Blog struct {
	Id       int64		`json:"id"`
	Title string		`orm:"size(128)"		json:"title"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"		json:"createTime"`
	User  *User			`orm:"null;rel(fk)"	json:"user"`
	AdmireCount int		`orm:"null" json:"admireCount"`
	ViewCount int		`orm:"null" json:"viewCount"`
}

type Comment struct {
	Id      int64		`json:"id"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"		json:"createTime"`
	Blog *Blog		`orm:"rel(fk)"		json:"blog"`
	Comment *Comment		`orm:"rel(fk)"		json:"comment"`
}

func AddBlog(p Blog) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&p)
	if err != nil {
		logs.Debug(err)
	}
	return id
}

func GetBlog(id int64) (p *Blog, err error) {
	o := orm.NewOrm()
	logs.Debug(id)
	blog := Blog{Id: id}
	
	err = o.Read(&blog)
	return &blog,err
}

func GetAllBlogs() map[string]*Blog {
	return BlogList
}

func UpdateBlog(uid string, pp *Blog) (p *Blog, err error) {
	return nil, errors.New("Blog Not Exist")
}

func DeleteBlog(uid string) {
	delete(BlogList, uid)
}
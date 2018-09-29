package models

import(
	"time"

	"github.com/astaxie/beego/orm"
	"strconv"
)

func init(){
	orm.RegisterModel(new(Category),new(Topic))
}

type Category struct {
	Id int64     `orm:"auto"`
	Title string
	Created time.Time //`orm:"index"`
	Views int64 //`orm:"index"`
	TopicTime time.Time //`orm:"index"`
	TopicCount int64
	TopiclastUserId int64
}

type Topic struct{
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time  `orm:"index"`
	Update time.Time  `orm:"index"`
	Views int64
	Author string
	ReplyTime time.Time  `orm:"index"`
	ReplyCount int64
	RepleyLastUserId int64
	Category string
}

func AddTopic(title string,content string) error{
	o := orm.NewOrm()
	topic := &Topic{
		Title:title,
		Content:content,
		Created:time.Now(),
		Update :time.Now(),
	}
	_, err := o.Insert(topic)
	if err != nil{
		return err
	}

	return nil
}

func AddCategory(name string) error{
	o := orm.NewOrm()
	cate := &Category{Title :name}
	qs := o.QueryTable("category")
	err := qs.Filter("title",name).One(cate)
	if err == nil {
		return err
	}

	_ , err = o.Insert(cate)
	if err != nil{
		return err
	}
	return nil
}

func GetAllTopics() ([]*Topic ,error) {
	o := orm.NewOrm()
	topics := make([]*Topic,0)
	qs:= o.QueryTable("topic")
	_ , err := qs.All(&topics)
	return topics,err
}

func GetAllCategories() ([]*Category, error){
	o := orm.NewOrm()

	cates := make([]*Category,0)

	qs := o.QueryTable("category")
	_ ,err := qs.All(&cates)
	return cates,err
}

func DelCategory(id string)error{
	cid ,err := strconv.ParseInt(id,10,64)
	if err!= nil{
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id:cid}
	_ , err = o.Delete(cate)
	return err
}
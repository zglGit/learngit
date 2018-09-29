package models

import "github.com/astaxie/beego/orm"


func init(){
	orm.RegisterModel(new(User),new(Profile),new(Post),new(Tag))
}

type User struct {
	Id  int64
	Name string
	ProfileId  int64
}

type Profile struct {
	Id   int64
	Age  int64
	UserId int64
}

type Post struct {
	Id   int
	Title string
	UserId int64
}


type Tag struct {
	Id   int
	Name  string
}

func QueryPostbyUserId(id int64)([]*Post,error){
	var posts []*Post
	o := orm.NewOrm()
	n , err := o.QueryTable("post").Filter("UserId",id).All(&posts)
	if err != nil && n>0 {
		return nil,err
	}
	return posts , err
}

func QueryUserByPostTitle(title string)(*User,error){
	 user := &User{}
	 post := &Post{}
	 o := orm.NewOrm()
	 err := o.QueryTable("post").Filter("Title",title).One(post)
	if err != nil {
		return nil,err
	}
	user.Id = post.UserId
	err = o.Read(user)
	if err != nil {
		return nil,err
	}
	return user,nil
}

func QueryUserByUserId(id int64) (*User,error) {
	user :=&User{}
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("Id",id).One(user)
	if err != nil {
		return nil,err
	}
	return user, nil
}

func QueryProfileByUser(id int64) (*Profile,error){
	profile := &Profile{}
	o := orm.NewOrm()
	err := o.QueryTable("profile").Filter("UserId",id).One(profile)
	if err != nil {
		return nil,err
	}
	return profile,nil
}



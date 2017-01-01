package models

import (
  // "strings"
  // "fmt"
  "errors"
  "time"
	"gopkg.in/mgo.v2/bson"
)

type Tag struct {
  Text        string `json: "tag"`
}

func InsertTag(id string, comment *Comment) error {
  if comment.Text == "" {
    //采用errors包的New方法 返回一个err的类型
    return errors.New("Comment text is null!")
  }
  //db.person.update({"name":"ryan"},{"$push":{"language":"english"}},true,true);//给数组的末尾添加一个值。
  c := DB.C("tags")
  err := c.Update(bson.M{"_id": objectId}, bson.M{"$push":bson.M{"comments":comment}})
  return err
  // return err
}

// type Category struct {
// 	Id_         bson.ObjectId `bson:"_id"`
// 	Name        string
// 	Title       string
// 	Content     string
// 	CreatedTime time.Time
// 	UpdatedTime time.Time
// 	Views       int
// }
//
// func (category *Category) CreatCategory() error {
// 	//category.Id_ = bson.NewObjectId()
// 	c := DB.C("category")
// 	err := c.Insert(category)
// 	SetAppCategories()
// 	return err
// }
//
// func (category *Category) UpdateCategory() error {
// 	c := DB.C("category")
// 	err := c.UpdateId(category.Id_, category)
// 	SetAppCategories()
// 	return err
// }

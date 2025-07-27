package part3

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"homework/gorms"
)

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
type Userx struct {
	gorm.Model
	Name      string
	CharCount uint32
	Postxs    []Postx `gorm:"foreignKey:UserxID"`
}
type Postx struct {
	gorm.Model
	Title         string
	UserxID       uint       `gorm:"index:,"`
	Commentxs     []Commentx `gorm:"foreignKey:"PostxID"`
	CommentStatus string
}

type Commentx struct {
	gorm.Model
	Content string
	PostxID uint `gorm:"index:,"`
}

func GormRun1() {
	gorms.DB.AutoMigrate(&Userx{})
	gorms.DB.AutoMigrate(&Postx{})
	gorms.DB.AutoMigrate(&Commentx{})

	/*gorms.DB.Transaction(func(tx *gorm.DB) error {
	/*u := Userx{
		Name: "lisi",
	}
	err := tx.Save(&u).Error
	if err != nil {
		fmt.Printf("save User info error \n")
		panic(err)
	}

	pos := []Postx{
		{Title: "T4", UserxID: u.Model.ID},
		{Title: "T5", UserxID: u.Model.ID},
		{Title: "T6", UserxID: u.Model.ID},
	}
	tx.Save(&pos)
	if err != nil {
		fmt.Printf("save Post info error \n")
		panic(err)
	}*/

	/*coms := []Commentx{
		{Content: "C41", PostxID: 4},
		{Content: "C42", PostxID: 4},
		{Content: "C43", PostxID: 4},
		{Content: "C44", PostxID: 4},
		{Content: "C45", PostxID: 4},
		{Content: "C51", PostxID: 5},
		{Content: "C52", PostxID: 5},
		{Content: "C53", PostxID: 5},
		{Content: "C61", PostxID: 6},
		{Content: "C62", PostxID: 6},
		{Content: "C63", PostxID: 6},
	}
	tx.CreateInBatches(&coms, 11)
	return nil*/
	//})*/

}

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func QueryUserInfo(uid uint) string {
	u := &Userx{}
	err := gorms.DB.Model(&Userx{}).Find(&u, "id = ?", uid).Error
	if err != nil {
		panic(err)
	}

	pos := []Postx{}
	err = gorms.DB.Model(&Postx{}).Find(&pos, "userx_id = ?", uid).Error
	if err != nil {
		panic(err)
	}

	pIds := make([]uint, len(pos))
	for idx, p := range pos {
		pIds[idx] = p.Model.ID
	}

	coms := []Commentx{}
	err = gorms.DB.Model(&Commentx{}).Find(&coms, "postx_id in ?", pIds).Error
	if err != nil {
		panic(err)
	}

	// group by postx_id
	cGroup := map[uint][]Commentx{}
	for _, v := range coms {
		commentxes, exist := cGroup[v.PostxID]
		if !exist {
			cGroup[v.PostxID] = []Commentx{v}
		} else {
			commentxes = append(commentxes, v)
			cGroup[v.PostxID] = commentxes
		}
	}

	// range pos
	for idx, p := range pos {
		pos[idx].Commentxs = cGroup[p.Model.ID]
	}

	u.Postxs = pos

	uData, _ := json.Marshal(u)
	fmt.Printf("user json info %v \n", string(uData))
	return string(uData)
}
func QueryMaxCommentUserInfo() {
	sql := "select count(u.id) cc, u.id" +
		"	from userxes u" +
		"  left join postxes p on p.userx_id = u.id " +
		"  left join commentxes c on c.postx_id = p.id" +
		"  group by u.id\norder by cc desc" +
		"  limit 1"

	//u := Userx{}
	//gorms.DB.Raw(sql).Scan(&u)
	//fmt.Printf("评论最多的文章用户id:%v \n", u.Model.ID)
	id := 0
	cc := 0
	row := gorms.DB.Raw(sql).Row()
	row.Scan(&cc, &id)
	//fmt.Println(err)
	fmt.Printf("评论最多的文章用户id:%v, cc:%v \n", id, cc)
}

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func (p *Postx) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Printf(" Postx BeforeSave hook ...\n")
	titles := []string{}
	errr := tx.Model(&Postx{}).Raw("select title from postxes where userx_id = ?", p.UserxID).Scan(&titles).Error
	if errr == nil {
		var length = 0
		for _, v := range titles {
			length += len(v)
		}

		errrr := tx.Model(&Userx{}).Where("id = ?", p.UserxID).Update("CharCount", length).Error
		if errrr != nil {
			fmt.Printf("update user charCount failed\n")
		}
	}
	return nil
}
func TriggerPostxHook() {
	p := Postx{
		Title:         "Ttttt",
		UserxID:       1,
		CommentStatus: "",
	}
	gorms.DB.Save(&p)
}

func (c *Commentx) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Printf(" Commentx AfterDelete hook ...\n")
	left := 0
	errr := tx.Model(&Commentx{}).Raw("select count(1) from commentxes where postx_id = ?", c.PostxID).Scan(&left).Error
	if errr == nil {

		errrr := tx.Model(&Postx{}).Where("id = ?", c.PostxID).Update("comment_status", "无评论").Error
		if errrr != nil {
			fmt.Printf("update Postx comment_status failed\n")
		}
	}
	return nil
}

func TriggerCommentxHook() {
	c := Commentx{
		Model:   gorm.Model{ID: 3},
		PostxID: 1,
	}
	gorms.DB.Model(&Commentx{}).Delete(&c)
}

package part4

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" form:"username" json:"username"`
	Password string `gorm:"not null" form:"password" json:"password"`
	Email    string `gorm:"unique;not null" form:"email" json:"email"`
}
type Post struct {
	gorm.Model
	Title   string `gorm:"not null" form:"title" json:"title"`
	Content string `gorm:"not null" form:"content" json:"content"`
	UserID  uint   `form:"userID" json:"userID"`
	User    User
}
type PostUpdate struct {
	ID      uint
	Title   string `gorm:"not null" form:"title" json:"title"`
	Content string `gorm:"not null" form:"content" json:"content"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" form:"content" json:"content"`
	UserID  uint   `form:"userID" json:"userID"`
	User    User
	PostID  uint `form:"postID" json:"postID"`
	Post    Post
}

func init() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
}

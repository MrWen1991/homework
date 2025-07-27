package part4

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func postCreate(c *gin.Context) {
	post := Post{}
	if !bindParam(c, post) {
		return
	}

	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "data save has failed"})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post created successfully"})

}

func bindParam(c *gin.Context, post Post) bool {
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param error"})
		fmt.Println(err.Error())
		return false
	}
	return true
}

func postRead(c *gin.Context) {
	post := Post{}
	id := c.PostForm("id")

	if err := db.Debug().Model(&Post{}).Preload("User").Where("id = ?", id).Find(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query post has failed"})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, &post)

}
func postUpdate(c *gin.Context) {
	postUpdate := PostUpdate{}
	if err := c.ShouldBind(&postUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param error"})
		fmt.Println(err.Error())
		return
	}

	err := db.Debug().Model(&Post{}).
		Where("id = ?", postUpdate.ID).
		Updates(map[string]interface{}{
			"content": postUpdate.Content,
			"title":   postUpdate.Title,
		}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update post has failed"})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
}
func postDelete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param id is absent"})
		return
	}
	idp, err := strconv.ParseUint(id, 10, 0)
	if err != nil {

	}
	var post = Post{Model: gorm.Model{ID: uint(idp)}}
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete post has failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete post success"})
}

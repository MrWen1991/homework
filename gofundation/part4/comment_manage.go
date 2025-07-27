package part4

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentCreate(c *gin.Context) {
	var comment = Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param parse error"})
		fmt.Println(err.Error())
		return
	}
	if err := db.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create Comment error"})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create Comment successfully"})

}

func CommentsQuery(c *gin.Context) {
	pId := c.Query("post_id")
	if pId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post_id is empty"})
		return
	}
	postId, err := strconv.ParseUint(pId, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post_id is not Integer"})
		return
	}

	coms := []Comment{}
	err = db.Model(&Comment{}).Where("post_id = ?", postId).Find(&coms).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query error"})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, coms)

}

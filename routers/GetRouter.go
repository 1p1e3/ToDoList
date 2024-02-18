package routers

import (
	"ToDoList/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// GetRouter 查询计划接口
// get: "/todo"
func GetRouter(db *gorm.DB, r *gin.Engine) {
	r.GET("/todo", func(ctx *gin.Context) {
		// 默认直接查询全部，将全部的计划(除了软删除之外)返回给前端
		// 由前端对这些计划进行区分展示
		todolist := []models.ToDoList{}
		db.Find(&todolist)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "ok",
			"data":    todolist,
		})
	})
}

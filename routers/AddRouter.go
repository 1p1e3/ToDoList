package routers

import (
	"ToDoList/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// AddRouter 添加计划接口
// post: "/todo"
func AddRouter(db *gorm.DB, r *gin.Engine) {
	r.POST("/todo", func(ctx *gin.Context) {
		// 解析请求体，并将请求体 json 映射到 map
		m := make(map[string]string)
		json.NewDecoder(ctx.Request.Body).Decode(&m)
		// 获取 map 中的 name 参数，name 参数表示计划的名称
		name := m["name"]
		// 插入一条记录，新创建的计划的 Status 字段为 1，表示已创建状态
		db.Create(&models.ToDoList{
			Name:   name,
			Status: 1,
		})
		// 响应
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "ok",
		})
	})
}

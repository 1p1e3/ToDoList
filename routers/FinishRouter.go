package routers

import (
	"ToDoList/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// FinishRouter 完成计划接口
// post: "/finish"
func FinishRouter(db *gorm.DB, r *gin.Engine) {
	r.POST("/finish", func(ctx *gin.Context) {
		// 解析请求体，并将请求体 json 映射到 map
		m := make(map[string]int)
		json.NewDecoder(ctx.Request.Body).Decode(&m)
		// 获取 map 中的 id 参数，id 参数表示要完成的计划的 id
		id := m["id"]
		// 将计划的 Status 改为 2 表示已完成状态
		db.Model(&models.ToDoList{}).Where("id = ?", id).Update("Status", 2)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "ok",
		})
	})
}

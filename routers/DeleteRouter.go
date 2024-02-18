package routers

import (
	"ToDoList/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"net/http"
)

// DeleteRouter 删除计划接口
// post: "/delete"
func DeleteRouter(db *gorm.DB, r *gin.Engine) {
	r.POST("/delete", func(ctx *gin.Context) {
		// 解析请求体，并将请求体 json 映射到 map
		m := make(map[string]int)
		json.NewDecoder(ctx.Request.Body).Decode(&m)
		// 获取 map 中的 id 参数，id 参数表示要删除的计划的 id
		id := m["id"]
		db.Delete(&models.ToDoList{}, id)
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "ok",
		})
	})
}

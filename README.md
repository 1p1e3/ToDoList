# ToDoList(simple & ugly)
### 1. 说明
go + go mod + gin + gorm + mysql 小练习

前端在这儿: [ToDoList-FrontEnd](https://github.com/1p1e3/ToDoList-FrontEnd)

### 2. 目录结构
- `logs`： 存放日志文件
- `models`：模型包
- `routers`：路由包、接口包、api 包
- `utils`：工具包
- `go.mod`
- `main.go`


### 3. 文件
- `models` 包
  - `ToDoListModel.go`：定义 ToDoList 结构体，用于数据表映射
- `routers` 包
  - `AddRouter.go`：定义添加计划接口
  - `DeleteRouter.go`：定义删除计划接口
  - `FinishRouter.go`：定义完成计划接口
  - `GetRouter.go`：定义查询计划接口
- `utils`
  - `Cors.go`：解决前后端跨域，网上 copy 的代码
  - `CreateLogFile.go`：创建日志存放目录(logs)和日志文件

### 4. 接口列表
|  路径   |  方法  |    参数及类型     |                      响应                       |
|:-----:|:----:|:------------:|:---------------------------------------------:|
| /todo | POST | name: string |               {"message": "ok"}               |
|/delete| POST |   id: int    |               {"message": "ok"}               |
|/finish| POST |   id: int    |               {"message": "ok"}               |
| /todo | GET  |      没有      |   {"message": "ok", "data": 保存数据表中所有记录的切片}    |

package models

import (
	"gorm.io/gorm"
)

// ToDoList 嵌套 gorm 预定义的 gorm.Model
type ToDoList struct {
	gorm.Model
	Name   string // 计划的名称
	Status int    // 计划的状态，改为布尔值表示更好，布尔占 1 个字节，int 占 8 个字节
}

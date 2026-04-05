package model

import "time"

// Document 对应数据库 documents 表
type Document struct {
	ID        uint64    `json:"id"         gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title"      gorm:"size:512;not null"`
	Content   string    `json:"content"    gorm:"type:text;not null"`
	URL       string    `json:"url"        gorm:"size:1024;default:''"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Document) TableName() string {
	return "documents"
}

// CreateDocRequest 创建文档请求体
type CreateDocRequest struct {
	Title   string `json:"title"   binding:"required,min=1,max=512"`
	Content string `json:"content" binding:"required,min=1"`
	URL     string `json:"url"`
}

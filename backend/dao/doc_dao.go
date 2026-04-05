package dao

import (
	"search-engine/model"

	"gorm.io/gorm"
)

// DocDAO 文档数据访问对象
type DocDAO struct {
	db *gorm.DB
}

func NewDocDAO(db *gorm.DB) *DocDAO {
	return &DocDAO{db: db}
}

// Create 新增文档
func (d *DocDAO) Create(doc *model.Document) error {
	return d.db.Create(doc).Error
}

// GetByID 按 ID 查询文档
func (d *DocDAO) GetByID(id uint64) (*model.Document, error) {
	var doc model.Document
	if err := d.db.First(&doc, id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

// GetAll 查询所有文档（用于启动时构建索引）
func (d *DocDAO) GetAll() ([]*model.Document, error) {
	var docs []*model.Document
	if err := d.db.Order("id asc").Find(&docs).Error; err != nil {
		return nil, err
	}
	return docs, nil
}

// Delete 按 ID 删除文档
func (d *DocDAO) Delete(id uint64) error {
	return d.db.Delete(&model.Document{}, id).Error
}

// Update 更新文档字段
func (d *DocDAO) Update(doc *model.Document) error {
	return d.db.Save(doc).Error
}

// Count 返回文档总数
func (d *DocDAO) Count() (int64, error) {
	var count int64
	err := d.db.Model(&model.Document{}).Count(&count).Error
	return count, err
}

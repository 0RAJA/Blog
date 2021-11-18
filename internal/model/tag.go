package model

import (
	"Blog/pkg/app"
	"github.com/jinzhu/gorm"
)

// Tag 标签
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// Count 统计标签数
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Get 获取
func (t Tag) Get(db *gorm.DB) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND is_del = ? AND state = ?", t.ID, 0, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tag, nil
}

// List 获取指定范围的标签
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// Create 新增标签
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Update 更新标签
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除标签
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.ID, 0).Delete(&t).Error
}

// IsExit 判断标签是否存在
func (t Tag) IsExit(db *gorm.DB) bool {
	return db.Where("name = ?", t.Name).Find(&t).Error == nil
}

/*
Model：指定运行 DB 操作的模型实例，默认解析该结构体的名字为表名，格式为大写驼峰转小写下划线驼峰。若情况特殊，也可以编写该结构体的 TableName 方法用于指定其对应返回的表名。
Where：设置筛选条件，接受 map,struct 或 string 作为条件。
Offset：偏移量，用于指定开始返回记录之前要跳过的记录数。
Limit：限制检索的记录数。
Find：查找符合筛选条件的记录。
Updates：更新所选字段。
Delete：删除数据。
Count：统计行为，用于统计模型的记录数。
*/

package dao

import (
	"Blog/internal/model"
	"Blog/pkg/app"
)

/*
我们主要是在 dao 层进行了数据访问对象的封装，并针对业务所需的字段进行了处理
*/

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Model: &model.Model{CreatedBy: createBy},
		Name:  name,
		State: state,
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
		Name:  name,
		State: state,
	}
	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := &model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

package service

/*
业务接口校验:
	required	必填
	gt			大于
	gte			大于等于
	lt			小于
	lte			小于等于
	min			最小值
	max			最大值
	oneof		参数集内的其中之一
	len			长度要求与 len 给定的一致
在结构体中应用到了两个 tag 标签，分别是 form 和 binding，它们分别代表着表单的映射字段名和入参校验的规则内容，其主要功能是实现参数绑定和参数检验。
*/

type CountArticleRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
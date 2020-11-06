package model

import "github.com/catherine.li/go_blog/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State string `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

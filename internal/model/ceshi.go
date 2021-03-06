package model

import (
	"time"
)

type Ceshi struct {
	Id         int       `gorm:"AUTO_INCREMENT;column:id;type:INT;primary_key" json:"id"`
	Uri        string    `gorm:"column:uri" json:"uri"`
	Tag 	   int 		 `json:"tag" gorm:"column:tag"`
	BackName   string    `gorm:"column:back_name" json:"backName"`
	IsDelete   int64     `gorm:"column:is_delete" json:"isDelete"`
	CreateTime int64     `gorm:"create_time" json:"createTime"`
	DeleteTime int64     `gorm:"delete_time" json:"deleteTime"`
	ModifyTime time.Time `gorm:"column:modify_time" json:"modifyTime"`
}

func (Ceshi) TableName() string {
	return "ceshi"
}

type CeshiWith struct {
	Ceshi
	CeshiDemo []CeshiDemo1 `gorm:"ForeignKey:uri;References:uri"`
}

type CeshiEs struct {
	Id         int       `json:"id"`
	Uri        string    `json:"uri"`
	Tag 	   int 		 `json:"tag"`
	BackName   string    `json:"backName"`
	IsDelete   int64     `json:"isDelete"`
	CreateTime int64     `json:"createTime"`
	DeleteTime int64     `json:"deleteTime"`
}

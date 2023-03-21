package base

import "time"

//基础结构体
type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` //主键id
	CreatedAt time.Time `json:"created_at"`           //创建时间
	UpdatedAt time.Time `json:"update_at"`            //更新时间
}

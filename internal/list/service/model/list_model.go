package model

import "time"

type ListModel struct {
	ID          uint64     `gorm:"type:bigint(20) auto_increment; primaryKey; comment:'id'" json:"id"`
	TodoContext string     `gorm:"type:varchar(128); not null; default:''; comment:'待办事项'" json:"todoContext"`
	TodoStatus  uint8      `gorm:"type:tinyint(4); not null; default:'0'; comment:'待办事项状态 0: 待处理 1: 已完成'" json:"todoStatus"`
	DeletedAt   *time.Time `gorm:"comment:'删除时间'" json:"-"`
	CreatedAt   time.Time  `gorm:"default:current_timestamp; comment:'创建时间'" json:"-"`
	UpdatedAt   time.Time  `gorm:"default:current_timestamp;comment:'更新时间'" json:"-"`
}

func (l *ListModel) TableName() string {
	return "list"
}

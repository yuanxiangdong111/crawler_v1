package dao

import (
    "time"
)

type User struct {
    ID int64 // 主键
    // 通过在字段后面的标签说明，定义golang字段和表字段的关系
    // 例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
    Username string `gorm:"column:username"`
    Password string `gorm:"column:password"`
    // 创建时间，时间戳
    CreateTime int64 `gorm:"column:createtime"`
}

type GalaxyBoardPerson struct {
    ID        uint64    `gorm:"primary_key;column:id"` // 自增id
    Oid       uint32    `gorm:"column:oid"`            // 机构id
    Uid       string    `gorm:"column:uid"`            // 用户uid
    Name      string    `gorm:"column:name"`           // 用户名称
    Data      string    `gorm:"column:data"`           // 用户白板个性化数据
    CreatedAt time.Time `gorm:"column:created_at"`     // 创建时间
    UpdatedAt time.Time `gorm:"column:updated_at"`     // 更新时间
}

// type ResData struct {
//     FiledData string `gorm:"column:field_data"`
// }

func (GalaxyBoardPerson) TableName() string {
    return "galaxy_board_person"
}

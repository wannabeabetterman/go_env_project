package model

// User
type User struct {
	ID       uint64 `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"id"`
	Username string `gorm:"column:username;type:varchar(32);not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(900);not null" json:"password"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

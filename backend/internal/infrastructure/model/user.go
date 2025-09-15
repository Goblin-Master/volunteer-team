package model

type User struct {
	ID       int    `gorm:"primaryKey;column:id;type:int;comment:'主键ID'"`
	UserID   int64  `gorm:"column:user_id;type:bigint;unique,comment:'用户ID'"`
	Ctime    int64  `gorm:"column:ctime;type:bigint;comment:'创建时间'"`
	Utime    int64  `gorm:"column:utime;type:bigint;comment:'更新时间'"`
	Account  string `gorm:"column:account;type:varchar(20);unique;comment:'账号'"`
	Password string `gorm:"column:password;type:varchar(20);not null;comment:'密码'"`
	Email    string `gorm:"column:email;type:varchar(20);unique;comment:'邮箱'"`
	Username string `gorm:"column:username;type:varchar(20);not null;comment:'用户名'"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null;comment:'头像'"`
	Role     int8   `gorm:"column:role;type:tinyint;default:2;comment:'角色'"`
}

func (u *User) TableName() string {
	return "user"
}

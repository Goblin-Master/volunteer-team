package model

type Order struct {
	ID                 int    `gorm:"column:id;type:bigint;comment:'主键ID';primaryKey"`                                                                             // 主键
	OrderID            int64  `gorm:"column:order_id;type:bigint;not null;comment:'订单ID';unique,uniqueIndex:idx_order_state,priority:1"`                           //订单id
	UserID             int64  `gorm:"column:user_id;type:bigint;not null;comment:'用户ID'"`                                                                          // 逻辑外键
	Ctime              int64  `gorm:"column:ctime;type:bigint;not null;comment:'创建时间';index:idx_state_time,priority:2;"`                                           // 创建时间
	Utime              int64  `gorm:"column:utime;type:bigint;comment:'更新时间'"`                                                                                     // 更新时间
	Username           string `gorm:"column:username;type:varchar(20);not null;comment:'报修人姓名'"`                                                                   // 报修人姓名
	StudentID          string `gorm:"column:student_id;type:varchar(20);not null;comment:'学号';index"`                                                              // 学号
	CampusLocation     string `gorm:"column:campus_location;type:varchar(10);not null;comment:'校区'"`                                                               // 校区
	Department         string `gorm:"column:department;type:varchar(30);not null;comment:'学院'"`                                                                    // 学院
	PhoneNumber        string `gorm:"column:phone_number;type:varchar(11);not null;comment:'手机号'"`                                                                 // 手机号
	WechatID           string `gorm:"column:wechat_id;type:varchar(30);comment:'微信号'"`                                                                             // 微信号
	Address            string `gorm:"column:address;type:varchar(255);not null;comment:'详细地址'"`                                                                    // 详细地址
	DeviceModel        string `gorm:"column:device_model;type:varchar(100);not null;comment:'设备型号'"`                                                               // 设备型号
	OSVersion          string `gorm:"column:os_version;type:varchar(20);not null;comment:'操作系统版本'"`                                                                // 系统版本
	ProblemDescription string `gorm:"column:problem_description;type:text;not null;comment:'问题描述'"`                                                                // 问题描述
	Notes              string `gorm:"column:notes;type:text;comment:'备注'"`                                                                                         // 备注
	State              int8   `gorm:"column:state;type:tinyint;default:1;comment:'订单处理状态';index:idx_state_time,priority:1,uniqueIndex:idx_order_state,priority:2"` //订单状态
}

// TableName 设置表名
func (o *Order) TableName() string {
	return "order"
}

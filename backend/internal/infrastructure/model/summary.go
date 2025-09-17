package model

// Summary 维修小结表
type Summary struct {
	ID                 int    `gorm:"primaryKey;column:id;type:bigint;comment:'主键ID'"`                                        // 主键
	OrderID            int    `gorm:"column:order_id;type:bigint;not null;uniqueIndex:idx_order_id;comment:'关联订单ID'"`         // 逻辑外键 → order.id
	UserID             int64  `gorm:"column:user_id;type:bigint;not null;index:idx_user_ctime,priority:1;comment:'用户ID（冗余）'"` // 冗余用户 ID，方便查询
	Ctime              int64  `gorm:"column:ctime;type:bigint;not null;index:idx_user_ctime,priority:2;comment:'创建时间'"`       // 创建时间
	Utime              int64  `gorm:"column:utime;type:bigint;not null;index;comment:'更新时间'"`                                 // 更新时间
	ProblemType        string `gorm:"column:problem_type;type:varchar(50);not null;comment:'问题分类'"`                           // 如：硬件故障 / 软件异常 / 网络问题
	ProblemDescription string `gorm:"column:problem_description;type:text;not null;comment:'问题描述'"`                           // 工程师填写的问题小结
	RepairSummary      string `gorm:"column:repair_summary;type:text;not null;comment:'维修小结'"`                                // 维修过程、解决方案、更换配件等
	ReceiverName       string `gorm:"column:receiver_name;type:varchar(30);not null;comment:'负责人姓名'"`                         // 接待人
}

func (s *Summary) TableName() string {
	return "summary"
}

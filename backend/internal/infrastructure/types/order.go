package types

type CreateOrderReq struct {
	Username           string `json:"username"`
	StudentID          string `json:"student_id"`
	CampusLocation     string `json:"campus_location"`
	Department         string `json:"department"` //学院
	PhoneNumber        string `json:"phone_number"`
	WechatID           string `json:"wechat_id"`
	Address            string `json:"address"`             //详细地址
	DeviceModel        string `json:"device_model"`        //设备型号
	OSVersion          string `json:"os_version"`          //操作系统类型
	ProblemDescription string `json:"problem_description"` //具体问题描述
	Notes              string `json:"notes"`               //备注
}

type OrderItem struct {
	ID                 int    `json:"id"`
	Ctime              int64  `json:"ctime"`
	ProblemDescription string `json:"problem_description"`
}

type OrderListResp struct {
	Orders []OrderItem `json:"orders"`
}
type OrderDetailReq struct {
	ID int `form:"id"`
}
type OrderDetailResp struct {
	Username           string `json:"username"`
	StudentID          string `json:"student_id"`
	CampusLocation     string `json:"campus_location"`
	Department         string `json:"department"` //学院
	PhoneNumber        string `json:"phone_number"`
	WachatID           string `json:"wachat_id"`
	Address            string `json:"address"`             //详细地址
	DeviceModel        string `json:"device_model"`        //设备型号
	OSVersion          string `json:"os_version"`          //操作系统类型
	ProblemDescription string `json:"problem_description"` //具体问题描述
	Notes              string `json:"notes"`               //备注
	Ctime              int64  `json:"ctime"`               //创建时间
}

type FinishOrderReq struct {
	ID int `form:"id"`
}

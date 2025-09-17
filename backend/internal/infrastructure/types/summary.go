package types

type CreateSummaryReq struct {
	OrderID            int    `json:"order_id" binding:"required"`
	ProblemType        string `json:"problem_type" binding:"required"`
	ProblemDescription string `json:"problem_description" binding:"required"`
	RepairSummary      string `json:"repair_summary" binding:"required"`
	ReceiverName       string `json:"receiver_name" binding:"required"`
}

type SummaryItem struct {
	ID                 int    `json:"id"`
	OrderID            int    `json:"order_id" `
	ProblemType        string `json:"problem_type"`
	ProblemDescription string `json:"problem_description" `
	RepairSummary      string `json:"repair_summary" `
	ReceiverName       string `json:"receiver_name" `
	Utime              int64  `json:"utime" `
}
type SummaryListResp struct {
	Summaries []SummaryItem `json:"summaries"`
}
type SummaryDetailReq struct {
	ID int `form:"id" binding:"required"`
}
type SummaryDetailResp struct {
	ID                 int    `json:"id"`
	OrderID            int    `json:"order_id" `
	ProblemType        string `json:"problem_type"`
	ProblemDescription string `json:"problem_description" `
	RepairSummary      string `json:"repair_summary" `
	ReceiverName       string `json:"receiver_name" `
	Utime              int64  `json:"utime" `
}

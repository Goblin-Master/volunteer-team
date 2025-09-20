package types

type CreateSummaryReq struct {
	OrderID            string `json:"order_id" binding:"required"`
	ProblemType        string `json:"problem_type" binding:"required"`
	ProblemDescription string `json:"problem_description" binding:"required"`
	RepairSummary      string `json:"repair_summary" binding:"required"`
	ReceiverName       string `json:"receiver_name" binding:"required"`
}

type SummaryItem struct {
	SummaryID          int64  `json:"summary_id,string"`
	OrderID            int64  `json:"order_id,string"`
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
	SummaryID string `form:"summary_id" binding:"required"`
}
type SummaryDetailResp struct {
	SummaryID          int64  `json:"summary_id,string"`
	OrderID            int64  `json:"order_id,string"`
	ProblemType        string `json:"problem_type"`
	ProblemDescription string `json:"problem_description" `
	RepairSummary      string `json:"repair_summary" `
	ReceiverName       string `json:"receiver_name" `
	Utime              int64  `json:"utime" `
}

type UpdateSummaryReq struct {
	SummaryID          string `json:"summary_id" binding:"required"`
	ProblemType        string `json:"problem_type"`
	ProblemDescription string `json:"problem_description" `
	RepairSummary      string `json:"repair_summary" `
	ReceiverName       string `json:"receiver_name" `
}

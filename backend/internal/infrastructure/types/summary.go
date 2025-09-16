package types

type CreateSummaryReq struct {
	OrderID            int    `json:"order_id" binding:"required"`
	ProblemType        string `json:"problem_type" binding:"required"`
	ProblemDescription string `json:"problem_description" binding:"required"`
	RepairSummary      string `json:"repair_summary" binding:"required"`
	ReceiverName       string `json:"receiver_name" binding:"required"`
}

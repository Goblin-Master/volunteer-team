package handler

import "volunteer-team/backend/internal/logic"

type SummaryHandler struct {
	summaryLogic *logic.SummaryLogic
}

func NewSummaryHandler() *SummaryHandler {
	return &SummaryHandler{
		summaryLogic: logic.NewSummaryLogic(),
	}
}

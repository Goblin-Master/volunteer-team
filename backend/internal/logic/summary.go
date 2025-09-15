package logic

import "volunteer-team/backend/internal/repo"

type ISummaryLogic interface{}
type SummaryLogic struct {
	summaryRepo *repo.SummaryRepo
}

func NewSummaryLogic() *SummaryLogic {
	return &SummaryLogic{
		summaryRepo: repo.NewSummaryRepo(),
	}
}

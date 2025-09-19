package handler

import (
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/middleware"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/response"
	"volunteer-team/backend/internal/logic"

	"github.com/gin-gonic/gin"
)

type SummaryHandler struct {
	summaryLogic *logic.SummaryLogic
}

func NewSummaryHandler() *SummaryHandler {
	return &SummaryHandler{
		summaryLogic: logic.NewSummaryLogic(),
	}
}

func (sh *SummaryHandler) CreateSummary(c *gin.Context) {
	cr := middleware.GetBind[types.CreateSummaryReq](c)
	global.Log.Info(cr)
	resp, err := sh.summaryLogic.CreateSummary(c.Request.Context(), jwtx.GetUserID(c), cr)
	response.Response(c, resp, err)
}

func (sh *SummaryHandler) GetSummaryList(c *gin.Context) {
	resp, err := sh.summaryLogic.GetSummaryList(c.Request.Context())
	response.Response(c, resp, err)
}

func (sh *SummaryHandler) GetSummaryDetail(c *gin.Context) {
	cr := middleware.GetBind[types.SummaryDetailReq](c)
	global.Log.Info(cr)
	resp, err := sh.summaryLogic.GetSummaryDetail(c.Request.Context(), cr.ID)
	response.Response(c, resp, err)
}

func (sh *SummaryHandler) UpdateSummary(c *gin.Context) {
	cr := middleware.GetBind[types.UpdateSummaryReq](c)
	global.Log.Info(cr)
	resp, err := sh.summaryLogic.UpdateSummary(c.Request.Context(), cr)
	response.Response(c, resp, err)
}

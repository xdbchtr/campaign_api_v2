package handler

import (
	"campaign_api/campaign"
	"campaign_api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *CampaignHandler {
	return &CampaignHandler{service}
}

func (h *CampaignHandler) GetCampaigns(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get Campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("List of Campaigns", http.StatusBadRequest, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

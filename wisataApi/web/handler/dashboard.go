package handler

import (
	"net/http"
	"wisataapi/category"
	"wisataapi/itinerary"
	"wisataapi/tourist"
	"wisataapi/user"

	"github.com/gin-gonic/gin"
)

type dashboardHandler struct {
	itineraryService itinerary.Service
	userService      user.Service
	categoryService  category.Service
	touristService   tourist.Service
}

func NewDashboardHandler(itineraryService itinerary.Service, userService user.Service, categoryService category.Service, touristService tourist.Service) *dashboardHandler {
	return &dashboardHandler{itineraryService, userService, categoryService, touristService}
}

func (h *dashboardHandler) Index(c *gin.Context) {

	users, err := h.userService.GetAllUsersByadmin()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	category, err := h.categoryService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	tourist, err := h.touristService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	itinerary, err := h.itineraryService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{"user": len(users), "category": len(category), "tourist": len(tourist), "itinerary": len(itinerary)})
}

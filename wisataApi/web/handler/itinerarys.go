package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"wisataapi/itinerary"
	"wisataapi/user"

	"github.com/gin-gonic/gin"
)

type itineraryHandler struct {
	itineraryService itinerary.Service
	userService      user.Service
}

func NewItineraryHandler(itineraryService itinerary.Service, userService user.Service) *itineraryHandler {
	return &itineraryHandler{itineraryService, userService}
}

func (h *itineraryHandler) Index(c *gin.Context) {
	users, err := h.itineraryService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	fmt.Println(users)
	c.HTML(http.StatusOK, "itinerarys_index.html", gin.H{"itinerarys": users})
}

func (h *itineraryHandler) Detail(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	itinerarys, err := h.itineraryService.FindByIDUser(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	timeline, err := h.itineraryService.FindByIDTimeline(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userProfile := user.Profile{
		ID: itinerarys.UserId,
	}

	user, err := h.userService.Profile(userProfile)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	fmt.Println(timeline)

	c.HTML(http.StatusOK, "itinerarys_detail.html", gin.H{
		"initiallocal": itinerarys.InitialLocal,
		"startday":     itinerarys.StartDay,
		"endday":       itinerarys.EndDay,
		"starttime":    itinerarys.StartTime,
		"endtime":      itinerarys.EndTime,
		"timeline":     timeline,
		"name":         user.Name,
	})
}

package handler

import (
	"fmt"
	"net/http"
	"wisataapi/helper"
	itinerary "wisataapi/itinerary"

	"github.com/gin-gonic/gin"
)

type itineraryHandler struct {
	service itinerary.Service
}

func NewItineraryHandler(service itinerary.Service) *itineraryHandler {
	return &itineraryHandler{service}
}

func (h *itineraryHandler) FindAllItinerary(c *gin.Context) {

	itinerarys, err := h.service.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get itinerary", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of itinerary", http.StatusOK, "success", itinerary.FormatItinerarys(itinerarys))
	c.JSON(http.StatusOK, response)
}

func (h *itineraryHandler) FindAllItineraryByUser(c *gin.Context) {
	var input itinerary.FindAllItineraryByUser

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get itinerary", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	itinerarys, err := h.service.FindAllByIDUser(input.ID)
	if err != nil {
		response := helper.APIResponse("Error to get itinerary", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of itinerary by user", http.StatusOK, "success", itinerary.FormatItinerarys(itinerarys))
	c.JSON(http.StatusOK, response)
}

func (h *itineraryHandler) SaveItinerary(c *gin.Context) {
	var input itinerary.CreateItineraryInput
	fmt.Println(input)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create itinerary", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newItinerary, err := h.service.Save(input)
	if err != nil {
		response := helper.APIResponse("Failed to create itinerary", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create itinerary", http.StatusOK, "success", itinerary.FormatItinerary(newItinerary))
	c.JSON(http.StatusOK, response)
}

func (h *itineraryHandler) UpdateItinerary(c *gin.Context) {
	var input itinerary.UpdateItineraryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update itinerary", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedItinerary, err := h.service.Update(input)
	if err != nil {
		response := helper.APIResponse("Failed to update itinerary", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update itinerary", http.StatusOK, "success", itinerary.FormatItinerary(updatedItinerary))
	c.JSON(http.StatusOK, response)
}

func (h *itineraryHandler) DeleteItinerary(c *gin.Context) {
	var inputID itinerary.DeleteItineraryInput

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete itinerary", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fmt.Println(inputID)
	deleteItinerary, err := h.service.Delete(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete itinerary", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete itinerary", http.StatusOK, "success", itinerary.FormatItinerary(deleteItinerary))
	c.JSON(http.StatusOK, response)
}

func (h *itineraryHandler) SaveTimeline(c *gin.Context) {
	var input itinerary.CreateTimelineInput
	fmt.Println(input)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create Timeline", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTimeline, err := h.service.SaveTimeline(input)
	if err != nil {
		response := helper.APIResponse("Failed to create Timeline", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create Timeline", http.StatusOK, "success", itinerary.FormatTimeline(newTimeline))
	c.JSON(http.StatusOK, response)
}

func (h *itineraryHandler) UpdateTimeline(c *gin.Context) {
	var input itinerary.UpdateTimelineInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update Timeline", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedItinerary, err := h.service.UpdateTimeline(input)
	if err != nil {
		response := helper.APIResponse("Failed to update Timeline", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update Timeline", http.StatusOK, "success", itinerary.FormatTimeline(updatedItinerary))
	c.JSON(http.StatusOK, response)
}

func (h *itineraryHandler) DeleteTimeline(c *gin.Context) {
	var inputID itinerary.DeleteItineraryInput

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete timeline", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fmt.Println(inputID)
	deleteItinerary, err := h.service.DeleteTimeline(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete timeline", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete timeline", http.StatusOK, "success", itinerary.FormatTimeline(deleteItinerary))
	c.JSON(http.StatusOK, response)
}

func (h *itineraryHandler) FindByIDTimeline(c *gin.Context) {
	var inputID itinerary.DeleteItineraryInput

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get timeline", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fmt.Println(inputID)
	deleteItinerary, err := h.service.FindByIDTimeline(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get timeline", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get timeline", http.StatusOK, "success", itinerary.FormatTimelines(deleteItinerary))
	c.JSON(http.StatusOK, response)
}

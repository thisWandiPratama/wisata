package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"wisataapi/helper"
	tourist "wisataapi/tourist"

	"github.com/gin-gonic/gin"
)

type touristHandler struct {
	service tourist.Service
}

func NewTouristHandler(service tourist.Service) *touristHandler {
	return &touristHandler{service}
}

func (h *touristHandler) FindAllTourist(c *gin.Context) {

	campaigns, err := h.service.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get tourist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of tourist", http.StatusOK, "success", tourist.FormatTourists(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *touristHandler) SaveTourist(c *gin.Context) {
	file, err := c.FormFile("image_primary")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image primary", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	pathAvatar := fmt.Sprintf("images/%s", strings.Replace(file.Filename, " ", "", -1))

	err = c.SaveUploadedFile(file, pathAvatar)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	value3 := c.PostForm("category_id")
	number3, err := strconv.ParseInt(value3, 10, 32)
	category_id := int(number3)

	input := tourist.CreateTouristInput{
		CategoryID:   category_id,
		Name:         c.PostForm("name"),
		Description:  c.PostForm("description"),
		Address:      c.PostForm("address"),
		Email:        c.PostForm("email"),
		Phone:        c.PostForm("phone"),
		Website:      c.PostForm("website"),
		Latitude:     c.PostForm("latitude"),
		Longitude:    c.PostForm("longitude"),
		LinkVideo:    c.PostForm("link_video"),
		ImagePrimary: pathAvatar,
	}

	newCampaign, err := h.service.Save(input)
	if err != nil {
		response := helper.APIResponse("Failed to create tourist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create tourist", http.StatusOK, "success", tourist.FormatTourist(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *touristHandler) UpdateTourist(c *gin.Context) {
	file, err := c.FormFile("image_primary")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image primary", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	pathAvatar := fmt.Sprintf("images/%s", strings.Replace(file.Filename, " ", "", -1))

	err = c.SaveUploadedFile(file, pathAvatar)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	value3 := c.PostForm("category_id")
	number3, err := strconv.ParseInt(value3, 10, 32)
	category_id := int(number3)

	value1 := c.PostForm("tourist_id")
	number1, err := strconv.ParseInt(value1, 10, 32)
	tourist_id := int(number1)

	input := tourist.UpdateTouristInput{
		ID:           tourist_id,
		CategoryID:   category_id,
		Name:         c.PostForm("name"),
		Description:  c.PostForm("description"),
		Address:      c.PostForm("address"),
		Email:        c.PostForm("email"),
		Phone:        c.PostForm("phone"),
		Website:      c.PostForm("website"),
		Latitude:     c.PostForm("latitude"),
		Longitude:    c.PostForm("longitude"),
		LinkVideo:    c.PostForm("link_video"),
		ImagePrimary: pathAvatar,
	}

	updatedCampaign, err := h.service.Update(input)
	if err != nil {
		response := helper.APIResponse("Failed to update category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update category", http.StatusOK, "success", tourist.FormatTourist(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *touristHandler) DeleteTourist(c *gin.Context) {
	var inputID tourist.DeleteTouristInput

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete tourist", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fmt.Println(inputID)
	deleteCampaign, err := h.service.Delete(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete tourist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete tourist", http.StatusOK, "success", tourist.FormatTourist(deleteCampaign))
	c.JSON(http.StatusOK, response)
}

// gallery
func (h *touristHandler) FindAllGalleryTourist(c *gin.Context) {
	var inputID tourist.FindAllGalleryTouristByID

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete categori", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	campaigns, err := h.service.FindAllGallery(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Error to get tourist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List image galleri of tourist", http.StatusOK, "success", tourist.FormatGalleryTourists(campaigns))
	c.JSON(http.StatusOK, response)
}

// save gallery
func (h *touristHandler) SaveGalleryTourist(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image gallery", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	pathAvatar := fmt.Sprintf("images/%s", strings.Replace(file.Filename, " ", "", -1))

	err = c.SaveUploadedFile(file, pathAvatar)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image gallery", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	value3 := c.PostForm("tourist_id")
	number3, err := strconv.ParseInt(value3, 10, 32)
	tourist_id := int(number3)

	input := tourist.Gallery{
		TouristID: tourist_id,
		Avatar:    pathAvatar,
	}

	newCampaign, err := h.service.SaveGallery(input)
	if err != nil {
		response := helper.APIResponse("Failed to create tourist image galllery", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create tourist image galllery", http.StatusOK, "success", tourist.GalleryTouristFormatter(newCampaign))
	c.JSON(http.StatusOK, response)
}

// delete gallery tourist
func (h *touristHandler) DeleteGalleryTourist(c *gin.Context) {
	var inputID tourist.DeleteGalleryTouristInput

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete tourist image gallery", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deleteCampaign, err := h.service.DeleteGallery(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete tourist image gallery", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete tourist image gallery", http.StatusOK, "success", tourist.FormatGalleryTourist(deleteCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *touristHandler) FindAllByCategory(c *gin.Context) {
	var inputID tourist.AllTouristByCategory

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get tourist by categori", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fmt.Println(inputID)
	deleteCampaign, err := h.service.FindAllByCategory(inputID.CategoryID)
	if err != nil {
		response := helper.APIResponse("Failed to get tourist by categori", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get tourist by categori", http.StatusOK, "success", tourist.FormatTourists(deleteCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *touristHandler) Search1(c *gin.Context) {
	var inputID tourist.Search

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get tourist by search", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	campaigns, err := h.service.Search(inputID.Name)
	if err != nil {
		response := helper.APIResponse("Error to get tourist", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of tourist", http.StatusOK, "success", tourist.FormatTourists(campaigns))
	c.JSON(http.StatusOK, response)
}

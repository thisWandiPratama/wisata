package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	category "wisataapi/category"
	"wisataapi/helper"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service category.Service
}

func NewCampaignHandler(service category.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) FindAll(c *gin.Context) {

	campaigns, err := h.service.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of categori", http.StatusOK, "success", category.FormatCategorys(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) Save(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

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

	input := category.CreateCategoryInput{
		Name:   c.PostForm("name"),
		Avatar: pathAvatar,
	}

	newCampaign, err := h.service.Save(input)
	if err != nil {
		response := helper.APIResponse("Failed to create categori", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create categori", http.StatusOK, "success", category.FormatCategory(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) Update(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

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

	value3 := c.PostForm("id")
	number3, err := strconv.ParseInt(value3, 10, 32)
	logika := int(number3)

	input := category.UpdateCategoryInput{
		ID:     logika,
		Name:   c.PostForm("name"),
		Avatar: pathAvatar,
	}

	updatedCampaign, err := h.service.Update(input)
	if err != nil {
		response := helper.APIResponse("Failed to update category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update category", http.StatusOK, "success", category.FormatCategory(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) Delete(c *gin.Context) {
	var inputID category.DeleteCategoryInput

	err := c.ShouldBindJSON(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete categori", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fmt.Println(inputID)
	deleteCampaign, err := h.service.Delete(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete categori", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete categori", http.StatusOK, "success", category.FormatCategory(deleteCampaign))
	c.JSON(http.StatusOK, response)
}

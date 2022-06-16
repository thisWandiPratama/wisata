package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"wisataapi/category"
	"wisataapi/helper"
	"wisataapi/user"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.Service
}

func NewCategoryHandler(categoryService category.Service) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (h *categoryHandler) Index(c *gin.Context) {
	users, err := h.categoryService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "category_index.html", gin.H{"users": users})
}

func (h *categoryHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "category_new.html", nil)
}

func (h *categoryHandler) Create(c *gin.Context) {

	file, err := c.FormFile("avatar")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	path := fmt.Sprintf("images/%s", file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	name := c.PostForm("name")

	createCategory := category.CreateCategoryInput{}
	createCategory.Name = name
	createCategory.Avatar = path

	_, err = h.categoryService.Save(createCategory)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/categorys")
}
func (h *categoryHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	file, err := c.FormFile("avatar")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	path := fmt.Sprintf("images/%s", file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	name := c.PostForm("name")

	updateCategory := category.UpdateCategoryInput{}
	updateCategory.ID = id
	updateCategory.Name = name
	updateCategory.Avatar = path

	_, err = h.categoryService.Update(updateCategory)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/categorys")
}

func (h *categoryHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	registeredUser, err := h.categoryService.FindByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := user.FormUpdateUserInput{}
	input.ID = registeredUser.ID
	input.Name = registeredUser.Name
	input.Email = registeredUser.Avatar

	c.HTML(http.StatusOK, "category_edit.html", input)
}

func (h *categoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	_, err := h.categoryService.Delete(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/categorys")
}

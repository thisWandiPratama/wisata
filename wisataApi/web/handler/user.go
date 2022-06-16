package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"wisataapi/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsersByadmin()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}

func (h *userHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "user_new.html", nil)
}

func (h *userHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	registerInput := user.RegisterUserInput{}
	registerInput.Name = name
	registerInput.Email = email
	registerInput.Password = password
	registerInput.Phone = phone

	_, err := h.userService.RegisterUser(registerInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	registeredUser, err := h.userService.GetUserByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := user.FormUpdateUserInput{}
	input.ID = registeredUser.ID
	input.Name = registeredUser.Name
	input.Email = registeredUser.Email
	input.Phone = registeredUser.Phone

	c.HTML(http.StatusOK, "user_edit.html", input)
}

func (h *userHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input user.FormUpdateUserInput

	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")

	input.ID = id
	input.Name = name
	input.Email = email
	input.Phone = phone

	_, err := h.userService.UpdateUser(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) NewAvatar(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	c.HTML(http.StatusOK, "user_avatar.html", gin.H{"ID": id})
}

func (h *userHandler) CreateAvatar(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	file, err := c.FormFile("avatar")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userID := id

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	_, err := h.userService.Delete(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/users")
}

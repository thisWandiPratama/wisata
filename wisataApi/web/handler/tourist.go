package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"wisataapi/category"
	"wisataapi/helper"
	"wisataapi/tourist"

	"github.com/gin-gonic/gin"
)

type touristHandler struct {
	touristService  tourist.Service
	categoryService category.Service
}

func NewTouristHandler(touristService tourist.Service, categoryService category.Service) *touristHandler {
	return &touristHandler{touristService, categoryService}
}

func (h *touristHandler) Index(c *gin.Context) {
	users, err := h.touristService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "tourist_index.html", gin.H{"users": users})
}

func (h *touristHandler) New(c *gin.Context) {

	category, err := h.categoryService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "tourist_new.html", gin.H{"categorys": category})
}

func (h *touristHandler) Create(c *gin.Context) {

	file, err := c.FormFile("imageprimary")
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
	idParam := c.PostForm("categoryid")
	categoryid, _ := strconv.Atoi(idParam)

	name := c.PostForm("name")
	description := c.PostForm("description")
	address := c.PostForm("address")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	website := c.PostForm("website")
	latitude := c.PostForm("latitude")
	longitude := c.PostForm("longitude")
	linkvideo := c.PostForm("linkvideo")

	fmt.Println(categoryid)

	createTourist := tourist.CreateTouristInput{}
	createTourist.CategoryID = categoryid
	createTourist.Name = name
	createTourist.Description = description
	createTourist.Address = address
	createTourist.Email = email
	createTourist.Phone = phone
	createTourist.Website = website
	createTourist.Latitude = latitude
	createTourist.Longitude = longitude
	createTourist.LinkVideo = linkvideo
	createTourist.ImagePrimary = path

	_, err = h.touristService.Save(createTourist)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/tourists")
}
func (h *touristHandler) Update(c *gin.Context) {
	file, err := c.FormFile("imageprimary")
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
	idParam := c.PostForm("categoryid")
	categoryid, _ := strconv.Atoi(idParam)

	idParam22 := c.Param("id")
	id, _ := strconv.Atoi(idParam22)

	name := c.PostForm("name")
	description := c.PostForm("description")
	address := c.PostForm("address")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	website := c.PostForm("website")
	latitude := c.PostForm("latitude")
	longitude := c.PostForm("longitude")
	linkvideo := c.PostForm("linkvideo")

	fmt.Println(categoryid)

	createTourist := tourist.UpdateTouristInput{}
	createTourist.ID = id
	createTourist.CategoryID = categoryid
	createTourist.Name = name
	createTourist.Description = description
	createTourist.Address = address
	createTourist.Email = email
	createTourist.Phone = phone
	createTourist.Website = website
	createTourist.Latitude = latitude
	createTourist.Longitude = longitude
	createTourist.LinkVideo = linkvideo
	createTourist.ImagePrimary = path

	_, err = h.touristService.Update(createTourist)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/tourists")
}

func (h *touristHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	category, err := h.categoryService.FindAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	tourists, err := h.touristService.FindByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "tourist_edit.html", gin.H{
		"ID":          tourists.ID,
		"categorys":   category,
		"name":        tourists.Name,
		"description": tourists.Description,
		"categoryid":  tourists.CategoryID,
		"address":     tourists.Address,
		"email":       tourists.Email,
		"phone":       tourists.Phone,
		"website":     tourists.Website,
		"latitude":    tourists.Latitude,
		"longitude":   tourists.Longitude,
		"linkvideo":   tourists.LinkVideo,
	})
}

func (h *touristHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	_, err := h.touristService.Delete(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/tourists")
}

func (h *touristHandler) Detail(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	fmt.Println(id)

	tourist, err := h.touristService.FindByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	touristgalerry, err := h.touristService.FindAllGallery(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	fmt.Println(touristgalerry)

	c.HTML(http.StatusOK, "tourist_detail.html", gin.H{
		"name":        tourist.Name,
		"description": tourist.Description,
		"address":     tourist.Address,
		"email":       tourist.Email,
		"phone":       tourist.Phone,
		"website":     tourist.Website,
		"latitude":    tourist.Latitude,
		"longitude":   tourist.Longitude,
		"video":       tourist.LinkVideo,
		"image":       tourist.ImagePrimary,
		"gallery":     touristgalerry,
	})
}

func (h *touristHandler) NewGallery(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	gallery, err := h.touristService.FindAllGallery(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	fmt.Println("gallery")
	fmt.Println(id)
	fmt.Println(gallery)
	c.HTML(http.StatusOK, "tourist_avatar.html", gin.H{"gallerys": gallery, "ID": id})
}

func (h *touristHandler) Gallery(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	fmt.Println(id)

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

	saveCategory := tourist.Gallery{}
	saveCategory.TouristID = id
	saveCategory.Avatar = path

	_, err = h.touristService.SaveGallery(saveCategory)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/tourists/newgallery/"+idParam)
}

package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"wisataapi/auth"
	category "wisataapi/category"
	"wisataapi/handler"
	"wisataapi/helper"
	"wisataapi/itinerary"
	"wisataapi/tourist"
	"wisataapi/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "monalisa:monalisa@tcp(127.0.0.1:3306)/wisata?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	categoryRepository := category.NewRepository(db)
	touristRepository := tourist.NewRepository(db)
	itineraryRepository := itinerary.NewRepository(db)

	userService := user.NewService(userRepository)
	categoryService := category.NewService(categoryRepository)
	touristService := tourist.NewService(touristRepository)
	itineraryService := itinerary.NewService(itineraryRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	categoryHandler := handler.NewCampaignHandler(categoryService)
	touristHandler := handler.NewTouristHandler(touristService)
	itineraryHandler := handler.NewItineraryHandler(itineraryService)

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("wisataapi", cookieStore))

	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	api := router.Group("/api/v1")
	// users
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/updateavatar", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/profile", authMiddleware(authService, userService), userHandler.FetchUser)

	//category
	api.GET("/all_category", categoryHandler.FindAll)
	api.POST("/add_category", authMiddleware(authService, userService), categoryHandler.Save)
	api.PUT("/update_category", authMiddleware(authService, userService), categoryHandler.Update)
	api.POST("/delete_category", authMiddleware(authService, userService), categoryHandler.Delete)

	//tourist
	api.GET("/all_tourist", touristHandler.FindAllTourist)
	api.POST("/search_tourist", touristHandler.Search1)
	api.POST("/all_tourist_by_categori", touristHandler.FindAllByCategory)
	api.POST("/add_tourist", authMiddleware(authService, userService), touristHandler.SaveTourist)
	api.PUT("/update_tourist", authMiddleware(authService, userService), touristHandler.UpdateTourist)
	api.POST("/delete_tourist", authMiddleware(authService, userService), touristHandler.DeleteTourist)

	// gallery image tourist
	api.POST("/all_image_gallery_tourist", touristHandler.FindAllGalleryTourist)
	api.POST("/add_image_gallery_tourist", authMiddleware(authService, userService), touristHandler.SaveGalleryTourist)
	api.POST("/delete_image_gallery_tourist", authMiddleware(authService, userService), touristHandler.DeleteGalleryTourist)

	// itinerary
	api.GET("/all_itinerary", itineraryHandler.FindAllItinerary)
	api.POST("/all_itinerary_by_user", itineraryHandler.FindAllItineraryByUser)
	api.POST("/add_itinerary", itineraryHandler.SaveItinerary)
	api.PUT("/update_itinerary", itineraryHandler.UpdateItinerary)
	api.POST("/delete_itinerary", itineraryHandler.DeleteItinerary)

	// timeline
	api.POST("/timeline_by_itinerary_id", itineraryHandler.FindByIDTimeline)
	api.POST("/add_timeline", itineraryHandler.SaveTimeline)
	api.PUT("/update_timeline", itineraryHandler.UpdateTimeline)
	api.POST("/delete_timeline", itineraryHandler.DeleteTimeline)
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

func authAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userIDSession := session.Get("userID")

		if userIDSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

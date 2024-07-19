package webserver

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Epyklab/trident/agent/utils"
	logger "github.com/Epyklab/trident/agent/utils/logging"
	"github.com/gin-gonic/gin"
)

func loggingMiddleware(lineChan chan<- string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		messsage := fmt.Sprintf("%s %s %s %s %s",
			c.PostForm("username"),
			c.PostForm("password"),
			c.Request.Proto,
			c.ClientIP(),
			c.Request.Proto)
		entry := logger.Entry(messsage,
			"webserver")

		lineChan <- string(entry)

	}
}

func Router(lineChan chan string) *gin.Engine {
	router := setupRouter(lineChan)

	// Define routes and associate handlers
	router.GET("/", homeHandler)
	router.GET("/login", loginHandler)
	router.POST("/login", loginHandler) // Handle POST request on the login page
	router.GET("/search", searchHandler)
	router.POST("/upload", uploadHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	return router
}

func setupRouter(lineChan chan string) *gin.Engine {
	conf, _ := utils.ParseConfig()

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// Apply the logging middleware
	router.Use(loggingMiddleware(lineChan))

	// Load HTML templates
	router.LoadHTMLGlob(conf.Webfiles)

	return router
}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"title": "Home"})
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		// Example: Log the username when a POST request is made to the login page
		username := c.PostForm("username")
		log.Printf("Login attempt with username: %s", username)
	}
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "Login"})
}

func searchHandler(c *gin.Context) {
	query := c.Query("q")
	// Log the search query directly from the handler as an example
	log.Printf("Search query: %s", query)
	c.HTML(http.StatusOK, "search.html", gin.H{"title": "Search"})
}

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// Define the directory and filename where the file will be stored
	filename := filepath.Base(file.Filename)
	targetPath := filepath.Join("/opt/trident/uploads", filename)

	if err := c.SaveUploadedFile(file, targetPath); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", filename))
}

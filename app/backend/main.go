package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *gorm.DB
var err error

func main() {
	// SQLite veritabanına bağlan
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}

	// Kullanıcı tablosunu oluştur
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}
	// Gin router oluştur
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(config))

	// Tüm kullanıcıları getir
	router.GET("/users", getAllUsers)

	// Belirli bir kullanıcıyı getir
	router.GET("/users/:id", getUserByID)

	// Yeni kullanıcı oluştur
	router.POST("/users", createUser)

	// Belirli bir kullanıcıyı güncelle
	router.PUT("/users/:id", updateUser)

	// Belirli bir kullanıcıyı sil
	router.DELETE("/users/:id", deleteUser)

	// API'yi 8080 portunda başlat
	router.Run(":8080")
}

// Tüm kullanıcıları getir
func getAllUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Belirli bir kullanıcıyı getir
func getUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Yeni kullanıcı oluştur
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// Belirli bir kullanıcıyı güncelle
func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

// Belirli bir kullanıcıyı sil
func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user User
	result := db.Delete(&user, id)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

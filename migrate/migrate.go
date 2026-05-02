package main

import (
	"github.com/anirbandas/go-crud/initializers"
	"github.com/anirbandas/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

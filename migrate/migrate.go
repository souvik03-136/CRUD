package migrate

import (
	"github.com/souvik03-136/CRUD/initializers"
	"github.com/souvik03-136/CRUD/models"
)

func init() {
	initializers.LoadEnvVariavles()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

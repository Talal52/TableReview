// routes/routes.go
package routes

import (
	"fmt"
	"net/http"
	"tablereviews/database"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *database.DB) {
	router.POST("/add-review", func(c *gin.Context) {
		var newReview database.Review
		if err := c.ShouldBindJSON(&newReview); err != nil {
			fmt.Println("error binding JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received JSON payload: %+v\n", newReview)

		if err := db.AddReview(newReview); err != nil {
			fmt.Println("error adding review:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add review"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Review added successfully"})
	})

}

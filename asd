// // routes/routes.go
// package routes

// import (
// 	"tablereviews/database"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func SetupRoutes(router *gin.Engine, db *database.DB) {
// 	// ... existing routes

// 	// Add a new route for getting a review by email
// 	router.GET("/get-review/:email", func(c *gin.Context) {
// 		email := c.Param("email")
// 		review, err := db.GetReviewByEmail(email)
// 		if err != nil {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, review)
// 	})

// 	// Add a new route for deleting a review by email
// 	router.DELETE("/delete-review/:email", func(c *gin.Context) {
// 		email := c.Param("email")
// 		err := db.DeleteReviewByEmail(email)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
// 	})
// }






// /////////////////////////

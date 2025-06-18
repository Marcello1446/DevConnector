package main

import (
	"DevConnector/database"
	"DevConnector/internal/handlers/CommentHandlers"
	"DevConnector/internal/handlers/PostHandlers"
	"DevConnector/internal/middleware/IDMatching"

	//"DevConnector/internal/handlers/PostHandlers"
	"DevConnector/internal/handlers/ProfileHandlers"
	"DevConnector/internal/middleware/Duplicates"
	"DevConnector/internal/middleware/JWT"
	"DevConnector/internal/validators"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
}

func main() {
	InitEnv()
	database.InitDB()

	e := echo.New()
	profile := e.Group("/profile")
	post := e.Group("/post")
	comment := e.Group("/comment")

	e.POST("/register", ProfileHandlers.CreateProfile, validators.ValidateProfileCreating, Duplicates.CheckForDuplicates)
	e.POST("/login", ProfileHandlers.Login, validators.ValidateProfileLoggining)
	e.GET("/validate", ProfileHandlers.ValidateExistance, JWT.RequireProfileAuth)
	e.POST("/logout", ProfileHandlers.Logout)
	e.GET("/follow/:id", ProfileHandlers.FollowProfile)
	e.GET("/unfollow/:id", ProfileHandlers.UnfollowProfile)
	e.GET("/posts", PostHandlers.GetAllPosts)
	e.GET("/comments/:id", CommentHandlers.GetComments)
	profile.PATCH("/update/:id", ProfileHandlers.UpdateProfile)
	profile.GET("/get/:id", ProfileHandlers.GetProfileById)
	profile.GET("/getfollowers", ProfileHandlers.GetFollowers)
	profile.GET("/getfollowings", ProfileHandlers.GetFollowings)
	profile.POST("/post", PostHandlers.CreatePost, JWT.RequireProfileAuth)
	post.GET("/:id", PostHandlers.GetPost)
	post.GET("/like/:id", PostHandlers.LikePost, JWT.RequireProfileAuth)
	post.GET("/dislike/:id", PostHandlers.UnlikePost, JWT.RequireProfileAuth)
	post.PATCH("/update/:id", PostHandlers.UpdatePost, JWT.RequireProfileAuth, IDMatching.CheckForID)
	comment.POST("/:id", CommentHandlers.AddComment, JWT.RequireProfileAuth)
	comment.DELETE("/:postID/:commentID", CommentHandlers.DeleteComment, JWT.RequireProfileAuth)

	err := e.Start(":8080")

	if err != nil {
		log.Fatal("Failed to start a server")
	}
}

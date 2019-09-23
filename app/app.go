package app

import (
	"bitbucket.org/kyicy/seifer/app/router"
	"github.com/labstack/echo/v4"
)

// RegisterRoute function
func RegisterRoute(e *echo.Echo) {
	e.POST("/user_story", router.CreateUserStory)
	e.POST("/user_story/similar", router.SimilarUserStories)
}

//ExpandRoute is func
func ExpandRoute(e *echo.Echo) {
	e.POST("/expand", router.CreateUserStory)
	e.POST("/expand/similar", router.SimilarUserStories)
}

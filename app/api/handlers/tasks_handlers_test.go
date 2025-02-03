package handlers

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/pkg/session"
	"github.com/golang-documented-todo-api/app/repository"
	"github.com/stretchr/testify/assert"
)

func fakeAuth(c *fiber.Ctx) error {
	session.StoreSession(c, repository.SelectUserBySessionIDRow{
		Username: "Eduarte",
	})
	return c.Next()
}

func Test_tasks_handlers(t *testing.T) {
	mockTasks := new(taskServiceMock)
	app := fiber.New()
	app.Post("", fakeAuth, PostTasks(mockTasks))

	// test if it'll reject a post that is not application.json
	req, _ := http.NewRequest("POST", "", nil)
	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, 400, res.StatusCode)
}

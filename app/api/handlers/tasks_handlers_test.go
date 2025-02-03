package handlers

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_tasks_handlers(t *testing.T) {
	app := fiber.New()
	app.Post("", PostTasks())

	// test if it'll reject a post that is not application.json
	req, _ := http.NewRequest("POST", "", nil)
	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, 400, res.StatusCode)

}

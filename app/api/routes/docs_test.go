package routes

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestDocsRoutes(t *testing.T) {
	app := fiber.New()
	DocsRouter(app)
	req, _ := http.NewRequest("GET", "/docs/openapi.yaml", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

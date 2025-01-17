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
	// test openapi.yaml
	req, _ := http.NewRequest("GET", "/docs/openapi.yaml", nil)
	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	// test reference
	req, _ = http.NewRequest("GET", "/docs/reference", nil)
	res, err = app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

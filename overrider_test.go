package overrider

import (
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var app *fiber.App

func TestMain(m *testing.M) {
	setupFiber()
	code := m.Run()

	os.Exit(code)
}

func setupFiber() {
	app = fiber.New()

	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("get")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("post")
	})

	app.Put("/", func(c *fiber.Ctx) error {
		return c.SendString("put")
	})

	app.Patch("/", func(c *fiber.Ctx) error {
		return c.SendString("patch")
	})

	app.Delete("/", func(c *fiber.Ctx) error {
		return c.SendString("delete")
	})
}

func TestNew(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		method       string
		respBody     string
		expectedCode int
	}{
		{
			description:  "send GET get HTTP status 200",
			route:        "/",
			method:       "GET",
			respBody:     "get",
			expectedCode: 200,
		},
		{
			description:  "send PUT get HTTP status 200",
			route:        "/?_method=PUT",
			method:       "POST",
			respBody:     "put",
			expectedCode: 200,
		},
		{
			description:  "send PATCH get HTTP status 200",
			route:        "/?_method=PATCH",
			method:       "POST",
			respBody:     "patch",
			expectedCode: 200,
		},
		{
			description:  "send DELETE get HTTP status 200",
			route:        "/?_method=DELETE",
			method:       "POST",
			respBody:     "delete",
			expectedCode: 200,
		},
	}
	for _, tt := range tests {

		req := httptest.NewRequest(tt.method, tt.route, nil)
		resp, err := app.Test(req)
		require.NoError(t, err)

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		require.NoError(t, err)

		assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
		assert.Equal(t, tt.respBody, string(body))
	}
}

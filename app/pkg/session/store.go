package session

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-documented-todo-api/app/repository"
)

type sessionStoreKey = uint8

const (
	storeKey sessionStoreKey = iota
)

func StoreSession(c *fiber.Ctx, sessionData repository.SelectUserBySessionIDRow) {
	c.Locals(storeKey, sessionData)
}

func GetStoredSession(c *fiber.Ctx) (repository.SelectUserBySessionIDRow, bool) {
	session, ok := c.Locals(storeKey).(repository.SelectUserBySessionIDRow)
	return session, ok
}

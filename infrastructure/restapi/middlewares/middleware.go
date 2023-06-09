// Package middlewares contains the middlewares for the restapi api
package middlewares

import (
	"blockchain-go/utils/lists"

	secureDomain "blockchain-go/domain/security"

	authConst "blockchain-go/utils/constant/auth"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

func AuthJWTMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not provided"})
		}

		claims := &secureDomain.Claims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secureDomain.PublicKey, nil
		})

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		ctx.Locals(authConst.Authorized, claims)

		return ctx.Next()
	}
}

// AuthRoleMiddleware is a function that validates the role of user
func AuthRoleMiddleware(validRoles []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

		if authData.Role == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You are not authorized"})
		}

		if !lists.Contains(validRoles, authData.Role) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You are not authorized for this path"})
		}

		return ctx.Next()
	}
}

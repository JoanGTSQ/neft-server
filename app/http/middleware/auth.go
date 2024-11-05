package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

// AuthMiddleware verifica si el usuario está autenticado usando un token
func AuthMiddleware() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Header("token")
		_, err := facades.Auth(ctx).Parse(token)
		if err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"error": "Usuario no válido",
			})
			return
		}

		var user models.User
		err = facades.Auth(ctx).User(&user)
		if err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"error": "Usuario no válido",
			})
			// Abortamos si no se pudo obtener el usuario
			return
		}

		// Almacena el usuario en el contexto
		ctx.WithValue("user", user)

		// Log para verificar que el usuario se ha almacenado
		facades.Log().Debug("Usuario autenticado:", user)

		ctx.Request().Next() // Continuar con la siguiente función middleware o controlador
	}
}

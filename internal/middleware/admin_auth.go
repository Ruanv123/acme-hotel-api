package middleware

import (
	"fmt"
	"net/http"

	"github.com/ruanv123/acme-hotel-api/internal/service"
)

func AdminMiddleware(authService service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := extractTokenFromHeader(r)
			if tokenString == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			user, err := authService.VerifyTokenAdmin(tokenString)
			if err != nil {
				fmt.Print(err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := service.WithUserContext(r.Context(), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

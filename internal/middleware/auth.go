package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ruanv123/acme-hotel-api/internal/service"
)

func AuthMiddleware(authService service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := extractTokenFromHeader(r)
			if tokenString == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			user, err := authService.VerifyToken(tokenString)
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

func extractTokenFromHeader(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

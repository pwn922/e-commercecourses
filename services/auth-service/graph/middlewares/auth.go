package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/pwn922/auth-service/internal/services/jwt"
)

// contextKey es un tipo para evitar colisiones en el contexto.
type contextKey string

// userIDKey es la clave utilizada para almacenar y recuperar el ID de usuario del contexto.
const UserIDKey = contextKey("userID") // Cambiado a exportado

// AuthMiddleware contiene el servicio JWT para la autenticación.
type AuthMiddleware struct {
	JWTService *jwt.JWTService
}

// NewAuthMiddleware crea una nueva instancia de AuthMiddleware.
func NewAuthMiddleware(jwtService *jwt.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		JWTService: jwtService,
	}
}

// CORS Middleware
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Cambia '*' por tu dominio específico en producción
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Si la solicitud es un preflight (OPTIONS), solo retornamos una respuesta 200
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Authenticate es un middleware que valida el token JWT y establece el ID de usuario en el contexto.
func (a *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		println("Authenticate")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := a.JWTService.VerifyAccessToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Almacena el userID en el contexto
		ctx = context.WithValue(ctx, UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

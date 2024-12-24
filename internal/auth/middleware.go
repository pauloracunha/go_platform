package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Permissions define todas as permissões disponíveis no sistema
var Permissions = map[string]map[string]map[string]string{
	"user": {
		"register": {
			"permission":  "user_register",
			"description": "Cria novos usuários no sistema",
		},
		"update": {
			"permission":  "user_update",
			"description": "Edita informações de um usuário existente",
		},
		"delete": {
			"permission":  "user_delete",
			"description": "Remove um usuário do sistema",
		},
		"view": {
			"permission":  "user_view",
			"description": "Visualiza informações dos usuários",
		},
	},
}

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Configura o contexto para usar informações do token
		next.ServeHTTP(w, r)
	})
}

func PermissionCheck(permission string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userPermissions := claims["permissions"].([]interface{})
				hasPermission := false
				for _, p := range userPermissions {
					if p == permission {
						hasPermission = true
						break
					}
				}
				if !hasPermission {
					http.Error(w, "Permissão negada", http.StatusForbidden)
					return
				}
			} else {
				http.Error(w, "Token inválido", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

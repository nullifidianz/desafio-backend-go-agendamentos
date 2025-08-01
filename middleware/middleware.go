package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nullifidianz/desafio-backend-go-agendamentos/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Token ausente"})
			return
		}

		strToken := strings.TrimPrefix(header, "Bearer ")
		claims, err := auth.ValidarToken(strToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Token inválido"})
			return
		}

		c.Set("usuario_id", claims.UsuarioId)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func Autorizar(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleCtx, existe := c.Get("role")
		if !existe {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"erro": "Acesso negado"})
			return
		}

		role := roleCtx.(string)
		for _, permitido := range roles {
			if role == permitido {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"erro": "Permissão insuficiente"})
	}
}

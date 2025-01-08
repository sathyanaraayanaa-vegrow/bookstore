package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract JWT token from headers
		type User struct {
			UserID string `json:"user_id"`
			Role   string `json:"role"`
		}
		user := &User{}
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: User Not Logged In. LOGIN to continue."})
			c.Abort()
			return
		}

		// jsonBody := []byte(fmt.Sprintf(`{"token": "%s"}`, tokenString))
		// bodyReader := bytes.NewBuffer(jsonBody)

		req, err := http.NewRequest(http.MethodPost, os.Getenv("RailApiUrl"), nil)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Invalid Request to Authorise"})
			c.Abort()
			return
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", tokenString)

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Invalid Request to Authorise"})
			c.Abort()
			return
		}
		if res.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Could not validate the Token."})
			c.Abort()
			return
		}
		derr := json.NewDecoder(res.Body).Decode(user)
		if derr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Invalid Request to Authorise"})
			c.Abort()
			return
		}
		u, e := strconv.Atoi(user.UserID)
		if e != nil {
			return
		}
		c.Set("userID", u)
		c.Set("roles", user.Role)
		// fmt.Println(user.UserID, user.Role)
		// Validate JWT token
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	// Add your JWT secret key here for token validation
		// 	return []byte(os.Getenv("JWT")), nil
		// })
		// if err != nil || !token.Valid {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Invalid or expired token"})
		// 	c.Abort()
		// 	return
		// }

		// // Extract JWT claims
		// claims, ok := token.Claims.(jwt.MapClaims)
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Invalid token claims"})
		// 	c.Abort()
		// 	return
		// }
		// // Store user ID in Gin context
		// userID, ok := claims["sub"].(string)
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: User ID missing in token"})
		// 	c.Abort()
		// 	return
		// }
		// user, err := strconv.Atoi(userID)
		// if err != nil {
		// 	return
		// }
		// c.Set("userID", user)

		// roles, ok := claims["role"].(string)
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"Msg": "Error: Roles missing in token"})
		// 	c.Abort()
		// 	return
		// }
		// c.Set("roles", roles)

		c.Next()
	}
}

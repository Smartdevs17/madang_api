package middleware

import (
	"fmt"
	"madang_api/config"
	"madang_api/models"
	"madang_api/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	fmt.Println("Middleware running successfully")

	//Get the token from the request header
	requestToken := c.GetHeader("Authorization")

	if requestToken == "" {
		utils.AbortResponse(c, http.StatusUnauthorized, "access token required")
		return
	}

	tokenString := strings.Split(requestToken, "Bearer ")[1]
	// fmt.Println(tokenString)
	if tokenString == "" {
		utils.AbortResponse(c, http.StatusUnauthorized, "access token required")
		return
	}

	fmt.Println(tokenString)

	//Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		utils.AbortResponse(c, http.StatusUnauthorized, "invalid access token")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

		}

		//find the user with the token
		var user models.User
		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Attach to the req
		c.Set("user", user)

		//Continue
		c.Next()
	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)

	}
}

package pkg

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
)

func AuthenticationMiddlware(c *gin.Context){
	token := c.GetHeader("Authorization")
	fmt.Println("tokenStringtokenString",token)
	if token == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	prefix:="Bearer "
	if !strings.HasPrefix(token,prefix){
		c.JSON(401, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}
	tokenString:=token[len(prefix):]

	claims,err:=ParseToken(tokenString)
	fmt.Println("claimsclaims",claims,err)
	if err != nil{
		c.JSON(401,gin.H{"error":"Unauthorized"})
		c.Abort()
		return
	}
	
	c.Set("username", claims.Username)
	c.Next()
}
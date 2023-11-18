package pkg

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
// In-memory user storage (replace with a database in production)
var users = make(map[string]User) 

func LoginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fields"})
		return
	}
	// Check if username and password are provided
	if user.Username == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Username and password are mandatory"})
		return
	}

	// Authenticate user (validate username and password)
	fmt.Println(users)

	
	userfound,exists:=users[user.Username]
	if !exists{
		c.JSON(400,gin.H{"error":"user not found,please register"})
		return
	}
	err:=bcrypt.CompareHashAndPassword([]byte(userfound.Password),[]byte(user.Password))
	fmt.Println("hashedPassword",err)
	if err != nil{
		c.JSON(400,gin.H{"error":"password is incorrect"})
		return
	}
	// For simplicity, assume valid credentials and generate a token
	token, err := GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


func RegisterHandler(c *gin.Context){
	// the sructure of data you need to recive on post
	var user User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid fields"})
		return
	}
	// Check if username and password are provided
	if user.Username == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Username and password are mandatory"})
		return
	}
	//check for the username
	if _,exists:=users[user.Username];exists{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Username already exists"})
		return
	}

	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err!=nil{
		c.JSON(500,gin.H{"error":"cant hash passowrd"})
		return
	}
	user.Password=string(hashedPassword)
	users[user.Username]=user
	c.JSON(200,gin.H{"message":"User successfully registerd"})

}

func ProtectedHandler(c *gin.Context){
	username, _ := c.Get("username")
	c.JSON(200,gin.H{"message":fmt.Sprintf("Hello %s",username)})
}
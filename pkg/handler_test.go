package pkg

func TestLoginHandler(t *testing.T) {
    // Create a mock gin context
    ginContextMock := mockGinContext{}
    user=User{
        "Username":"ajayan",
        "Password":"AS@12345"
    }
    ginContextMock.BindJSON(user)

    // Call the LoginHandler function
    LoginHandler(&ginContextMock)

    // Assert that the correct response was sent
    if ginContextMock.Response.Code != http.StatusOK {
        t.Errorf("Expected status code to be 200, but got %d", ginContextMock.Response.Code)
    }
}

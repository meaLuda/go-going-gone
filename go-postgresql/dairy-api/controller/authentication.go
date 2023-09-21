package controller

import(
	"dairyapi/model"
	"github.com/gin-gonic/gin"
    "net/http"
)


// PI Controllers
func Register(context *gin.Context){
	var input model.AuthenticationInput

	if err := context.ShouldBingJson(&input); err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	user := model.User{
        Username: input.Username,
        Password: input.Password,
    }

	savedUser,err := user.Save()

	if err != nil{
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	context.JSON(http.StatusCreated,gin.H{"user":savedUser})
}
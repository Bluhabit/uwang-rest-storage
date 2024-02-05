package profile

import (
	"github.com/Bluhabit/uwang-rest-storage/models"
	"github.com/Bluhabit/uwang-rest-storage/repositories/profile"
	"github.com/gin-gonic/gin"
)

func UploadProfilePicture(context *gin.Context) {
	profileRepository := profile.Init()
	var response = models.BaseResponse[string]{}
	session_id := context.GetString("session_id")

	if len(session_id) < 1 {
		context.JSON(401, response.BadRequest("", "User belum login."))
		return
	}

	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(200, response.BadRequest("", "File tidak boleh kosong."))
		return
	}
	uploadFile := profileRepository.UploadProfilePicture(session_id, file)
	context.JSON(200, uploadFile)
}

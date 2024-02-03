package routes

import (
	ctx "context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func UploadProfilePicture(context *gin.Context) {
	c := ctx.Background()
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(200, gin.H{
			"status_code": 1002,
			"data":        false,
			"message":     "file tidak ditemukan",
		})
	}

	user_id := context.PostForm("user_id")
	if len(user_id) < 1 {
		context.JSON(200, gin.H{
			"status_code": 1002,
			"data":        false,
			"message":     "file tidak ditemukan",
		})
	}
	endpoint := "storage.bluhabit.id"
	accessKeyID := "RNyTQZeXrsPhP3xQElCr"
	secretAccessKey := "pnuDPidX1yUO8NpiGtUVBE2ZCDHbaKIMftZPcjAi"
	useSSL := true

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		context.JSON(200, gin.H{
			"status_code": 1002,
			"data":        false,
			"message":     "file tidak ditemukan",
		})
	}

	f, err := os.OpenFile(fmt.Sprintf("./uploads/%s.png", user_id), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		context.JSON(200, gin.H{
			"status_code": 1002,
			"data":        false,
			"message":     "file tidak ditemukan",
		})
		return
	}
	defer f.Close()

	reader, err := file.Open()
	if err != nil {
		context.JSON(200, gin.H{
			"status_code": 1002,
			"data":        false,
			"message":     "file tidak ditemukan",
		})
	}
	defer reader.Close()
	_, err = io.Copy(f, reader)
	if err != nil {
		context.JSON(200, gin.H{
			"status_code": 1002,
			"data":        false,
			"message":     "file tidak ditemukan",
		})
	}

	bucketName := "uwang-dev"
	objectName := fmt.Sprintf("profile-picture/%s.png",user_id)
	filePath := fmt.Sprintf("uploads/%s.png", user_id)
	contentType := "image/png"

	info, err := minioClient.FPutObject(c, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
		context.JSON(200, gin.H{
			"status_code": 1002,
			"data":        false,
			"message":     "file tidak ditemukan",
		})
		return
	}


	// cari cara hapus file di uploads
	err = os.Remove(fmt.Sprintf("./uploads/%s.png", user_id))
	if err != nil {
		context.JSON(200, gin.H{
			"status_code": 1002,
			"data":        false,
			"message":     err,
		})
	}

	context.JSON(200, gin.H{
		"status_code": 1001,
		"data":        info,
		"message":     "Berhasil",
	})
}

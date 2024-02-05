package profile

import (
	ctx "context"
	"fmt"
	"github.com/Bluhabit/uwang-rest-storage/common"
	"github.com/Bluhabit/uwang-rest-storage/models"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
)

type ProfileRespository struct {
	db    *gorm.DB
	cache *redis.Client
	minio *minio.Client
}

func Init() *ProfileRespository {
	dbConn := common.GetDbConnection()
	redis := common.GetRedisConnection()
	minio := common.GetMinioClient()

	return &ProfileRespository{
		db:    dbConn,
		cache: redis,
		minio: minio,
	}

}

func (repo *ProfileRespository) UploadProfilePicture(session_id string, file *multipart.FileHeader) models.BaseResponse[string] {
	var response = models.BaseResponse[string]{}

	redis_key := common.CreateRedisKeyUserSession(session_id)
	session := repo.cache.HGetAll(ctx.Background(), redis_key)
	user := session.Val()
	userId := user["user_id"]
	fmt.Println(session)
	if len(userId) < 1 {
		return response.BadRequest("", "Sesi tidak ditemukan.")
	}
	f, err := os.OpenFile(fmt.Sprintf("./data/%s.png", userId), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return response.BadRequest("", "Gagal menyimpan file")
	}
	defer f.Close()

	reader, err := file.Open()
	if err != nil {
		return response.BadRequest("", "Gagal menyimpan file [1].")
	}
	defer reader.Close()
	_, err = io.Copy(f, reader)
	if err != nil {
		return response.BadRequest("", "Gagal menyimpan file [2].")
	}

	bucketName := "uwang-dev"
	objectName := fmt.Sprintf("profile-picture/%s.png", userId)
	filePath := fmt.Sprintf("data/%s.png", userId)
	contentType := "image/png"

	fileInfo, err := repo.minio.FPutObject(ctx.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return response.BadRequest("", "Gagal menyimpan file [3].")
	}

	// cari cara hapus file di uploads
	err = os.Remove(fmt.Sprintf("./data/%s.png", userId))
	if err != nil {
		return response.BadRequest("", "Gagal menyimpan file [4].")
	}

	fmt.Print(fileInfo)
	return response.Success(fmt.Sprintf("%s.png", userId), "Berhasil menyimpan foto.")
}

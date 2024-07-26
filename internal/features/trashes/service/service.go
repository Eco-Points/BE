package service

import (
	"eco_points/internal/features/trashes"
	"eco_points/internal/utils"
	"log"
	"mime/multipart"
)

type trashService struct {
	qry trashes.QueryInterface
}

func NewTrashService(q trashes.QueryInterface) trashes.ServiceInterface {
	return &trashService{
		qry: q,
	}
}

func (t *trashService) AddTrash(tData trashes.TrashEntity, file *multipart.FileHeader) error {

	src, err := file.Open()
	if err != nil {
		log.Println("error when open image", err)
		return err
	}
	defer src.Close()

	imageUrl, err := utils.UploadToCloudinary(src, file.Filename)
	if err != nil {
		log.Println("error when upload image", err)
		return err
	}
	tData.Image = imageUrl

	err = t.qry.AddTrash(tData)
	if err != nil {
		log.Println("Error Add Trash", err)
	}
	return nil
}

package service

import (
	"eco_points/internal/features/trashes"
	"eco_points/internal/utils"
	"log"
	"mime/multipart"
)

type trashService struct {
	qry   trashes.QueryTrashInterface
	cloud utils.CloudinaryUtilityInterface
}

func NewTrashService(q trashes.QueryTrashInterface, c utils.CloudinaryUtilityInterface) trashes.ServiceTrashInterface {
	return &trashService{
		qry:   q,
		cloud: c,
	}
}

func (t *trashService) AddTrash(tData trashes.TrashEntity, file *multipart.FileHeader) error {

	src, err := file.Open()
	if err != nil {
		log.Println("error when open image", err)
		return err
	}
	defer src.Close()

	imageUrl, err := t.cloud.UploadToCloudinary(src, file.Filename)
	if err != nil {
		log.Println("error when upload image", err)
		return err
	}
	tData.Image = imageUrl

	err = t.qry.AddTrash(tData)
	if err != nil {
		log.Println("Error Add Trash", err)
		return err
	}
	return nil
}

func (s *trashService) GetTrash(ttype string) (trashes.ListTrashEntity, error) {
	var result trashes.ListTrashEntity
	var err error
	if ttype != "" {
		result, err = s.qry.GetTrashbyType(ttype)
		if err != nil {
			log.Println("Error Get Trash by Type ", ttype, err)
			return trashes.ListTrashEntity{}, err

		}
	} else {
		result, err = s.qry.GetTrashLimit()
		if err != nil {
			log.Println("Error Get All ", err)
			return trashes.ListTrashEntity{}, err
		}
	}
	return result, err
}

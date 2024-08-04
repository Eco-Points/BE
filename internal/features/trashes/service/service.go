package service

import (
	"eco_points/internal/features/trashes"
	"eco_points/internal/utils"
	"log"
	"mime/multipart"
)

type TrashService struct {
	qry   trashes.QueryTrashInterface
	cloud utils.CloudinaryUtilityInterface
}

func NewTrashService(q trashes.QueryTrashInterface, c utils.CloudinaryUtilityInterface) trashes.ServiceTrashInterface {
	return &TrashService{
		qry:   q,
		cloud: c,
	}
}

func (t *TrashService) AddTrash(tData trashes.TrashEntity, file *multipart.FileHeader) error {

	src, err := t.cloud.FileOpener(file)
	if err != nil {
		log.Println("error when open image", err)
		return err
	}

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

func (s *TrashService) GetTrash(ttype string) (trashes.ListTrashEntity, error) {
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

func (s *TrashService) DeleteTrash(id uint) error {
	err := s.qry.DeleteTrash(id)
	if err != nil {
		log.Println("Error Delete Trash ", err)
		return err
	}

	return nil
}

func (s *TrashService) UpdateTrash(id uint, data trashes.TrashEntity) error {

	err := s.qry.UpdateTrash(id, data)
	if err != nil {
		log.Println("Error Update Trash ", err)
		return err
	}

	return nil
}

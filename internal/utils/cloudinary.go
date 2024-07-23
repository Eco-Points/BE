package utils

import (
	"context"
	"eco_points/config"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file io.Reader, filename string) (string, error) {
	cld, err := cloudinary.NewFromURL(config.ImportSetting().CldKey)

	if err != nil {
		return "", err
	}

	// upload file to cloudinary
	uploadParams := uploader.UploadParams{
		Folder:   "user_picture",
		PublicID: filename,
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", err
	}

	publicURL := uploadResult.URL
	return publicURL, nil
}

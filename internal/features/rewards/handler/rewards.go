package handler

import (
	"eco_points/internal/features/rewards"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type RewardHandler struct {
	srv rewards.RServices
	tu  utils.JwtUtilityInterface
	cu  utils.CloudinaryUtilityInterface
}

func NewRewardHandler(s rewards.RServices, t utils.JwtUtilityInterface, c utils.CloudinaryUtilityInterface) rewards.RHandler {
	return &RewardHandler{
		srv: s,
		tu:  t,
		cu:  c,
	}
}

func (rh *RewardHandler) AddReward() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := rh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("error from jwt")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "error", "Unauthorized", nil))
		}

		// Get image from form data
		image, err := c.FormFile("image")
		if err != nil {
			log.Println("error get image file")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid image file", nil))
		}

		// Open the image file
		src, err := image.Open()
		if err != nil {
			log.Println("error open the image file")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to open image file", nil))
		}
		defer src.Close()

		// Upload image to Cloudinary
		imageURL, err := rh.cu.UploadToCloudinary(src, image.Filename)
		if err != nil {
			log.Println("error upload image to cloudinary")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to upload image", nil))
		}

		// Bind the request
		var req CreateOrUpdateRewardRequest
		err = c.Bind(&req)
		if err != nil {
			log.Println("invalid reqeust body")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid request body", nil))
		}

		// Convert request to reward model
		newReward := ToRewardModel(req, imageURL)

		// Add reward to database
		err = rh.srv.AddReward(newReward)
		if err != nil {
			log.Println("failed add reward to database")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Failed to add reward", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Reward was successfully created", nil))

	}
}

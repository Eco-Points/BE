package handler

import (
	"eco_points/internal/features/rewards"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type RewardHandler struct {
	srv rewards.RServices
	tu  utils.JwtUtilityInterface
}

func NewRewardHandler(s rewards.RServices, t utils.JwtUtilityInterface) rewards.RHandler {
	return &RewardHandler{
		srv: s,
		tu:  t,
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

		// Bind the request
		var req CreateOrUpdateRewardRequest
		err = c.Bind(&req)
		if err != nil {
			log.Println("invalid request body")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid request body", nil))
		}

		// Convert request to reward model
		newReward := ToRewardModel(req, "")

		// Add reward to database
		err = rh.srv.AddReward(newReward, src, image.Filename)
		if err != nil {
			log.Println("failed add reward to database")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Failed to add reward", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Reward was successfully created", nil))
	}
}

func (rh *RewardHandler) UpdateReward() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := rh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("error from jwt")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "error", "Unauthorized", nil))
		}

		// Get reward ID from URL parameter
		rewardIDStr := c.Param("id")
		rewardID, err := strconv.ParseUint(rewardIDStr, 10, 32)
		if err != nil {
			log.Println("invalid reward ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid reward ID", nil))
		}

		// Get image from form data (optional)
		var src multipart.File
		var filename string
		image, err := c.FormFile("image")
		if err == nil {
			src, err = image.Open()
			if err != nil {
				log.Println("error open the image file")
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to open image file", nil))
			}
			defer src.Close()
			filename = image.Filename
		}

		// Bind the request
		var req CreateOrUpdateRewardRequest
		err = c.Bind(&req)
		if err != nil {
			log.Println("invalid request body")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid request body", nil))
		}

		// Convert request to reward model
		updatedReward := ToRewardModel(req, "")

		// Update reward in database
		err = rh.srv.UpdateReward(uint(rewardID), updatedReward, src, filename)
		if err != nil {
			log.Println("failed to update reward in database")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to update reward", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Reward was successfully updated", nil))
	}
}

func (rh *RewardHandler) DeleteReward() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := rh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("error from jwt")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "error", "Unauthorized", nil))
		}

		// Get reward ID from URL parameter
		rewardIDStr := c.Param("id")
		rewardID, err := strconv.ParseUint(rewardIDStr, 10, 32)
		if err != nil {
			log.Println("invalid reward ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid reward ID", nil))
		}

		// Delete reward from database
		err = rh.srv.DeleteReward(uint(rewardID))
		if err != nil {
			log.Println("failed to delete reward from database")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to delete reward", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Reward was successfully deleted", nil))
	}
}

func (rh *RewardHandler) GetRewardByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get reward ID from URL parameter
		rewardIDStr := c.Param("id")
		rewardID, err := strconv.ParseUint(rewardIDStr, 10, 32)
		if err != nil {
			log.Println("invalid reward ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid reward ID", nil))
		}

		// Get reward from database
		reward, err := rh.srv.GetRewardByID(uint(rewardID))
		if err != nil {
			log.Println("failed to get reward from database")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to get reward", nil))
		}

		rewardResponse := ToRewardResponse(reward)

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Reward was successfully retrieved", rewardResponse))
	}
}

func (rh *RewardHandler) GetAllRewards() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || limit <= 0 {
			limit = 10 // default limit
		}

		offset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil || offset < 0 {
			offset = 0 // default offset
		}

		rewards, totalItems, err := rh.srv.GetAllRewards(limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "Failed to get rewards", nil))
		}

		responseData := ToRewardsResponse(rewards)
		meta := helpers.Meta{
			TotalItems:   totalItems,
			ItemsPerPage: limit,
			CurrentPage:  offset/limit + 1,
			TotalPages:   (totalItems + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "Successfully get all rewards", responseData, meta))
	}
}

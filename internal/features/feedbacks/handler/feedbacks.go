package handler

import (
	"eco_points/internal/features/feedbacks"
	"eco_points/internal/helpers"
	"eco_points/internal/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type FeedbackHandler struct {
	srv feedbacks.FServices
	tu  utils.JwtUtilityInterface
}

func NewFeedbackHandler(s feedbacks.FServices, t utils.JwtUtilityInterface) feedbacks.FHandler {
	return &FeedbackHandler{
		srv: s,
		tu:  t,
	}
}

func (fh *FeedbackHandler) AddFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := fh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("error from jwt")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "error", "Unauthorized", nil))
		}

		var req CreateOrUpdateFeedbackRequest
		contentType := c.Request().Header.Get("Content-Type")

		// Handle multipart/form-data
		if strings.Contains(contentType, "multipart/form-data") {
			if err := c.Request().ParseMultipartForm(10 << 20); err != nil {
				log.Println("invalid form data")
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid form data", nil))
			}

			// Convert form values to appropriate types
			ratingStr := c.FormValue("rating")
			rating, err := strconv.ParseUint(ratingStr, 10, 32)
			if err != nil {
				log.Println("invalid rating")
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid rating", nil))
			}
			req.Rating = uint(rating)

			req.Review = c.FormValue("review")
		} else if contentType == "application/json" {
			if err := c.Bind(&req); err != nil {
				log.Println("invalid JSON request body")
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid JSON request body", nil))
			}
		} else {
			return c.JSON(http.StatusUnsupportedMediaType, helpers.ResponseFormat(http.StatusUnsupportedMediaType, "failed", "Unsupported Media Type", nil))
		}

		newFeedback := ToFeedbackModel(req)
		newFeedback.UserID = uint(userID)

		err := fh.srv.AddFeedback(newFeedback)
		if err != nil {
			log.Println("failed to add feedback to database")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Failed to add feedback", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "Feedback was successfully created", nil))
	}
}

func (fh *FeedbackHandler) UpdateFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := fh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("error from jwt")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "error", "Unauthorized", nil))
		}

		feedbackIDStr := c.Param("id")
		feedbackID, err := strconv.ParseUint(feedbackIDStr, 10, 32)
		if err != nil {
			log.Println("invalid feedback ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid feedback ID", nil))
		}

		var req CreateOrUpdateFeedbackRequest
		contentType := c.Request().Header.Get("Content-Type")

		// Handle multipart/form-data
		if strings.Contains(contentType, "multipart/form-data") {
			if err := c.Request().ParseMultipartForm(10 << 20); err != nil {
				log.Println("invalid form data")
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid form data", nil))
			}

			// Convert form values to appropriate types
			ratingStr := c.FormValue("rating")
			rating, err := strconv.ParseUint(ratingStr, 10, 32)
			if err != nil {
				log.Println("invalid rating")
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid rating", nil))
			}
			req.Rating = uint(rating)

			req.Review = c.FormValue("review")
		} else if contentType == "application/json" {
			if err := c.Bind(&req); err != nil {
				log.Println("invalid JSON request body")
				return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid JSON request body", nil))
			}
		} else {
			return c.JSON(http.StatusUnsupportedMediaType, helpers.ResponseFormat(http.StatusUnsupportedMediaType, "failed", "Unsupported Media Type", nil))
		}

		updatedFeedback := ToFeedbackModel(req)

		err = fh.srv.UpdateFeedback(uint(feedbackID), updatedFeedback)
		if err != nil {
			log.Println("failed to update feedback in database")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to update feedback", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Feedback was successfully updated", nil))
	}
}

func (fh *FeedbackHandler) DeleteFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := fh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("error from jwt")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "error", "Unauthorized", nil))
		}

		feedbackIDStr := c.Param("id")
		feedbackID, err := strconv.ParseUint(feedbackIDStr, 10, 32)
		if err != nil {
			log.Println("invalid feedback ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid feedback ID", nil))
		}

		err = fh.srv.DeleteFeedback(uint(feedbackID))
		if err != nil {
			log.Println("failed to delete feedback from database")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to delete feedback", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Feedback was successfully deleted", nil))
	}
}

func (fh *FeedbackHandler) GetFeedbackByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		feedbackIDStr := c.Param("id")
		feedbackID, err := strconv.ParseUint(feedbackIDStr, 10, 32)
		if err != nil {
			log.Println("invalid feedback ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "Invalid feedback ID", nil))
		}

		feedback, err := fh.srv.GetFeedbackByID(uint(feedbackID))
		if err != nil {
			log.Println("failed to get feedback from database")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to get feedback", nil))
		}

		feedbackResponse := ToFeedbackResponse(feedback)

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "Feedback was successfully retrieved", feedbackResponse))
	}
}

func (fh *FeedbackHandler) GetAllFeedbacks() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || limit <= 0 {
			limit = 10 // default limit
		}

		offset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil || offset < 0 {
			offset = 0 // default offset
		}

		feedbacks, totalItems, err := fh.srv.GetAllFeedbacks(limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "Failed to get feedbacks", nil))
		}

		responseData := ToFeedbacksResponse(feedbacks)
		meta := helpers.Meta{
			TotalItems:   totalItems,
			ItemsPerPage: limit,
			CurrentPage:  offset/limit + 1,
			TotalPages:   (totalItems + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "Successfully get all feedbacks", responseData, meta))
	}
}

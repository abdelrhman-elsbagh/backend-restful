package api

import (
	"awesomeProject/db"
	"awesomeProject/helpers"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	Context *gin.Context
	Queries *db.Queries
}

func NewHandler(c *gin.Context, queries *db.Queries) *Handler {
	return &Handler{
		Context: c,
		Queries: queries,
	}
}

func (h *Handler) CreateNewUser() {
	var createUserParams db.CreateUserParams
	if err := h.Context.ShouldBindJSON(&createUserParams); err != nil {
		h.Context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user, err := h.Queries.CreateUser(context.Background(), createUserParams)

	if err != nil {
		h.Context.JSON(http.StatusInternalServerError, gin.H{"error": "error " + err.Error()})
		return
	}

	if user.ID == 0 {
		h.Context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user this phone number used before"})
		return
	}

	h.Context.JSON(http.StatusCreated, user)
}

func (h *Handler) GenerateOTP() {
	var input struct {
		PhoneNumber string `json:"phone_number"`
	}
	if err := h.Context.ShouldBindJSON(&input); err != nil {
		h.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.Queries.GetUserByPhoneNumber(context.Background(), input.PhoneNumber)
	if err != nil {
		h.Context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	otp := helpers.GenerateRandomOTP()
	user.Otp = otp
	expTime := time.Now().UTC().Add(1 * time.Minute) // Set expiration time to 1 minute from now

	var UpdateUserOTPParams db.UpdateUserOTPParams
	UpdateUserOTPParams.ID = user.ID
	UpdateUserOTPParams.OtpExpirationTime = expTime
	UpdateUserOTPParams.Otp = otp

	err = h.Queries.UpdateUserOTP(context.Background(), UpdateUserOTPParams)
	if err != nil {
		h.Context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	updatedUser, _ := h.Queries.GetUserByPhoneNumber(context.Background(), input.PhoneNumber)

	h.Context.JSON(http.StatusOK, gin.H{"otp": otp, "expiration_time": updatedUser.OtpExpirationTime})
}

func (h *Handler) VerifyOTP() {
	var input struct {
		PhoneNumber string `json:"phone_number"`
		Otp         string `json:"otp"`
	}

	if err := h.Context.ShouldBindJSON(&input); err != nil {
		h.Context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Queries.GetUserByPhoneNumber(context.Background(), input.PhoneNumber)
	if err != nil {
		h.Context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid phone number"})
		return
	}

	// check if OTP is expired
	if user.OtpExpirationTime.After(time.Now().Add(1 * time.Minute).UTC()) {
		h.Context.JSON(http.StatusUnauthorized, gin.H{"error": "OTP expired"})
		return
	}

	// check if OTP is correct
	if user.Otp != input.Otp {
		h.Context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	h.Context.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}

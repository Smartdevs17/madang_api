package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Avatar    string    `json:"avatar"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Register handles user registration
func (controller *UserController) Register(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := utils.ValidateStruct(c, body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	// Generate OTP
	otp := utils.GenerateOTP()
	var user models.User
	user.Name = body.Name
	user.Email = body.Email
	user.Password = body.Password
	user.Role = body.Role
	user.EmailVerificationOTP = otp

	err := controller.UserService.RegisterUser(&user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to register user", err.Error())
		return
	}

	// newUser := UserResponse{
	// 	ID:        user.ID,
	// 	Name:      user.Name,
	// 	Email:     user.Email,
	// 	Phone:     user.Phone,
	// 	Avatar:    user.Avatar,
	// 	Role:      user.Role,
	// 	Active:    user.Active,
	// 	Token:     user.Token,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", otp)
}

// Validate email
func (controller *UserController) ValidateEmail(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
		Otp   string `json:"otp"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := utils.ValidateStruct(c, body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	user, err := controller.UserService.VerifyEmailOTP(body.Email, body.Otp)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to validate email", err.Error())
		return
	}

	loggedInUser := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Active:    user.Active,
		Token:     user.Token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Email validated successfully", loggedInUser)
}

// Login User
func (controller *UserController) Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := utils.ValidateStruct(c, body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	user, err := controller.UserService.LoginUser(body.Email, body.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to login user", err.Error())
		return
	}

	loggedInUser := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Active:    user.Active,
		Token:     user.Token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "User logged in successfully", loggedInUser)
}

// GetUserByID get user by ID
func (controller *UserController) GetUserByID(c *gin.Context) {
	userID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	user, err := controller.UserService.GetUserByID(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Failed to get user", err.Error())
		return
	}

	userResponse := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Active:    user.Active,
		Token:     user.Token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "User retrieved successfully", userResponse)
}

// GetAllUsers get all users
func (controller *UserController) GetAllUsers(c *gin.Context) {
	users, err := controller.UserService.GetAllUsers()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get users", err.Error())
		return
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponse := UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Avatar:    user.Avatar,
			Role:      user.Role,
			Active:    user.Active,
			Token:     user.Token,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		userResponses = append(userResponses, userResponse)
	}

	utils.SuccessResponse(c, http.StatusOK, "Users retrieved successfully", userResponses)
}

// UpdateUser update user
func (controller *UserController) UpdateUser(c *gin.Context) {
	userID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}
	var body struct {
		Name        string `json:"name"`
		Phone       string `json:"phone"`
		Avatar      string `json:"avatar"`
		DeviceId    string `json:"device_id"`
		DeviceToken string `json:"device_token"`
		Active      bool   `json:"active"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	// if err := utils.ValidateStruct(c, body); err != nil {
	// 	utils.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
	// 	return
	// }
	// Get user from database
	user, err := controller.UserService.GetUserByID(userID)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get user", err.Error())
		return
	}
	// Update user fields
	user.Name = body.Name
	user.Phone = body.Phone
	user.Avatar = body.Avatar
	user.DeviceId = body.DeviceId
	user.DeviceToken = body.DeviceToken
	user.Active = body.Active

	// Update user in database
	err = controller.UserService.UpdateUser(user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}
	// Return updated user
	userResponse := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Active:    user.Active,
		Token:     user.Token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	// Return updated user
	utils.SuccessResponse(c, http.StatusOK, "User updated successfully", userResponse)
}

// DeleteUser delete user
func (controller *UserController) DeleteUser(c *gin.Context) {
	userID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}
	err := controller.UserService.DeleteUser(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Failed to delete user", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User deleted successfully", nil)
}

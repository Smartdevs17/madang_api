package services

import (
	"errors"
	"madang_api/config"
	"madang_api/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// RegisterUser creates a new user record
func (s *UserService) RegisterUser(user *models.User) error {
	var existingUser models.User
	config.DB.Where("email = ?", user.Email).First(&existingUser)
	//check if email already exists
	if existingUser.Email != "" {
		return errors.New("email already exists")
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	newUser := models.User{Name: user.Name, Email: user.Email, Phone: user.Phone, Password: user.Password, Avatar: user.Avatar, Role: user.Role, Active: user.Active, EmailVerificationOTP: user.EmailVerificationOTP}
	result := config.DB.Create(&newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

// Implement email otp validation
func (s *UserService) VerifyEmailOTP(email string, otp string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if user.EmailVerificationOTP != otp {
		return nil, errors.New("invalid OTP")
	}

	//generate for the user a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, err
	}
	//store the token in the user model
	user.Token = tokenString
	user.EmailVerified = true
	user.Active = true
	user.EmailVerificationOTP = ""
	result = config.DB.Save(&user)
	return &user, nil
}

// LoginUser authenticates a user and returns the user object if successful with token generated and stored in the user model
func (s *UserService) LoginUser(email string, password string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	//check if email is verified
	if !user.EmailVerified {
		return nil, errors.New("email not verified")
	}
	//check if password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	//generate for the user a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, err
	}
	//store the token in the user model
	user.Token = tokenString
	return &user, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser updates a user record
func (s *UserService) UpdateUser(user *models.User) error {
	result := config.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user record
func (s *UserService) DeleteUser(id uint) error {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	config.DB.Delete(&user)
	return nil
}

// GetAllUsers retrieves a list of users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

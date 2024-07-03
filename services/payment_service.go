package services

import (
	"madang_api/config"
	"madang_api/models"
)

type PaymentService struct{}

func (s *PaymentService) GetPayments() ([]models.Payment, error) {
	var payments []models.Payment
	result := config.DB.Find(&payments)
	return payments, result.Error
}

func (s *PaymentService) GetPayment(id uint) (models.Payment, error) {
	var payment models.Payment
	result := config.DB.First(&payment, id)
	return payment, result.Error
}

func (s *PaymentService) CreatePayment(payment *models.Payment) (*models.Payment, error) {
	result := config.DB.Create(payment)
	if result.Error != nil {
		return nil, result.Error
	}

	return payment, nil
}

func (s *PaymentService) UpdatePayment(payment *models.Payment) (*models.Payment, error) {
	result := config.DB.Save(payment)
	if result.Error != nil {
		return nil, result.Error
	}

	return payment, nil
}

func (s *PaymentService) DeletePayment(id uint) error {
	result := config.DB.Delete(&models.Payment{}, id)
	return result.Error
}

func (s *PaymentService) GetRestaurantPayments(restaurantID uint) ([]models.Payment, error) {
	var payments []models.Payment
	result := config.DB.Where("restaurant_id = ?", restaurantID).Find(&payments)
	return payments, result.Error
}

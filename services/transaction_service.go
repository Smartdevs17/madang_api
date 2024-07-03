package services

import (
	"madang_api/config"
	"madang_api/models"
)

type TransactionService struct{}

func (s *TransactionService) GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := config.DB.Find(&transactions)
	return transactions, result.Error
}

func (s *TransactionService) GetTransaction(id uint) (models.Transaction, error) {
	var transaction models.Transaction
	result := config.DB.First(&transaction, id)
	return transaction, result.Error
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	result := config.DB.Create(transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	return transaction, nil
}

func (s *TransactionService) UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	result := config.DB.Save(transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	return transaction, nil
}

func (s *TransactionService) DeleteTransaction(id uint) error {
	result := config.DB.Delete(&models.Transaction{}, id)
	return result.Error
}

func (s *TransactionService) GetRestaurantTransactions(restaurantID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := config.DB.Where("restaurant_id = ?", restaurantID).Find(&transactions)
	return transactions, result.Error
}

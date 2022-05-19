package utils

import (
	"EasyEcommerce-backend/internal/mysql/models"
	"bytes"
	"crypto/rand"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"math/big"
	"sort"
)

func CreateRandomNumber() string {
	var numbers = []byte{1, 2, 3, 4, 5, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()

	for i := 1; i <= 12; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			logger.Panic(err)
		}
		container += fmt.Sprintf("%d", numbers[random.Int64()])
	}
	return container
}

func CreateRandomString() string {
	var numbers = []byte{1, 2, 3, 4, 5, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()

	for i := 1; i <= 6; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			logger.Panic(err)
		}
		container += fmt.Sprintf("%d", numbers[random.Int64()])
	}
	return container
}

func SortCreateTime(data []models.SaleData) []models.SaleData {
	sort.Slice(data, func(i, j int) bool {
		return data[i].CreatedAt.After(data[j].CreatedAt)
	})
	return data
}

func SortCreateTimeASC(data []models.SaleData) []models.SaleData {
	sort.Slice(data, func(i, j int) bool {
		return data[i].CreatedAt.Before(data[j].CreatedAt)
	})
	return data
}

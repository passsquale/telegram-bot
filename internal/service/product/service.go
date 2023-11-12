package product

import (
	"fmt"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	if idx < 0 || idx >= len(allProducts) {
		return nil, fmt.Errorf("Invalid index %d", idx)
	}
	return &allProducts[idx], nil
}

package mypackage

import (
	"errors"
)

type PackageService struct {
}

func NewService() *PackageService {
	return &PackageService{}
}
func (s *PackageService) List() []Package {
	return allPackage
}

func (s *PackageService) Get(idx int) (*Package, error) {
	if idx < 0 || idx >= len(allPackage) {
		return nil, errors.New("Invalid index")
	}
	return &allPackage[idx], nil
}

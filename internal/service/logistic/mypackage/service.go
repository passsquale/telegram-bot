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
		return nil, errors.New("invalid index")
	}
	return &allPackage[idx], nil
}
func (s *PackageService) New(title string) (*Package, error) {
	if title == "" || len(title) > 20 {
		return nil, errors.New("invalid title")
	}
	newPackage := Package{Title: title}
	allPackage = append(allPackage, newPackage)
	return &newPackage, nil
}

package mypackage

import (
	"errors"
)

type PackageService interface {
	List() []Package
}

type DummyPackageService struct {
}

func NewService() *DummyPackageService {
	return &DummyPackageService{}
}
func (s *DummyPackageService) List() []Package {
	return allPackages
}

func (s *DummyPackageService) Get(idx int) (*Package, error) {
	if idx < 0 || idx >= len(allPackages) {
		return nil, errors.New("invalid index")
	}
	return &allPackages[idx], nil
}
func (s *DummyPackageService) Create(title string) (*Package, error) {
	if title == "" || len(title) > 20 {
		return nil, errors.New("invalid title")
	}
	newPackage := Package{title}
	allPackages = append(allPackages, newPackage)
	return &newPackage, nil
}
func (s *DummyPackageService) Delete(idx int) (bool, error) {
	if idx < 0 || idx >= len(allPackages) {
		return false, errors.New("invalid index")
	}
	allPackages = append(allPackages[:idx], allPackages[idx+1:]...)
	return true, nil
}

package service

import (
	"parqueadero/internal/entrada_vehiculo/domain"
	"parqueadero/internal/entrada_vehiculo/repository"
)

type ServiceImpl struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) *ServiceImpl {
	return &ServiceImpl{Repo: repo}
}

func (s *ServiceImpl) Create(entrada domain.EntradaVehiculo) error {
	return s.Repo.Save(entrada)
}

func (s *ServiceImpl) GetByID(id int64) (*domain.EntradaVehiculo, error) {
	return s.Repo.FindByID(id)
}

func (s *ServiceImpl) Update(entrada domain.EntradaVehiculo) error {
	return s.Repo.Update(entrada)
}

func (s *ServiceImpl) Delete(id int64) error {
	return s.Repo.Delete(id)
}

func (s *ServiceImpl) List() ([]domain.EntradaVehiculo, error) {
	return s.Repo.FindAll()
}

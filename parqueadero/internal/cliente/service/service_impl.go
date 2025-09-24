package service

import (
	"parqueadero/internal/cliente/domain"
	"parqueadero/internal/cliente/repository"
)

type ServiceImpl struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) *ServiceImpl {
	return &ServiceImpl{Repo: repo}
}

func (s *ServiceImpl) Create(cliente domain.Cliente) error {
	return s.Repo.Save(cliente)
}

func (s *ServiceImpl) GetByID(id int64) (*domain.Cliente, error) {
	return s.Repo.FindByID(id)
}

func (s *ServiceImpl) Update(cliente domain.Cliente) error {
	return s.Repo.Update(cliente)
}

func (s *ServiceImpl) Delete(id int64) error {
	return s.Repo.Delete(id)
}

func (s *ServiceImpl) List() ([]domain.Cliente, error) {
	return s.Repo.FindAll()
}

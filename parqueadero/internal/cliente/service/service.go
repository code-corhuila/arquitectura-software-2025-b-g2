package service

import "parqueadero/internal/cliente/domain"

type Service interface {
	Create(cliente domain.Cliente) error
	GetByID(id int64) (*domain.Cliente, error)
	Update(cliente domain.Cliente) error
	Delete(id int64) error
	List() ([]domain.Cliente, error)
}

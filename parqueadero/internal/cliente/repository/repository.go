package repository

import "parqueadero/internal/cliente/domain"

type Repository interface {
	Save(cliente domain.Cliente) error
	FindByID(id int64) (*domain.Cliente, error)
	Update(cliente domain.Cliente) error
	Delete(id int64) error
	FindAll() ([]domain.Cliente, error)
}

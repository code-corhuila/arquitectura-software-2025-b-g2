package repository

import "parqueadero/internal/entrada_vehiculo/domain"

type Repository interface {
	Save(entrada domain.EntradaVehiculo) error
	FindByID(id int64) (*domain.EntradaVehiculo, error)
	Update(entrada domain.EntradaVehiculo) error
	Delete(id int64) error
	FindAll() ([]domain.EntradaVehiculo, error)
}

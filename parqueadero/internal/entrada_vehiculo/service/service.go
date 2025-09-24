package service

import "parqueadero/internal/entrada_vehiculo/domain"

type Service interface {
	Create(entrada domain.EntradaVehiculo) error
	GetByID(id int64) (*domain.EntradaVehiculo, error)
	Update(entrada domain.EntradaVehiculo) error
	Delete(id int64) error
	List() ([]domain.EntradaVehiculo, error)
}

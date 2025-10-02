package repository

import (
	"database/sql"
	"parqueadero/internal/entrada_vehiculo/domain"
)

type SQLiteRepository struct {
	DB *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{DB: db}
}

func (r *SQLiteRepository) Save(entrada domain.EntradaVehiculo) error {
	_, err := r.DB.Exec("INSERT INTO entradas_vehiculo (cliente_id, placa, fecha_hora) VALUES (?, ?, ?)", entrada.ClienteID, entrada.Placa, entrada.FechaHora)
	return err
}

func (r *SQLiteRepository) FindByID(id int64) (*domain.EntradaVehiculo, error) {
	row := r.DB.QueryRow("SELECT id, cliente_id, placa, fecha_hora FROM entradas_vehiculo WHERE id = ?", id)
	var e domain.EntradaVehiculo
	err := row.Scan(&e.ID, &e.ClienteID, &e.Placa, &e.FechaHora)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *SQLiteRepository) Update(entrada domain.EntradaVehiculo) error {
	_, err := r.DB.Exec("UPDATE entradas_vehiculo SET cliente_id = ?, placa = ?, fecha_hora = ? WHERE id = ?", entrada.ClienteID, entrada.Placa, entrada.FechaHora, entrada.ID)
	return err
}

func (r *SQLiteRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM entradas_vehiculo WHERE id = ?", id)
	return err
}

func (r *SQLiteRepository) FindAll() ([]domain.EntradaVehiculo, error) {
	rows, err := r.DB.Query("SELECT id, cliente_id, placa, fecha_hora FROM entradas_vehiculo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entradas []domain.EntradaVehiculo
	for rows.Next() {
		var e domain.EntradaVehiculo
		if err := rows.Scan(&e.ID, &e.ClienteID, &e.Placa, &e.FechaHora); err != nil {
			return nil, err
		}
		entradas = append(entradas, e)
	}
	return entradas, nil
}

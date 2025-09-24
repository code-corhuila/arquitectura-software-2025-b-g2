package repository

import (
	"database/sql"
	"parqueadero/internal/cliente/domain"
)

type SQLiteRepository struct {
	DB *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{DB: db}
}

func (r *SQLiteRepository) Save(cliente domain.Cliente) error {
	_, err := r.DB.Exec("INSERT INTO clientes (nombre, telefono) VALUES (?, ?)", cliente.Nombre, cliente.Telefono)
	return err
}

func (r *SQLiteRepository) FindByID(id int64) (*domain.Cliente, error) {
	row := r.DB.QueryRow("SELECT id, nombre, telefono FROM clientes WHERE id = ?", id)
	var c domain.Cliente
	err := row.Scan(&c.ID, &c.Nombre, &c.Telefono)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *SQLiteRepository) Update(cliente domain.Cliente) error {
	_, err := r.DB.Exec("UPDATE clientes SET nombre = ?, telefono = ? WHERE id = ?", cliente.Nombre, cliente.Telefono, cliente.ID)
	return err
}

func (r *SQLiteRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM clientes WHERE id = ?", id)
	return err
}

func (r *SQLiteRepository) FindAll() ([]domain.Cliente, error) {
	rows, err := r.DB.Query("SELECT id, nombre, telefono FROM clientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var clientes []domain.Cliente
	for rows.Next() {
		var c domain.Cliente
		if err := rows.Scan(&c.ID, &c.Nombre, &c.Telefono); err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}
	return clientes, nil
}

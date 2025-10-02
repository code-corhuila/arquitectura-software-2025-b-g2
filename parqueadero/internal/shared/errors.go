package shared

import "errors"

var (
	ErrNotFound     = errors.New("recurso no encontrado")
	ErrInvalidInput = errors.New("entrada inválida")
	// ...otros errores comunes
)

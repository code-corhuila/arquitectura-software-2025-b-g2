package domain

type Cliente struct {
	ID       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Telefono string `json:"telefono"`
	// ...otros campos
}

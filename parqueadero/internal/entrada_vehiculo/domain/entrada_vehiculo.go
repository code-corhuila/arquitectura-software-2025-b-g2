package domain

type EntradaVehiculo struct {
	ID        int64  `json:"id"`
	ClienteID int64  `json:"cliente_id"`
	Placa     string `json:"placa"`
	FechaHora string `json:"fecha_hora"`
	// ...otros campos
}

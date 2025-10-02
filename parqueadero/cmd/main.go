package main

import (
	"fmt"
	"net/http"
	clienteApi "parqueadero/internal/cliente/api"
	clienteRepo "parqueadero/internal/cliente/repository"
	clienteService "parqueadero/internal/cliente/service"
	entradaApi "parqueadero/internal/entrada_vehiculo/api"
	entradaRepo "parqueadero/internal/entrada_vehiculo/repository"
	entradaService "parqueadero/internal/entrada_vehiculo/service"
	"parqueadero/internal/shared"
)

func main() {
	db, err := shared.NewDB("parqueadero.db")
	if err != nil {
		panic(err)
	}

	// Cliente
	cRepo := clienteRepo.NewSQLiteRepository(db)
	cService := clienteService.NewService(cRepo)
	cHandler := clienteApi.Handler{Service: cService}

	http.HandleFunc("/clientes/create", cHandler.CreateCliente)
	http.HandleFunc("/clientes/get", cHandler.GetCliente)
	http.HandleFunc("/clientes/update", cHandler.UpdateCliente)
	http.HandleFunc("/clientes/delete", cHandler.DeleteCliente)
	http.HandleFunc("/clientes/list", cHandler.ListClientes)

	// EntradaVehiculo
	eRepo := entradaRepo.NewSQLiteRepository(db)
	eService := entradaService.NewService(eRepo)
	eHandler := entradaApi.Handler{Service: eService}

	http.HandleFunc("/entradas/create", eHandler.CreateEntradaVehiculo)
	http.HandleFunc("/entradas/get", eHandler.GetEntradaVehiculo)
	http.HandleFunc("/entradas/update", eHandler.UpdateEntradaVehiculo)
	http.HandleFunc("/entradas/delete", eHandler.DeleteEntradaVehiculo)
	http.HandleFunc("/entradas/list", eHandler.ListEntradasVehiculo)

	fmt.Println("Servidor Parqueadero iniciado en :8080")
	http.ListenAndServe(":8080", nil)
}

package api

import (
	"encoding/json"
	"net/http"
	"parqueadero/internal/entrada_vehiculo/domain"
	"parqueadero/internal/entrada_vehiculo/service"
	"strconv"
)

type Handler struct {
	Service service.Service
}

func (h *Handler) CreateEntradaVehiculo(w http.ResponseWriter, r *http.Request) {
	var e domain.EntradaVehiculo
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input"))
		return
	}
	if err := h.Service.Create(e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating entry"))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetEntradaVehiculo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}
	e, err := h.Service.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Entry not found"))
		return
	}
	json.NewEncoder(w).Encode(e)
}

func (h *Handler) UpdateEntradaVehiculo(w http.ResponseWriter, r *http.Request) {
	var e domain.EntradaVehiculo
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input"))
		return
	}
	if err := h.Service.Update(e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error updating entry"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteEntradaVehiculo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}
	if err := h.Service.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error deleting entry"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListEntradasVehiculo(w http.ResponseWriter, r *http.Request) {
	entradas, err := h.Service.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error listing entries"))
		return
	}
	json.NewEncoder(w).Encode(entradas)
}

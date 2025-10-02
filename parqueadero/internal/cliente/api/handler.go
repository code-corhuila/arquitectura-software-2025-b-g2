package api

import (
	"encoding/json"
	"net/http"
	"parqueadero/internal/cliente/domain"
	"parqueadero/internal/cliente/service"
	"strconv"
)

type Handler struct {
	Service service.Service
}

func (h *Handler) CreateCliente(w http.ResponseWriter, r *http.Request) {
	var c domain.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input"))
		return
	}
	if err := h.Service.Create(c); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating client"))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetCliente(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}
	c, err := h.Service.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Client not found"))
		return
	}
	json.NewEncoder(w).Encode(c)
}

func (h *Handler) UpdateCliente(w http.ResponseWriter, r *http.Request) {
	var c domain.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid input"))
		return
	}
	if err := h.Service.Update(c); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error updating client"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteCliente(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}
	if err := h.Service.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error deleting client"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListClientes(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.Service.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error listing clients"))
		return
	}
	json.NewEncoder(w).Encode(clientes)
}

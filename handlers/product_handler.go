package handlers

import (
	"encoding/json"
	"learnGo/models"
	"learnGo/services"
	"learnGo/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		utils.JSONResponse(w, http.StatusMethodNotAllowed, "method not allowed", nil)
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		log.Println("ERROR GetAll products:", err)
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "success", products)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.ProductModel
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	if err := h.service.Create(&product); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "success", product)
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		utils.JSONResponse(w, http.StatusMethodNotAllowed, "method not allowed", nil)
	}
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid product id", nil)
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, "product not found", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "success", product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid product id", nil)
		return
	}

	var product models.ProductModel
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	product.ID = id
	if err := h.service.Update(&product); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "success", product)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid product id", nil)
		return
	}

	if err := h.service.Delete(id); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "failed to delete product", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "success", nil)
}

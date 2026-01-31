package handlers

import (
	"encoding/json"
	"learnGo/models"
	"learnGo/services"
	"learnGo/utils"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategories - GET /api/category | POST /api/category
func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		utils.JSONResponse(w, http.StatusMethodNotAllowed, "method not allowed", nil)
	}
}

// GET /api/category
func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "success", categories)
}

// POST /api/category
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.CategoryModel

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	if err := h.service.Create(&category); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "success", category)
}

// HandleCategoryByID - GET/PUT/DELETE /api/category/{id}
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
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

// GET /api/category/{id}
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid category id", nil)
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, "category not found", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "success", category)
}

// PUT /api/category/{id}
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid category id", nil)
		return
	}

	var category models.CategoryModel
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	category.ID = id
	if err := h.service.Update(&category); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "success", category)
}

// DELETE /api/category/{id}
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "invalid category id", nil)
		return
	}

	if err := h.service.Delete(id); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, "failed to delete category", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "success", nil)
}

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavolubacheski/API/CRUD-API/models"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/users", usersGetHandler).Methods("GET")
	r.HandleFunc("/users", usersPostHandler).Methods("POST")
	r.HandleFunc("/users/{cpf_cnpj}", usersGetByCPFHandler).Methods("GET")
	r.HandleFunc("/users/{cpf_cnpj}", usersDeleteByCPFHandler).Methods("DELETE")
	r.HandleFunc("/users/{cpf_cnpj}", usersUpdateByCPFHandler).Methods("PUT")
	return r
}

func ContentTypeJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func usersGetHandler(w http.ResponseWriter, r *http.Request) {
	ContentTypeJson(w)
	users, err := models.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  "INTERNAL SERVER ERROR",
			Status: 500,
		})
		return
	}
	json.NewEncoder(w).Encode(users)
}

func usersPostHandler(w http.ResponseWriter, r *http.Request) {
	ContentTypeJson(w)
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  "INVALID JSON",
			Status: 400,
		})
		return
	}
	_, err := models.NewUser(user)
	if err != nil {
		fmt.Println("Database Error:", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  "UNPROCESSABLE ENTITY",
			Status: 422,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{
		Message: "User created successfully",
		Status:  201,
	})
}

func usersGetByCPFHandler(w http.ResponseWriter, r *http.Request) {
	ContentTypeJson(w)
	params := mux.Vars(r)
	cpf := params["cpf_cnpj"]
	user, err := models.GetUserByCPF(cpf)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  "User not found",
			Status: http.StatusNotFound,
		})
		return
	}
	json.NewEncoder(w).Encode(user)
}

func usersDeleteByCPFHandler(w http.ResponseWriter, r *http.Request) {
	ContentTypeJson(w)
	params := mux.Vars(r)
	cpf := params["cpf_cnpj"]
	rowsAffected, err := models.DeleteUserByCPF(cpf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  "Error deleting user",
			Status: http.StatusInternalServerError,
		})
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{Error: "User not found", Status: http.StatusNotFound})
		return
	}

	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{
		Message: "User deleted successfully",
		Status:  200,
	})
}

func usersUpdateByCPFHandler(w http.ResponseWriter, r *http.Request) {
	ContentTypeJson(w)
	params := mux.Vars(r)
	cpf := params["cpf_cnpj"]

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{Error: "INVALID JSON", Status: http.StatusBadRequest})
		return
	}

	user.CPF_CNPJ = cpf

	rowsAffected, err := models.UpdateUser(user)
	if err != nil {
		fmt.Println("Database Error:", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  "UNPROCESSABLE ENTITY",
			Status: 422,
		})
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{Error: "User not found to update", Status: http.StatusNotFound})
		return
	}

	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{
		Message: "User updated successfully",
		Status:  http.StatusOK,
	})
}

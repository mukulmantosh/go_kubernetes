package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *MuxServer) addUser(w http.ResponseWriter, r *http.Request) {
	var userData UserParam
	var user User
	json.NewDecoder(r.Body).Decode(&userData)

	user.Name = userData.Name
	user.Email = userData.Email
	user.Age = userData.Age

	s.db.Create(&user)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *MuxServer) listUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	s.db.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (s *MuxServer) updateUser(w http.ResponseWriter, r *http.Request) {
	var userData UserParam
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user User
	s.db.First(&user, userId)
	user.Name = userData.Name
	user.Email = userData.Email
	user.Age = userData.Age

	s.db.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User updated successfully",
	})

	return

}

func (s *MuxServer) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var user User
	s.db.First(&user, userId)
	s.db.Delete(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User deleted successfully",
	})
	return
}

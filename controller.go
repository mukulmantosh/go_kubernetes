package main

import (
	"encoding/json"
	"net/http"
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
	json.NewEncoder(w).Encode(user)
}

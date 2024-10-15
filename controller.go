package main

import (
	"encoding/json"
	"net/http"
)

func (s *MuxServer) addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	s.db.Create(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

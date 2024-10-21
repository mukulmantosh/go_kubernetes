package main

func (s *MuxServer) routes() {
	s.gorilla.HandleFunc("/user", s.addUser).Methods("POST")
	s.gorilla.HandleFunc("/users", s.listUsers).Methods("GET")
	s.gorilla.HandleFunc("/user/{id}", s.updateUser).Methods("PUT")
	s.gorilla.HandleFunc("/user/{id}", s.deleteUser).Methods("DELETE")

}

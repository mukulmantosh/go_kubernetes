package main

func (s *MuxServer) routes() {
	s.gorilla.HandleFunc("/user", s.addUser).Methods("POST")

}

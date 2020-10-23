package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/MeguMan/buyer-exp-test/internal/app/store"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Data struct {
	Email string
	Link string
}

type server struct {
	router *mux.Router
	store  store.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/follow", s.follow()).Methods("POST")
}

func (s *server) follow() func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		d := Data{}
		u := model.User{}
		ad := model.Ad{}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		u.Email = d.Email
		ad.Link = d.Link
		ad.Price = 123 //Need price parsing

		ad.CheckPrice()

		ur := s.store.User()
		adr := s.store.Ad()
		userAd := s.store.UserAd()

		u.ID, err = ur.Create(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Print(err)
			return
		}
		ad.ID, err = adr.Create(&ad)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Print(err)
			return
		}
		err = userAd.Create(u.ID, ad.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Print(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, u, ad)
	}
}

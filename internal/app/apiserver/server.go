package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/MeguMan/buyer-exp-test/internal/app/emailsender"
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/MeguMan/buyer-exp-test/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"log"
	"net/http"
)

type Data struct {
	Email string
	Link  string
}

type server struct {
	router *mux.Router
	store  store.Store
	es     emailsender.Sender
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(store store.Store, emailSender emailsender.Sender) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
		es:     emailSender,
	}
	s.configureRouter()
	c := cron.New()
	c.AddFunc("* * * * * *", s.store.Ad().CheckPrice)
	c.Start()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/follow", s.Follow()).Methods("POST")
}

func (s *server) Follow() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d := Data{}
		u := model.User{}
		a := model.Ad{}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		u.Email = d.Email
		a.Link = d.Link

		if err := u.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		if err := a.Validate(); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ur := s.store.User()
		adr := s.store.Ad()
		userAd := s.store.UserAd()

		_, err = ur.Create(&u)
		if err != nil {
			log.Print(err)
			return
		}
		_, err = adr.Create(&a)
		if err != nil {
			log.Print(err)
			return
		}
		err = userAd.Create(u.ID, a.ID)
		if err != nil {
			log.Print(err)
			return
		}

		err = s.es.SendEmail(&u, &a)
		if err != nil {
			log.Print(err)
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, u, a)
	}
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeremy-spitzig/udemy-web-dev-with-go/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(resp http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)

	resp.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(resp, "%s\n", uj)
}

func (uc UserController) CreateUser(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(req.Body).Decode(&u)

	u.Id = "007"

	uj, _ := json.Marshal(u)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	fmt.Fprintf(resp, "%s\n", uj)
}

func (uc UserController) DeleteUser(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	resp.WriteHeader(http.StatusOK)
	fmt.Fprint(resp, "Write code to delete user\n")
}

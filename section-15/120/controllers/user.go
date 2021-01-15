package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jeremy-spitzig/udemy-web-dev-with-go/models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	data map[string]models.User
}

func NewUserController() *UserController {
	data, err := load()
	if err != nil {
		log.Panicln(err)
	}
	return &UserController{data}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	u, ok := uc.data[id]

	// Fetch user
	if !ok {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = uuid.NewV4().String()

	uc.data[u.Id] = u

	uc.persist()

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	_, ok := uc.data[id]

	// Delete user
	if !ok {
		w.WriteHeader(404)
		return
	}

	delete(uc.data, id)

	uc.persist()

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}

func (uc UserController) persist() error {
	f, err := os.OpenFile("./users.db", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for _, user := range uc.data {
		w.Write([]string{
			user.Id,
			user.Name,
			user.Gender,
			strconv.Itoa(user.Age),
		})
	}
	w.Flush()
	return nil
}

func load() (map[string]models.User, error) {
	_, err := os.Stat("./users.db")
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]models.User{}, nil
		}
		return map[string]models.User{}, err
	}

	f, err := os.Open("./users.db")
	if err != nil {
		return map[string]models.User{}, err
	}
	defer f.Close()

	data := map[string]models.User{}
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return map[string]models.User{}, err
			}
		}
		age, err := strconv.Atoi(record[3])
		if err != nil {
			return map[string]models.User{}, err
		}
		u := models.User{
			Id:     record[0],
			Name:   record[1],
			Gender: record[2],
			Age:    age,
		}
		data[u.Id] = u
	}
	return data, nil
}

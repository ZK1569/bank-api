package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	errs "github.com/zk1569/test-mongodb/src/errors"
	service "github.com/zk1569/test-mongodb/src/services"
)

var (
	lock *sync.Mutex
)

func init() {
	lock = &sync.Mutex{}
}

type Bank struct {
	path        string
	bankService *service.Bank
}

var singleBankInstance *Bank

func GetBankInstance() *Bank {
	if singleBankInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleBankInstance = &Bank{
			path:        "/bank",
			bankService: service.GetBankInstance(),
		}
	}

	return singleBankInstance
}

func (self *Bank) Mount(r chi.Router) {
	r.Route(self.path, func(r chi.Router) {
		r.Post("/", self.createAccount)
		r.Delete("/{UserID}", self.deleteAccount)
		r.Post("/{UserID}/money", self.addMoney)
		r.Delete("/{UserID}/money", self.removeMoney)
		r.Get("/{UserID}", self.getAccount)
		r.Get("/{UserID}/total", self.getTotal)
	})
}

type NewAccountBody struct {
	Name      string `json:"name"`
	Firstname string `json:"firstname"`
	Age       int    `json:"age"`
}

func (self *Bank) createAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var newAccount NewAccountBody
	if err := json.NewDecoder(r.Body).Decode(&newAccount); err != nil {
		jsonResponse(w, http.StatusBadRequest, errs.BadRequest)
		return
	}

	ctx := r.Context()
	userID, err := self.bankService.CreateAccount(
		ctx,
		newAccount.Name,
		newAccount.Firstname,
		newAccount.Age,
	)
	if err != nil {
		jsonResponseError(w, err)
		return
	}

	jsonResponse(w, http.StatusCreated, userID)
	return
}

func (self *Bank) deleteAccount(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "UserID")

	ctx := r.Context()
	err := self.bankService.DeleteAccount(ctx, userID)

	if err != nil {
		jsonResponseError(w, err)
		return
	}

	jsonResponse(w, http.StatusOK, nil)
	return
}
func (self *Bank) getAccount(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "UserID")

	ctx := r.Context()
	user, err := self.bankService.GetAccount(ctx, userID)

	if err != nil {
		jsonResponseError(w, err)
		return
	}

	jsonResponse(w, http.StatusOK, user)
	return
}

func (self *Bank) getTotal(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "UserID")

	ctx := r.Context()
	user, err := self.bankService.GetAccount(ctx, userID)

	if err != nil {
		jsonResponseError(w, err)
		return
	}

	total := 0
	for _, i := range user.History {
		total += i.Money
	}

	jsonResponse(w, http.StatusOK, total)
	return
}

type BodyMoney struct {
	Amount int `json:"amount"`
}

func (self *Bank) addMoney(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "UserID")

	var body BodyMoney
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		jsonResponse(w, http.StatusBadRequest, errs.BadRequest)
		return
	}

	ctx := r.Context()
	err := self.bankService.AddMoney(ctx, userID, body.Amount)

	if err != nil {
		jsonResponseError(w, err)
		return
	}

	jsonResponse(w, http.StatusOK, nil)
	return
}
func (self *Bank) removeMoney(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "UserID")

	var body BodyMoney
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		jsonResponse(w, http.StatusBadRequest, errs.BadRequest)
		return
	}

	ctx := r.Context()
	err := self.bankService.RemoveMoney(ctx, userID, body.Amount)

	if err != nil {
		jsonResponseError(w, err)
		return
	}

	jsonResponse(w, http.StatusOK, nil)
	return
}

// --- END Handler ---

func jsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}

	return writeJSON(w, status, &envelope{Data: data})
}

func jsonResponseError(w http.ResponseWriter, err error) error {

	type envelope struct {
		Error string `json:"error"`
	}

	switch err {
	default:
		log.Printf("❌ Error: %v \n", err)
		return writeJSON(w, http.StatusInternalServerError, &envelope{Error: "Internal error"})
	}

}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

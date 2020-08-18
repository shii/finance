package routes

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/shii/finance/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

var store *cache.Cache

func SetupRoutes(cache *cache.Cache) *http.ServeMux {
	store = cache
	route := http.NewServeMux()
	route.HandleFunc("/v1/transfer", accountDeposit)
	route.HandleFunc("/v1/balance", accountBalance)

	return route
}

func accountDeposit(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var transferRequest TransferRequest

	err = json.Unmarshal(b, &transferRequest)
	if err != nil {
		fmt.Printf("%v", err)
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("%v", transferRequest)

	// validations
	to := model.GetKey(transferRequest.AccountTo)
	var credit, debit int64

	if acntTo, found := store.Get(to); !found {
		http.Error(w, fmt.Sprintf("Accoount %d not found", transferRequest.AccountTo), http.StatusBadRequest)
	} else {
		x := acntTo.(model.Account)
		credit = x.AvailableAmount
	}

	from := model.GetKey(transferRequest.AccountFrom)
	if acnt, found := store.Get(from); !found {
		http.Error(w, fmt.Sprintf("Accoount %d not found", transferRequest.AccountFrom), http.StatusBadRequest)
	} else {
		x := acnt.(model.Account)
		debit = x.AvailableAmount
		if x.AvailableAmount < transferRequest.Amount {
			http.Error(w, "Insufficient funds", http.StatusBadRequest)
		} else {
			credit = credit + transferRequest.Amount // 600+100
			creditAcnt := model.GenerateAccount(transferRequest.AccountTo, credit)
			model.UpdateAccount(store, creditAcnt)

			debit = debit - transferRequest.Amount // 1000-100
			debitAcnt := model.GenerateAccount(transferRequest.AccountFrom, debit)
			model.UpdateAccount(store, debitAcnt)
		}
	}
}

func accountBalance(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "405 Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	params := r.URL.Query()
	id, _ := strconv.Atoi(params.Get("id"))
	key := model.GetKey(int64(id))
	if acnt, found := store.Get(key); found {
		account := acnt.(model.Account)
		err := json.NewEncoder(w).Encode(account)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	} else {
		http.Error(w, fmt.Sprintf("Accoount %d not found", id), http.StatusNotFound)
	}
}

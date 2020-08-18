package model

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type Account struct {
	Id              int64 `json:"id"`
	AvailableAmount int64 `json:"available_amount"`
}

func GenerateAccount(id int64, amount int64) Account {

	return Account{
		Id:              id,
		AvailableAmount: amount,
	}
}

func SetupCache() *cache.Cache {
	c := cache.New(5*time.Minute, 30*time.Second)

	account := GenerateAccount(1234, 1000)

	key := GetKey(account.Id)
	if _, found := c.Get(key); !found {
		c.Set(key, account, cache.NoExpiration)
	}

	account2 := GenerateAccount(7890, 600)

	key2 := GetKey(account2.Id)
	if _, found := c.Get(key2); !found {
		c.Set(key2, account2, cache.NoExpiration)
	}

	return c
}

func UpdateAccount(store *cache.Cache, acnt Account) {
	key := GetKey(acnt.Id)
	store.Delete(key)
	store.Set(key, acnt, cache.NoExpiration)
}

func GetKey(id int64) string {
	return "acc_" + string(id)
}

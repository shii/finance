package model

import (
	"testing"
)

func TestCanUpdateAccount(t *testing.T) {
	cache := SetupCache()

	var credit, debit int64
	var transferAmount int64 = 100

	to := GetKey(7890)
	if acntTo, found := cache.Get(to); found {

		x := acntTo.(Account)
		credit = x.AvailableAmount
		if credit != 600 {
			t.Errorf("Credit account amount should be %d; want 1", credit)
		}

	}

	from := GetKey(1234)
	if acnt, found := cache.Get(from); found {
		x := acnt.(Account)
		debit = x.AvailableAmount
		if debit != 1000 {
			t.Errorf("Debit account amount should be %d; want 1", debit)
		}

	}

	credit = credit + transferAmount // 600+100
	creditAcnt := GenerateAccount(7890, credit)
	UpdateAccount(cache, creditAcnt)

	debit = debit - transferAmount // 1000-100
	debitAcnt := GenerateAccount(1234, debit)
	UpdateAccount(cache, debitAcnt)

	toX := GetKey(7890)
	if acntTo, found := cache.Get(toX); found {

		x := acntTo.(Account)
		credit = x.AvailableAmount
		if credit != 700 {
			t.Errorf("Credit account amount should be %d; want 1", credit)
		}

	}

	fromX := GetKey(1234)
	if acnt, found := cache.Get(fromX); found {
		x := acnt.(Account)
		debit = x.AvailableAmount
		if debit != 900 {
			t.Errorf("Debit account amount should be %d; want 1", debit)
		}

	}
}

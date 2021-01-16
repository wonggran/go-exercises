package account

// Open(initialDeposit int64) *Account
// (*Account) Close() (payout int64, ok bool)
// (*Account) Balance() (balance int64, ok bool)
// (*Account) Deposit(amount int64) (newBalance int64, ok bool)

import (
	"sync"
)

type Account struct {
	balance   int64
	mutex     sync.Mutex
	willClose bool
}

func isValidDeposit(balance, amount int64) bool {
	return balance+amount >= 0
}

func (acc *Account) Deposit(amount int64) (int64, bool) {
	acc.mutex.Lock()

	if !acc.willClose {
		if isValidDeposit(acc.balance, amount) {
			acc.balance += amount
			defer acc.mutex.Unlock()
			return acc.balance, true
		} else {
			defer acc.mutex.Unlock()
			return amount, false
		}
	} else {
		defer acc.mutex.Unlock()
		return 0, false
	}
}

func (acc *Account) Balance() (int64, bool) {
	acc.mutex.Lock()

	if !acc.willClose {
		defer acc.mutex.Unlock()
		return acc.balance, true
	} else {
		defer acc.mutex.Unlock()
		return 0, false
	}
}

/* What happens if I close an account? The *Account will be nil but only after the mutex is unlocked.
However, we'll have the willClose flag to be true. So other threads may deposit, call balance, or even close
it again because the mutex will be unlocked but if it's going to be closed we'll release the lock and let
the closing finish. */
func (acc *Account) Close() (int64, bool) {
	acc.mutex.Lock()

	if !acc.willClose {
		closingBalance := acc.balance
		acc.balance = 0
		acc.willClose = true
		defer acc.mutex.Unlock()

		return closingBalance, true
	} else {
		defer acc.mutex.Unlock()
		/* We shouldn't be able to close an already closing account so do nothing but return that it's not ok. */
		return 0, false
	}

}

func Open(initialDeposit int64) *Account {
	if isValidDeposit(0, initialDeposit) {
		newAcc := Account{balance: initialDeposit, willClose: false}
		return &newAcc
	}
	return nil
}

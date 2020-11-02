package main

// Storage holds all key:values
var Storage = map[string]string{}

// empty type defined in order to not alloc memory
type empty struct{}

// Transaction links to a key:value storage
type Transaction struct {
	storage map[string]string // transaction key value storage
	deleted map[string]empty  // transaction deleted keys
	next    *Transaction
}

// TransactionList lists all active transactions (linked list)
type TransactionList struct {
	head *Transaction
}

// Set sets the given key the specified value.
// The key can also be updated.
func (tl *TransactionList) Set(key, value string) {
	if tl.head != nil {
		tl.head.Set(key, value)
	} else {
		Storage[key] = value
	}
}

// Get returns the current value of the specified key
func (tl *TransactionList) Get(key string) (value string, err error) {
	if tl.head == nil {
		if value, ok := Storage[key]; ok {
			return value, nil
		}

		return value, KeyNotFound.Error(key)
	}

	return tl.head.Get(key)
}

// Delete deletes the current specified key
func (tl *TransactionList) Delete(key string) bool {
	if tl.head == nil {
		if _, found := Storage[key]; found {
			delete(Storage, key)
			return true
		}
		return false
	}

	return tl.head.Delete(key)
}

// Set sets key:value for the current transaction storage
func (t *Transaction) Set(key, value string) {
	t.storage[key] = value

	// if this key is in the deleted storage it must be removed
	if _, found := t.deleted[key]; found {
		delete(t.deleted, key)
	}
}

// Get returns the current value of the specified key for the current transaction
func (t *Transaction) Get(key string) (value string, err error) {
	if _, found := t.deleted[key]; found {
		return value, KeyNotFound.Error(key)
	}

	value, found := t.storage[key]
	if !found {
		if t.next == nil {
			if value, ok := Storage[key]; ok {
				return value, nil
			}

			return value, KeyNotFound.Error(key)
		}

		if t.next != nil {
			return t.next.Get(key)
		}
	}

	return value, nil
}

// Delete deletes a key from the current transaction storage
func (t *Transaction) Delete(key string) bool {
	t.deleted[key] = empty{}
	if _, found := t.storage[key]; found {
		delete(t.storage, key)
		return true
	}

	return false
}

// Begin starts a new Transaction
func (tl *TransactionList) Begin() {
	tmp := Transaction{
		storage: make(map[string]string),
		deleted: make(map[string]empty),
		next:    tl.head,
	}

	tl.head = &tmp
}

// Rollback aborts current active transaction
func (tl *TransactionList) Rollback() error {
	if tl.head == nil {
		return NoActiveTransactions.Error()
	}

	tl.head = tl.head.next
	return nil
}

// Commit commits current active transaction
func (tl *TransactionList) Commit() error {
	if tl.head == nil {
		return NoActiveTransactions.Error()
	}

	if tl.head.next == nil {
		for k, v := range tl.head.storage {
			Storage[k] = v
		}

		for k := range tl.head.deleted {
			if _, found := Storage[k]; found {
				delete(Storage, k)
			}
		}

		tl.head = nil
	} else {
		for k, v := range tl.head.storage {
			tl.head.next.Set(k, v)
		}

		for k := range tl.head.deleted {
			tl.head.next.Delete(k)
		}

		tl.head = tl.head.next
	}

	return nil
}

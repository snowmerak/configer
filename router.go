package main

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger/v3"
	"github.com/snowmerak/lux/context"
)

var db *badger.DB

func init() {
	if len(os.Args) < 3 {
		panic("Please provide a path to the database")
	}

	dbPath := os.Args[2]

	var err error
	db, err = badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func rootGetRouter(lc *context.LuxContext) error {
	if err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(lc.GetPathVariable("name")))
		if err != nil {
			lc.SetConflict()
			return fmt.Errorf("txn.Get: %w", err)
		}
		return item.Value(func(val []byte) error {
			if val == nil {
				lc.SetBadRequest()
				return fmt.Errorf("item.Value: val is nil")
			}
			if err := lc.ReplyBinary(val); err != nil {
				lc.SetInternalServerError()
				return fmt.Errorf("lc.ReplyBinary: %w", err)
			}
			return nil
		})
	}); err != nil {
		return fmt.Errorf("db.View: %w", err)
	}
	return nil
}

func rootPostRouter(lc *context.LuxContext) error {
	if err := db.Update(func(txn *badger.Txn) error {
		body, err := lc.GetBody()
		if err != nil {
			lc.SetBadRequest()
			return fmt.Errorf("lc.GetBody: %w", err)
		}
		if err := txn.Set([]byte(lc.GetPathVariable("name")), body); err != nil {
			lc.SetInternalServerError()
			return fmt.Errorf("txn.Set: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("db.Update: %w", err)
	}
	return nil
}

func rootPutRouter(lc *context.LuxContext) error {
	if err := db.Update(func(txn *badger.Txn) error {
		value := lc.GetURLQuery("value")
		if value == "" {
			lc.SetBadRequest()
			return fmt.Errorf("lc.GetBody: value is empty")
		}
		if err := txn.Set([]byte(lc.GetPathVariable("name")), []byte(value)); err != nil {
			lc.SetInternalServerError()
			return fmt.Errorf("txn.Set: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("db.Update: %w", err)
	}
	return nil
}

func rootDeleteRouter(lc *context.LuxContext) error {
	if err := db.Update(func(txn *badger.Txn) error {
		if err := txn.Delete([]byte(lc.GetPathVariable("name"))); err != nil {
			lc.SetInternalServerError()
			return fmt.Errorf("txn.Delete: %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("db.Update: %w", err)
	}
	return nil
}

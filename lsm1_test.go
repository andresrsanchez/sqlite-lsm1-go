//go:build cgo
// +build cgo

package lsm1

// import "C"
import (
	"os"
	"testing"
)

func TestOpenLSM(t *testing.T) {
	name := "test"
	lsm, err := OpenLSM(name)
	defer remove()
	if err != nil || lsm == nil {
		t.Fatal("Failed to open database:", err)
	}
	_, err = OpenLSM(name)
	if err != nil || lsm == nil {
		t.Fatal("Error reopening the database")
	}
}
func TestCloseLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove()
	err := lsm.Close()
	if err != nil {
		t.Fatal("Failed to close database:", err)
	}
	err = lsm.Close()
	if err != nil {
		t.Fatal("Failed to close database:", err)
	}
}

func TestInsertLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove()
	err := lsm.Insert("key", "value")
	if err != nil {
		t.Fatal("Failed to insert into database:", err)
	}
	val, _ := lsm.Single("key")
	if val != "value" {
		t.Fatal("Failed to insert into database:", err)
	}
	err = lsm.Insert("key", "value2")
	if err != nil {
		t.Fatal("Failed to insert into database:", err)
	}
	val, _ = lsm.Single("key")
	if val != "value2" {
		t.Fatal("Failed to insert into database:", err)
	}
}

func TestDeleteLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove()
	err := lsm.Delete("nonexistant")
	if err != nil {
		t.Fatal("Failed to delete key:", err)
	}
	lsm.Insert("key", "value")
	err = lsm.Delete("key")
	val, _ := lsm.Single("key")
	if val != "" {
		t.Fatal("Failed to delete key:", err)
	}
	err = lsm.Delete("")
	if err != nil {
		t.Fatal("Failed to delete key:", err)
	}
}

func TestDeleteRangeLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove()
	err := lsm.DeleteRange("", "")
	if err != nil {
		t.Fatal("Failed to delete range keys:", err)
	}
	lsm.Insert("1", "1")
	lsm.Insert("2", "2")
	lsm.Insert("3", "3")
	lsm.Insert("4", "4")
	err = lsm.DeleteRange("2", "4")
	if err != nil {
		t.Fatal("Failed to delete range keys:", err)
	}
	m, _ := lsm.All()
	if len(m) != 3 {
		t.Fatal("Failed to delete range keys:", err)
	}
}

func TestSingleLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove()
	val, err := lsm.Single("nonexistant")
	if err != nil || val != "" {
		t.Fatal("Failed to get key:", err)
	}
	lsm.Insert("key", "value")
	val, err = lsm.Single("key")
	if err != nil || val != "value" {
		t.Fatal("Failed to get key:", err)
	}
}

func TestAllLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove()
	m, err := lsm.All()
	if err != nil || len(m) != 0 {
		t.Fatal("Failed to get all keys:", err)
	}
	lsm.Insert("key", "value")
	lsm.Insert("key1", "value")
	lsm.Insert("key2", "value")
	m, err = lsm.All()
	if err != nil || len(m) != 3 {
		t.Fatal("Failed to get all keys:", err)
	}
}

func TestRangeLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove()
	m, err := lsm.Range("non", "existant")
	if err != nil || len(m) != 0 {
		t.Fatal("Failed to get range keys:", err)
	}
	lsm.Insert("1", "1")
	lsm.Insert("2", "2")
	lsm.Insert("3", "3")
	lsm.Insert("4", "4")

	m, err = lsm.Range("", "")
	if err != nil || len(m) != 4 {
		t.Fatal("Failed to get range keys:", err)
	}
	m, err = lsm.Range("2", "3")
	if err != nil || len(m) != 2 {
		t.Fatal("Failed to get range keys:", err)
	}
	if v, ok := m["2"]; v != "2" || !ok {
		t.Fatal("Failed to get range keys:", err)
	}
	if v, ok := m["3"]; v != "3" || !ok {
		t.Fatal("Failed to get range keys:", err)
	}
	m, err = lsm.Range("", "3")
	if err != nil || len(m) != 3 {
		t.Fatal("Failed to get range keys:", err)
	}
	m, err = lsm.Range("2", "")
	if err != nil || len(m) != 3 {
		t.Fatal("Failed to get range keys:", err)
	}
}

func TestReverseRangeLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove()
	m, err := lsm.ReverseRange("non", "existant")
	if err != nil || len(m) != 0 {
		t.Fatal("Failed to get range keys:", err)
	}
	lsm.Insert("1", "1")
	lsm.Insert("2", "2")
	lsm.Insert("3", "3")
	lsm.Insert("4", "4")

	m, err = lsm.ReverseRange("", "")
	if err != nil || len(m) != 4 {
		t.Fatal("Failed to get range keys:", err)
	}
	m, err = lsm.ReverseRange("3", "2")
	if err != nil || len(m) != 2 {
		t.Fatal("Failed to get range keys:", err)
	}
	if v, ok := m["2"]; v != "2" || !ok {
		t.Fatal("Failed to get range keys:", err)
	}
	if v, ok := m["3"]; v != "3" || !ok {
		t.Fatal("Failed to get range keys:", err)
	}
	m, err = lsm.ReverseRange("", "3")
	if err != nil || len(m) != 2 {
		t.Fatal("Failed to get range keys:", err)
	}
	m, err = lsm.ReverseRange("2", "")
	if err != nil || len(m) != 2 {
		t.Fatal("Failed to get range keys:", err)
	}
}

func TestTxLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove()
	lsm.Begin()
	lsm.Insert("j", "ten")
	lsm.Insert("k", "eleven")
	lsm.Commit()

	m, _ := lsm.All()
	if len(m) != 2 {
		t.Fatal("Error in transactions")
	}

	lsm.Begin()
	lsm.Delete("j")
	lsm.Delete("k")
	lsm.Rollback()

	m, _ = lsm.All()
	if len(m) != 2 {
		t.Fatal("Error in transactions")
	}

	lsm.Begin()
	lsm.Insert("l", "twelve")
	lsm.Rollback()
	lsm.Insert("m", "thirteen")
	lsm.Commit()

	m, _ = lsm.All()
	if len(m) != 3 {
		t.Fatal("Error in transactions")
	}

	lsm.Begin()
	lsm.Begin()
	lsm.Begin()
	lsm.Delete("j")
	lsm.Commit()
	lsm.Commit()
	lsm.Begin()
	lsm.Begin()
	lsm.Begin()
	lsm.Delete("k")
	lsm.Rollback()
	lsm.Rollback()
	lsm.Delete("m")
	lsm.Commit()

	m, _ = lsm.All()
	if len(m) != 1 {
		t.Fatal("Error in transactions")
	}
}

func remove() {
	os.Remove("test")
	os.Remove("test-log")
	os.Remove("test-shm")
}

//go:build cgo
// +build cgo

package lsm1

// import "C"
import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"testing"
)

// func TestOpenLSM(t *testing.T) {
// 	name := "test"
// 	lsm, err := OpenLSM(name)
// 	defer remove(lsm)
// 	if err != nil || lsm == nil {
// 		t.Fatal("Failed to open database:", err)
// 	}
// }

// func TestCloseLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.Close()
// 	if err != nil {
// 		t.Fatal("Failed to close database:", err)
// 	}
// 	err = lsm.Close()
// 	if err != nil {
// 		t.Fatal("Failed to close database:", err)
// 	}
// }

// func TestInsertLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.Insert("key", "value")
// 	if err != nil {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// 	val, _ := lsm.Single("key")
// 	if val != "value" {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// 	err = lsm.Insert("key", "value2")
// 	if err != nil {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// 	val, _ = lsm.Single("key")
// 	if val != "value2" {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// }

// func TestDeleteLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.Delete("nonexistant")
// 	if err != nil {
// 		t.Fatal("Failed to delete key:", err)
// 	}
// 	lsm.Insert("key", "value")
// 	err = lsm.Delete("key")
// 	val, _ := lsm.Single("key")
// 	if val != "" {
// 		t.Fatal("Failed to delete key:", err)
// 	}
// 	err = lsm.Delete("")
// 	if err != nil {
// 		t.Fatal("Failed to delete key:", err)
// 	}
// }

// func TestDeleteRangeLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.DeleteRange("", "")
// 	if err != nil {
// 		t.Fatal("Failed to delete range keys:", err)
// 	}
// 	lsm.Insert("1", "1")
// 	lsm.Insert("2", "2")
// 	lsm.Insert("3", "3")
// 	lsm.Insert("4", "4")
// 	err = lsm.DeleteRange("2", "4")
// 	if err != nil {
// 		t.Fatal("Failed to delete range keys:", err)
// 	}
// 	m, _ := lsm.All()
// 	if len(m) != 3 {
// 		t.Fatal("Failed to delete range keys:", err)
// 	}
// }

// func TestSingleLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	val, err := lsm.Single("nonexistant")
// 	if err != nil || val != "" {
// 		t.Fatal("Failed to get key:", err)
// 	}
// 	lsm.Insert("key", "value")
// 	val, err = lsm.Single("key")
// 	if err != nil || val != "value" {
// 		t.Fatal("Failed to get key:", err)
// 	}
// }

// func TestAllLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	m, err := lsm.All()
// 	if err != nil || len(m) != 0 {
// 		t.Fatal("Failed to get all keys:", err)
// 	}
// 	lsm.Insert("key", "value")
// 	lsm.Insert("key1", "value")
// 	lsm.Insert("key2", "value")
// 	m, err = lsm.All()
// 	if err != nil || len(m) != 3 {
// 		t.Fatal("Failed to get all keys:", err)
// 	}
// }

// func TestRangeLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	m, err := lsm.Range("non", "existant")
// 	if err != nil || len(m) != 0 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	lsm.Insert("1", "1")
// 	lsm.Insert("2", "2")
// 	lsm.Insert("3", "3")
// 	lsm.Insert("4", "4")

// 	m, err = lsm.Range("", "")
// 	if err != nil || len(m) != 4 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.Range("2", "3")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["2"]; v != "2" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["3"]; v != "3" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.Range("", "3")
// 	if err != nil || len(m) != 3 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.Range("2", "")
// 	if err != nil || len(m) != 3 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// }

// func TestReverseRangeLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	m, err := lsm.ReverseRange("non", "existant")
// 	if err != nil || len(m) != 0 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	lsm.Insert("1", "1")
// 	lsm.Insert("2", "2")
// 	lsm.Insert("3", "3")
// 	lsm.Insert("4", "4")

// 	m, err = lsm.ReverseRange("", "")
// 	if err != nil || len(m) != 4 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.ReverseRange("3", "2")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["2"]; v != "2" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["3"]; v != "3" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.ReverseRange("", "3")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.ReverseRange("2", "")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// }
// //go:build cgo
// // +build cgo

// package lsm1

// // import "C"
// import (
// 	"os"
// 	"testing"
// )

// func TestOpenLSM(t *testing.T) {
// 	name := "test"
// 	lsm, err := OpenLSM(name)
// 	defer remove(lsm)
// 	if err != nil || lsm == nil {
// 		t.Fatal("Failed to open database:", err)
// 	}
// }

// func TestCloseLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.Close()
// 	if err != nil {
// 		t.Fatal("Failed to close database:", err)
// 	}
// 	err = lsm.Close()
// 	if err != nil {
// 		t.Fatal("Failed to close database:", err)
// 	}
// }

// func TestInsertLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.Insert("key", "value")
// 	if err != nil {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// 	val, _ := lsm.Single("key")
// 	if val != "value" {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// 	err = lsm.Insert("key", "value2")
// 	if err != nil {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// 	val, _ = lsm.Single("key")
// 	if val != "value2" {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// }

// func TestDeleteLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.Delete("nonexistant")
// 	if err != nil {
// 		t.Fatal("Failed to delete key:", err)
// 	}
// 	lsm.Insert("key", "value")
// 	err = lsm.Delete("key")
// 	val, _ := lsm.Single("key")
// 	if val != "" {
// 		t.Fatal("Failed to delete key:", err)
// 	}
// 	err = lsm.Delete("")
// 	if err != nil {
// 		t.Fatal("Failed to delete key:", err)
// 	}
// }

// func TestDeleteRangeLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.DeleteRange("", "")
// 	if err != nil {
// 		t.Fatal("Failed to delete range keys:", err)
// 	}
// 	lsm.Insert("1", "1")
// 	lsm.Insert("2", "2")
// 	lsm.Insert("3", "3")
// 	lsm.Insert("4", "4")
// 	err = lsm.DeleteRange("2", "4")
// 	if err != nil {
// 		t.Fatal("Failed to delete range keys:", err)
// 	}
// 	m, _ := lsm.All()
// 	if len(m) != 3 {
// 		t.Fatal("Failed to delete range keys:", err)
// 	}
// }

// func TestSingleLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	val, err := lsm.Single("nonexistant")
// 	if err != nil || val != "" {
// 		t.Fatal("Failed to get key:", err)
// 	}
// 	lsm.Insert("key", "value")
// 	val, err = lsm.Single("key")
// 	if err != nil || val != "value" {
// 		t.Fatal("Failed to get key:", err)
// 	}
// }

// func TestAllLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	m, err := lsm.All()
// 	if err != nil || len(m) != 0 {
// 		t.Fatal("Failed to get all keys:", err)
// 	}
// 	lsm.Insert("key", "value")
// 	lsm.Insert("key1", "value")
// 	lsm.Insert("key2", "value")
// 	m, err = lsm.All()
// 	if err != nil || len(m) != 3 {
// 		t.Fatal("Failed to get all keys:", err)
// 	}
// }

// func TestRangeLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	m, err := lsm.Range("non", "existant")
// 	if err != nil || len(m) != 0 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	lsm.Insert("1", "1")
// 	lsm.Insert("2", "2")
// 	lsm.Insert("3", "3")
// 	lsm.Insert("4", "4")

// 	m, err = lsm.Range("", "")
// 	if err != nil || len(m) != 4 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.Range("2", "3")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["2"]; v != "2" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["3"]; v != "3" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.Range("", "3")
// 	if err != nil || len(m) != 3 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.Range("2", "")
// 	if err != nil || len(m) != 3 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// }

// func TestReverseRangeLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	m, err := lsm.ReverseRange("non", "existant")
// 	if err != nil || len(m) != 0 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	lsm.Insert("1", "1")
// 	lsm.Insert("2", "2")
// 	lsm.Insert("3", "3")
// 	lsm.Insert("4", "4")

// 	m, err = lsm.ReverseRange("", "")
// 	if err != nil || len(m) != 4 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.ReverseRange("3", "2")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["2"]; v != "2" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["3"]; v != "3" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.ReverseRange("", "3")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.ReverseRange("2", "")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// }
// //go:build cgo
// // +build cgo

// package lsm1

// // import "C"
// import (
// 	"os"
// 	"testing"
// )

// func TestOpenLSM(t *testing.T) {
// 	name := "test"
// 	lsm, err := OpenLSM(name)
// 	defer remove(lsm)
// 	if err != nil || lsm == nil {
// 		t.Fatal("Failed to open database:", err)
// 	}
// }

// func TestCloseLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.Close()
// 	if err != nil {
// 		t.Fatal("Failed to close database:", err)
// 	}
// 	err = lsm.Close()
// 	if err != nil {
// 		t.Fatal("Failed to close database:", err)
// 	}
// }

// func TestInsertLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.Insert("key", "value")
// 	if err != nil {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// 	val, _ := lsm.Single("key")
// 	if val != "value" {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// 	err = lsm.Insert("key", "value2")
// 	if err != nil {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// 	val, _ = lsm.Single("key")
// 	if val != "value2" {
// 		t.Fatal("Failed to insert into database:", err)
// 	}
// }

// func TestDeleteLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.Delete("nonexistant")
// 	if err != nil {
// 		t.Fatal("Failed to delete key:", err)
// 	}
// 	lsm.Insert("key", "value")
// 	err = lsm.Delete("key")
// 	val, _ := lsm.Single("key")
// 	if val != "" {
// 		t.Fatal("Failed to delete key:", err)
// 	}
// 	err = lsm.Delete("")
// 	if err != nil {
// 		t.Fatal("Failed to delete key:", err)
// 	}
// }

// func TestDeleteRangeLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	err := lsm.DeleteRange("", "")
// 	if err != nil {
// 		t.Fatal("Failed to delete range keys:", err)
// 	}
// 	lsm.Insert("1", "1")
// 	lsm.Insert("2", "2")
// 	lsm.Insert("3", "3")
// 	lsm.Insert("4", "4")
// 	err = lsm.DeleteRange("2", "4")
// 	if err != nil {
// 		t.Fatal("Failed to delete range keys:", err)
// 	}
// 	m, _ := lsm.All()
// 	if len(m) != 3 {
// 		t.Fatal("Failed to delete range keys:", err)
// 	}
// }

// func TestSingleLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	val, err := lsm.Single("nonexistant")
// 	if err != nil || val != "" {
// 		t.Fatal("Failed to get key:", err)
// 	}
// 	lsm.Insert("key", "value")
// 	val, err = lsm.Single("key")
// 	if err != nil || val != "value" {
// 		t.Fatal("Failed to get key:", err)
// 	}
// }

// func TestAllLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	m, err := lsm.All()
// 	if err != nil || len(m) != 0 {
// 		t.Fatal("Failed to get all keys:", err)
// 	}
// 	lsm.Insert("key", "value")
// 	lsm.Insert("key1", "value")
// 	lsm.Insert("key2", "value")
// 	m, err = lsm.All()
// 	if err != nil || len(m) != 3 {
// 		t.Fatal("Failed to get all keys:", err)
// 	}
// }

// func TestRangeLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	m, err := lsm.Range("non", "existant")
// 	if err != nil || len(m) != 0 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	lsm.Insert("1", "1")
// 	lsm.Insert("2", "2")
// 	lsm.Insert("3", "3")
// 	lsm.Insert("4", "4")

// 	m, err = lsm.Range("", "")
// 	if err != nil || len(m) != 4 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.Range("2", "3")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["2"]; v != "2" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["3"]; v != "3" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.Range("", "3")
// 	if err != nil || len(m) != 3 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.Range("2", "")
// 	if err != nil || len(m) != 3 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// }

// func TestReverseRangeLSM(t *testing.T) {
// 	lsm, _ := OpenLSM("test")
// 	defer remove(lsm)
// 	m, err := lsm.ReverseRange("non", "existant")
// 	if err != nil || len(m) != 0 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	lsm.Insert("1", "1")
// 	lsm.Insert("2", "2")
// 	lsm.Insert("3", "3")
// 	lsm.Insert("4", "4")

// 	m, err = lsm.ReverseRange("", "")
// 	if err != nil || len(m) != 4 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.ReverseRange("3", "2")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["2"]; v != "2" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	if v, ok := m["3"]; v != "3" || !ok {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.ReverseRange("", "3")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// 	m, err = lsm.ReverseRange("2", "")
// 	if err != nil || len(m) != 2 {
// 		t.Fatal("Failed to get range keys:", err)
// 	}
// }

func TestTxLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	lsm.Tx(func() error {
		lsm.Insert("j", "ten")
		lsm.Insert("k", "eleven")
		return nil
	})
	defer remove(lsm)
	m, _ := lsm.All()
	if len(m) != 2 {
		t.Fatal("Error in transactions")
	}

	lsm.Tx(func() error {
		lsm.Insert("m", "ten")
		return fmt.Errorf("an error")
	})
	m, _ = lsm.All()
	if len(m) != 2 {
		t.Fatal("Error in transactions")
	}
}

func TestBatchTxLSM(t *testing.T) {
	lsm, _ := OpenLSM("test")
	defer remove(lsm)
	m := make(map[string]string)
	for i := 0; i < 1000; i++ {
		k := rand.Int() % int(10*1000000)
		key := fmt.Sprintf("vsz=%05d-k=%010d", 128, k)
		val := `\x8c-ͣ\t\xf4\x06G\xb1\x17z\x16\x18Bڢ0}\x05\xfe\x818\xe7\xa1W\x04٦_\"t\xe4#\x98\x04\xca\xc5y!\xf5\xcb\x1c\x19\xffW\xc3o\x12\x12\xa3\x0ew/\xbe\xd93D\x8c\xe7\xddXf\xf0\xdcD^cO\xb2\x10w\xfa\xf7\xfdzp\x94\x8c\xffL\r\xc2\xdc\x1c\xd7\xe5?E}\xf3\u05fa\xa3ĭ42+\x01\xa3\xcaS\xb7\xec\xf7\xdc\xf9\t\x88|Hc\xd4LX\xfe\x1bd[?\x83\xf5<\u007f\xecϬ\`
		m[key] = val
	}
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			lsm.Tx(func() error {
				for k, v := range m {
					err := lsm.Insert(k, v)
					if err != nil {
						fmt.Println("WTF")
						return err
					}
				}
				return nil
			})
			wg.Done()
		}()
	}
	wg.Wait()
}

func remove(l *LSMTable) {
	err := l.Close()
	if err != nil {
		panic(err)
	}
	os.Remove("test")
	os.Remove("test-log")
	os.Remove("test-shm")

}

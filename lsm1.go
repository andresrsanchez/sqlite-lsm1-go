package lsm1

/*
#include "lsm.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type LSMTable struct {
	db            *C.lsm_db
	isOpen        bool
	ntransactions int
}

var errMSG map[C.int]string = map[C.int]string{
	C.LSM_ERROR:    "Error",
	C.LSM_BUSY:     "Busy",
	C.LSM_NOMEM:    "Out of memory",
	C.LSM_READONLY: "Database is read-only",
	C.LSM_IOERR:    "Unspecified IO error",
	C.LSM_CORRUPT:  "Database is corrupt",
	C.LSM_FULL:     "Storage device is full",
	C.LSM_CANTOPEN: "Cannot open database",
	C.LSM_PROTOCOL: "Protocol error",
	C.LSM_MISUSE:   "Misuse",
	C.LSM_MISMATCH: "Mismatch",
}

func getError(code C.int) error {
	if val, ok := errMSG[code]; ok {
		return fmt.Errorf(val)
	}
	return fmt.Errorf("lsm error")
}

func OpenLSM(name string) (*LSMTable, error) {
	var env *C.lsm_env
	var db *C.lsm_db
	ok := C.lsm_new(env, &db)
	if ok != C.LSM_OK {
		return nil, getError(ok)
	}
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	ok = C.lsm_open(db, cn)
	if ok != C.LSM_OK {
		return nil, getError(ok)
	}
	return &LSMTable{db: db, isOpen: true}, nil
}

func (l *LSMTable) Close() error {
	if !l.isOpen {
		return nil
	}
	ok := C.lsm_close(l.db)
	if ok != C.LSM_OK {
		return getError(ok)
	}
	l.isOpen = false
	return nil
}

func (l *LSMTable) Insert(k, v string) error {
	var ck *C.char = C.CString(k)
	var cv *C.char = C.CString(v)
	defer func() {
		C.free(unsafe.Pointer(ck))
		C.free(unsafe.Pointer(cv))
	}()
	ok := C.lsm_insert(l.db, unsafe.Pointer(ck), C.int(len(k)), unsafe.Pointer(cv), C.int(len(v)))
	if ok != C.LSM_OK {
		return getError(ok)
	}
	return nil
}

func (l *LSMTable) Delete(k string) error {
	var ck *C.char = C.CString(k)
	defer C.free(unsafe.Pointer(ck))
	ok := C.lsm_delete(l.db, unsafe.Pointer(ck), C.int(len(k)))
	if ok != C.LSM_OK {
		return getError(ok)
	}
	return nil
}

//implement end = "" to delete a range from start, same as start == ""
//start end not included wtf
func (l *LSMTable) DeleteRange(start, end string) error {
	var cstart *C.char = C.CString(start)
	var cend *C.char = C.CString(end)
	defer func() {
		C.free(unsafe.Pointer(cstart))
		C.free(unsafe.Pointer(cend))
	}()
	ok := C.lsm_delete_range(l.db, unsafe.Pointer(cstart), C.int(len(start)), unsafe.Pointer(cend), C.int(len(end)))
	if ok != C.LSM_OK {
		return getError(ok)
	}
	return nil
}

func (l *LSMTable) Single(k string) (string, error) {
	var csr *C.lsm_cursor
	ok := C.lsm_csr_open(l.db, &csr)
	defer C.lsm_csr_close(csr)
	if ok != C.LSM_OK {
		return "", getError(ok)
	}
	var ck *C.char = C.CString(k)
	defer C.free(unsafe.Pointer(ck))
	ok = C.lsm_csr_seek(csr, unsafe.Pointer(ck), C.int(len(k)), C.LSM_SEEK_EQ)
	if ok != C.LSM_OK {
		return "", getError(ok)
	} else if C.lsm_csr_valid(csr) == C.LSM_OK {
		return "", nil
	}
	var cvalue *C.char
	defer C.free(unsafe.Pointer(cvalue))
	var clen C.int
	ok = C.lsm_csr_value(csr, (*unsafe.Pointer)(unsafe.Pointer(&cvalue)), &clen)
	if ok != C.LSM_OK {
		return "", getError(ok)
	}
	r := C.GoStringN(cvalue, clen)
	return r, nil
}

func (l *LSMTable) All() (map[string]string, error) {
	var csr *C.lsm_cursor
	ok := C.lsm_csr_open(l.db, &csr)
	defer C.lsm_csr_close(csr)
	if ok != C.LSM_OK {
		return nil, getError(ok)
	}
	r := make(map[string]string)
	for ok := C.lsm_csr_first(csr); ok == C.LSM_OK && C.lsm_csr_valid(csr) != C.LSM_OK; ok = C.lsm_csr_next(csr) {
		k, v, err := getKV(csr)
		if err != nil {
			return nil, err
		}
		r[k] = v
	}
	return r, nil
}

func (l *LSMTable) Range(start, end string) (map[string]string, error) {
	var csr *C.lsm_cursor
	ok := C.lsm_csr_open(l.db, &csr)
	defer C.lsm_csr_close(csr)
	if ok != C.LSM_OK {
		return nil, getError(ok)
	}
	r := make(map[string]string)
	var ck *C.char = C.CString(start)
	defer C.free(unsafe.Pointer(ck))
	var res C.int
	for ok := C.lsm_csr_seek(csr, unsafe.Pointer(ck), C.int(len(start)), C.LSM_SEEK_GE); ok == C.LSM_OK && C.lsm_csr_valid(csr) != C.LSM_OK; ok = C.lsm_csr_next(csr) {
		if end != "" { //weird
			cv := C.CString(end)
			ok = C.lsm_csr_cmp(csr, unsafe.Pointer(cv), C.int(len(end)), &res)
			if ok != C.LSM_OK {
				return nil, getError(ok)
			} else if res > 0 {
				break
			}
			C.free(unsafe.Pointer(cv))
		}
		k, v, err := getKV(csr)
		if err != nil {
			return nil, err
		}
		r[k] = v
	}
	return r, nil
}

func (l *LSMTable) ReverseRange(start, end string) (map[string]string, error) {
	var csr *C.lsm_cursor
	ok := C.lsm_csr_open(l.db, &csr)
	defer C.lsm_csr_close(csr)
	if ok != C.LSM_OK {
		return nil, getError(ok)
	}
	r := make(map[string]string)
	var ck *C.char = C.CString(start)
	defer C.free(unsafe.Pointer(ck))
	var res C.int
	var valid C.int
	if start == "" {
		valid = C.lsm_csr_last(csr)
	} else {
		valid = C.lsm_csr_seek(csr, unsafe.Pointer(ck), C.int(len(start)), C.LSM_SEEK_LE)
	}
	for ok := valid; ok == C.LSM_OK && C.lsm_csr_valid(csr) != C.LSM_OK; ok = C.lsm_csr_prev(csr) {
		if end != "" { //weird
			cv := C.CString(end)
			ok = C.lsm_csr_cmp(csr, unsafe.Pointer(cv), C.int(len(end)), &res)
			if ok != C.LSM_OK {
				return nil, getError(ok)
			} else if res < 0 {
				break
			}
			C.free(unsafe.Pointer(cv))
		}
		k, v, err := getKV(csr)
		if err != nil {
			return nil, err
		}
		r[k] = v
	}
	return r, nil
}

func getKV(csr *C.lsm_cursor) (string, string, error) {
	var cstartkey, cstartval *C.char
	var cstartlen, cstartvallen C.int
	ok := C.lsm_csr_key(csr, (*unsafe.Pointer)(unsafe.Pointer(&cstartkey)), &cstartlen)
	if ok != C.LSM_OK {
		return "", "", getError(ok)
	}
	ok = C.lsm_csr_value(csr, (*unsafe.Pointer)(unsafe.Pointer(&cstartval)), &cstartvallen)
	if ok != C.LSM_OK {
		return "", "", getError(ok)
	}
	key := C.GoStringN(cstartkey, cstartlen)
	val := C.GoStringN(cstartval, cstartvallen)
	return key, val, nil
}

func (l *LSMTable) Checkpoint() (int, error) {
	var pnKB C.int
	ok := C.lsm_checkpoint(l.db, &pnKB)
	if ok != C.LSM_OK {
		return 0, getError(ok)
	}
	return int(pnKB), nil
}

func (l *LSMTable) Work(nKB int) (int, error) {
	var nWrite C.int
	ok := C.lsm_work(l.db, 1, C.int(nKB), &nWrite)
	if ok != C.LSM_OK {
		return 0, getError(ok)
	}
	return int(nWrite), nil
}

func (l *LSMTable) Flush() error {
	ok := C.lsm_flush(l.db)
	if ok != C.LSM_OK {
		return getError(ok)
	}
	return nil
}

func (l *LSMTable) Begin() error {
	if l.ntransactions < 0 { //lol
		l.ntransactions = 0
	}
	l.ntransactions += 1
	ok := C.lsm_begin(l.db, C.int(l.ntransactions))
	if ok != C.LSM_OK {
		return getError(ok)
	}
	return nil
}

func (l *LSMTable) Commit() error {
	if l.ntransactions < 0 {
		return fmt.Errorf("no transactions on course")
	}
	l.ntransactions -= 1
	ok := C.lsm_commit(l.db, C.int(l.ntransactions))
	if ok != C.LSM_OK {
		return getError(ok)
	}
	return nil
}

func (l *LSMTable) Rollback() error {
	if l.ntransactions < 0 {
		return fmt.Errorf("no transactions on course")
	}
	l.ntransactions -= 1
	ok := C.lsm_rollback(l.db, C.int(l.ntransactions))
	if ok != C.LSM_OK {
		return getError(ok)
	}
	return nil
}

package gls

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var gls struct {
	m map[int64]map[interface{}]interface{}
	sync.RWMutex
}

func init() {
	gls.m = make(map[int64]map[interface{}]interface{})
}

func GetGoId() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)
	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}
	return int64(id)
}

func Get(key interface{}) interface{} {
	gls.RLock()
	defer gls.RUnlock()
	goId := GetGoId()
	return gls.m[goId][key]
}

func Set(key interface{}, v interface{}) {
	gls.Lock()
	defer gls.Unlock()
	goId := GetGoId()
	if _, ok := gls.m[goId][key]; !ok {
		gls.m[goId] = make(map[interface{}]interface{})
	}
	gls.m[goId][key] = v
}

func Clean() {
	gls.Lock()
	defer gls.Unlock()
	delete(gls.m, GetGoId())
}

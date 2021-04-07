package gls

import (
	"fmt"
	"strconv"
	"testing"
)

var key = "test"

func GO(fn func()) {
	go fn()
}

func TestGls(t *testing.T) {
	GO(func() {
		defer Clean()
		Set(key, "aaa")
		testPrint()
	})
	GO(func() {
		defer Clean()
		Set(key, "bbb")
		testPrint()
	})
	GO(func() {
		defer Clean()
		Set(key, "ccc")
		testPrint()
	})
}


func BenchmarkGls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GO(func() {
			defer Clean()
			Set(key, "ccc"+strconv.FormatInt(GetGoId(), 10))
			testGet()
		})
	}
}

func testPrint() {
	fmt.Println(Get(key), GetGoId())
}

func testGet() {
	Get(key)
}


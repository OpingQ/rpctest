package main

import (
	"fmt"
	"os"
	"path/filepath"
	"rpctest/framework"
	"testing"

	"github.com/BurntSushi/toml"
)

func Test_main(t *testing.T) {
	// s := "123456789"
	// s = Reverse(s)
	// fmt.Println(s)
	// s2 := []int{1, 2, 3, 4, 5, 6}
	// s3 := ReverseIn(s2)
	// fmt.Println(s3)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dir = dir + `\environment\local.toml`
	var config framework.Config
	toml.DecodeFile(dir, &config)
	fmt.Println(config)
}

func Reverse(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseIn(ro interface{}) interface{} {
	if r, isConvert := ro.([]interface{}); isConvert {
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return r
	}
	return nil
}

type KeyStruct struct {
	Name  string
	Price int
}

func Test_testMain(t *testing.T) {
	ks := KeyStruct{
		Name: "test",
	}
	mapKey := make(map[KeyStruct]int, 0)
	mapKey[ks] = 19
	ks2 := KeyStruct{
		Name: "test",
	}
	// 若組成結構一樣將是屬同一個KEY值
	if value, hasKey := mapKey[ks2]; hasKey {
		fmt.Println(value)
	} else {
		fmt.Println("無值")
	}
}

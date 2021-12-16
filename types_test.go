package winbooks_test

import (
	"encoding/json"
	"log"
	"testing"

	winbooks "github.com/omniboost/go-winbooks-web"
)

func TestValueString(t *testing.T) {
	// t1 := winbooks.ValueString("")
	// b, err := json.Marshal(t1)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	// var t2 *winbooks.ValueString
	// b, err = json.Marshal(t2)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	// b = []byte(`
	// 	{"T": null}
	// `)
	// t3 := struct {
	// 	T *winbooks.ValueString
	// }{}
	// err = json.Unmarshal(b, &t3)
	// if err != nil {
	// 	t.Error(err)
	// }
	// b, err = json.Marshal(t3)
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(string(b))

	t1 := winbooks.ValueNullString{nil}
	b, err := json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))

	var s winbooks.ValueString
	s = ""
	t1 = winbooks.ValueNullString{&s}
	b, err = json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))
}

package winbooks_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSubaccountGetAll(t *testing.T) {
	req := client.NewSubaccountGetAll()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

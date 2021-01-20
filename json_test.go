package env

import (
	"fmt"
	"testing"
)

func TestJSON(t *testing.T) {
	fmt.Printf("Running Test Group: %s\n", t.Name())

	// test set up
	SetVar("TEST_JSON_RAW", "{\"key\":\"my value\"}")

	type Test struct {
		Key string `json:"key"`
	}

	// tests
	t.Run("TestJSONRaw", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetJSONRaw("TEST_JSON_RAW")
		if err != nil {
			t.Errorf("Expected to receive json.RawMessage, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
		//fmt.Println(string(v))
	})

	t.Run("TestUnmarshal", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		s := Test{}
		err := Unmarshal("TEST_JSON_RAW", &s)
		if err != nil {
			t.Errorf("Expected to unmarshal JSON to struct, received error: %v\n", err)
		}
		if s.Key != "my value" {
			t.Errorf("Expected to read 'my value' from struct field, received %s", s.Key)
		}
		//fmt.Printf("%+v", s)
	})
}

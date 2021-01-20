package env

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestBasic(t *testing.T) {
	fmt.Printf("Running Test Group: %s\n", t.Name())

	// test set up
	SetVar("TEST_STRING", "This is a string")
	SetVar("TEST_BOOL", "true")
	SetVar("TEST_UUID", "d1b0eaca-3b7c-4fef-88fe-75c9db4bf90a")
	SetVar("TEST_RUNE", "0b£")
	SetVar("TEST_BYTES", "test")
	SetVar("TEST_DURATION", "3h")

	SetVar("TEST_INT", "42")
	SetVar("TEST_INT8", "127")
	SetVar("TEST_INT16", "32767")
	SetVar("TEST_INT32", "2147483647")
	SetVar("TEST_INT64", "9223372036854775807")

	SetVar("TEST_UINT", "1")
	SetVar("TEST_UINT8", "255")
	SetVar("TEST_UINT16", "65535")
	SetVar("TEST_UINT32", "4294967295")
	SetVar("TEST_UINT64", "18446744073709551615")

	SetVar("TEST_FLOAT32", "42.42")
	SetVar("TEST_FLOAT64", "42.42")

	SetVar("TEST_COMPLEX64", "3.4E+38")
	SetVar("TEST_COMPLEX128", "+1.7E+308")

	// tests
	t.Run("TestGetString", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetString("TEST_STRING")
		if err != nil {
			t.Errorf("Expected to recieve string, received error: %v\n", err)
		}
		if v == "" {
			t.Error("Expected to receive value, received empty string")
		}
	})

	t.Run("TestGetBool", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetBool("TEST_BOOL")
		if err != nil {
			t.Errorf("Expected to recieve bool, received error: %v\n", err)
		}
		if v != true {
			t.Errorf("Expected to receive true, received: %v\n", v)
		}
	})

	t.Run("TestGetUUID", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUUID("TEST_UUID")
		if err != nil {
			t.Errorf("Expected to recieve uuid, received error: %v\n", err)
		}
		if v == uuid.Nil {
			t.Errorf("Expected to receive UUID, received: %v\n", v)
		}
	})

	t.Run("TestGetRune", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetRune("TEST_RUNE")
		if err != nil {
			t.Errorf("Expected to recieve rune, received error: %v\n", err)
		}
		if v == nil {
			t.Errorf("Expected to receive []rune, received: %v\n", v)
		}
	})

	t.Run("TestGetBytes", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetBytes("TEST_BYTES")
		if err != nil {
			t.Errorf("Expected to recieve bytes, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive []byte, received nothing\n")
		}
	})

	t.Run("TestGetDuration", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetDuration("TEST_DURATION")
		if err != nil {
			t.Errorf("Expected to recieve duration, received error: %v\n", err)
		}
		if v != time.Duration(3*time.Hour) {
			t.Errorf("Expected to receive duration of 3hr, received %v\n", v)
		}
	})

	t.Run("TestGetInt", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetInt("TEST_INT")
		if err != nil {
			t.Errorf("Expected to recieve int, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetInt8", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetInt8("TEST_INT8")
		if err != nil {
			t.Errorf("Expected to recieve int8, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetInt16", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetInt16("TEST_INT16")
		if err != nil {
			t.Errorf("Expected to recieve int16, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetInt32", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetInt32("TEST_INT32")
		if err != nil {
			t.Errorf("Expected to recieve int32, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetInt64", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetInt64("TEST_INT64")
		if err != nil {
			t.Errorf("Expected to recieve int64, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetUint", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUint("TEST_UINT")
		if err != nil {
			t.Errorf("Expected to recieve uint, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetUint8", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUint8("TEST_UINT8")
		if err != nil {
			t.Errorf("Expected to recieve uint8, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetUint16", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUint16("TEST_UINT16")
		if err != nil {
			t.Errorf("Expected to recieve uint16, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetUint32", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUint32("TEST_UINT32")
		if err != nil {
			t.Errorf("Expected to recieve uint32, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetUint64", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUint64("TEST_UINT64")
		if err != nil {
			t.Errorf("Expected to recieve uint64, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive integer != 0, received %v\n", v)
		}
	})

	t.Run("TestGetFloat32", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetFloat32("TEST_FLOAT32")
		if err != nil {
			t.Errorf("Expected to recieve float32, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive float32 != 0, received %v\n", v)
		}
	})

	t.Run("TestGetFloat64", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetFloat64("TEST_FLOAT64")
		if err != nil {
			t.Errorf("Expected to recieve float64, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive float64 != 0, received %v\n", v)
		}
	})

	t.Run("TestGetComplex64", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetComplex64("TEST_COMPLEX64")
		if err != nil {
			t.Errorf("Expected to recieve complex64, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive complex64 != 0, received %v\n", v)
		}
	})

	t.Run("TestGetComplex128", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetComplex128("TEST_COMPLEX128")
		if err != nil {
			t.Errorf("Expected to recieve complex128, received error: %v\n", err)
		}
		if v == 0 {
			t.Errorf("Expected to receive complex128 != 0, received %v\n", v)
		}
	})

}

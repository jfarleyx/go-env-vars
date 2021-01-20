package env

import (
	"fmt"
	"testing"
)

func TestArrays(t *testing.T) {
	fmt.Printf("Running Test Group: %s\n", t.Name())

	// test set up
	SetVar("TEST_SLICE_STRING", "A,B,C")
	SetVar("TEST_SLICE_UUID", "e5ef78a4-6299-4428-8f0a-f20c8a8b2d62,c8a53726-5adc-4d3d-a765-064d0ddb4e9f")
	SetVar("TEST_SLICE_DURATION", "3h,10m,100s")
	SetVar("TEST_SLICE_INT", "1,2,3")
	SetVar("TEST_SLICE_INT8", "1,2,3")
	SetVar("TEST_SLICE_INT16", "1,2,3")
	SetVar("TEST_SLICE_INT32", "1,2,3")
	SetVar("TEST_SLICE_INT64", "1,2,3")
	SetVar("TEST_SLICE_UINT", "1,2,3")
	SetVar("TEST_SLICE_UINT8", "1,2,3")
	SetVar("TEST_SLICE_UINT16", "1,2,3")
	SetVar("TEST_SLICE_UINT32", "1,2,3")
	SetVar("TEST_SLICE_UINT64", "1,2,3")
	SetVar("TEST_SLICE_FLOAT32", "2.3,3.4,4.5")
	SetVar("TEST_SLICE_FLOAT64", "2.3,3.4,4.5")
	SetVar("TEST_SLICE_COMPLEX64", "3.4E+35,3.4E+36,3.4E+37")
	SetVar("TEST_SLICE_COMPLEX128", "+1.7E+306,+1.7E+307")

	// tests
	t.Run("TestGetStringSlice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetStringSlice("TEST_SLICE_STRING")
		if err != nil {
			t.Errorf("Expected to recieve slice of string, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetUUIDSlice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetStringSlice("TEST_SLICE_UUID")
		if err != nil {
			t.Errorf("Expected to recieve slice of uuid, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetDurationSlice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetDurationSlice("TEST_SLICE_DURATION")
		if err != nil {
			t.Errorf("Expected to recieve slice of time.Duration, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetIntSlice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetIntSlice("TEST_SLICE_INT")
		if err != nil {
			t.Errorf("Expected to recieve slice of int, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetInt8Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetInt8Slice("TEST_SLICE_INT8")
		if err != nil {
			t.Errorf("Expected to recieve slice of int8, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetInt16Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetInt16Slice("TEST_SLICE_INT16")
		if err != nil {
			t.Errorf("Expected to recieve slice of int16, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetInt32Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetInt32Slice("TEST_SLICE_INT32")
		if err != nil {
			t.Errorf("Expected to recieve slice of int32, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetInt64Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetInt64Slice("TEST_SLICE_INT64")
		if err != nil {
			t.Errorf("Expected to recieve slice of int64, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetUintSlice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUintSlice("TEST_SLICE_UINT")
		if err != nil {
			t.Errorf("Expected to recieve slice of uint, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetUint8Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUint8Slice("TEST_SLICE_UINT8")
		if err != nil {
			t.Errorf("Expected to recieve slice of uint8, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetUint16Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUint16Slice("TEST_SLICE_UINT16")
		if err != nil {
			t.Errorf("Expected to recieve slice of uint16, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetUint32Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUint32Slice("TEST_SLICE_UINT32")
		if err != nil {
			t.Errorf("Expected to recieve slice of uint32, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetUint64Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetUint64Slice("TEST_SLICE_UINT64")
		if err != nil {
			t.Errorf("Expected to recieve slice of uint64, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetFloat32Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetFloat32Slice("TEST_SLICE_FLOAT32")
		if err != nil {
			t.Errorf("Expected to recieve slice of float32, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetFloat64Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetFloat64Slice("TEST_SLICE_FLOAT64")
		if err != nil {
			t.Errorf("Expected to recieve slice of float64, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetComplex64Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetComplex64Slice("TEST_SLICE_COMPLEX64")
		if err != nil {
			t.Errorf("Expected to recieve slice of complex64, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})

	t.Run("TestGetComplex128Slice", func(t *testing.T) {
		fmt.Printf("  %s\n", t.Name())
		v, err := GetComplex128Slice("TEST_SLICE_COMPLEX128")
		if err != nil {
			t.Errorf("Expected to recieve slice of complex128, received error: %v\n", err)
		}
		if len(v) == 0 {
			t.Error("Expected to receive slice of values, received nothing")
		}
	})
}

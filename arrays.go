package env

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// GetStringSlice gets a comma delimited list of strings and returns them as a slice of string.
func GetStringSlice(name string) ([]string, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var sa []string
	for idx := range vals {
		sa = append(sa, vals[idx])
	}

	return sa, nil
}

// GetUUIDSlice gets a comma delimted list of UUID's and returns them as a slice of UUID.
func GetUUIDSlice(name string) ([]uuid.UUID, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ua []uuid.UUID
	for idx := range vals {
		i, err := uuid.Parse(vals[idx])
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ua = append(ua, i)
	}

	return ua, nil
}

// GetIntSlice gets a comma delimited list of integers and returns them as a slice of int.
func GetIntSlice(name string) ([]int, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []int
	for idx := range vals {
		i, err := strconv.Atoi(vals[idx])
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, i)
	}

	return ia, nil
}

// GetInt8Slice gets a comma delimited list of int8 and returns them as a slice of int8.
func GetInt8Slice(name string) ([]int8, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []int8
	for idx := range vals {
		i, err := strconv.ParseInt(vals[idx], 10, 8)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, int8(i))
	}

	return ia, nil
}

// GetInt16Slice gets a comma delimited list of int16 and returns them as a slice of int16.
func GetInt16Slice(name string) ([]int16, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []int16
	for idx := range vals {
		i, err := strconv.ParseInt(vals[idx], 10, 16)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, int16(i))
	}

	return ia, nil
}

// GetInt32Slice gets a comma delimited list of int32 and returns them as a slice of int32.
func GetInt32Slice(name string) ([]int32, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []int32
	for idx := range vals {
		i, err := strconv.ParseInt(vals[idx], 10, 32)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, int32(i))
	}

	return ia, nil
}

// GetInt64Slice gets a comma delimited list of int64 and returns them as a slice of int64.
func GetInt64Slice(name string) ([]int64, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []int64
	for idx := range vals {
		i, err := strconv.ParseInt(vals[idx], 10, 64)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, i)
	}

	return ia, nil
}

// GetUintSlice gets a comma delimited list of uint and returns them as a slice of uint.
func GetUintSlice(name string) ([]uint, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []uint
	for idx := range vals {
		i, err := strconv.ParseUint(vals[idx], 10, 64)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, uint(i))
	}

	return ia, nil
}

// GetUint8Slice gets a comma delimited list of uint8 and returns them as a slice of uint8.
func GetUint8Slice(name string) ([]uint8, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []uint8
	for idx := range vals {
		i, err := strconv.ParseUint(vals[idx], 10, 8)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, uint8(i))
	}

	return ia, nil
}

// GetUint16Slice gets a comma delimited list of uint16 and returns them as a slice of uint16.
func GetUint16Slice(name string) ([]uint16, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []uint16
	for idx := range vals {
		i, err := strconv.ParseUint(vals[idx], 10, 16)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, uint16(i))
	}

	return ia, nil
}

// GetUint32Slice gets a comma delimited list of uint32 and returns them as a slice of uint32.
func GetUint32Slice(name string) ([]uint32, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []uint32
	for idx := range vals {
		i, err := strconv.ParseUint(vals[idx], 10, 32)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, uint32(i))
	}

	return ia, nil
}

// GetUint64Slice gets a comma delimited list of uint64 and returns them as a slice of uint64.
func GetUint64Slice(name string) ([]uint64, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []uint64
	for idx := range vals {
		i, err := strconv.ParseUint(vals[idx], 10, 64)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, i)
	}

	return ia, nil
}

// GetFloat32Slice gets a comma delimited list of float32 values and returns them as a slice of float32.
func GetFloat32Slice(name string) ([]float32, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []float32
	for idx := range vals {
		i, err := strconv.ParseFloat(vals[idx], 32)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, float32(i))
	}

	return ia, nil
}

// GetFloat64Slice gets a comma delimited list of float64 values and returns them as a slice of float64.
func GetFloat64Slice(name string) ([]float64, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []float64
	for idx := range vals {
		i, err := strconv.ParseFloat(vals[idx], 64)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, float64(i))
	}

	return ia, nil
}

// GetComplex64Slice gets a comma delimited list of complex64 values and returns them as a slice of complex64.
func GetComplex64Slice(name string) ([]complex64, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []complex64
	for idx := range vals {
		i, err := strconv.ParseComplex(vals[idx], 64)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, complex64(i))
	}

	return ia, nil
}

// GetComplex128Slice gets a comma delimited list of complex128 values and returns them as a slice of complex128.
func GetComplex128Slice(name string) ([]complex128, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []complex128
	for idx := range vals {
		i, err := strconv.ParseComplex(vals[idx], 128)
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, complex128(i))
	}

	return ia, nil
}

// GetDurationSlice gets a comma delimited list of time.Duration values and returns them as a slice of time.Duration.
func GetDurationSlice(name string) ([]time.Duration, error) {
	vals, err := getList(name)
	if err != nil {
		return nil, err
	}

	var ia []time.Duration
	for idx := range vals {
		i, err := time.ParseDuration(vals[idx])
		if err != nil {
			return nil, NewError(InvalidType, name, err.Error())
		}
		ia = append(ia, i)
	}

	return ia, nil
}

func getList(name string) ([]string, error) {
	v, err := lookup(name)
	if err != nil {
		return nil, err
	}

	return strings.Split(v, ","), nil
}

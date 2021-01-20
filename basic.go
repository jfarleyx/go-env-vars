package env

import (
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// SetVar sets an environment variable and its value.
func SetVar(name string, value string) error {
	return os.Setenv(name, value)
}

// lookup looks up an env var and returns an error if the env var is missing or
// the value is empty.
func lookup(name string) (string, error) {
	v, ok := os.LookupEnv(name)
	if !ok {
		return "", NewError(NotFound, name, "")
	}
	if v == "" {
		return "", NewError(EmptyValue, name, "")
	}

	return v, nil
}

// GetString returns the environment variable value as a string. Returns
// an error for missing env var or empty value.
func GetString(name string) (string, error) {
	v, err := lookup(name)
	if err != nil {
		return "", err
	}

	return v, nil
}

// GetBool returns the environment variable value as a boolean. Returns
// an error for missing env var or empty value.
func GetBool(name string) (bool, error) {
	i, err := lookup(name)
	if err != nil {
		return false, err
	}

	v, err := strconv.ParseBool(i)
	if err != nil {
		return false, NewError(InvalidType, name, err.Error())
	}

	return v, nil
}

// GetUUID returns the environment variable value as a UUID. Returns
// an error for missing env var or empty value.
func GetUUID(name string) (uuid.UUID, error) {
	i, err := lookup(name)
	if err != nil {
		return uuid.Nil, err
	}

	v, err := uuid.Parse(i)
	if err != nil {
		return uuid.Nil, NewError(InvalidType, name, err.Error())
	}

	return v, nil
}

// GetRune returns the environment variable value as a Rune type. Returns
// and error for missing env var or empty value.
func GetRune(name string) ([]rune, error) {
	i, err := lookup(name)
	if err != nil {
		return nil, err
	}

	return []rune(i), nil
}

// GetBytes returns the environment variable value as a byte slice. Returns
// and error for missing env var or empty value.
func GetBytes(name string) ([]byte, error) {
	i, err := lookup(name)
	if err != nil {
		return nil, err
	}

	return []byte(i), nil
}

// GetDuration returns the environment variable value as a time.Duration type.  Returns
// and error for missing env var or empty value.
func GetDuration(name string) (time.Duration, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := time.ParseDuration(i)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return v, nil
}

// GetInt returns the environment variable value as an int. Returns
// an error for missing env var or empty value.
func GetInt(name string) (int, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.Atoi(i)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return v, nil
}

// GetInt8 returns the environment variable value as an int8. Returns
// an error for missing env var or empty value.
func GetInt8(name string) (int8, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(i, 10, 8)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return int8(v), nil
}

// GetInt16 returns the environment variable value as an int16. Returns
// an error for missing env var or empty value.
func GetInt16(name string) (int16, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(i, 10, 16)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return int16(v), nil
}

// GetInt32 returns the environment variable value as an int32. Returns
// an error for missing env var or empty value.
func GetInt32(name string) (int32, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(i, 10, 32)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return int32(v), nil
}

// GetInt64 returns the environment variable value as an int64. Returns
// an error for missing env var or empty value.
func GetInt64(name string) (int64, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return v, nil
}

// GetUint returns the environment variable value as an uint. Returns
// an error for missing env var or empty value.
func GetUint(name string) (uint, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.Atoi(i)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return uint(v), nil
}

// GetUint8 returns the environment variable value as an uint8. Returns
// an error for missing env var or empty value.
func GetUint8(name string) (uint8, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseUint(i, 10, 8)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return uint8(v), nil
}

// GetUint16 returns the environment variable value as an uint16. Returns
// an error for missing env var or empty value.
func GetUint16(name string) (uint16, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseUint(i, 10, 16)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return uint16(v), nil
}

// GetUint32 returns the environment variable value as an uint32. Returns
// an error for missing env var or empty value.
func GetUint32(name string) (uint32, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseUint(i, 10, 32)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return uint32(v), nil
}

// GetUint64 returns the environment variable value as an uint64. Returns
// an error for missing env var or empty value.
func GetUint64(name string) (uint64, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseUint(i, 10, 64)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return uint64(v), nil
}

// GetFloat32 returns the environment variable value as a float32. Returns
// an error for missing env var or empty value.
func GetFloat32(name string) (float32, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseFloat(i, 32)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return float32(v), nil
}

// GetFloat64 returns the environment variable value as a float64. Returns
// an error for missing env var or empty value.
func GetFloat64(name string) (float64, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseFloat(i, 64)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return float64(v), nil
}

// GetComplex64 returns the environment variable value as a complex64 type. Returns
// an error for missing env var or empty value.
func GetComplex64(name string) (complex64, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseComplex(i, 64)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return complex64(v), nil
}

// GetComplex128 returns the environment variable value as a complex128 type. Returns
// an error for missing env var or empty value.
func GetComplex128(name string) (complex128, error) {
	i, err := lookup(name)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseComplex(i, 128)
	if err != nil {
		return 0, NewError(InvalidType, name, err.Error())
	}

	return complex128(v), nil
}

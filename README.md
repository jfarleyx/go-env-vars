# go-env-vars
[![GoDoc](https://godoc.org/github.com/jfarleyx/go-env-vars?status.svg)](http://godoc.org/github.com/jfarleyx/go-env-vars)
[![Go Report](https://goreportcard.com/badge/github.com/jfarleyx/go-env-vars)](https://goreportcard.com/report/github.com/jfarleyx/go-env-vars)

Go module to fetch environment variables by desired Go data type. The module can return env vars as the following data types:

- string
- bool
- v4 UUID
- Rune
- []Byte
- time.Duration
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64
- float32, float64
- complex64, complex128

The module can return env vars containing comma delimited lists of data as the following types:

- []string
- []v4 UUID
- []time.Duration
- []int, []int8, []int16, []int32, []int64
- []uint, []uint8, []uint16, []uint32, []uint64
- []float32, []float64
- []complex64, []complex128

Last, but not least, env vars containing JSON can be returned as:

- json.RawMessage
- User defined struct

## Requirements

Tested using Go version 1.15

## Install

``` go get github.com/jfarleyx/go-env-vars ```

## Usage
Fetching environment variables and converting them to the appropriate data type can be a tedious task when there are a lot of them. With go-env-vars, it can be much easier and faster to get your env vars as the data type you want. 

```
import (
    "github.com/jfarleyx/go-env-vars"
)

...

// Set an environment variable
env.SetVar("TEST_BOOL", "true")

// Get env var as a boolean
v, err := env.GetBool("TEST_BOOL")
if err != nil {    
    fmt.Printf("error type: %s; error message: %s", err.Type, err.Err.Error())
    // evaluate the error type to determine root cause (e.g. missing env var, empty env var, invalid type)
    if err.Type == env.EmptyValue {
        ...
    }
}

////////////////////////////////////////////////////////////////////////////////
env.SetVar("TEST_UUID", "d1b0eaca-3b7c-4fef-88fe-75c9db4bf90a")

v, err := env.GetUUID("TEST_UUID")
if err != nil {
    ...
}

////////////////////////////////////////////////////////////////////////////////
env.SetVar("TEST_DURATION", "3h")

v, err := env.GetDuration("TEST_DURATION")
if err != nil {
    ...
}

////////////////////////////////////////////////////////////////////////////////
// Convert comma delimited values to slices of native Go types
env.SetVar("TEST_SLICE_INT", "1,2,3")

v, err := env.GetIntSlice("TEST_SLICE_INT")
if err != nil {
    ...
}

////////////////////////////////////////////////////////////////////////////////
// Convert JSON formatted env variable values
env.SetVar("TEST_JSON_RAW", "{\"key\":\"my value\"}")

// Get JSON data as json.RawMessage
v, err := env.GetJSONRaw("TEST_JSON_RAW")
if err != nil {
    ...
}

// Or unmarshal the data to a struct
type Test struct {
    Key string `json:"key"`
}

s := Test{}
err := env.Unmarshal("TEST_JSON_RAW", &s)
if err != nil {
    ...
}
```

### Errors
The go-env-vars module methods return a custom error struct that defines the type of error as one of the following:


```
// ErrorType defines the type of error the env module returns.
type ErrorType string

const (
	// NotFound ErrorType indicates the requested env variable was not found.
	NotFound ErrorType = "NotFound"
	// EmptyValue ErrorType indicates the requested env variable's value was an empty string.
	EmptyValue ErrorType = "EmptyValue"
	// InvalidType ErrorType indicates the requested env variable's value was not of the requested data type.
	InvalidType ErrorType = "InvalidType"
)

// Error wraps a Go error and defines an ErrorType.
type Error struct {
	// Type is the ErrorType assigned to this Error.
	Type ErrorType
	// Msg is the message assigned to this Error.
	Err error
}
```

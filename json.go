package env

import "encoding/json"

// GetJSONRaw returns the environment variable's JSON formatted string as a json.RawMessage type. Returns
// an error for missing env var or empty value.
func GetJSONRaw(name string) (json.RawMessage, error) {
	i, err := lookup(name)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(i), nil
}

// Unmarshal unmarshals the JSON formatted env variable value to v.
func Unmarshal(name string, v interface{}) error {
	i, err := lookup(name)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(i), &v)
	if err != nil {
		return NewError(InvalidType, name, err.Error())
	}

	return nil
}

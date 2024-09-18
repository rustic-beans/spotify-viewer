package graphql

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

// StringMap represents a map[string]string type for GraphQL
type StringMap map[string]string

// MarshalStringMap marshals the map[string]string to JSON
func MarshalStringMap(m map[string]string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		json.NewEncoder(w).Encode(m)
	})
}

// UnmarshalStringMap unmarshals JSON to a map[string]string
func UnmarshalStringMap(v interface{}) (map[string]string, error) {
	switch v := v.(type) {
	case map[string]string:
		return v, nil
	case map[string]interface{}:
		result := make(map[string]string)

		for k, val := range v {
			str, ok := val.(string)
			if !ok {
				return nil, fmt.Errorf("value for key %s is not a string", k)
			}

			result[k] = str
		}

		return result, nil
	case string:
		var result map[string]string
		err := json.Unmarshal([]byte(v), &result)

		return result, err
	default:
		return nil, fmt.Errorf("cannot unmarshal %T into map[string]string", v)
	}
}

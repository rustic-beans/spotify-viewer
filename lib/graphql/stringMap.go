package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

// StringMap represents a map[string]string type for GraphQL
type StringMap map[string]string

func (s StringMap) MarshalGQLContext(_ context.Context, w io.Writer) error {
	if err := json.NewEncoder(w).Encode(s); err != nil {
		return err
	}

	return nil
}

func (s *StringMap) UnmarshalGQLContext(_ context.Context, v interface{}) error {
	switch v := v.(type) {
	case map[string]string:
		*s = v
	case map[string]interface{}:
		result := make(map[string]string)

		for k, val := range v {
			str, ok := val.(string)
			if !ok {
				return fmt.Errorf("value for key %s is not a string", k)
			}

			result[k] = str
		}

		*s = result
	case string:
		if err := json.Unmarshal([]byte(v), &s); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot unmarshal %T into map[string]string", v)
	}

	return nil
}

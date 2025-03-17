package models

import (
	"context"
	"encoding/json"
	"io"

	"github.com/cockroachdb/errors"
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
				return errors.Newf("value for key %s is not a string", k)
			}

			result[k] = str
		}

		*s = result
	case string:
		if err := json.Unmarshal([]byte(v), &s); err != nil {
			return errors.Wrapf(err, "failed unmarshalling %s into map[string]string", v)
		}
	default:
		return errors.Newf("cannot unmarshal %T into map[string]string", v)
	}

	return nil
}

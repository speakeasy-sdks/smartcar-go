// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SecurityActionActionEnum string

const (
	SecurityActionActionEnumLock   SecurityActionActionEnum = "LOCK"
	SecurityActionActionEnumUnlock SecurityActionActionEnum = "UNLOCK"
)

func (e SecurityActionActionEnum) ToPointer() *SecurityActionActionEnum {
	return &e
}

func (e *SecurityActionActionEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LOCK":
		fallthrough
	case "UNLOCK":
		*e = SecurityActionActionEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SecurityActionActionEnum: %v", v)
	}
}

type SecurityAction struct {
	Action *SecurityActionActionEnum `json:"action,omitempty"`
}

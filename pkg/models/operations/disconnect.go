// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DisconnectRequest struct {
	VehicleID *string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

// DisconnectStatusEnum - Revoke application access
type DisconnectStatusEnum string

const (
	DisconnectStatusEnumSuccess DisconnectStatusEnum = "success"
)

func (e DisconnectStatusEnum) ToPointer() *DisconnectStatusEnum {
	return &e
}

func (e *DisconnectStatusEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "success":
		*e = DisconnectStatusEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DisconnectStatusEnum: %s", s)
	}
}

type DisconnectResponse struct {
	ContentType string
	// Revoke application access
	Status      *DisconnectStatusEnum
	StatusCode  int
	RawResponse *http.Response
}
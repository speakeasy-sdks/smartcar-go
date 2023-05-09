// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CompassDirectionEnum - The direction the vehicle is traveling.
type CompassDirectionEnum string

const (
	CompassDirectionEnumN  CompassDirectionEnum = "N"
	CompassDirectionEnumNe CompassDirectionEnum = "NE"
	CompassDirectionEnumE  CompassDirectionEnum = "E"
	CompassDirectionEnumSe CompassDirectionEnum = "SE"
	CompassDirectionEnumS  CompassDirectionEnum = "S"
	CompassDirectionEnumSw CompassDirectionEnum = "SW"
	CompassDirectionEnumW  CompassDirectionEnum = "W"
	CompassDirectionEnumNw CompassDirectionEnum = "NW"
)

func (e CompassDirectionEnum) ToPointer() *CompassDirectionEnum {
	return &e
}

func (e *CompassDirectionEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "N":
		fallthrough
	case "NE":
		fallthrough
	case "E":
		fallthrough
	case "SE":
		fallthrough
	case "S":
		fallthrough
	case "SW":
		fallthrough
	case "W":
		fallthrough
	case "NW":
		*e = CompassDirectionEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CompassDirectionEnum: %v", v)
	}
}

// Compass - returns the compass heading of a Tesla.
type Compass struct {
	// The direction the vehicle is traveling.
	Direction *CompassDirectionEnum `json:"direction,omitempty"`
	// The direction the vehicle is traveling (in degrees).
	Heading *float32 `json:"heading,omitempty"`
}
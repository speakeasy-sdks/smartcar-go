// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/shared"
	"net/http"
)

type GetBatteryLevelRequest struct {
	VehicleID string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetBatteryLevelResponse struct {
	// return EV Battery Level reading
	BatteryLevel *shared.BatteryLevel
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/shared"
	"net/http"
)

type SetChargingLimitRequest struct {
	VehicleID   string              `pathParam:"style=simple,explode=false,name=vehicle_id"`
	ChargeLimit *shared.ChargeLimit `request:"mediaType=application/json"`
}

type SetChargingLimitResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// return Success Response
	SuccessResponse *shared.SuccessResponse
}

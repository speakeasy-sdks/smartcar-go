// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/shared"
	"net/http"
)

type GetEngineOilRequest struct {
	VehicleID *string `pathParam:"style=simple,explode=false,name=vehicle_id"`
}

type GetEngineOilResponse struct {
	ContentType string
	// return engine oil reading
	EngineOil   *shared.EngineOil
	StatusCode  int
	RawResponse *http.Response
}

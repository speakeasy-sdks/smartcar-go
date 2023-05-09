// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/shared"
	"net/http"
)

type GetTeslaExteriorTemperatureRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetTeslaExteriorTemperatureResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// returns the exterior temperature of a Tesla.
	Temperature *shared.Temperature
}
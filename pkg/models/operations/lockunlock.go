// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/shared"
	"net/http"
)

type LockUnlockRequest struct {
	VehicleID      string                 `pathParam:"style=simple,explode=false,name=vehicle_id"`
	SecurityAction *shared.SecurityAction `request:"mediaType=application/json"`
}

type LockUnlockResponse struct {
	ContentType string
	// return Compatibility
	SecurityResponse *shared.SecurityResponse
	StatusCode       int
	RawResponse      *http.Response
}

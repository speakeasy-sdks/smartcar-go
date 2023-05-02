// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package smartcar

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/operations"
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/shared"
	"github.com/speakeasy-sdks/smartcar-go/pkg/utils"
	"net/http"
)

// evs - Operations about electric vehicles
type evs struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newEvs(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *evs {
	return &evs{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetBatteryCapacity - EV Battery Capacity
// __Description__
//
// Returns the total capacity of an electric vehicle's battery.
//
// __Permission__
//
// `read_battery`
//
// __Response body__
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  capacity|   number|  The total capacity of the vehicle's battery (in kilowatt-hours). 	|

func (s *evs) GetBatteryCapacity(ctx context.Context, vehicleID string) (*operations.GetBatteryCapacityResponse, error) {
	request := operations.GetBatteryCapacityRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/battery/capacity", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetBatteryCapacityResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.BatteryCapacity
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.BatteryCapacity = out
		}
	}

	return res, nil
}

// GetBatteryLevel - EV Battery Level
// __Description__
//
// Retrieve EV battery level of a vehicle.
//
// __Permission__
//
// `read_battery`
//
// __Response body__
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  `percentRemaining`|   number|  An EV battery’s state of charge (in percent). 	|
// |   `range`|   number	|   The estimated remaining distance the vehicle can travel (in kilometers by default or in miles using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers).	|

func (s *evs) GetBatteryLevel(ctx context.Context, vehicleID string) (*operations.GetBatteryLevelResponse, error) {
	request := operations.GetBatteryLevelRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/battery", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetBatteryLevelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.BatteryLevel
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.BatteryLevel = out
		}
	}

	return res, nil
}

// GetChargingLimit - EV Charging Limit
// __Description__
//
// Returns the current charge limit of an electric vehicle.

func (s *evs) GetChargingLimit(ctx context.Context, vehicleID string) (*operations.GetChargingLimitResponse, error) {
	request := operations.GetChargingLimitRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/charge/limit", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetChargingLimitResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ChargeLimit
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ChargeLimit = out
		}
	}

	return res, nil
}

// GetChargingStatus - EV Charging Status
// __Description__
//
// Returns the current charge status of an electric vehicle.
//
// __Permission__
//
// `read_charge`
//
// __Response body__
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  `isPluggedIn` 	|   boolean	|  Indicates whether a charging cable is currently plugged into the vehicle’s charge port. 	|
// |   `state`	|   string	|   Indicates whether the vehicle is currently charging. Options: `CHARGING` `FULLY_CHARGED` `NOT_CHARGING`	|

func (s *evs) GetChargingStatus(ctx context.Context, vehicleID string) (*operations.GetChargingStatusResponse, error) {
	request := operations.GetChargingStatusRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/charge", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetChargingStatusResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ChargeStatus
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ChargeStatus = out
		}
	}

	return res, nil
}

// SetChargingLimit - Set EV Charging Limit
// __Description__
//
// Returns the current charge limit of an electric vehicle.

func (s *evs) SetChargingLimit(ctx context.Context, vehicleID string, chargeLimit *shared.ChargeLimit) (*operations.SetChargingLimitResponse, error) {
	request := operations.SetChargingLimitRequest{
		VehicleID:   vehicleID,
		ChargeLimit: chargeLimit,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/charge/limit", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "ChargeLimit", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.SetChargingLimitResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SuccessResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.SuccessResponse = out
		}
	}

	return res, nil
}

// StartStopCharge - Start or stop charging an electric vehicle. Please contact us to request early access.
// __Description__
//
// Returns the current charge status of an electric vehicle.
//
// __Permission__
//
// `read_charge`
//
// __Response body__
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  `isPluggedIn` 	|   boolean	|  Indicates whether a charging cable is currently plugged into the vehicle’s charge port. 	|
// |   `state`	|   string	|   Indicates whether the vehicle is currently charging. Options: `CHARGING` `FULLY_CHARGED` `NOT_CHARGING`	|

func (s *evs) StartStopCharge(ctx context.Context, vehicleID string, chargeAction *shared.ChargeAction) (*operations.StartStopChargeResponse, error) {
	request := operations.StartStopChargeRequest{
		VehicleID:    vehicleID,
		ChargeAction: chargeAction,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/charge", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "ChargeAction", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.StartStopChargeResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SuccessResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.SuccessResponse = out
		}
	}

	return res, nil
}

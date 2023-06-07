// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package smartcar

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/operations"
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/shared"
	"github.com/speakeasy-sdks/smartcar-go/pkg/utils"
	"io"
	"net/http"
)

type tesla struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newTesla(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *tesla {
	return &tesla{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// GetAmmeter - Retrieve charging ammeter time for a Tesla.
// __Description__
//
// When the vehicle is plugged in, this endpoint returns the amperage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.
func (s *tesla) GetAmmeter(ctx context.Context, vehicleID string) (*operations.GetTeslaAmmeterResponse, error) {
	request := operations.GetTeslaAmmeterRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/tesla/charge/ammeter", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTeslaAmmeterResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ChargeAmmeter
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ChargeAmmeter = out
		}
	}

	return res, nil
}

// GetChargeTime - Retrieve charging completion time for a Tesla.
// __Description__
//
// When the vehicle is charging, this endpoint returns the date and time the vehicle expects to complete this charging session. When the vehicle is not charging, this endpoint results in a vehicle state error.
func (s *tesla) GetChargeTime(ctx context.Context, vehicleID string) (*operations.GetTeslaChargeTimeResponse, error) {
	request := operations.GetTeslaChargeTimeRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/tesla/charge/completion", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTeslaChargeTimeResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ChargeTime
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ChargeTime = out
		}
	}

	return res, nil
}

// GetCompass - Retrieve compass heading for a Tesla.
// __Description__
//
// This endpoint returns the compass heading of a Tesla. The value is in degrees, with 0 degrees being North.
func (s *tesla) GetCompass(ctx context.Context, vehicleID string) (*operations.GetTeslaCompassResponse, error) {
	request := operations.GetTeslaCompassRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/tesla/compass", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTeslaCompassResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Compass
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Compass = out
		}
	}

	return res, nil
}

// GetExteriorTemperature - Retrieve exterior temperature for a Tesla.
// __Description__
//
// This endpoint returns the exterior temperature of a Tesla, in celsius by default.
func (s *tesla) GetExteriorTemperature(ctx context.Context, id string) (*operations.GetTeslaExteriorTemperatureResponse, error) {
	request := operations.GetTeslaExteriorTemperatureRequest{
		ID: id,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{id}/tesla/thermometer/exterior", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTeslaExteriorTemperatureResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Temperature
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Temperature = out
		}
	}

	return res, nil
}

// GetInteriorTemperature - Retrieve interior temperature for a Tesla.
// __Description__
//
// This endpoint returns the interior temperature of a Tesla, in celsius by default.
func (s *tesla) GetInteriorTemperature(ctx context.Context, id string) (*operations.GetTeslaInteriorTemperatureResponse, error) {
	request := operations.GetTeslaInteriorTemperatureRequest{
		ID: id,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{id}/tesla/thermometer/interior", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTeslaInteriorTemperatureResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Temperature
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Temperature = out
		}
	}

	return res, nil
}

// GetSpeedometer - Retrieve speed for a Tesla.
// __Description__
//
// This endpoint returns the speed of a Tesla (in kilometers/hour by default or in miles/hour using the sc-unit-system).
func (s *tesla) GetSpeedometer(ctx context.Context, vehicleID string) (*operations.GetTeslaSpeedResponse, error) {
	request := operations.GetTeslaSpeedRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/tesla/speedometer", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTeslaSpeedResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Speed
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.Speed = out
		}
	}

	return res, nil
}

// GetVoltage - Retrieve charging voltmeter time for a Tesla.
// __Description__
//
// When the vehicle is plugged in, this endpoint returns the voltage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.
func (s *tesla) GetVoltage(ctx context.Context, vehicleID string) (*operations.GetTeslaVoltageResponse, error) {
	request := operations.GetTeslaVoltageRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/tesla/charge/voltmeter", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTeslaVoltageResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ChargeVoltage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ChargeVoltage = out
		}
	}

	return res, nil
}

// GetWattmeter - Retrieve charging wattmeter time for a Tesla.
// __Description__
//
// When the vehicle is plugged in, this endpoint returns the wattage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.
func (s *tesla) GetWattmeter(ctx context.Context, vehicleID string) (*operations.GetTeslaWattmeterResponse, error) {
	request := operations.GetTeslaWattmeterRequest{
		VehicleID: vehicleID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/tesla/charge/wattmeter", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetTeslaWattmeterResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ChargeWattage
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.ChargeWattage = out
		}
	}

	return res, nil
}

// SetAmmeter - Set charging ammeter time for a Tesla.
// __Description__
//
// When the vehicle is plugged in, this endpoint sets the amperage of the charger measured by the vehicle. When the vehicle is not plugged in, this endpoint results in a vehicle state error.
func (s *tesla) SetAmmeter(ctx context.Context, vehicleID string, chargeAmmeter *shared.ChargeAmmeter) (*operations.SetTeslaAmmeterResponse, error) {
	request := operations.SetTeslaAmmeterRequest{
		VehicleID:     vehicleID,
		ChargeAmmeter: chargeAmmeter,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/tesla/charge/ammeter", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "ChargeAmmeter", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.SetTeslaAmmeterResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SuccessResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.SuccessResponse = out
		}
	}

	return res, nil
}

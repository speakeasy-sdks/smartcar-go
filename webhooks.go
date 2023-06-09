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

type webhooks struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newWebhooks(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *webhooks {
	return &webhooks{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// Subscribe - Subscribe Webhook
// __Description__
//
// Subscribe to a webhook for a vehicle.
//
// __Permission__
//
// `required: webhook:read`
//
// __Response body__
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|
func (s *webhooks) Subscribe(ctx context.Context, vehicleID string, webhookID string, webhookInfo *shared.WebhookInfo) (*operations.SubscribeResponse, error) {
	request := operations.SubscribeRequest{
		VehicleID:   vehicleID,
		WebhookID:   webhookID,
		WebhookInfo: webhookInfo,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/webhooks/{webhookId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "WebhookInfo", "json")
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

	res := &operations.SubscribeResponse{
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

// Unsubscribe - Unsubscribe Webhook
// __Description__
//
// Delete a webhook for a vehicle.
//
// __Permission__
//
// `required: webhook:read`
//
// __Response body__
//
// |  Name 	|Type   	|Boolean   	|
// |---	|---	|---	|
// |  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|
func (s *webhooks) Unsubscribe(ctx context.Context, vehicleID string, webhookID string) (*operations.UnsubscribeResponse, error) {
	request := operations.UnsubscribeRequest{
		VehicleID: vehicleID,
		WebhookID: webhookID,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/vehicles/{vehicle_id}/webhooks/{webhookId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
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

	res := &operations.UnsubscribeResponse{
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

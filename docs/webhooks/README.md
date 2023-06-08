# Webhooks

### Available Operations

* [Subscribe](#subscribe) - Subscribe Webhook
* [Unsubscribe](#unsubscribe) - Unsubscribe Webhook

## Subscribe

__Description__

Subscribe to a webhook for a vehicle.

__Permission__

`required: webhook:read`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/smartcar-go"
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/operations"
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/shared"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.Subscribe(ctx, "repellendus", "sapiente", &shared.WebhookInfo{
        Vehicleid: smartcar.String("dc6ea99e-57d1-4e41-b129-27e7eb58713e"),
        Webhookid: smartcar.String("9b6ae692-60cc-4b3e-89d8-71e7549cf805"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

## Unsubscribe

__Description__

Delete a webhook for a vehicle.

__Permission__

`required: webhook:read`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  status|   string|  If the request is successful, Smartcar will return “success” (HTTP 200 status).|

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/smartcar-go"
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/operations"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.Unsubscribe(ctx, "quo", "odit")
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/smartcar-go"
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
        sdk.WithVehicleID(sdk.String("36ab27d0-fd9d-4455-823a-ce30af709ffc")),
    )

    ctx := context.Background()    
    req := operations.GetLocationRequest{}

    res, err := s.Vehicles.GetLocation(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Location != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->
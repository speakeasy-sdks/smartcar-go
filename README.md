# github.com/speakeasy-sdks/smartcar-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/smartcar-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Compatibility](docs/compatibility/README.md)

* [ListCompatibility](docs/compatibility/README.md#listcompatibility) - Compatibility

### [Evs](docs/evs/README.md)

* [GetBatteryCapacity](docs/evs/README.md#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](docs/evs/README.md#getbatterylevel) - EV Battery Level
* [GetChargingStatus](docs/evs/README.md#getchargingstatus) - EV Charging Status

### [Vehicles](docs/vehicles/README.md)

* [Disconnect](docs/vehicles/README.md#disconnect) - Revoke Access
* [Get](docs/vehicles/README.md#get) - Vehicle Info
* [GetEngineOil](docs/vehicles/README.md#getengineoil) - Engine Oil Life
* [GetFuelTank](docs/vehicles/README.md#getfueltank) - Fuel Tank (US Only)
* [GetLocation](docs/vehicles/README.md#getlocation) - Location
* [GetOdometer](docs/vehicles/README.md#getodometer) - Odometer
* [GetPermissions](docs/vehicles/README.md#getpermissions) - Application Permissions
* [GetTirePressure](docs/vehicles/README.md#gettirepressure) - Tire pressure
* [ListVehicles](docs/vehicles/README.md#listvehicles) - All Vehicles
* [LockUnlock](docs/vehicles/README.md#lockunlock) - Unlock Vehicle
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)

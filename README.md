<div align="center">
    <img src="https://user-images.githubusercontent.com/6267663/232771888-a65b182b-9ae7-42f3-9bbe-85658a61b9e3.svg" width="350px">
    <h1>Go SDK</h1>
   <p>Build and scale your mobility business</p>
   <a href="https://smartcar.com/docs/api/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/smartcar-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/smartcar-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/smartcar-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/smartcar-go?sort=semver&style=for-the-badge" /></a>
</div>

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
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Vehicles.GetLocation(ctx, "36ab27d0-fd9d-4455-823a-ce30af709ffc")
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

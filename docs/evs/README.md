# Evs

## Overview

Operations about electric vehicles

### Available Operations

* [GetBatteryCapacity](#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](#getbatterylevel) - EV Battery Level
* [GetChargingStatus](#getchargingstatus) - EV Charging Status

## GetBatteryCapacity

__Description__

Returns the total capacity of an electric vehicle's battery.

__Permission__

`read_battery`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  capacity|   number|  The total capacity of the vehicle's battery (in kilowatt-hours). 	|

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
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
        sdk.WithVehicleID(sdk.String("36ab27d0-fd9d-4455-823a-ce30af709ffc")),
    )

    ctx := context.Background()    
    req := operations.GetBatteryCapacityRequest{}

    res, err := s.Evs.GetBatteryCapacity(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BatteryCapacity != nil {
        // handle response
    }
}
```

## GetBatteryLevel

__Description__

Retrieve EV battery level of a vehicle.

__Permission__

`read_battery`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  `percentRemaining`|   number|  An EV battery’s state of charge (in percent). 	|
|   `range`|   number	|   The estimated remaining distance the vehicle can travel (in kilometers by default or in miles using the [sc-unit-system](https://smartcar.com/docs/api?version=v2.0&language=cURL#request-headers).	|

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
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
        sdk.WithVehicleID(sdk.String("36ab27d0-fd9d-4455-823a-ce30af709ffc")),
    )

    ctx := context.Background()    
    req := operations.GetBatteryLevelRequest{}

    res, err := s.Evs.GetBatteryLevel(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BatteryLevel != nil {
        // handle response
    }
}
```

## GetChargingStatus

__Description__

Returns the current charge status of an electric vehicle.

__Permission__

`read_charge`

__Response body__

|  Name 	|Type   	|Boolean   	|
|---	|---	|---	|
|  `isPluggedIn` 	|   boolean	|  Indicates whether a charging cable is currently plugged into the vehicle’s charge port. 	|
|   `state`	|   string	|   Indicates whether the vehicle is currently charging. Options: `CHARGING` `FULLY_CHARGED` `NOT_CHARGING`	|

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
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
        sdk.WithVehicleID(sdk.String("36ab27d0-fd9d-4455-823a-ce30af709ffc")),
    )

    ctx := context.Background()    
    req := operations.GetChargingStatusRequest{}

    res, err := s.Evs.GetChargingStatus(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeStatus != nil {
        // handle response
    }
}
```

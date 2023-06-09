# Evs

## Overview

Operations about electric vehicles

### Available Operations

* [GetBatteryCapacity](#getbatterycapacity) - EV Battery Capacity
* [GetBatteryLevel](#getbatterylevel) - EV Battery Level
* [GetChargingLimit](#getcharginglimit) - EV Charging Limit
* [GetChargingStatus](#getchargingstatus) - EV Charging Status
* [SetChargingLimit](#setcharginglimit) - Set EV Charging Limit
* [StartStopCharge](#startstopcharge) - Start or stop charging an electric vehicle. Please contact us to request early access.

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
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.GetBatteryCapacity(ctx, "unde")
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
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.GetBatteryLevel(ctx, "nulla")
    if err != nil {
        log.Fatal(err)
    }

    if res.BatteryLevel != nil {
        // handle response
    }
}
```

## GetChargingLimit

__Description__

Returns the current charge limit of an electric vehicle.

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.GetChargingLimit(ctx, "corrupti")
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeLimit != nil {
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
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.GetChargingStatus(ctx, "illum")
    if err != nil {
        log.Fatal(err)
    }

    if res.ChargeStatus != nil {
        // handle response
    }
}
```

## SetChargingLimit

__Description__

Returns the current charge limit of an electric vehicle.

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.SetChargingLimit(ctx, "vel", &shared.ChargeLimit{
        Limit: smartcar.Float32(1),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

## StartStopCharge

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
	"github.com/speakeasy-sdks/smartcar-go/pkg/models/shared"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Evs.StartStopCharge(ctx, "error", &shared.ChargeAction{
        Action: shared.ChargeActionActionStart.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SuccessResponse != nil {
        // handle response
    }
}
```

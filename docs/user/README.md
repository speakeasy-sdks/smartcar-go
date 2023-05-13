# User

### Available Operations

* [GetInfo](#getinfo) - User Info

## GetInfo

__Description__

Returns the id of the vehicle owner who granted access to your application. This should be used as the static unique identifier for storing the access token and refresh token pair in your database. Note: A single user can own multiple vehicles, and multiple users can own the same vehicle.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/smartcar-go"
)

func main() {
    s := smartcar.New(
        smartcar.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.User.GetInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.UserInfo != nil {
        // handle response
    }
}
```

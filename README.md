# Go logging
This interface presents methods that logger should implement to be
compatible with go components. This is also intended
to be actually imported package in business logic. For detailed
documentation use go doc

## Structured Logging
This interface is mainly intended for structured logging to be compatible  with `FEA`.

## Implementing logging interface
Implementation should implement all functions of `Logger` interface. It should also
contain `init` function, which will register backend to `Logger`. Example
```
func init() {
	Register("test", NewTestLogger)
}
```
Where `NewTestLogger` is function that creates new instance of logger.

## How to use
### Setting specific implemenation into use
Following lines take default logger implemenation into use.
```
import (
	_ "github.com/astota/go-logger"
	"github.com/astota/go-logging"
)

	...
	logging.UseLogger("default-logger")
	...
```

### Using logger in code
Creating completely new logger:
```
	logger := logging.NewLogger()
```

Getting already created logger from `context.Context`. That logger contains
example fields related to request:
```
	func someFucntion(ctx context.Context, some, other, parameters string) {
	...
	logger := logging.GetLogger(ctx)
```

Adding fields to already existing logger:
```
	logger = logger.AddFields(
		logging.Fields{
			"some_field": "some_value",
		}
	)
```

Log lines using existing logger:
```
	orderValue = 10.0
	logger.Infof("Order values is %f", orderValue)
```

Chaining commands:
```
	orderValue = 10.0

	logging.GetLogger(ctx).AddFields(
		logging.Fields{
			"order_value": orderValue,
		}
	).Info("order created")
```
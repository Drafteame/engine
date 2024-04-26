# engine

Decoupled and generic lambda event manager that helps to build lambda functions in a more structured way.

## Installation

```bash
go get github.com/Drafteame/engine
```

## Usage

### Plain lambda

```go
package main

import (
	"context"
	"fmt"
	
	"github.com/Drafteame/engine"
	"github.com/Drafteame/engine/decorators"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, req Request) (Response, error) {
	fmt.Println("Hello, world!")
	return Response{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	engine.New(handler).
		Use(decorators.PanicRecover[Request, Response]()).
		Run()
}
```

### SQS lambda

```go
package main

import (
    "context"
    "fmt"
	
    "github.com/Drafteame/engine"
    "github.com/Drafteame/engine/decorators"
	
    "github.com/aws/aws-lambda-go/events"
)

func handler(ctx context.Context, event events.SQSEvent) (events.SQSEventResponse, error) {
    for _, record := range event.Records {
        fmt.Println(record.Body)
    }
    return events.SQSEventResponse{}, nil
}

func main() {
    engine.New(handler).
        Use(decorators.PanicRecover[events.SQSEvent, events.SQSEventResponse]()).
        Run()
}
```

### API gateway V1 lambda

```go
package main

import (
	"fmt"
	"net/http"
	
	"github.com/Drafteame/engine"
	"github.com/Drafteame/engine/decorators"
	"github.com/Drafteame/engine/handlers/apigatewayv1"
)

var s *http.ServerMux

func init() {
	s = http.NewServeMux()
	s.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
}

func main() {
	engine.New(apigatewayv1.NewHandler(s)).
		Use(decorators.PanicRecover[apigatewayv1.HTTPRequest, apigatewayv1.HTTPResponse]()).
		Run()
}
```

### API gateway V2 lambda

```go
package main

import (
	"fmt"
	"net/http"
    
    "github.com/Drafteame/engine"
    "github.com/Drafteame/engine/decorators"
    "github.com/Drafteame/engine/handlers/apigatewayv2"
)

var s *http.ServerMux

func init() {
	s = http.NewServeMux()
	s.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
}

func main() {
    engine.New(apigatewayv2.NewHandler(s)).
        Use(decorators.PanicRecover[apigatewayv2.HTTPRequest, apigatewayv2.HTTPResponse]()).
        Run()
}
```

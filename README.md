## GOTH STACK - Go, Templ, HTMX 

## Echo
- compatible with http/net
Echo - func (context) error

## Fiber
- not compatible with core http/net since uses fast api
Fiber - func (context) error

Chi

## Gin
- does not return an error so no central error handling
Gin - func(context)

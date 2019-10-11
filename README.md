# ztaylor.me/log

An opinionated log library

### Global scope is removed

Protects client in cases of reference shadowing

### Enhanced log levels, color scheme

Added log level `Trace`, defined *above* level `Error`. This is easier to follow, and doesn't impact other log messages. It would be considered wrong to commit code producing `Trace` logs.

## Change log

### v0.0.4 - current version
- fix log roller

### v0.0.3
- log writing functions now take variadic `...interface{}`

### v0.0.2
- `DefaultFormatter` sorts log data by key name

### v0.0.1
- correct color coding
- change terminal color dependencies
- add `io.Closer` to `log.Service`

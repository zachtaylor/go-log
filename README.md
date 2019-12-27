# ztaylor.me/log

An opinionated log library

### Global scope is removed

Protects client in cases of reference shadowing

### Enhanced log levels, color scheme

Added log level `Trace`, defined *above* level `Error`. This is easier to follow, and doesn't impact other log messages.

## Change log

### v0.0.5 2019-12-27
- add `Entry.Source`, updated formatter
- add `cmd/ztaylor.me.log.test/main`

### v0.0.4 2019-10-11
- fix log roller

### v0.0.3 2019-09-12
- log writing functions now take variadic `...interface{}`

### v0.0.2 2019-08-07
- `DefaultFormatter` sorts log data by key name

### v0.0.1 2019-4-18
- correct color coding
- change terminal color dependencies
- add `io.Closer` to `log.Service`

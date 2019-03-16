# ztaylor.me/log

An opinionated log library

### Global scope is removed

Protects client in cases of reference shadowing

### Enhanced log levels, color scheme

Added log level `Trace`, defined *above* level `Error`. This is easier to follow, and doesn't impact other log messages. It would be considered wrong to commit code producing `Trace` logs.

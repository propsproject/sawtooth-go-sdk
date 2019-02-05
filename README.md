## Changes

Changes to the logger so it's output is compatible with loggly

* Changed the default logger to https://github.com/uber-go/zap
* JSON logging by default
* 'ts' changed to 'timestamp' in ISO8601
* added a function called SetDefaultKeyValues to be used in logs output

```
var logger *logging.Logger = logging.Get()
logger.SetDefaultKeyValues(zap.String("app", "myapp"), zap.String("name", "microservicename"), zap.String("env", "development"))
```

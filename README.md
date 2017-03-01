# Srslog Hooks for Logrus <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:"/>

Near drop-in replacement for the `logrus_syslog` hook to utilize `sysrlog` instead of built-in `log/syslog`.

Why?
1. `log/syslog` does not compile for most operating systems
2. `sysrlog` support TCP+TLS (but not client-certificate)
3. No need for complexity to just deliver log entries

## Usage

```go
import (
    syslog "github.com/RackSec/srslog"
    "github.com/Sirupsen/logrus"
    logrus_srslog "github.com/grutz/logrus_srslog"
)

func main() {
    log       := logrus.New()
    hook, err := logrus_srslog.NewSrslogHook("udp", "localhost:514", syslog.LOG_INFO, "", "")

    if err == nil {
        log.Hooks.Add(hook)
    }
}
```

When using TCP+TLS, include the certificate file:

```go
import (
    syslog "github.com/RackSec/srslog"
    "github.com/Sirupsen/logrus"
    logrus_srslog "github.com/grutz/logrus_srslog"
)

func main() {
    log       := logrus.New()
    hook, err := logrus_srslog.NewSrslogHook("tcp+tls", "localhost:1999", syslog.LOG_INFO, "", "certificate.pem")

    if err == nil {
        log.Hooks.Add(hook)
    }
}
```
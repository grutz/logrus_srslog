package logrus_srslog

import (
	"fmt"
	syslog "github.com/RackSec/srslog"
	"github.com/sirupsen/logrus"
	"os"
)

// SrslogHook to send logs via srslog.
type SrslogHook struct {
	Writer        *syslog.Writer
	SyslogNetwork string
	SyslogRaddr   string
}

// NewSrslogHook creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewSrslogHook("udp", "localhost:514", syslog.LOG_DEBUG, "", "")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewSrslogHook(network, raddr string, priority syslog.Priority, tag string, cert string) (*SrslogHook, error) {
	var (
		w *syslog.Writer
		err error
	)
	if cert != "" {
		network = "tcp+tls"	// When cert is defined the network is always tcp+tls
		w, err = syslog.DialWithTLSCertPath(network, raddr, priority, tag, cert)
	} else {
		w, err = syslog.Dial(network, raddr, priority, tag)
	}

	return &SrslogHook{w, network, raddr}, err
}

// Fire called when Logrus has log entries to send
func (hook *SrslogHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	switch entry.Level {
	case logrus.PanicLevel:
		return hook.Writer.Crit(line)
	case logrus.FatalLevel:
		return hook.Writer.Crit(line)
	case logrus.ErrorLevel:
		return hook.Writer.Err(line)
	case logrus.WarnLevel:
		return hook.Writer.Warning(line)
	case logrus.InfoLevel:
		return hook.Writer.Info(line)
	case logrus.DebugLevel:
		return hook.Writer.Debug(line)
	default:
		return nil
	}
}

// Levels returns all levels as supported by Srslog
func (hook *SrslogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

package zerolog

import (
	"testing"
	"tsf-cron/pkg/core/log"

	"github.com/stretchr/testify/assert"
)

func TestZeroLog(t *testing.T) {
	zerolog := New()
	assert.NotNil(t, zerolog)
	assert.Equal(t, "zerolog", zerolog.String())
	
	log.New(New())
	log.Info("Hello")
	log.Infof("Hello %s", "World")
	log.Debug("Hello")
	log.Debugf("Hello %s", "World")
	log.Warn("Hello")
	log.Warnf("Hello %s", "World")
	log.Error("Hello")
	log.Errorf("Hello %s", "World")
}

func TestMoreMsgLog(t *testing.T) {
	log.New(New())
	log.Info("Hello World", "foo", "bar", "num", 1)
	log.Debug("Hello World", "foo", "bar", "num", 1)
	log.Warn("Hello World", "foo", "bar", "num", 1)
	log.Error("Hello World", "foo", "bar", "num", 1)
}
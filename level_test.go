package golog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelString(t *testing.T) {
	for _, name := range levelNames {
		level, err := LogLevel(name)
		assert.NoError(t, err)
		assert.Equal(t, name, level.String())
	}
}

func TestLogLevel(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		level, err := LogLevel("INFO")
		assert.NoError(t, err)
		assert.Equal(t, INFO, level)
	})

	t.Run("MiscCase", func(t *testing.T) {
		level, err := LogLevel("iNfO")
		assert.NoError(t, err)
		assert.Equal(t, INFO, level)
	})

	t.Run("Invalid", func(t *testing.T) {
		_, err := LogLevel("XXX")
		assert.Error(t, err)
	})
}

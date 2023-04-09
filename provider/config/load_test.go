package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadYaml(t *testing.T) {
	conf, err := Load(filepath.Join("./../../", "conf", "config.yaml"))
	assert.NoError(t, err)
	assert.NotNil(t, conf)

	t.Log(conf)
}

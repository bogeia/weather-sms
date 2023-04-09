package sentence

import (
	"path/filepath"
	"testing"

	"github.com/bogeia/weather-sms/provider/config"
	"github.com/stretchr/testify/assert"
)

func TestGetShaDiaoInfo(t *testing.T) {
	conf, err := config.Load(filepath.Join("./../../", "conf", "config.yaml"))
	assert.NoError(t, err)
	assert.NotNil(t, conf)

	resp, err := GetShaDiaoInfo(conf.Sentence)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	t.Log(resp)
}

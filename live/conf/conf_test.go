package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfNew(t *testing.T) {
	conf, err := New("conf.example.yaml")

	c := &Conf{
		App: App{
			Name:         "api-server",
			Addr:         "9011",
			Debug:        true,
			ReadTimeout:  "10s",
			WriteTimeout: "10s",
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, c, conf)
}

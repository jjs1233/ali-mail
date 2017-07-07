package aliMail

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewGun(t *testing.T) {
	gc := &GunConfig{
		User:     "test",
		Password: "test",
		Host:     "test:24",
	}
	g := NewGun(gc)
	require.NotNil(t, g.Auth)
}

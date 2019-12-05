package client_test

import (
	"testing"

	"github.com/ivost/myvault/pkg/client"
	"github.com/ivost/shared/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	conf := config.New("../config/config.yaml")
	require.NotNil(t, conf)
	c := client.New(conf)
	require.NotNil(t, c)
	rsp, err := c.Health()
	_, _ = rsp, err
}

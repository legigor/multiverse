package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_can_connect_to_kubernetes_cluster(t *testing.T) {

	err := Execute()
	assert.NoError(t, err)
}

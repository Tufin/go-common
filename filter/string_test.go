package filter_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/go-common/filter"
)

func TestEndsWith(t *testing.T) {

	require.True(t, filter.EndsWith("/k8s/runtime.e9f08fdf10bf61f58ac0.js", ".css", ".js"))
}

func TestEndsWith_False(t *testing.T) {

	require.False(t, filter.EndsWith("/k8s/runtime.e9f08fdf10bf61f58ac0.jss", ".css", ".js"))
}

package filter_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/go-common/filter"
)

func TestIsUIStaticFile(t *testing.T) {

	require.True(t, filter.IsUIStaticFile("/k8s/runtime.e9f08fdf10bf61f58ac0.js"))
}

func TestIsUIStaticFile_False(t *testing.T) {

	require.False(t, filter.IsUIStaticFile("/k8s/pods"))
}

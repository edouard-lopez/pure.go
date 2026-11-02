package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_returns_fallback_when_unset(t *testing.T) {
	actual := Get("PURE_TEST_UNSET_VARIABLE", "default")

	assert.Equal(t, "default", actual)
}

func Test_Get_returns_env_value_when_set(t *testing.T) {
	t.Setenv("PURE_TEST_VARIABLE", "custom")

	actual := Get("PURE_TEST_VARIABLE", "default")

	assert.Equal(t, "custom", actual)
}

func Test_Get_returns_empty_value_when_set_to_empty(t *testing.T) {
	t.Setenv("PURE_TEST_VARIABLE", "")

	actual := Get("PURE_TEST_VARIABLE", "default")

	assert.Equal(t, "", actual)
}

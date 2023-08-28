package colorize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Primary_colorize(t *testing.T) {
	label := "primary-text"
	expected := "\x1b[34m" + label + "\x1b[0m"

	actual := Primary("primary-text")

	assert.Equal(t, expected, actual)
}
func Test_Info_colorize(t *testing.T) {
	label := "info-text"
	expected := "\x1b[36m" + label + "\x1b[0m"

	actual := Info("info-text")

	assert.Equal(t, expected, actual)
}
func Test_Mute_colorize(t *testing.T) {
	label := "mute-text"
	expected := "\x1b[90m" + label + "\x1b[0m"

	actual := Mute("mute-text")

	assert.Equal(t, expected, actual)
}
func Test_Success_colorize(t *testing.T) {
	label := "success-text"
	expected := "\x1b[35m" + label + "\x1b[0m"

	actual := Success("success-text")

	assert.Equal(t, expected, actual)
}
func Test_Normal_colorize(t *testing.T) {
	label := "normal-text"
	expected := "\x1b[39m" + label + "\x1b[0m"

	actual := Normal("normal-text")

	assert.Equal(t, expected, actual)
}
func Test_Danger_colorize(t *testing.T) {
	label := "danger-text"
	expected := "\x1b[31m" + label + "\x1b[0m"

	actual := Danger("danger-text")

	assert.Equal(t, expected, actual)
}
func Test_Light_colorize(t *testing.T) {
	label := "light-text"
	expected := "\x1b[37m" + label + "\x1b[0m"

	actual := Light("light-text")

	assert.Equal(t, expected, actual)
}
func Test_Warning_colorize(t *testing.T) {
	label := "warning-text"
	expected := "\x1b[33m" + label + "\x1b[0m"

	actual := Warning("warning-text")

	assert.Equal(t, expected, actual)
}
func Test_Dark_colorize(t *testing.T) {
	label := "dark-text"
	expected := "\x1b[30m" + label + "\x1b[0m"

	actual := Dark("dark-text")

	assert.Equal(t, expected, actual)
}

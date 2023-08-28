package go_version

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var FakeGoVersionOutput = "go version go1.21.0 linux/amd64"

func Test_Parse_Go_version(t *testing.T) {
	output := FakeGoVersionOutput
	expectedVersion := "1.21.0"

	actual := Parse(output)

	assert.Equal(t, expectedVersion, actual)
}

type MockedRuntime struct {
	mock.Mock
}

func (m *MockedRuntime) Output() ([]byte, error) {
	args := m.Called()
	return []byte(args.String(0)), errors.New("Missing go binary")
}

func Test_Get_a_version(t *testing.T) {
	mockRuntime := new(MockedRuntime)
	expectedVersion := "1.21.0"
	mockRuntime.On("Output").Return(FakeGoVersionOutput)
	GetGoVersion = mockRuntime.Output // redefine go_version.GetGoVersion to use the mock

	actual := Get()

	assert.Equal(t, expectedVersion, actual)

	mockRuntime.AssertExpectations(t)
}

func Test_Get_no_version(t *testing.T) {
	mockRuntime := new(MockedRuntime)
	expectedVersion := ""
	mockRuntime.On("Output").Return(expectedVersion)
	GetGoVersion = mockRuntime.Output // redefine go_version.GetGoVersion to use the mock

	actual := Get()

	assert.Equal(t, expectedVersion, actual)

	mockRuntime.AssertExpectations(t)
}

package go_version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedRuntime struct {
	mock.Mock
}

func (m *MockedRuntime) Version() string {
	args := m.Called()
	return args.String(0)
}

func Test_Get(t *testing.T) {
	mockRuntime := new(MockedRuntime)
	expectedVersion := "0.0.1"
	mockRuntime.On("Version").Return(expectedVersion)
	getGoVersion = mockRuntime.Version // redefine getGoVersion to use the mock

	result := Get()

	assert.Equal(t, expectedVersion, result)

	mockRuntime.AssertExpectations(t)
}

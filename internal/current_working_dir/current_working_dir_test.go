package current_working_dir

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_Current_Directory(t *testing.T) {
	err := os.Chdir(`/tmp`)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	actual := Get()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func Test_ReplaceByTilde(t *testing.T) {
	expected := "~"

	actual := ReplaceByTilde("/home/" + os.Getenv("USER"))

	assert.Equal(t, expected, actual)
}

func Test_Get_Home_Directory_Replace_By_Tilde(t *testing.T) {
	err := os.Chdir(os.Getenv("HOME"))
	if err != nil {
		t.Fatal(err)
	}

	current_working_dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	expected := ReplaceByTilde(current_working_dir)

	actual := Get()

	assert.Equal(t, expected, actual)
}

package current_working_dir

import (
	"os"
	"testing"

	"github.com/carlmjohnson/be"
)

func Test_Get_Current_Directory(t *testing.T) {
	err := os.Chdir(`/tmp`)
	if err != nil {
		t.Fatal(err)
	}

	expect, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	actual := Get()
	if err != nil {
		t.Fatal(err)
	}

	be.Equal(t, actual, expect)
}

func Test_ReplaceByTilde(t *testing.T) {
	expect := "~"

	actual := ReplaceByTilde("/home/" + os.Getenv("USER"))

	be.Equal(t, actual, expect)
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
	expect := ReplaceByTilde(current_working_dir)

	actual := Get()

	be.Equal(t, actual, expect)
}

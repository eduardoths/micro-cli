package utils_test

import (
	"testing"

	"github.com/eduardoths/micro-cli/utils"
)

func TestMergePaths(t *testing.T) {
	type testCase struct {
		it   string
		in   []string
		want string
	}

	tc := []testCase{
		{
			it:   "should return the inputed path",
			in:   []string{"."},
			want: ".",
		},
		{
			it:   "should return the inputed path if it ends with '/'",
			in:   []string{"./"},
			want: "./",
		},
		{
			it:   "should merge two paths",
			in:   []string{"..", "file.go"},
			want: "../file.go",
		},
		{
			it:   "should merge two paths that end/start with '/'",
			in:   []string{"/usr/local/", "/bin/educli"},
			want: "/usr/local/bin/educli",
		},
		{
			it:   "should merge one path that is root",
			in:   []string{"/"},
			want: "/",
		},
		{
			it:   "should merge three paths that are root",
			in:   []string{"/", "/", "/"},
			want: "/",
		},
		{
			it:   "should end with slash",
			in:   []string{"/", "test/"},
			want: "/test/",
		},
		{
			it:   "should mark current directory with a dot",
			in:   []string{"", "/src/structs", "file.go"},
			want: "./src/structs/file.go",
		},
	}

	for _, c := range tc {
		t.Run(c.it, func(t *testing.T) {
			actual := utils.MergePaths(c.in...)
			if c.want != actual {
				t.Errorf("TestMergePaths failed.\nGot:\t\t%s\nwant:\t%s", actual, c.want)
				t.Logf("Case: %s", c.it)
			}
		})
	}
}

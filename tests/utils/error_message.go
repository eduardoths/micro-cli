package utils

import "testing"

func Error(t *testing.T, want any, got any) {
	t.Helper()
	t.Errorf("Test %s failed:\nexpected:\t%v\ngot:\t\t\t\t%v", t.Name(), want, got)
}

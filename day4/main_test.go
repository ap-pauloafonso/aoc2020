package main

import "testing"

func AssertAreEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v != want = %v", got, want)
	}
}
func TestValidByr(t *testing.T) {
	var testsData = []struct {
		n    string // input
		want bool   // expected result
	}{
		{"2002", true},
		{"2003", false},
	}

	for _, a := range testsData {

		AssertAreEqual(t, validByr(a.n), a.want)
	}

}
func TestValidHgt(t *testing.T) {
	var testsData = []struct {
		n    string // input
		want bool   // expected result
	}{
		{"60in", true},
		{"190cm", true},
		{"190in", false},
		{"190", false},
	}

	for _, a := range testsData {
		AssertAreEqual(t, validHgt(a.n), a.want)
	}
}

func TestValidHcl(t *testing.T) {
	var testsData = []struct {
		n    string // input
		want bool   // expected result
	}{
		{"#123abc", true},
		{"#123abz", false},
		{"123abc", false},
	}

	for _, a := range testsData {
		AssertAreEqual(t, validHcl(a.n), a.want)
	}
}

func TestValidEcl(t *testing.T) {
	var testsData = []struct {
		n    string // input
		want bool   // expected result
	}{
		{"brn", true},
		{"wat", false},
	}

	for _, a := range testsData {
		AssertAreEqual(t, validEcl(a.n), a.want)
	}
}

func TestValidPid(t *testing.T) {
	var testsData = []struct {
		n    string // input
		want bool   // expected result
	}{
		{"000000001", true},
		{"0123456789", false},
	}

	for _, a := range testsData {
		AssertAreEqual(t, validPid(a.n), a.want)
	}
}

package problem67maximumpathsumtwo

import (
	"testing"
	"io/ioutil"
)

func TestProblem67Solution (t *testing.T) {
	buf, _ := ioutil.ReadFile("./fixture.txt")
	fixture := string(buf)

	expected := 7273

	solver := New(fixture)
	result, err := solver.solve()

	if err != nil {
		t.Errorf("Encountered error whilst solving: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d got %d", expected, result)
	}
}
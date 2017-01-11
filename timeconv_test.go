package timeconversion

import "testing"

var (
	timeTests = []struct {
		input string
		want  string
	}{
		{"10:05:47AM", "10:05:47"},
		{"07:05:45PM", "19:05:45"},
	}
)

// TestConvertTime runs tests against ConvertTime.
func TestConvertTime(t *testing.T) {
	for _, value := range timeTests {
		if got := ConvertTime(value.input); got != value.want {
			t.Fatalf("ConvertTime(%s) = %q, want %q", value.input, got, value.want)
		}
	}
}

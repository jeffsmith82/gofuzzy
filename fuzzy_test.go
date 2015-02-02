package gofuzzy

import "testing"

func TestLd(t *testing.T) {

	cases := []struct {
		in1  string
		in2  string
		want int
	}{
		{"Kitten", "sitting", 3},
		{"Test", "testing", 4},
		{"", "", 0},
		{"世界", "Hello", 6},
	}
	for _, c := range cases {
		got := Ld(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Ld(%q,%q) == %d, want %d\n", c.in1, c.in2, got, c.want)
		}
	}
}

func TestSoundex(t *testing.T) {

	cases := []struct {
		in, want  string
		wanterror bool
	}{
		{"Robert", "R163", false},
		{"Rupert", "R163", false},
		{"Rubin", "R150", false},
		{"ashcroft", "A261", false},
		{"ashcraft", "A261", false},
		{"moses", "M220", false},
		{"O'Mally", "O540", false},
		{"d jay", "D200", false},
		{"R2-D2", "R300", false},
		{"12p2", "initial character must be a letter", true},
		{"naïve", "non-ASCII letters unsupported", true},
		{"", "no letters present", true},
		{"bump\t", "ASCII control characters disallowed", true},
	}
	for _, c := range cases {
		got, err := Soundex(c.in)
		if err != nil && err.Error() != c.want {
			t.Errorf("Soundex(%q) = %q, err want %q\n", c.in, got, err.Error(), c.want)
		} else if got != c.want && c.wanterror != true {
			t.Errorf("Soundex(%q) == %q, want %q\n", c.in, got, c.want)
		}
	}
}

func BenchmarkLd(b *testing.B){

	for n := 0; n < b.N; n++ {
		Ld("Testing","tester")
	}
}

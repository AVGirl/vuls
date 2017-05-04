package oval

import "testing"

func TestParseCvss2(t *testing.T) {
	type out struct {
		score  float64
		vector string
	}
	var tests = []struct {
		in  string
		out out
	}{
		{
			in: "5/AV:N/AC:L/Au:N/C:N/I:N/A:P",
			out: out{
				score:  5.0,
				vector: "AV:N/AC:L/Au:N/C:N/I:N/A:P",
			},
		},
		{
			in: "",
			out: out{
				score:  0,
				vector: "",
			},
		},
	}
	for _, tt := range tests {
		s, v := Redhat{}.parseCvss2(tt.in)
		if s != tt.out.score || v != tt.out.vector {
			t.Errorf("\nexpected: %f, %s\n  actual: %f, %s",
				tt.out.score, tt.out.vector, s, v)
		}
	}
}

func TestParseCvss3(t *testing.T) {
	type out struct {
		score  float64
		vector string
	}
	var tests = []struct {
		in  string
		out out
	}{
		{
			in: "5.6/CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L",
			out: out{
				score:  5.6,
				vector: "AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L",
			},
		},
		{
			in: "",
			out: out{
				score:  0,
				vector: "",
			},
		},
	}
	for _, tt := range tests {
		s, v := Redhat{}.parseCvss3(tt.in)
		if s != tt.out.score || v != tt.out.vector {
			t.Errorf("\nexpected: %f, %s\n  actual: %f, %s",
				tt.out.score, tt.out.vector, s, v)
		}
	}
}

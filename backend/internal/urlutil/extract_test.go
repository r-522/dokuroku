package urlutil

import "testing"

func TestExtract_UT_URL_001(t *testing.T) {
	got := Extract("URLなし")
	if len(got) != 0 {
		t.Fatalf("got %v", got)
	}
}

func TestExtract_UT_URL_002(t *testing.T) {
	got := Extract("https://go.dev/ と https://pkg.go.dev/net/http。 https://go.dev/")
	want := []string{"https://go.dev/", "https://pkg.go.dev/net/http"}
	if len(got) != len(want) {
		t.Fatalf("got %v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v", got)
		}
	}
}

package ondavital

import (
	"testing"
)

func TestSearch(t *testing.T) {

	tests := make(map[string]string, 3)
	tests["die hard"] = "Jungla de cristal"
	tests["rapido y furioso"] = "A todo gas"
	tests["white chicks"] = "Dos rubias de pelo en pecho"
	tests["bomberman"] = "" //dont crash

	for k, v := range tests {
		title, _ := Search(k)
		if title != v {
			_ = "breakpoint"
			t.Errorf("Search(%s)== expected: %s, got: %s\n", k, v, title)
		}
	}
}

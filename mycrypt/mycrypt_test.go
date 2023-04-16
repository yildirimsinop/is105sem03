package mycrypt

import (
	"reflect"
	"testing"
)

// Kjevik;SN39040;18.03.2022 01:50;6
// var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

// Testen feiler, siden æ ligger tre posisjoner til høyre fra x
// Testen skal passere, hvis man tester for w istedenfor x
func TestKrypter(t *testing.T) {
	wanted := []rune("æ")
	state := Krypter([]rune("w"), ALF_SEM03, 4)
	if !reflect.DeepEqual(state, wanted) {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

// Tester posisjon for rune 'æ' i ALF_SEM03
// (begynner på 0 teller fra venstre til høyre)
func TestSokIAlfabetet(t *testing.T) {
	wanted := 26
	got := sokIAlfabetet('æ', ALF_SEM03)
	if got != wanted {
		t.Errorf("Feil, fikk %d, ønsket %d.", got, wanted)
	}
}

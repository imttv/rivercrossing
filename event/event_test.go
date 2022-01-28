package event

import "testing"

func TestViewEvent(t *testing.T) {
	wanted := "| Kylling Rev Korn HS -båthavn vest-\\________/__________________-Båthavn øst- |"
	event := ViewEvent()
	if event != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", event, wanted)
	}
}

func TestBåtVestKylling(t *testing.T) {
	wanted := event
	if event != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", event, wanted)
	}
}

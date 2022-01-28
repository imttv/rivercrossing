package event

var event = "| Kylling Rev Korn HS -båthavn vest-\\________/__________________-Båthavn øst- |"

func ViewEvent() string {
	return event
}

func BåtVestKylling() {
	event = "| Rev Korn -båthavn vest-\\_Kylling HS_/__________________-båthavn øst- |"
}

func BåtØstKylling() {
	event = "| Rev Korn -båthavn vest-__________________\\_Kylling HS_/-båthavn øst- |"
}

func LandØstKylling() {
	event = "| Rev Korn -båthavn vest-__________________\\__HS__/-båthavn øst- Kylling |"
}

func BåtTilVest() {
	event = "| Rev Korn -båthavn vest-\\__HS__/__________________-båthavn øst- Kylling |"
}

package main

type state struct {
	hunger    int
	happiness int
	age       int
	wisdom    int
	fitness   int
	name      string
	gender    string
	lesson    string
	mood      string
	aging     int
}

func initState() state {
	return state{
		hunger:    50,
		happiness: 50,
		age:       0,
		wisdom:    0,
		fitness:   50,
	}
}

func (s state) describeMood() string {
	if s.happiness < 10 {
		return "suicidal"
	}
	if s.happiness < 30 {
		return "depressed"
	}
	if s.happiness < 40 {
		return "sad"
	}
	if s.happiness < 50 {
		return "down"
	}
}

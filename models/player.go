package models

type Player struct {
	Name    string
	Surname string
	Age     int
	Role    Role
}

func NewPlayer(name, surname string, age int, role Role) Player {
	return Player{
		name,
		surname,
		age,
		role,
	}
}

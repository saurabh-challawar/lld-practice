package player

type Player struct {
	id   int
	name string
}

func NewPlayer(id int, name string) *Player {
	return &Player{
		id:   id,
		name: name,
	}
}

func (player *Player) GetId() int {
	return player.id
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) SetId(id int) *Player {
	player.id = id
	return player
}

func (player *Player) SetName(name string) *Player {
	player.name = name
	return player
}

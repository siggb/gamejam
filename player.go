package main

type Vec2 struct {
	X float32
	Y float32
}

const (
	minPlayerPosY float32 = 20
	maxPlayerPosY float32 = 365
)

var availableColors []string = []string{"red", "blue", "green", "yellow"}

var currentPlayerNumber = 0

type Player struct {
	Id      string
	Name    string
	Pos     Vec2
	Vel     Vec2
	TargetY float32
	Color   string
}

func NewPlayer(socketId string, playerName string) (player *Player) {
	player = &Player{
		Id:      socketId,
		Name:    playerName,
		Vel:     Vec2{X: 250, Y: 0},
		Pos:     Vec2{X: 80, Y: float32(currentPlayerNumber+1) * 80.0},
		TargetY: 100,
		Color:   availableColors[currentPlayerNumber],
	}

	currentPlayerNumber++

	return
}

func (self *Player) Move(msg string) {
	switch msg {
	case "down":
		self.Vel.Y = 100
	case "up":
		self.Vel.Y = -100
	case "stop":
		self.Vel.Y = 0
	}
}

func (self *Player) Update(dt float32) {
	self.Pos.X += self.Vel.X * dt
	self.Pos.Y += self.Vel.Y * dt

	if self.Pos.Y < minPlayerPosY {
		self.Pos.Y = minPlayerPosY
	}

	if self.Pos.Y > maxPlayerPosY {
		self.Pos.Y = maxPlayerPosY
	}

}
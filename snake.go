package main

var (
	DIRECTION_LEFT  = Point{X: -1, Y: 0}
	DIRECTION_RIGHT = Point{X: 1, Y: 0}
	DIRECTION_UP    = Point{X: 0, Y: -1}
	DIRECTION_DOWN  = Point{X: 0, Y: 1}
)

type Point struct {
	X int
	Y int
}

type Snake struct {
	Body        []Point
	Direction   Point
	GrowCounter int
}

func NewSnake() *Snake {
	return &Snake{
		Body: []Point{
			{X: SCREEN_WIDTH / TILE_SIZE / 2, Y: SCREEN_HEIGTH / TILE_SIZE / 2},
		},
		Direction: DIRECTION_RIGHT,
	}
}

func (s *Snake) Move() {
	newHead := Point{
		X: s.Body[0].X + s.Direction.X,
		Y: s.Body[0].Y + s.Direction.Y,
	}
	s.Body = append([]Point{newHead}, s.Body...)

	if s.GrowCounter > 0 {
		s.GrowCounter--
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}

func (s *Snake) IsDirectionVertical() bool {
	return s.Direction.X == 0
}

func (s *Snake) IsDirectionHorizontal() bool {
	return s.Direction.Y == 0
}

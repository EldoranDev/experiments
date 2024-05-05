package screen

const WIDTH = 64
const HEIGHT = 32

type Screen struct {
	data [WIDTH * HEIGHT]bool
}

func New() *Screen {
	return &Screen{}
}

func (s *Screen) Clear() {
	for i := range WIDTH * HEIGHT {
		s.data[i] = false
	}
}

func (s *Screen) Set(x, y uint16) uint8 {
	i := y*WIDTH + x
	if s.data[i] {
		s.data[i] = false
		return 1
	}

	s.data[i] = true

	return 0
}

func (s *Screen) IsSet(x, y int) bool {
	return s.data[y*WIDTH+x]
}

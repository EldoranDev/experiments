package timer

type Timer struct {
	Sound uint8
	Delay uint8
}

func New() *Timer {
	return &Timer{}
}

func (t *Timer) Tick() {
	if t.Sound > 0 {
		t.Sound--
	}

	if t.Delay > 0 {
		t.Delay--
	}
}

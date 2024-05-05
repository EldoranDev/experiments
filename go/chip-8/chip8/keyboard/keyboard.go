package keyboard

type Keyboard struct {
	PressedKey uint8
}

func New() *Keyboard {
	return &Keyboard{}
}

func (k *Keyboard) IsPressed(key uint8) bool {
	return k.PressedKey == key
}

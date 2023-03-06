package wisdom

import "math/rand"

func NewWisdom() *Wisdom {
	return &Wisdom{
		store: []string{
			"The fool doth think he is wise, but the wise man knows himself to be a fool\n\t\t\t(c) William Shakespeare",
			"It is better to remain silent at the risk of being thought a fool, than to talk and remove all doubt of it\n\t\t\t(c) Maurice Switzer",
			"Whenever you find yourself on the side of the majority, it is time to reform (or pause and reflect)\n\t\t\t(c) Mark Twain",
			"When someone loves you, the way they talk about you is different. You feel safe and comfortable\n\t\t\t(c) Jess C. Scott",
		},
	}
}

type Wisdom struct {
	store []string
}

func (w *Wisdom) Word() string {
	return w.store[rand.Int()%len(w.store)]
}

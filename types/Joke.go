package types

type Flag struct {
	NSFW      bool `json:"nsfw"`
	Religious bool `json:"religious"`
	Political bool `json:"political"`
	Racist    bool `json:"racist"`
	Sexist    bool `json:"sexist"`
	Explicit  bool `json:"explicit"`
}

type Joke struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup,omitempty"`
	Delivery string `json:"delivery,omitempty"`
	Joke     string `json:"joke,omitempty"`
	Flags    Flag   `json:"flags"`
}

func NewJoke(c, t, s, d, j string, f Flag) *Joke {
	return &Joke{
		Category: c,
		Type:     t,
		Setup:    s,
		Delivery: d,
		Joke:     j,
		Flags:    f,
	}
}

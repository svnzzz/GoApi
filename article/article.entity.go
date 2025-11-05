package article

type Article struct {
	ID             string  `json:"id"`
	Nome           string  `json:"nome"`
	Giacenza       int     `json:"giacenza"`
	PrezzoUnitario float32 `json:"prezzoUnitario"`
}

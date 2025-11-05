package article

type Article struct {
	ID             string  `json:"id"`
	Nome           string  `json:"nome"`
	Giacenza       int     `json:"giacenza"`
	PrezzoUnitario float32 `json:"prezzoUnitario"`
}

type EditArticleRequest struct {
	ID   string `json:"id"`
	Edit string `json:"edit"`
	Into string `json:"into"`
}

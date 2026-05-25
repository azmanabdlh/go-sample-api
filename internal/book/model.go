package book

type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	CoverImageURL string `json:"cover_image_url"`
	AuthorID      string `json:"author_id"`
}

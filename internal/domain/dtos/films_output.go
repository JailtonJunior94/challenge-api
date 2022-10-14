package dtos

type FilmsOutput struct {
	Title       string `json:"title"`
	EpisodeID   int    `json:"episode_id"`
	Director    string `json:"director"`
	ReleaseDate string `json:"release_date"`
}

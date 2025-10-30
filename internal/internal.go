package internal

import "github.com/21ess/animemaster/src/model"

// AnimeFetch adapter mode's target
type AnimeFetch interface {
	FetchRandomAnime(filter map[string]any) model.Anime
	FetchAllAnime(filter map[string]any) []model.Anime
}

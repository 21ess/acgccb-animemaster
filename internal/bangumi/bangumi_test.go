package bangumi

import (
	"testing"

	"github.com/21ess/animemaster/src/config"
)

func TestBangumiFetch_FetchRandomAnime(t *testing.T) {
	config.LoadConfig()
	filter := map[string]any{
		"tag":       []string{"京阿尼", "校园", "推理", "TV"},
		"type":      []int{2},
		"meta_tags": []string{},
		"air_date":  []string{">=2008-07-01", "<2024-10-01"},
		"rating":    []string{},
		"rank":      []string{},
		"nsfw":      false,
	}

	bangumi := &FetchAdapter{
		client: NewBangumi(),
	}

	bangumi.FetchRandomAnime(filter)
}

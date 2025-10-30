package bangumi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/21ess/animemaster/src/log"
)

type SubjectReq struct {
	Keyword string         `json:"keyword"`
	Sort    string         `json:"sort,omitempty"`
	Filter  map[string]any `json:"filter,omitempty"`
}

type PagedSubjectRsp struct {
	Total  int       `json:"total"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
	Data   []Subject `json:"data"`
}

type Bangumi struct {
	Client *http.Client
	Prefix string
}

func NewBangumi() *Bangumi {
	return &Bangumi{
		Client: &http.Client{},
		Prefix: "https://api.bgm.tv",
	}
}

func (b *Bangumi) SearchSubject(keyword string, sort string, filter map[string]any, params url.Values) (*PagedSubjectRsp, error) {
	apiURL := b.Prefix + "/v0/search/subjects"
	reqBody := SubjectReq{
		Keyword: keyword,
		Sort:    sort,
		Filter:  filter,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Log.Error(err.Error())
		return nil, err
	}
	if params.Encode() != "" {
		apiURL += "?" + params.Encode()
	}
	request, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Log.Error(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+os.Getenv("BANGUMI_TOKEN"))
	request.Header.Set("User-Agent", os.Getenv("USER_AGENT"))

	resp, err := b.Client.Do(request)
	if err != nil {
		log.Log.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		var pagedRsp PagedSubjectRsp
		if err := json.Unmarshal(body, &pagedRsp); err != nil {
			log.Log.Error(err.Error())
			return nil, err
		}

		return &pagedRsp, nil
	}

	log.Log.Error("Bangumi search subject api error", "Response Code", resp.StatusCode, "Body", string(body))
	return nil, fmt.Errorf("bangumi search subject api error: %s", string(body))
}

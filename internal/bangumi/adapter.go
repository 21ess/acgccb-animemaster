package bangumi

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/21ess/animemaster/src/log"
	"github.com/21ess/animemaster/src/model"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

type FetchAdapter struct {
	client *Bangumi
}

func (b *FetchAdapter) FetchRandomAnime(filter map[string]any) (*model.Anime, error) {
	pagedSubjectRsp, err := b.client.SearchSubject("", "rank", filter, url.Values{})
	if err != nil {
		log.Log.Error("bangumiFetch.client.SearchSubject err")
		return nil, err
	}

	var randSubject Subject
	if pagedSubjectRsp.Total == 0 {
		log.Log.Info("bangumi no such anime")
		return nil, errors.New("no anime found")
	} else if pagedSubjectRsp.Total <= pagedSubjectRsp.Limit {
		ind := rand.Int() % pagedSubjectRsp.Total
		randSubject = pagedSubjectRsp.Data[ind]
	} else {
		ind := rand.Int() % pagedSubjectRsp.Total
		if ind < pagedSubjectRsp.Limit {
			randSubject = pagedSubjectRsp.Data[ind]
		} else {
			params := url.Values{}
			limit := pagedSubjectRsp.Limit
			offset := ind / limit * limit
			ind = ind % limit
			params.Set("limit", strconv.Itoa(limit))
			params.Set("offset", strconv.Itoa(offset))
			subjectRsp, err := b.client.SearchSubject("", "rank", filter, params)
			if err != nil {
				return nil, err
			}
			randSubject = subjectRsp.Data[ind]
		}
	}

	fmt.Println(randSubject)
	return nil, nil
}

func (b *FetchAdapter) FetchAllAnime(filter map[string]any) []model.Anime {
	//TODO implement me
	panic("implement me")
}

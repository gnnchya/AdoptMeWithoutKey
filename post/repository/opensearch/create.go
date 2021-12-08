package opensearch

import (
	"context"
	"encoding/json"
	"github.com/gnnchya/AdoptMeWithoutKey/post/domain"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"io"
	"strings"
)

func (repo *Repository) CreateAdoptionPost(ctx context.Context, title *domain.CreateAdoptionPostStruct) error {
	out, err := json.Marshal(title)
	if err != nil {
		return err
	}

	var b strings.Builder
	b.WriteString(string(out))
	req := opensearchapi.IndexRequest{
		Index:      repo.AdoptionPostIndex,
		DocumentID: title.ID,
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}
	res, err := req.Do(ctx, repo.Client)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	return err
}

func (repo *Repository) CreateLostPetPost(ctx context.Context, title *domain.CreateLostPetPostStruct) error {
	out, err := json.Marshal(title)
	if err != nil {
		return err
	}

	var b strings.Builder
	b.WriteString(string(out))
	req := opensearchapi.IndexRequest{
		Index:      repo.LostPetPostIndex,
		DocumentID: title.ID,
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}
	res, err := req.Do(ctx, repo.Client)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	return err
}

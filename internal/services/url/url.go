package url

import (
	"context"

	"github.com/akanksh5/flickly/internal/stores"
)

type URLService struct {
	store stores.URLStore
}

func NewURLService(store stores.URLStore) *URLService {
	return &URLService{store: store}
}

func (s *URLService) GenerateShortenedURL(ctx context.Context, originalURL string) (string, error) {
	return s.store.GenerateShortenedURL(ctx, originalURL)
}

func (s *URLService) GetLongURL(ctx context.Context, shortURL string) (string, error) {
	return s.store.GetLongURL(ctx, shortURL)
}

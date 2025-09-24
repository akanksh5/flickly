package stores

import "context"

type URLStore interface {
	GenerateShortenedURL(ctx context.Context, originalURL string) (string, error)
	GetLongURL(ctx context.Context, shortURL string) (string, error)
}

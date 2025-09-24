package services

import "context"

type URLService interface {
	GenerateShortenedURL(ctx context.Context, originalURL string) (string, error)
	GetLongURL(ctx context.Context, shortURL string) (string, error)
}

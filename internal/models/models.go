package models

type URLMappingReq struct {
	OriginalURL string
}

type URLMappingResponse struct {
	OriginalURL  string
	ShortenedURL string
}

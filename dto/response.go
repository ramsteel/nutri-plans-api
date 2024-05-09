package dto

type BaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type SearchResponse struct {
	BaseResponse
	Metadata *MetadataResponse `json:"metadata,omitempty"`
}

type MetadataResponse struct {
	TotalData   int  `json:"total_data"`
	TotalCount  int  `json:"total_count"`
	NextOffset  int  `json:"next_offset"`
	HasLoadMore bool `json:"has_load_more"`
}

type PaginationResponse struct {
	BaseResponse
	Pagination *PaginationMetadata `json:"pagination"`
	Link       *Link               `json:"link"`
}

type Link struct {
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

type PaginationMetadata struct {
	CurrentPage int   `json:"current_page"`
	TotalPage   int   `json:"total_page"`
	TotalData   int64 `json:"total_data"`
}

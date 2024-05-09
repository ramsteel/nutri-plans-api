package dto

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

type SearchRequest struct {
	Limit  int `json:"limit" validate:"required,gt=0"`
	Offset int `json:"offset" validate:"required,gt=0"`
}

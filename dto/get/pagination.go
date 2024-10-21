package get

type PaginationResponse[T any] struct {
	Page     int `json:"page"`
	Size     int `json:"size"`
	Elements []T `json:"elements"`
}

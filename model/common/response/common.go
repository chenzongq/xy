package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type PageResultT[T any] struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
	Total    int64 `json:"total"`
	Data     []*T  `json:"data"`
}

func BuildPageResultT[T any](page, pageSize, total int64, data []*T) *PageResultT[T] {
	return &PageResultT[T]{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Data:     data,
	}
}

func BuildEmptyPageResultT[T any]() *PageResultT[T] {
	return &PageResultT[T]{
		Data: make([]*T, 0),
	}
}

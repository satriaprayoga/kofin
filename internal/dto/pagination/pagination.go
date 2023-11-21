package pagination

type ParamList struct {
	Page       int    `json:"pageIndex" valid:"Required"`
	PerPage    int    `json:"pageSize" valid:"Required"`
	Search     string `json:"query,omitempty"`
	InitSearch string `json:"init_search,omitempty"`
	SortField  Sort   `json:"sort,omitempty"`
}

type Sort struct {
	Order string `json:"order"`
	Key   string `json:"key"`
}

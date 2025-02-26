package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page" example:"1"`                // 页码
	PageSize int    `json:"pageSize" form:"pageSize" example:"10"`       // 每页大小
	Keyword  string `json:"keyword" form:"keyword" swaggerignore:"true"` //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

// GetTenantId Get role by id structure
type GetTenantId struct {
	TenantId uint `json:"tenantId,string" form:"tenantId"` // 角色ID
}

type Empty struct{}

type DateCond struct {
	StartCreatedAt *int64 `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *int64 `json:"endCreatedAt" form:"endCreatedAt"`
}

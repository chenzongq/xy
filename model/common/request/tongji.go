package request

type StatisticsRequest struct {
	From        uint  `json:"from" form:"from" binding:"required,ltefield=To"`
	To          uint  `json:"to" form:"to" binding:"required,gtefield=From"`
	CompanyId   int64 `json:"companyId" form:"companyId"`      // 服务公司ID/叫车公司ID
	CompanyType uint  `json:"companyType" form:"companyType" ` // 1 服务公司 2 叫车公司
	Unit        uint8 `json:"unit" form:"unit" example:"3"`    //1按小时统计 2按天统计 3按月统计 4按年统计
}

type PaymentDownDashboardRequest struct {
	From uint  `json:"from" form:"from" binding:"required,ltefield=To"`
	To   uint  `json:"to" form:"to" binding:"required,gtefield=From"`
	Unit uint8 `json:"unit" form:"unit" example:"3"` //1按小时统计 2按天统计 3按月统计 4按年统计
}
type PaymentDownDashboardCompareRequest struct {
	From       uint  `json:"from" form:"from" binding:"required,ltefield=To"`
	To         uint  `json:"to" form:"to" binding:"required,gtefield=From"`
	ServantIds []int `json:"servantIds" form:"servantIds" binding:"required"`
	Unit       uint8 `json:"unit" form:"unit" example:"3"` //1按小时统计 2按天统计 3按月统计 4按年统计
}
type PaymentDownDashboardAreaRequest struct {
	From uint   `json:"from" form:"from" binding:"required,ltefield=To"`
	To   uint   `json:"to" form:"to" binding:"required,gtefield=From"`
	Mode uint8  `json:"mode" form:"mode" ` // 救援模式 0非救援模式 1事故救援
	City string `json:"city" form:"city"`
	Unit uint8  `json:"unit" form:"unit" example:"3"` //1按小时统计 2按天统计 3按月统计 4按年统计
}

package global

import (
	"gorm.io/gorm"
	"time"
)

type GvaModel struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                       // 删除时间
	TenantId  uint           `json:"tenantId,string" ,gorm:"comment:租户ID"` //租户ID
}

// BaseModel 基础model字段
type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`                // 主键ID
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`       // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`       // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                      // 删除时间
	TenantId  *uint          `json:"tenantId,string" gorm:"comment:租户ID"` //租户ID
	Creator   *uint          `json:"creator" gorm:"comment:创建人ID"`        //创建人
	Updater   *uint          `json:"updater" gorm:"comment:更新人ID"`        //更新人
}

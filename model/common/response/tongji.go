package response

type RangeUnit uint8

const (
	// 这部分常量规定了按时间区间进行统计的枚举
	_     RangeUnit = iota
	Hour            // 按小时统计
	Day             // 按天统计
	Month           // 按月统计
	Year            // 按年统计
)

func (r RangeUnit) Format() string {
	switch r {
	case Hour:
		return "%m-%d %H"
	case Day:
		return "%Y-%m-%d"
	case Month:
		return "%Y-%m"
	case Year:
		return "%Y"
	}

	return ""
}

type RangeAble struct {
	Name     string  `json:"name"`     // 统计名称
	Desc     string  `json:"desc"`     // 统计描述，也可以是其他附加信息
	Efficacy string  `json:"efficacy"` // 产出占比
	Count    float32 `json:"count"`    // 统计值
	Unit     string  `json:"unit"`     // 单位
}

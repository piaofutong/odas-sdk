package common

// 按年龄段性别分布汇总统计数据
type TouristSexSummary struct {
	Total   int64 `json:"total"`   // 游客总数量
	Male    int64 `json:"male"`    // 男性游客总数量
	Female  int64 `json:"female"`  // 女性游客总数量
	Unknown int64 `json:"unknown"` // 未知性别游客总数量
}

// 游客统计汇总数据 Tourist Summary
type TouristSummary struct {
	Count int64 `json:"count"` // 游客总数量
}

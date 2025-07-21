package types

// ScenicSexAge represents scenic sex age statistics
type ScenicSexAge struct {
	Date     YMDTime `json:"date"`
	Sid      int     `json:"sid"`      // 原始供应商ID
	Lid      int     `json:"lid"`      // 产品ID
	Lname    string  `json:"lname"`    // 产品名称
	Country  string  `json:"country"`  // 国家
	Province string  `json:"province"` // 省份
	City     string  `json:"city"`     // 城市

	// 指标字段
	Total           int `json:"total"`                         // 总人数 - 只包含首次核销
	Unknown         int `json:"unknown"`                       // 性别年龄未知 - 只包含首次核销
	MaleAge0To7     int `json:"male_age_0to7"`                 // 男性0-7岁
	MaleAge8To17    int `json:"male_age_8to17"`                // 男性8-17岁
	MaleAge18To27   int `json:"male_age_18to27"`               // 男性18-27岁
	MaleAge28To40   int `json:"male_age_28to40"`               // 男性28-40岁
	MaleAge41To60   int `json:"male_age_41to60"`               // 男性41-60岁
	MaleAge61Plus   int `json:"male_age_61plus"`               // 男性60岁以上
	FemaleAge0To7   int `json:"female_age_0to7"`               // 女性0-7岁
	FemaleAge8To17  int `json:"female_age_8to17"`              // 女性8-17岁
	FemaleAge18To27 int `json:"female_age_18to27"`             // 女性18-27岁
	FemaleAge28To40 int `json:"female_age_28to40"`             // 女性28-40岁
	FemaleAge41To60 int `json:"female_age_41to60"`             // 女性41-60岁
	FemaleAge61Plus int `json:"female_age_61plus"`             // 女性60岁以上
}
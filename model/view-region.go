package model

type ProvinceView struct {
	ID           string `json:"id"`
	ProvinceName string `json:"provinceName"`
}

func (ProvinceView) TableName() string {
	return VIEW_PROVINCE
}

type RegencyView struct {
	ID           string `json:"id"`
	ProvinceID   string `json:"provinceId"`
	RegencyName  string `json:"regencyName"`
	ProvinceName string `json:"provinceName"`
}

func (RegencyView) TableName() string {
	return VIEW_REGENCY
}

type DistrictView struct {
	ID           string `json:"id"`
	ProvinceID   string `json:"provinceId"`
	RegencyID    string `json:"regencyId"`
	DistrictName string `json:"districtName"`
	RegencyName  string `json:"regencyName"`
	ProvinceName string `json:"provinceName"`
}

func (DistrictView) TableName() string {
	return VIEW_DISTRICT
}

type VillageView struct {
	ID           string `json:"id"`
	ProvinceID   string `json:"provinceId"`
	RegencyID    string `json:"regencyId"`
	DistrictID   string `json:"districtId"`
	VillageName  string `json:"villageName"`
	DistrictName string `json:"districtName"`
	RegencyName  string `json:"regencyName"`
	ProvinceName string `json:"provinceName"`
}

func (VillageView) TableName() string {
	return VIEW_VILLAGE
}

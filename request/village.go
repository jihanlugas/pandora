package request

type PageVillage struct {
	Paging
	VillageName string `json:"villageName" form:"villageName" query:"villageName"`
	ProvinceID  string `json:"provinceId" form:"provinceId" query:"provinceId"`
	RegencyID   string `json:"regencyId" form:"regencyId" query:"regencyId"`
	DistrictID  string `json:"districtId" form:"districtId" query:"districtId"`
}

type ListVillage struct {
	Listing
	VillageName string `json:"villageName" form:"villageName" query:"villageName"`
	ProvinceID  string `json:"provinceId" form:"provinceId" query:"provinceId"`
	RegencyID   string `json:"regencyId" form:"regencyId" query:"regencyId"`
	DistrictID  string `json:"districtId" form:"districtId" query:"districtId"`
}

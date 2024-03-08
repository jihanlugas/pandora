package request

type PageDistrict struct {
	Paging
	DistrictName string `json:"districtName" form:"districtName" query:"districtName"`
	ProvinceID   string `json:"provinceId" form:"provinceId" query:"provinceId"`
	RegencyID    string `json:"regencyId" form:"regencyId" query:"regencyId"`
}

type ListDistrict struct {
	Listing
	DistrictName string `json:"districtName" form:"districtName" query:"districtName"`
	ProvinceID   string `json:"provinceId" form:"provinceId" query:"provinceId"`
	RegencyID    string `json:"regencyId" form:"regencyId" query:"regencyId"`
}

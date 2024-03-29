package request

type PageRegency struct {
	Paging
	RegencyName string `json:"regencyName" form:"regencyName" query:"regencyName"`
	ProvinceID  string `json:"provinceId" form:"provinceId" query:"provinceId"`
}

type ListRegency struct {
	Listing
	RegencyName string `json:"regencyName" form:"regencyName" query:"regencyName"`
	ProvinceID  string `json:"provinceId" form:"provinceId" query:"provinceId"`
}

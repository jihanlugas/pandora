package request

type PageProvince struct {
	Paging
	ProvinceName string `json:"provinceName" form:"provinceName" query:"provinceName"`
}

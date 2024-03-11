package request

import "time"

type CreateKtp struct {
	Nik              string     `json:"nik" form:"nik" query:"nik" validate:"required"`
	Nama             string     `json:"nama" form:"nama" query:"nama" validate:"required"`
	TempatLahir      string     `json:"tempatLahir" form:"tempatLahir" query:"tempatLahir" validate:"required"`
	TanggalLahir     time.Time  `json:"tanggalLahir" form:"tanggalLahir" query:"tanggalLahir" validate:"required"`
	JenisKelamin     string     `json:"jenisKelamin" form:"jenisKelamin" query:"jenisKelamin" validate:"required"`
	ProvinceID       string     `json:"provinceId" form:"provinceId" query:"provinceId" validate:""`
	RegencyID        string     `json:"regencyId" form:"regencyId" query:"regencyId" validate:""`
	DistrictID       string     `json:"districtId" form:"districtId" query:"districtId" validate:""`
	VillageID        string     `json:"villageId" form:"villageId" query:"villageId" validate:""`
	Alamat           string     `json:"alamat" form:"alamat" query:"alamat" validate:""`
	Rtrw             string     `json:"rtrw" form:"rtrw" query:"rtrw" validate:""`
	Pekerjaan        string     `json:"pekerjaan" form:"pekerjaan" query:"pekerjaan" validate:""`
	StatusPerkawinan string     `json:"statusPerkawinan" form:"statusPerkawinan" query:"statusPerkawinan" validate:""`
	Kewarganegaraan  string     `json:"kewarganegaraan" form:"kewarganegaraan" query:"kewarganegaraan" validate:""`
	BerlakuHingga    *time.Time `json:"berlakuHingga" form:"berlakuHingga" query:"berlakuHingga" validate:""`
}

type UpdateKtp struct {
	Nik              string     `json:"nik" form:"nik" query:"nik" validate:"required"`
	Nama             string     `json:"nama" form:"nama" query:"nama" validate:"required"`
	TempatLahir      string     `json:"tempatLahir" form:"tempatLahir" query:"tempatLahir" validate:"required"`
	TanggalLahir     time.Time  `json:"tanggalLahir" form:"tanggalLahir" query:"tanggalLahir" validate:"required"`
	JenisKelamin     string     `json:"jenisKelamin" form:"jenisKelamin" query:"jenisKelamin" validate:"required"`
	ProvinceID       string     `json:"provinceId" form:"provinceId" query:"provinceId" validate:""`
	RegencyID        string     `json:"regencyId" form:"regencyId" query:"regencyId" validate:""`
	DistrictID       string     `json:"districtId" form:"districtId" query:"districtId" validate:""`
	VillageID        string     `json:"villageId" form:"villageId" query:"villageId" validate:""`
	Alamat           string     `json:"alamat" form:"alamat" query:"alamat" validate:""`
	Rtrw             string     `json:"rtrw" form:"rtrw" query:"rtrw" validate:""`
	Pekerjaan        string     `json:"pekerjaan" form:"pekerjaan" query:"pekerjaan" validate:""`
	StatusPerkawinan string     `json:"statusPerkawinan" form:"statusPerkawinan" query:"statusPerkawinan" validate:""`
	Kewarganegaraan  string     `json:"kewarganegaraan" form:"kewarganegaraan" query:"kewarganegaraan" validate:""`
	BerlakuHingga    *time.Time `json:"berlakuHingga" form:"berlakuHingga" query:"berlakuHingga" validate:""`
}

type PageKtp struct {
	Paging
	Nik              string `json:"nik" form:"nik" query:"nik"`
	Nama             string `json:"nama" form:"nama" query:"nama"`
	JenisKelamin     string `json:"jenisKelamin" form:"jenisKelamin" query:"jenisKelamin"`
	ProvinceID       string `json:"provinceId" form:"provinceId" query:"provinceId"`
	RegencyID        string `json:"regencyId" form:"regencyId" query:"regencyId"`
	DistrictID       string `json:"districtId" form:"districtId" query:"districtId"`
	VillageID        string `json:"villageId" form:"villageId" query:"villageId"`
	Alamat           string `json:"alamat" form:"alamat" query:"alamat"`
	Rtrw             string `json:"rtrw" form:"rtrw" query:"rtrw"`
	Pekerjaan        string `json:"pekerjaan" form:"pekerjaan" query:"pekerjaan"`
	StatusPerkawinan string `json:"statusPerkawinan" form:"statusPerkawinan" query:"statusPerkawinan"`
	Kewarganegaraan  string `json:"kewarganegaraan" form:"kewarganegaraan" query:"kewarganegaraan"`
	CreateBy         string `json:"createBy" form:"createBy" query:"createBy"`
}

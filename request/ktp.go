package request

import "time"

type CreateKtp struct {
	Nik              string     `json:"nik" form:"nik" query:"nik" validate:"required"`
	Nama             string     `json:"nama" form:"nama" query:"nama" validate:"required"`
	TempatLahir      string     `json:"tempatLahir" form:"tempatLahir" query:"tempatLahir" validate:"required"`
	TanggalLahir     time.Time  `json:"tanggalLahir" form:"tanggalLahir" query:"tanggalLahir" validate:"required"`
	JenisKelamin     string     `json:"jenisKelamin" form:"jenisKelamin" query:"jenisKelamin" validate:"required"`
	Alamat           string     `json:"alamat" form:"alamat" query:"alamat" validate:""`
	Rtrw             string     `json:"rtrw" form:"rtrw" query:"rtrw" validate:""`
	KelurahanDesa    string     `json:"kelurahanDesa" form:"kelurahanDesa" query:"kelurahanDesa" validate:""`
	Kecamatan        string     `json:"kecamatan" form:"kecamatan" query:"kecamatan" validate:""`
	KabupatenKota    string     `json:"kabupatenKota" form:"kabupatenKota" query:"kabupatenKota" validate:""`
	Provinsi         string     `json:"provinsi" form:"provinsi" query:"provinsi" validate:""`
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
	Alamat           string     `json:"alamat" form:"alamat" query:"alamat" validate:""`
	Rtrw             string     `json:"rtrw" form:"rtrw" query:"rtrw" validate:""`
	KelurahanDesa    string     `json:"kelurahanDesa" form:"kelurahanDesa" query:"kelurahanDesa" validate:""`
	Kecamatan        string     `json:"kecamatan" form:"kecamatan" query:"kecamatan" validate:""`
	KabupatenKota    string     `json:"kabupatenKota" form:"kabupatenKota" query:"kabupatenKota" validate:""`
	Provinsi         string     `json:"provinsi" form:"provinsi" query:"provinsi" validate:""`
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
	Alamat           string `json:"alamat" form:"alamat" query:"alamat"`
	KelurahanDesa    string `json:"kelurahanDesa" form:"kelurahanDesa" query:"kelurahanDesa"`
	Kecamatan        string `json:"kecamatan" form:"kecamatan" query:"kecamatan"`
	KabupatenKota    string `json:"kabupatenKota" form:"kabupatenKota" query:"kabupatenKota"`
	Provinsi         string `json:"provinsi" form:"provinsi" query:"provinsi"`
	StatusPerkawinan string `json:"statusPerkawinan" form:"statusPerkawinan" query:"statusPerkawinan"`
	Kewarganegaraan  string `json:"kewarganegaraan" form:"kewarganegaraan" query:"kewarganegaraan"`
	CreateBy         string `json:"createBy" form:"createBy" query:"createBy"`
}

package models

type RequestDetailPasar struct {
	PasarID	string	`json:"pasar"`
}

type RequestGetAllPasar struct {
	Limit	string	`json:"limit"`
}

type PasarTable struct {
	PasarID        string `json:"pasar" gorm:"column:id"`
	NamaPasar      string `json:"nama_pasar" gorm:"column:nama_pasar"`
	Alamat         string `json:"alamat" gorm:"column:alamat"`
	BackgroudColor string `json:"background_color" gorm:"column:backgroundcolor"`
	BranchCode     string `json:"branch_code" gorm:"column:branchcode"`
	Latitude       string `json:"latitude" gorm:"column:lat"`
	Longitude      string `json:"longitude" gorm:"column:lon"`
	Image          string `json:"image" gorm:"column:image"`
	UrlSid         string `json:"url" gorm:"column:url"`
	Kota           string `json:"kota" gorm:"column:kota"`
	Provinsi       string `json:"provinsi" gorm:"column:provinsi"`
	Kecamatan      string `json:"kecamatan" gorm:"column:kecamatan"`
	Subdomain      string `json:"subdomain" gorm:"column:subdomain"`
	Deskripsi      string `json:"deskripsi" gorm:"column:deskripsi"`
	Branchname     string `json:"branchname" gorm:"column:branchname"`
	Logo1          string `json:"logo1" gorm:"column:logo1"`
	Logo2          string `json:"logo2" gorm:"column:logo2"`
	Logo3          string `json:"logo3" gorm:"column:logo3"`
	Logo4          string `json:"logo4" gorm:"column:logo4"`
}

type DetailGetAllPasar struct {
	PasarID   string `json:"pasar"`
	NamaPasar string `json:"nama_pasar"`
	Alamat    string `json:"alamat"`
	Image     string `json:"image"`
	UrlSid    string `json:"url"`
	Kecamatan string `json:"kecamatan"`
	Kota      string `json:"kota"`
	Provinsi  string `json:"provinsi"`
}

type ResponseDetailPasar struct {
	NamaPasar 	string `json:"nama_pasar"`
	Alamat		string `json:"alamat"`
	Deskripsi	string `json:"deskripsi"`
	Distance	int `json:"distance"`
}

type ResponseSuccessGetDetailPasar struct {
	ResponseCode	string				`json:"RC"`
	Deskripsi		string				`json:"Deskripsi"`
	Data			ResponseDetailPasar `json:"Data"`
}

type ResponseSuccessGetAllPasar struct {
	ResponseCode	string				`json:"RC"`
	Deskripsi		string				`json:"Deskripsi"`
	Data			[]DetailGetAllPasar	`json:"Data"`
}


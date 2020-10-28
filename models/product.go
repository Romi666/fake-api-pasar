package models



type ProductTable struct{
	ProductID		string	`json:"id"`
	NamaProduk		string	`json:"nama_produk"`
	NamaPedagang	string	`json:"nama_pedagang"`
	AlamatPedagang	string	`json:"alamat_pedagang"`
	NoHP			string	`json:"no_hp"`
	Stock			float64	`json:"stock"`
	Harga			float64	`json:"harga"`
	Picture			string	`json:"picture"`
	Unit			string	`json:"unit"`
}

type HTTPResult struct {
	RC			string	`json:"RC"`
	Deskripsi	string	`json:"Deskripsi"`
	Data		[]ProductTable `json:"Data"`
}

type ResponseSuccessGetAllProduct struct {
	RC			string	`json:"RC"`
	Deskripsi	string	`json:"Deskripsi"`
	Data		[]ProductTable `json:"Data"`
}


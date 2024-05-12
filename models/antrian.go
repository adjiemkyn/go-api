package models

type Antrian struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	Nama      string `gorm:"type:varchar(255)" json:"nama_product"`
	Usia      string `gorm:"type:varchar(255)" json:"deskripsi"`
	Alamat    string `gorm:"type:varchar(255)" json:"alamat"`
	Tgl_lahir uint16 `gorm:"type:date" json:"tgl_lahir"`
	No_tlp    string `gorm:"varchar(15)" json:"no_tlp"`
}

package package_ats

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type pembeli struct {
	ID          	 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama        	 string             `bson:"nama,omitempty" json:"nama,omitempty"`
	TanggalLahir	 string             `bson:"tanggallahir,omitempty" json:"tanggallahir,omitempty"`
	Alamat      	 string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	nama_makanan     string             `bson:"institusi,omitempty" json:"institusi,omitempty"`
	jenis_makanan	 string             `bson:"bidang_studi,omitempty" json:"bidang_studi,omitempty"`
	//Publikasi    Publikasi          `bson:"publikasi,omitempty" json:"publikasi,omitempty"`
}

// Publikasi adalah struktur data untuk menyimpan informasi tentang publikasi.
// type Publikasi struct {
//	Judul    string `bson:"judul,omitempty" json:"judul,omitempty"`
//	Tanggal  time.Time `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
//	Penulis  string `bson:"penulis,omitempty" json:"penulis,omitempty"`
//	Category string `bson:"category,omitempty" json:"category,omitempty"`
//}

// makanan hewan adalah struktur data untuk menyimpan informasi tentang makanan hewan.
type makanan_hewan struct {
	ID       	 	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	nama_makanan    string             `bson:"judul,omitempty" json:"judul,omitempty"`
	jenis_makanan 	string             `bson:"institusi,omitempty" json:"institusi,omitempty"`
	expired  	 	string             `bson:"penulis,omitempty" json:"penulis,omitempty"`
	Datetime  	 	primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	komentar 	 	string             `bson:"ringkasan,omitempty" json:"ringkasan,omitempty"`
//	 string           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
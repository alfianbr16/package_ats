package package_ats_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/alfianbr16/package_ats/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertMakanan_hewan(t *testing.T) {
	biodata := map[string]interface{}{
		"nama_penjual":  	 "mahfud amin",
		"tanggal_lahir": 	 "27 08 1945",
		"alamat":        	 "planet bekasi",
		"jenis_makanan": 	 "kucing",
		"nama_makanan":	 	 "whiskase",
		"pembayaran":	 	 "cash",
		"expired":	     	 "23 01 1975",
		"tanggal_pembelian": "09 03 1973",
		"pembeli":	     	 "ganjar pranowo",
		"komentar":	     	 "kucing saya suka sekali dengan makanan ini",
	}

	insertedID := package_ats.InsertMakanan_hewan(
		"Dampak dari Krisis Air Bersih",
		"Universitas Ganjil Genap",
		"Markonah",
		primitive.NewDateTimeFromTime(time.Now()),
		"Dengan menggunakan pendekatan interdisipliner,menyadarkan masyarakat akan isu air bersih dan perlunya tindakan yang cepat dan efektif untuk menjaga sumber daya air yang vital bagi kehidupan dan keberlanjutan lingkungan",
		biodata,
	)
	fmt.Println("Inserted ID Makanan :", insertedID)
}

func TestGetAllMakanan_hewan(t *testing.T) {
	makanan_hewans := package_ats.GetAllMakanan_hewan()
	fmt.Println("All Makanan hewan:", makanan_hewans)
}

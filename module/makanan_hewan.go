package package_ats

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertMakanan_hewan(nama_makanan string, jenis_makanan string, pembeli string, datetime primitive.DateTime, komentar string, biodata map[string]interface{}) (insertedID interface{}) {
	var makanan_hewan = map[string]interface{}{
		"_id":       primitive.NewObjectID(),
		"nama_makanan":nama_makanan,
		"jenis_makanan": jenis_makanan,
		"pembeli":   pembeli,
		"datetime":  datetime,
		"komentar": komentar,
		"biodata":   biodata,
	}
	return InsertOneDoc("ATS", "makanan_hewan", makanan_hewan)
}

func GetAllMakanan_hewan() []interface{} {
	var makanan_hewans []interface{}
	collection := MongoConnect("ATS").Collection("makanan_hewan")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllmakanan_hewan:", err)
		return makanan_hewans
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var makanan_hewan interface{}
		err := cursor.Decode(&makanan_hewan)
		if err != nil {
			fmt.Printf("Error decoding makanan hewan: %v\n", err)
			continue
		}
		makanan_hewans = append(makanan_hewans, makanan_hewan)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("Cursor error: %v\n", err)
	}
	return makanan_hewans
}

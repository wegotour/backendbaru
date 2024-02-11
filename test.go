package backendbaru

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/wegotour/backendbaru/model"
	"github.com/wegotour/backendbaru/module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
)

var db = module.MongoConnect("MONGOSTRING", "serbaevent_db")

func TestGetUserFromEmail(t *testing.T) {
	email := "admin@gmail.com"
	hasil, err := module.GetUserFromEmail(email, db)
	if err != nil {
		t.Errorf("Error TestGetUserFromEmail: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

// Insert-Tiket
func TestInsertOneTiket(t *testing.T) {
	var doc model.Tiket
	doc.TujuanEvent = "Event Coldplay 10"
	doc.Jemputan = "Terminal Mangga Sari jakarta timur st.12 jalan soekarno hatta"
	doc.Keterangan = "Jam Jemputan 15:00"
	doc.Harga = "RP 120.0000"
	if doc.TujuanEvent == "" || doc.Jemputan == "" || doc.Keterangan == "" || doc.Harga == "" {
		t.Errorf("mohon untuk melengkapi data")
	} else {
		insertedID, err := module.InsertOneDoc(db, "tiket", doc)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil disimpan")
		} else {
			fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
		}
	}
}

type Userr struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email string             `bson:"email,omitempty" json:"email,omitempty"`
	Role  string             `bson:"role,omitempty" json:"role,omitempty"`
}

func TestGetAllDoc(t *testing.T) {
	hasil := module.GetAllDocs(db, "user", []Userr{})
	fmt.Println(hasil)
}

func TestInsertUser(t *testing.T) {
	var doc model.User
	doc.Email = "admin@gmail.com"
	password := "admin123"
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		t.Errorf("kesalahan server : salt")
	} else {
		hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
		user := bson.M{
			"email":    doc.Email,
			"password": hex.EncodeToString(hashedPassword),
			"salt":     hex.EncodeToString(salt),
			"role":     "admin",
		}
		_, err = module.InsertOneDoc(db, "user", user)
		if err != nil {
			t.Errorf("gagal insert")
		} else {
			fmt.Println("berhasil insert")
		}
	}
}

func TestGetUserByAdmin(t *testing.T) {
	id := "6569a026a943657839880665"
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Errorf("Error converting id to objectID: %v", err)
	}
	data, err := module.GetUserFromID(idparam, db)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		if data.Role == "pengguna" {
			datapengguna, err := module.GetPenggunaFromAkun(data.ID, db)
			if err != nil {
				t.Errorf("Error getting document: %v", err)
			} else {
				datapengguna.Akun = data
				fmt.Println(datapengguna)
			}
		}
	}
}

func TestSignUpPengguna(t *testing.T) {
	var doc model.Pengguna
	doc.NamaLengkap = "Sahijatea"
	doc.TanggalLahir = "30/08/2004"
	doc.JenisKelamin = "Perempuan"
	doc.NomorHP = "081234567890"
	doc.Alamat = "Wastukencana Blok No 32"
	doc.Akun.Email = "sahjatea@gmail.com"
	doc.Akun.Password = "sahijabandung"
	err := module.SignUpPengguna(db, doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
		fmt.Println("Data berhasil disimpan dengan nama :", doc.NamaLengkap)
	}
}

func TestLogIn(t *testing.T) {
	var doc model.User
	doc.Email = "wawan@gmail.com"
	doc.Password = "driverwawan"
	user, err := module.LogIn(db, doc)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		fmt.Println("Selamat datang Driver:", user)
	}
}

func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := module.GenerateKey()
	fmt.Println("ini private key :", privateKey)
	fmt.Println("ini public key :", publicKey)
	id := "6569a026a943657839880665"
	objectId, err := primitive.ObjectIDFromHex(id)
	role := "pengguna"
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	hasil, err := module.Encode(objectId, role, privateKey)
	fmt.Println("ini hasil :", hasil, err)
}

func TestUpdatePengguna(t *testing.T) {
	var doc model.Pengguna
	id := "6569adb773f4bc6c2069c0ed"
	objectId, _ := primitive.ObjectIDFromHex(id)
	id2 := "6569adb673f4bc6c2069c0eb"
	userid, _ := primitive.ObjectIDFromHex(id2)
	doc.NamaLengkap = "Shayeza aselole"
	doc.TanggalLahir = "30/08/2003"
	doc.JenisKelamin = "Perempuan"
	doc.NomorHP = "081234567890"
	doc.Alamat = "Wastukencana Blok No 32"
	if doc.NamaLengkap == "" || doc.TanggalLahir == "" || doc.JenisKelamin == "" || doc.NomorHP == "" || doc.Alamat == "" {
		t.Errorf("mohon untuk melengkapi data")
	} else {
		err := module.UpdatePengguna(objectId, userid, db, doc)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil diupdate")
		} else {
			fmt.Println("Data berhasil diupdate")
		}
	}
}

func TestWatoken(t *testing.T) {
	body, err := module.Decode("fca3dbba6c382d6e937d33837f7428c1211e01a9928cbbbc0b86bb8351c02407", " v4.public.eyJleHAiOiIyMDIzLTEyLTAxVDE4OjU4OjE1KzA4OjAwIiwiaWF0IjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsImlkIjoiNjU1YzNiOWExZDY1MjRmMmYxMjAwZmM2IiwibmJmIjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsInJvbGUiOiJwZW5nZ3VuYSJ9GIKgKcp8gj4lzPH_NFvpx3GR2kBZ2qsDquYMKQdQ1PFpvHKlDy-FeO1umIGCaMuYyACP5jd-Y0at1WCOrsNRCA")
	fmt.Println("isi : ", body, err)
}


// test Tiket
func TestInsertTiket(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "serbaevent_db")
	payload, err := module.Decode("fca3dbba6c382d6e937d33837f7428c1211e01a9928cbbbc0b86bb8351c02407", "v4.public.eyJleHAiOiIyMDIzLTEyLTAxVDE4OjU4OjE1KzA4OjAwIiwiaWF0IjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsImlkIjoiNjU1YzNiOWExZDY1MjRmMmYxMjAwZmM2IiwibmJmIjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsInJvbGUiOiJwZW5nZ3VuYSJ9GIKgKcp8gj4lzPH_NFvpx3GR2kBZ2qsDquYMKQdQ1PFpvHKlDy-FeO1umIGCaMuYyACP5jd-Y0at1WCOrsNRCA")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	var datatiket model.Tiket
	datatiket.TujuanEvent = "Event Coldplay 5 Jakarta"
	datatiket.Jemputan = "Terminal Bus Jakarta"
	datatiket.Keterangan = "Jemputan 15:00"
	datatiket.Harga = "Rp 120.000"
	err = module.InsertTiket(payload.Id, conn, datatiket)
	if err != nil {
		t.Errorf("Error insert : %v", err)
	} else {
		fmt.Println("Success!!!")
	}
}

func TestUpdateTiket(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "serbaevent_db")
	payload, err := module.Decode("fca3dbba6c382d6e937d33837f7428c1211e01a9928cbbbc0b86bb8351c02407", "v4.public.eyJleHAiOiIyMDIzLTEyLTAxVDE4OjU4OjE1KzA4OjAwIiwiaWF0IjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsImlkIjoiNjU1YzNiOWExZDY1MjRmMmYxMjAwZmM2IiwibmJmIjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsInJvbGUiOiJwZW5nZ3VuYSJ9GIKgKcp8gj4lzPH_NFvpx3GR2kBZ2qsDquYMKQdQ1PFpvHKlDy-FeO1umIGCaMuYyACP5jd-Y0at1WCOrsNRCA")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	if payload.Role != "admin" {
		t.Errorf("Error role: %v", err)
	}
	var datatiket model.Tiket
	datatiket.TujuanEvent = "Event Coldplay 3 surabaya"
	datatiket.Jemputan = "Terminal bus surabaya "
	datatiket.Keterangan = "jam jemputan 13:00"
	datatiket.Harga = "Rp 100.000"
	id := "6569a53d783c6970079a560b"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.UpdateTiket(objectId, payload.Id, conn, datatiket)
	if err != nil {
		t.Errorf("Error update : %v", err)
	} else {
		fmt.Println("Success!!!")
	}
}

func TestDeleteTiket(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "serbaevent_db")
	payload, err := module.Decode("fca3dbba6c382d6e937d33837f7428c1211e01a9928cbbbc0b86bb8351c02407", "v4.public.eyJleHAiOiIyMDIzLTEyLTAxVDE4OjU4OjE1KzA4OjAwIiwiaWF0IjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsImlkIjoiNjU1YzNiOWExZDY1MjRmMmYxMjAwZmM2IiwibmJmIjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsInJvbGUiOiJwZW5nZ3VuYSJ9GIKgKcp8gj4lzPH_NFvpx3GR2kBZ2qsDquYMKQdQ1PFpvHKlDy-FeO1umIGCaMuYyACP5jd-Y0at1WCOrsNRCA")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	id := "6569a53d783c6970079a560b"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.DeleteTiket(objectId, payload.Id, conn)
	if err != nil {
		t.Errorf("Error delete : %v", err)
	} else {
		fmt.Println("Success!!!")
	}
}

func TestGetAllTiket(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "serbaevent_db")
	data, err := module.GetAllTiket(conn)
	if err != nil {
		t.Errorf("Error get all : %v", err)
	} else {
		fmt.Println(data)
	}
}

func TestGetTiketFromID(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "serbaevent_db")
	id := "6569a025a943657839880661"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	tiket, err := module.GetTiketFromID(objectId, conn)
	if err != nil {
		t.Errorf("Error get Tiket : %v", err)
	} else {
		fmt.Println(tiket)
	}
}

func TestReturnStruct(t *testing.T) {
	id := "11b98454e034f3045021a8aa8eb84280"
	objectId, _ := primitive.ObjectIDFromHex(id)
	user, _ := module.GetUserFromID(objectId, db)
	data := model.User{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
	hasil := module.GCFReturnStruct(data)
	fmt.Println(hasil)
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Salt     string             `bson:"salt,omitempty,omitempty" json:"salt,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
}

type Password struct {
	Password    string `bson:"password,omitempty" json:"password,omitempty"`
	Newpassword string `bson:"newpass,omitempty" json:"newpass,omitempty"`
}

type Pengguna struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaLengkap  string             `bson:"namalengkap,omitempty" json:"namalengkap,omitempty"`
	TanggalLahir string             `bson:"tanggallahir,omitempty" json:"tanggallahir,omitempty"`
	JenisKelamin string             `bson:"jeniskelamin,omitempty" json:"jeniskelamin,omitempty"`
	NomorHP      string             `bson:"nomorhp,omitempty" json:"nomorhp,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Akun         User               `bson:"akun,omitempty" json:"akun,omitempty"`
}

type Admin struct {
	ID   primitive.ObjectID `bson:"_id,omitempty,omitempty" json:"_id,omitempty"`
	Akun User               `bson:"akun,omitempty" json:"akun,omitempty"`
}

type Tiket struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaTicket  string             `bson:"namaticket,omitempty" json:"namaticket,omitempty"`
	Harga       string             `bson:"harga,omitempty" json:"harga,omitempty"`
	NamaPembeli string             `bson:"namapembeli,omitempty" json:"namapembeli,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Alamat      string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	NoHP        string             `bson:"nohp,omitempty" json:"nohp,omitempty"`
	Quantity    string             `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Total       string             `bson:"total,omitempty" json:"total,omitempty"`
}


type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Role    string `json:"role,omitempty" bson:"role,omitempty"`
}

type Response struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Payload struct {
	Id   primitive.ObjectID `json:"id"`
	Role string             `json:"role"`
	Exp  time.Time          `json:"exp"`
	Iat  time.Time          `json:"iat"`
	Nbf  time.Time          `json:"nbf"`
}

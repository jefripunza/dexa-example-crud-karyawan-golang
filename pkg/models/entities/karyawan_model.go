package models

type Karyawan struct {
	ID            int    `db:"id"             json:"id"`
	CreateDate    string `db:"create_date"    json:"create_date"`
	PhoneNumber   string `db:"phone_number"   json:"phone_number"`
	Email         string `db:"email"          json:"email"`
	Password      string `db:"password"       json:"password"`
	Name          string `db:"name"           json:"name"`
	DOB           string `db:"dob"            json:"dob"`
	Address       string `db:"address"        json:"address"`
	Religion      string `db:"religion"       json:"religion"`
	Gender        string `db:"gender"         json:"gender"`
	MaritalStatus string `db:"marital_status" json:"marital_status"`
}

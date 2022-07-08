package queries

import (
	dto "dexa-training-crud-karyawan/pkg/DTO"
	"dexa-training-crud-karyawan/pkg/helpers"
	entities "dexa-training-crud-karyawan/pkg/models/entities"

	"math"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// KaryawanQueries struct for queries from Karyawan model.
type KaryawanQueries struct {
	*sqlx.DB
}

// query untuk menambahkan karyawan baru
func (q *KaryawanQueries) CreateKaryawan(karyawan *entities.Karyawan) error {
	// Define query string.
	query := `INSERT INTO karyawan (phone_number, email, password, name, dob, address, religion, gender, marital_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Send query to database.
	_, err := q.Exec(query, karyawan.PhoneNumber, karyawan.Email, karyawan.Password, karyawan.Name, karyawan.DOB, karyawan.Address, karyawan.Religion, karyawan.Gender, karyawan.MaritalStatus)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// query untuk menampilkan list data karyawan (pagination)
func (q *KaryawanQueries) GetKaryawanPagination(show int, page int, order_by string, keyword string, keyword_skip bool) ([]entities.Karyawan, int, int, int, int, error) {
	// Define karyawan variable.
	var karyawan []entities.Karyawan

	test, err := q.Query("SELECT * FROM karyawan")
	if err != nil {
		// Return empty object and error.
		return karyawan, 0, 0, 0, 0, err
	}
	TotalData := 0
	for test.Next() {
		TotalData += 1
	}
	// fmt.Println(TotalData)

	// Define query string.
	query := "SELECT id, create_date, phone_number, email, password, name, dob, address, religion, gender, marital_status FROM karyawan "
	if keyword_skip == false {
		query = query + " where CONCAT_WS(phone_number, email, name, address, religion, gender, marital_status) like '%" + keyword + "%' "
	}
	query = query + " order by id " + order_by + " LIMIT " + strconv.Itoa((page-1)*show) + ", " + strconv.Itoa(show) + " "
	// fmt.Println(query)

	// Send query to database.
	rows, err := q.Query(query)
	if err != nil {
		// Return empty object and error.
		return karyawan, 0, 0, 0, 0, err
	}

	// iterate over each row
	for rows.Next() {
		var rowKaryawan entities.Karyawan
		err = rows.Scan(&rowKaryawan.ID, &rowKaryawan.CreateDate, &rowKaryawan.PhoneNumber, &rowKaryawan.Email, &rowKaryawan.Password, &rowKaryawan.Name, &rowKaryawan.DOB, &rowKaryawan.Address, &rowKaryawan.Religion, &rowKaryawan.Gender, &rowKaryawan.MaritalStatus)
		karyawan = append(karyawan, rowKaryawan)
	}

	LastPage := int(math.Ceil(float64(TotalData / show)))
	NextPage := helpers.If(page+1 > LastPage, 0, page+1)
	// const PrevPage = page - 1 < 1 ? null : page - 1;
	PrevPage := helpers.If(page-1 < 1, 0, page-1)
	PrevPage = helpers.If(PrevPage > LastPage, LastPage, PrevPage)

	// Return query result.
	return karyawan, TotalData, LastPage, NextPage, PrevPage, nil
}

func (q *KaryawanQueries) UpdateKaryawan(id int, karyawan *entities.Karyawan) error {
	// Define query string.
	query := `UPDATE karyawan SET phone_number=?, email=?, password=?, name=?, dob=?, address=?, religion=?, gender=?, marital_status=? WHERE id=?`

	// Send query to database.
	_, err := q.Exec(query, karyawan.PhoneNumber, karyawan.Email, karyawan.Password, karyawan.Name, karyawan.DOB, karyawan.Address, karyawan.Religion, karyawan.Gender, karyawan.MaritalStatus, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (q *KaryawanQueries) DeleteKaryawan(payload *dto.RequestDelete) error {
	// Define query string.
	query := `DELETE FROM karyawan WHERE id=?`

	// Send query to database.
	_, err := q.Exec(query, payload.ID)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

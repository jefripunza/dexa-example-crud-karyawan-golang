package controllers

import (
	dto "dexa-training-crud-karyawan/pkg/DTO"
	message "dexa-training-crud-karyawan/pkg/message"
	entities "dexa-training-crud-karyawan/pkg/models/entities"
	utils "dexa-training-crud-karyawan/pkg/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// menambahkan karyawan baru
func CreateKaryawan(c *fiber.Ctx) error {
	// setup DTO
	karyawan := &entities.Karyawan{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(karyawan); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Create database connection.
	db, err := utils.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// execute!
	if err := db.CreateKaryawan(karyawan); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": message.SaveSuccess,
	})
}

// CreateBook func for creates a new book.
func ShowKaryawanPagination(c *fiber.Ctx) error {
	// ambil parameter dari endpoint
	show, err := strconv.Atoi(c.Params("show"))
	if err != nil {
		if show <= 0 {
			show = 20 // default
		}
	}
	// ---------------------------------------- //
	page, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		page = 1 // default
	}
	if page <= 0 {
		page = 1 // default
	}
	// ---------------------------------------- //
	order := c.Params("order_by")
	order_by := "asc" // default
	if order == "desc" {
		order_by = "desc"
	}
	// ---------------------------------------- //
	keyword := c.Params("keyword")
	keyword_skip := false
	if keyword == "'" {
		keyword_skip = true
	}
	// fmt.Println(show, page, order_by, keyword, keyword_skip) // for debug

	// Create database connection.
	db, err := utils.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// init jika error maka di handle ke default
	pagination := &dto.Pagination{}

	// Get all karyawan
	Rows, TotalData, LastPage, NextPage, PrevPage, err := db.GetKaryawanPagination(show, page, order_by, keyword, keyword_skip)
	if err != nil {
		// Return, if books not found.
		return c.Status(fiber.StatusNotFound).JSON(pagination)
	}

	// DTO Builder
	pagination.Data = Rows
	pagination.TotalData = TotalData
	pagination.CurrentPage = page
	pagination.LastPage = LastPage
	pagination.NextPage = NextPage
	pagination.PrevPage = PrevPage

	// Return status 200 OK.
	return c.JSON(pagination)
}

// CreateBook func for creates a new book.
func UpdateKaryawan(c *fiber.Ctx) error {
	// setup DTO
	karyawan := &entities.Karyawan{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(karyawan); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Create database connection.
	db, err := utils.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Update book by given ID.
	if err := db.UpdateKaryawan(karyawan.ID, karyawan); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": message.UpdateDataSuccess,
	})
}

// CreateBook func for creates a new book.
func DeleteKaryawan(c *fiber.Ctx) error {
	payload := &dto.RequestDelete{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&payload); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Create database connection.
	db, err := utils.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Update book by given ID.
	if err := db.DeleteKaryawan(payload); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	fmt.Println(err)

	// Return status 200 OK.
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": message.DeleteDataSuccess,
	})
}

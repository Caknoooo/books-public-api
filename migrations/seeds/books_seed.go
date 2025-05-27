package seeds

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

func BookSeeder(db *gorm.DB) error {
	csvFile, err := os.Open("./migrations/csv/books.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.TrimLeadingSpace = true

	// Baca semua baris
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Lewati header
	if len(records) <= 1 {
		return errors.New("no book data found")
	}

	// Cek dan buat tabel jika belum ada
	hasTable := db.Migrator().HasTable(&entity.Books{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Books{}); err != nil {
			return err
		}
	}

	for i, row := range records[1:] {
		if len(row) < 7 {
			continue // skip data yang tidak lengkap
		}

		// Ambil dan parsing data
		title := row[0]
		authors := row[1]
		description := row[2]
		categoriesStr := row[3]
		publisher := row[4]
		publishDateStr := row[5]
		priceStr := row[6]

		// Parsing tanggal
		publishDate, err := parseDate(publishDateStr)
		if err != nil {
			log.Printf("row %d: warning - %v, using zero date", i+2, err)
		}

		// Parsing harga
		price, err := parsePrice(priceStr)
		if err != nil {
			return errors.New("error parsing price on row " + strconv.Itoa(i+2) + ": " + err.Error())
		}

		// Parsing kategori
		categories := parseCategories(categoriesStr)
		categoriesJSON, err := json.Marshal(categories)
		if err != nil {
			return errors.New("error parsing categories on row " + strconv.Itoa(i+2) + ": " + err.Error())
		}

		book := entity.Books{
			Title:       title,
			Author:      authors,
			Description: description,
			Categories:  categoriesJSON,
			Publisher:   publisher,
			PublishDate: publishDate,
			Price:       price,
		}

		var existing entity.Books
		err = db.Where("title = ? AND author = ?", title, authors).First(&existing).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&book).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// parseDate mencoba beberapa format tanggal dan fallback jika gagal
func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return time.Time{}, nil // kosong = nol waktu
	}

	formats := []string{
		"Monday, January 2, 2006",
		"January 2, 2006",
		"January 2006",
		"January 2",
		"2006-01-02",
		"02 Jan 2006",
		"January 2 2006",
		"January",         // <-- menangani "December 1899"
		"January 2006",    // fallback lain
		"January 2, 1899", // tambahan jika pakai tahun lama
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	// kalau semua format gagal, return waktu nol dan log
	return time.Time{}, errors.New("invalid date format: " + dateStr)
}

// parsePrice mengubah string "$4.99" menjadi int dalam Rupiah (17.000x)
func parsePrice(priceStr string) (float64, error) {
	priceStr = strings.ReplaceAll(priceStr, "Price Starting at $", "")
	priceStr = strings.TrimSpace(priceStr)

	priceFloat, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}

	return priceFloat * 17000, nil
}

// parseCategories memecah kategori CSV menjadi slice string
func parseCategories(s string) []string {
	if s == "" {
		return []string{}
	}
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

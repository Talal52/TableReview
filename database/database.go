package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Review struct {
    Id          int    `json:"id"`
    Firstname   string `json:"firstname"`
    Lastname    string `json:"lastname"`
    Email       string `json:"email"`
    Description string `json:"description"`
    Content     string `json:"content"`
}


type DB struct {
	*sql.DB
}

func InitDB() (*DB, error) {
	os.Unsetenv("PGLOCALEDIR")
	db, err := sql.Open("postgres", "user=postgres password=12345 dbname=MyDatabase sslmode=disable")
	if err != nil {
		fmt.Println("error 1", err)
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) AddReview(newReview Review) error {
    _, err := db.Exec("INSERT INTO reviews ("id", \"firstname\", \"lastname\", \"email\", \"description\", \"content\") VALUES ($1, $2, $3, $4, $5, $6)",
        newReview.Id, newReview.Firstname, newReview.Lastname, newReview.Email, newReview.Description, newReview.Content)

    if err != nil {
        fmt.Println("error adding review:", err)
        return err
    }

    return nil
}


































// func (db *DB) GetUserByEmail(email string) (Review, error) {
// 	var review Review
// 	err := db.QueryRow("SELECT \"Id\", \"firstname\", \"lastname\", \"email\", \"description\", \"content\" FROM \"reviews\" WHERE \"email\" = $1", email).Scan(&review.Id, &review.Firstname, &review.Lastname, &review.Email, &review.Description,&review.Content)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return Review{}, err
// 	}
// 	return review, nil
// }

// Review struct definition...

// func (db *DB) GetReviewByEmail(email string) (Review, error) {
// 	var review Review
// 	err := db.QueryRow("SELECT id, firstname, lastname, email, description, content FROM reviews WHERE email = $1", email).
// 		Scan(&review.Id, &review.Firstname, &review.Lastname, &review.Email, &review.Description, &review.Content)

// 	if err != nil {
// 		fmt.Println("error getting review:", err)
// 		return Review{}, err
// 	}

// 	return review, nil
// }

// func (db *DB) DeleteReviewByEmail(email string) error {
// 	_, err := db.Exec("DELETE FROM reviews WHERE email = $1", email)
// 	if err != nil {
// 		fmt.Println("error deleting review:", err)
// 		return err
// 	}

// 	return nil
// }

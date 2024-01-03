//ini adalah bagian repository
package content

import (
	"hexa/business/content"

	"gorm.io/gorm"
)

// import "firebase.google.com/go/v4/db"

type PostgresRepository struct { //parents dependency
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository { //func dependency injection
	return &PostgresRepository{
		db: db,
	}
}

//bikin func seperti service untuk pengembalian si PostgreRepository
func (repo *PostgresRepository) FindContentbyID(id int) (content *content.Content, err error) {
	return content, nil
}

func (repo *PostgresRepository) FindAll() (contents []content.Content, err error) {
	result := repo.db.Find(&contents)
	if result.Error != nil {
		return nil, result.Error
	}
	return contents, nil
}

func (repo *PostgresRepository) InsertContent(content content.Content) (err error) {
	return
}

func (repo *PostgresRepository) UpdateContent(content content.Content, currentVersion int) (err error) {
	return
}

// func (repo *PostgresRepository) FindAll() (contents []content.Content, err error) {
// 	return
// }

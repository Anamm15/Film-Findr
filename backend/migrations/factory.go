package migrations

import (
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := ListUserSeeder(db); err != nil {
		return err
	}

	// if err := ListGenreSeeder(db); err != nil {
	// 	return err
	// }

	return nil
}

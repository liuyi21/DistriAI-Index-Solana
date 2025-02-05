package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.Migrator().DropTable(
		&Machine{},
		&Order{},
		&Reward{},
		&RewardMachine{},
	)
	db.AutoMigrate(
		&Log{},
		&Mailbox{},
		&Machine{},
		&Order{},
		&Reward{},
		&RewardMachine{},
	)
}

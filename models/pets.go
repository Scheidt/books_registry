package models

import "gorm.io/gorm"

type Pet struct{
	ID    			uint		`gorm:"primaryKey; autoIncrement" json: "id"`
	name    		*string		`json: "name"`
	age  			*int		`json: "age"`
	owner   		*string		`json: "owner"`
	size    		*string		`json: "size"`
	weight			*float32	`json: "weight"`
	paidThisMonth 	*bool		`json: "paid"`
}

func migratePets(db *gorm.DB) error{
	err:= db.AutoMigrate(&Pet{})
	return err
}

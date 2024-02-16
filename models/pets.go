package models

import "gorm.io/gorm"

type Pet struct{
	ID    			uint		`gorm:"primaryKey; autoIncrement" json: "id"`
	Name    		*string		`json: "name"`
	Age  			int			`json: "age"`
	Owner   		string		`json: "owner"`
	Size    		string		`json: "size"`
	Weight			float32		`json: "weight"`
	PaidThisMonth 	bool		`json: "paid"`
}

func MigratePets(db *gorm.DB) error{
	err:= db.AutoMigrate(&Pet{})
	return err
}

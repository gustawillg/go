package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//create
	// db.Create(&Product{
	// 	Name:  "test",
	// 	Price: 1000.00,
	// })

	//create batch
	// products := []Product{
	// 	{Name: "test", Price: 1000}
	// 	{Name: "test2", Price: 13000}
	// 	{Name: "test3", Price: 41000}
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "test")
	// fmt.Println(product)

	//select all
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//where
	// var products []Product
	// db.Where("name LIKE > ?", "%k").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//busca
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)

}

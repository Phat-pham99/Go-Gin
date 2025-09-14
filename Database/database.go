package DataBase

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(
		sqlite.Open("test.db"),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()
	// Migrate the schema
	db.AutoMigrate(&Product{})
	// var result int
	// Create
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "A69", Price: 200})

	// Read
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)       // find product with integer primary key
	products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // find product with code D42

	// Update - update product's price to 200
	result, err := gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 300)
	// Update - update multiple fields
	// result, err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// result, err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
	fmt.Println("result", result)
	fmt.Print("products", products)

}

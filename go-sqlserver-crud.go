
package main

import(
"fmt"
"gorm.io/gorm"
"gorm.io/driver/sqlserver"
"time"
)

// How to install gorm? go get -u gorm.io/gorm

type Product struct {
	gorm.Model 
	Code string `gorm:"column:code"`
	Price uint  `gorm:"column:price"`
	}

	var (
		// for windows authentication
		connectionString = "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
		// for SQL Server authentication
        //connectionString = "sqlserver://username:password@host:port?database=nameDB"
	)

 func main() { 
	// createProduct()
	//updateProduct()
	// deleteProduct()
	getProduct()
	getProductList()
}

func createProduct(){
	fmt.Println("createProduct")
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	   if err != nil {
		   fmt.Println("failed to connect database") 
		   panic("failed to connect database")
	   }
	  
	db.Exec("INSERT INTO products (created_at,code,price) VALUES (?,?,?)",time.Now(), "A01",100)
	db.Exec("INSERT INTO products (created_at,code,price) VALUES (?,?,?)",time.Now(), "A02",200)
}

func updateProduct(){
	fmt.Println("updateProduct")
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database") 
		panic("failed to connect database")
	}
   
   db.Exec("UPDATE products SET code = ? WHERE Id = ?","A03",27)
}

func deleteProduct(){
	fmt.Println("deleteProduct")
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database") 
		panic("failed to connect database")
	}
   
 db.Exec("DELETE products where id = ?",26)
}

func getProduct(){
	fmt.Println("getProduct")
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database") 
		panic("failed to connect database")
	}
    var product Product
	//db.Raw("select * from products where code = ?","A01").Scan(&product)  
	db.Exec("select * from products where code = ?","A03").Scan(&product)
	fmt.Println(product)
} 

func getProductList(){
	fmt.Println("getProductList")
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database") 
		panic("failed to connect database")
	}
   
	var products []Product
	//db.Raw("select * from products").Scan(&product)  
	db.Exec("select * from products").Scan(&products)
	//fmt.Println(products)
}



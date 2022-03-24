package main

import (
	"errors"
	"fmt"
	"learning-grom/database"
	"learning-grom/models"

	"gorm.io/gorm"
)

func main() {
	database.StarDB()
	//createUser("John@doe.com")
	//getUserById(1)
	//updateUserById(1, "john@gmail.com")
	//createProduct(1, "AA", "AX")
	//getUsersWithProducts()
	deleteProductById(1)
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating data: ", err)
		return
	}

	fmt.Println("New user created: ", User)
}

func getUserById(id uint) {
	db := database.GetDB()
	User := models.User{}

	err := db.First(&User, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User not found")
			return
		}
		print("Error getting data: ", err)
	}
	fmt.Printf("User found: %+v\n", User)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()
	User := models.User{}

	err := db.First(&User, "id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating data: ", err)
		return
	}
	fmt.Printf("User updated: %+v\n", User.Email)
}

func createProduct(userID uint, brand string, name string) {
	db := database.GetDB()
	Product := models.Product{
		UserID: userID,
		Brand:  brand,
		Name:   name,
	}
	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating data: ", err.Error())
		return
	}
	fmt.Println("New product created: ", Product)
}

func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}

	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting data: ", err.Error())
		return
	}

	fmt.Println("Users with products: ")
	fmt.Printf("%+v", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting data: ", err.Error())
		return
	}
	fmt.Println("Product deleted: ", id)
}

package repositories

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlserver"
	connections "connections"
	entities "entities"
	"time"
)

func SchemaMigration(){
	fmt.Println("repositories Schema Migration")
	
	db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.Migrator().CreateTable(&entities.User{})
}

func CreateNewUser(user entities.User, password string, is_email_verified bool) entities.User {
  fmt.Println("repository Create New User")

  db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})

  if err != nil {
	  panic("failed to connect database")
  }

  db.Exec("INSERT INTO users (created_at,email,password) VALUES (?,?,?)", time.Now(),user.Email,password)

  return user;
}

func UpdateUser(user entities.User){
	fmt.Println("repository Updateuser")
    db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})

  if err != nil {
	  panic("failed to connect database")
  }

  db.Exec("UPDATE users SET password = ? WHERE email = ?", user.Password,user.Email)
}

func GetAllUsers() []entities.User {
	fmt.Println("repository GetAllUsers")
	db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})

  if err != nil {
	  panic("failed to connect database")
  }

  var users []entities.User
  db.Raw("select * from users").Scan(&users)
  return users
}

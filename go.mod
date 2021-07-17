module main

go 1.15

require (
	connections v1.2.3
	controllers v1.2.3
	entities v1.2.3
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/swaggo/http-swagger v1.0.0 // indirect
	github.com/swaggo/swag v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210511113859-b0526f3d8744 // indirect
	golang.org/x/tools v0.1.1 // indirect
	gorm.io/driver/sqlserver v1.0.7 // indirect
	gorm.io/gorm v1.21.9 // indirect
	repositories v1.2.3
	services v1.2.3
	docs v1.2.3
)

replace entities v1.2.3 => ./data/entities

replace controllers v1.2.3 => ./presentation/controllers

replace services v1.2.3 => ./business/services

replace repositories v1.2.3 => ./data/repositories

replace connections v1.2.3 => ./data/connections

replace docs v1.2.3 => ./docs
module example.com/server

go 1.20

replace example.com/org => ../orgs

require (
	example.com/class v0.0.0-00010101000000-000000000000
	example.com/org v0.0.0-00010101000000-000000000000
	example.com/roles v0.0.0-00010101000000-000000000000
	example.com/users v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require (
	example.com/db v0.0.0-00010101000000-000000000000 // indirect
	example.com/model v0.0.0-00010101000000-000000000000 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	gorm.io/driver/postgres v1.5.2 // indirect
	gorm.io/gorm v1.25.1 // indirect
)

replace example.com/db => ../db

replace example.com/model => ../model

replace example.com/users => ../users

replace example.com/class => ../class

replace example.com/roles => ../roles

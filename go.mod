module github.com/aduryagin/cfop-backend

go 1.12

replace github.com/aduryagin/cfop-backend/db v0.0.0 => ./db

replace github.com/aduryagin/cfop-backend v0.0.0 => ./

require (
	github.com/99designs/gqlgen v0.8.2
	github.com/denisenkom/go-mssqldb v0.0.0-20190401154936-ce35bd87d4b3 // indirect
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/go-chi/chi v3.3.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/jinzhu/now v1.0.0 // indirect
	github.com/lib/pq v1.0.0 // indirect
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/rs/cors v1.6.0
	github.com/vektah/gqlparser v1.1.2
)

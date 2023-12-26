module fileServer

go 1.20

replace data v0.0.1 => ./data

replace cache v0.0.1 => ./cache

require (
	cache v0.0.1
	data v0.0.1
	github.com/go-sql-driver/mysql v1.7.1
)

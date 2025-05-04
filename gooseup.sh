[ -f sqlite.db ] || touch sqlite.db
goose -dir ./sql/schema/ sqlite3 ./sqlite.db up 

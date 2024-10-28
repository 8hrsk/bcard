package Database

import (
	"os"
	"vcard/types"
)

var DatabaseCredentials types.DatabaseCredentials

func InitDatabaseCredentials() types.DatabaseCredentials {
	DatabaseCredentials := types.DatabaseCredentials{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	return DatabaseCredentials
}

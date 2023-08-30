package main

import service "github.com/msegeya56/library_platform/book/cmd/service"

func main() {
	DB, err := connection.Connection()
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	err = connection.SyncDB(DB) // Apply migrations
	if err != nil {
		panic("failed to apply migrations: " + err.Error())
	}

	// Create the repository instance using the correct function name
	repo := repository.NewBookRepository(DB)



	
	service.Run(repo)
}

run: 
	nodemon --watch './*.go' --signal SIGTERM --exec go run .
	
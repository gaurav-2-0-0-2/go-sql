### A simple example to connect Go server to sql database (sqlite)

- Ensure you have sqlite installed. if not install it according to your OS
***I am using Debian, the below command can also work on ubuntu/mint and similar distros*** 
- install sqlite3: `sudo apt install sqlite3`
- clone the repo: `git clone https://github.com/gaurav-2-0-0-2/go-sql.git`
- go to db folder & create a database (users) with command: `sqlite3 users.db`
- build the code: `go build main.go`
- run binary: `./main`
- you can test it by running: `curl http://127.0.0.1:4000/users`

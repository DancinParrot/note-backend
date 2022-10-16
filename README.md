# note-backend

## PostgreSQL Database Setup

Install PostgreSQL: `sudo apt-get install postgresql`

### User Setup for Postgres
1. `sudo passwd postgres` and enter password.
2. Close and reopen terminal.

### Use psql cmd
1. `sudo service postgresql start` to start server.
2. `sudo -u postgres psql` to connect to postgres.

If doesn't work try using the postgres user.
1. `su - postgres`
2. `psql`

### Create user for postgres
`$ sudo -u postgres createuser <username>`

### Creating DB
`sudo -u postgres createdb <dbname>`

### Alter/Give user a password
1. `sudo -u postgres psql` to go to interactive psql cmd.
2. `alter user <username> with encrypted password '<password>';`

### Grant privileges on db
`grant all privileges on database <dbname> to <username> ;`

### Access postgresql
- Dbeaver
- PgAdmin

## References
https://harshityadav95.medium.com/postgresql-in-windows-subsystem-for-linux-wsl-6dc751ac1ff3

## Running the Go server
Simply execute `go run server.go` to run the backend server.

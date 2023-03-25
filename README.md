# Wait for

This is a tiny application that will wait for a service to become available. The service is available, the application will exit with a 0 exit code. If the service is not available after a defined period, the application will exit with a 1 exit code.

It supports the following services:

- MySQL
- Redis

## Configuration

The application is configured using environment variables. The following enviromment variables are available:

### General configuration

- `SERVICES` - an array of service types (seperated by commas) to wait for (`mysql`, `redis`)
- `INTERVAL` - the length of time to wait between checks in seconds (default: `2`)

### MySQL configuration

- `MYSQL_HOST` - the hostname or IP of a MySQL
- `MYSQL_PORT` - the MySQL port (default: 3306)
- `MYSQL_USERNAME` - the MySQL username (default: root)
- `MYSQL_PASSWORD` - the MySQL password
- `MYSQL_DATABASE` - the MySQL database to check presence of
- `MYSQL_TABLE` - the MySQL table (within above database) to check presence of

### RediS configuration

- `REDIS_HOST` - the hostname or IP of a Redis
- `REDIS_PORT` - the Redis port (default 6379)
- `REDIS_PASSWORD` - the Redis server password

# Waiter

Waiter is a tiny application that will wait for a set of services to become available and then exit cleanly when they are all available. This is ideal for with init-containers in Kubernetes so your application will only start when all its service dependencies are available.

It supports the following services:

- **MySQL** - wait for a MySQL server to become available and for a specific table to have been created.
- **Redis** - wait for a Redis server to be available.

## Usage

This is most useful when used as an init container within a Kubernetes pod. For example:

```yaml
initContainers:
  - name: waiter
    image: ghcr.io/adamcooke/waiter:v1
    env:
      - name: SERVICES
        value: mysql
      - name: MYSQL_HOST
        value: mysql.default.svc.cluster.local
      - name: MYSQL_USERNAME
        value: root
      - name: MYSQL_PASSWORD
        valueFrom:
          secretKeyRef:
            name: mysql-credentials
            key: password
      - name: MYSQL_DATABASE
        value: my_database
```

Alternatively, you can use this directly within Docker if you wish.

```
docker run --rm -e SERVICES=mysql -e MYSQL_HOST=mysql -e MYSQL_USERNAME=root -e MYSQL_PASSWORD=secret -e MYSQL_DATABASE=my_database ghcr.io/adamcooke/waiter:v1
```

## Configuration

All configuration is provided through environment variables. The following variables are available:

### General configuration

- `SERVICES` - an array of service types (seperated by commas) to wait for (options: `mysql`, `redis`)
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

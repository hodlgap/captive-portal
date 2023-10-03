# captive-portal
![logo](images/starbucks.png)

Purchase a cup of coffee to get access to the internet.

## Quick Start
Start dependencies
```bash
docker-compose up -d
```

Copy the example config file
```bash
cp config.example.yaml config.yaml
```

Modify config file
```yaml
# config.yaml
app:
  name: "captive-portal"
  env: "dev"  # Change to "prod" in production environment
  host: 127.0.0.1
  port: 8080
  graceful_timeout_second: 0.1
  log_level: "debug"

newrelic:
  license_key: "example"  # Change to your own license key

openwrt:
  encryption_key: "example"  # Change to the own encryption key of OpenNDS

redis:
  host: 127.0.0.1
  port: 6379
  db: 0
  password: ""  # Change to your own redis password or Refer to docker-compose.yml

db:
  host: 127.0.0.1
  port: 5432
  user: "postgres"
  name: "captive-portal"
  password: "example"  # Change to your own postgres password or Refer to docker-compose.yml
```

Run the application. The configuration file is in the current directory by default.
```bash
go run main.go
```

## Formatting
Format code
```bash
make format
```
Check formatting
```bash
make lint
```

## DB schema modifications
If you want to modify the database schema, you should modify the schema file directly. 

And then generate models from the database
```bash
make models
```
After you finished, dump the database schema
```bash
make dump-db
```

### Restore database
```bash
make restore-db
```

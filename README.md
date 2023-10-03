# captive-portal
![logo](images/starbucks.png)

Purchase a cup of coffee to get access to the internet.

## Quick Start
Start dependencies
```bash
docker-compose up -d
```

Run the application
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

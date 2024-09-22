# Starter REST API, with Golang, Fiber

### Start docker service
```bash
sudo systemctl start docker

make docker-dev-up
```

### Make database migration
```bash
# Install DBmate
sudo curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
sudo chmod +x /usr/local/bin/dbmate

# Start migration
dbmate new <migration-name>
dbmate up
dbmate down
```

### Start development
```bash
# install intial tool
make setup

# install libraries or dependencies
go mod download

# turn on docker containers
make docker-dev-up

# migrate schema (if needed)
dbmate load

# run go server with air hot reload
air
```

### Project Map / Structure Directory
```bash
├─ _docker-data       # Docker container data
├─ bin                # Binary compiled
├─ cmd                # Apps, main functions for each app
├─ configs            # Configs for service/dependency
├─ db                 # Database restore/backup/migration
│   ├─ backups        # Backup database files
│   ├─ migrations     # Database migrations
│   └─ schema         # SQL Table definitions
├─ docs               # Config generated swagger API Docs
├─ helper             # Other codes, can be imported anywhere
├─ internal           # Most logical, including app wrapper
│   ├─ bootstrap      # App components, call almost all here
│   ├─ controller     # Business logic, http handler/router
│   ├─ repository     # Repository layer, database integration
│   ├─ request        # Request schema from controller
│   ├─ response       # Response schema for controller
│   ├─ service        # Service layer
│   └─ web            # Web-Server stuff
├─ logs               # Log files
├─ script             # Automation scripts, CI/CD, etc.
├─ storage            # Static files
├─ test               # Unit test, integration test
└─ tmp                # Temporary files, for development

```
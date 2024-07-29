## Fiber template

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
│   ├─ bootstrap      # App components
│   ├─ controller     # Business logic, http handler
│   ├─ repository     # Repository layer, database integration
│   ├─ request        # Request schema from controller
│   ├─ response       # Response schema for controller
│   ├─ service        # Service layer
│   └─ web            # Web-Server stuff
├─ logs               # Log files
├─ script             # Automation scripts, including CI/CD
├─ storage            # Static files
├─ test               # Unit test, integration test
└─ tmp                # Temporary files, for development

```

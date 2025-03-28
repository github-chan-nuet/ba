# Structure
Guidelines:
- [Official Guidelines](https://go.dev/doc/modules/layout)
- [Unofficial best practices](https://github.com/golang-standards/project-layout)

# Starting Phishing Backend
## with docker (recommended)
```bash
docker build -f ./.docker/Dockerfile -t phishingbackend .
docker run -d -p 8080:8080 --env-file ./.docker/phishing_backend.dev.env --name phishingbackend_instance phishingbackend 

```

## without docker
```bash
# 1. download external libraries
go mod download
# 2. Compile and run application
go run .
# (alternative to 2., Build  application as binary)
go build -o myapp 
```

## verify that the server is running
open http://127.0.0.1:8080/api/health


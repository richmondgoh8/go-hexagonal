# go-hexagonal
Golang with Hexagonal Architecture, 
- Uses Zap as a logging mechanism
- JWT Middleware Auth for Authentication
- Connects to Postgres DB

# Reference Commands
- go mod init github.com/richmondgoh8/boilerplate

# Folder Structure
3 Primary Folders
- Core = Business Logic 
- Handlers = HTTP Handlers 
- Repositories = Actors (External Adapter i.e. Interaction with DB)

Handlers => Services => Ports <=> Repositories (DB/ Redis/ Storage)

# Effective Go
- Good package names are short and clear. They are lower case, with no under_scores or mixedCaps
- Source files are all lower case with underscore separating multiple words.
- Variables and Unexported Functions = camelCase
- Exported Functions & Constants = PascalCase
- Contexts are generally used to carry custom data among handlers.

# Quickstart
```cgo
# Runs in Docker
make full

# Runs only Postgres in Docker
make postgres
go mod tidy
make run
```

# Endpoints
```cgo
[GET] localhost:8080/token
[GET] localhost:8080/url/1
[PUT] localhost:8080/url/3?url=https://www.rlc4u.com&name=Richmond
```

# Local Docker Build Testing
```cgo
docker build --no-cache -t "test" .
docker build --no-cache --progress=plain -t "test" .
```
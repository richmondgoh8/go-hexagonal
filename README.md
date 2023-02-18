# go-hexagonal
Golang with Hexagonal Architecture, Uses Zap as a logging mechanism

# Reference Commands Commands
- go mod init github.com/richmondgoh8/boilerplate

# Folder Structure
3 Primary Folders
Core = Business Logic
Handlers = HTTP Handlers
Repositories = Actors (External Adapter i.e. Interaction with DB)

Handlers <=> Services
Ports <=> Actors

# Effective Go
- Good package names are short and clear. They are lower case, with no under_scores or mixedCaps
- Source files are all lower case with underscore separating multiple words.
- Variables and Unexported Functions = camelCase
- Exported Functions & Constants = PascalCase
- Contexts are generally used to carry custom data among handlers.
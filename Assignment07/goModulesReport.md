# 1. Blank Imports in Go: Use Cases

Blank imports in Go are used when a package is solely imported for its side effects, rather than for its exported functions or variables. Here are some common use cases:

* **Registering database drivers**: As in the example with the `mysql` package, blank imports are used to register the driver as a database driver in the `init()` function, without importing any other functions.
* **Initializing global variables**: Blank imports can be used to initialize global variables in a package, which are then used by other packages.
* **Side effects of imports**: Blank imports can be used to ensure that certain side effects occur when a package is imported, such as logging or setting up environment variables.

For example:
```go
import _ "github.com/go-sql-driver/mysql"
```
This imports the `mysql` package solely for its side effect of registering the driver, without importing any other functions.

**When to use blank imports**

Blank imports are useful when:

* You only need the side effects of a package, not its exported functions or variables.
* You want to ensure that certain initialization or setup occurs when a package is imported.
* You want to avoid polluting your import list with unused variables or functions.

# 2. What if repository name and module name doesn’t match, what are the consequences 

It is not mandatory to match the module name with the repository URL/Path. Although there are some directives about How a module name should be and the associated constraints with the path naming conventions at GO Modules Documentation

Consequences of dissimilar module name and user name
Go Runtime not able to find the go Module due to unavailability of module path.

Explanation:
When a new module is imported the go runtime looks up the repository location by sending a get request to modulePath?go-get=1. For the request above go runtime checks for metadata returned in response for the path of the module repository. After getting the path go performs a go get operation and clones the repository.

Therefore if the module name does not match path the GO runtime will not be able to find the repository.

# 3. Find what go.sum does and its use

The Go.sum file serves a key functionality in the dependency management. It serves following purposes: 

Content Validation:
- Go.mod file stores the cryptographic checksum of the module contents. In this way it makes sure the  downloaded dependencies match their intended state.
- Distributed Package Security:
Go implements distributed content validity management mechanism with go.mod file
- Module state Reproducibility:
Allows same Module to be built in same state as it intended

# 4. Go workspace examples and use case
Go workspaces are a feature introduced in Go 1.18 that simplify working with multiple interdependent modules simultaneously.

Workspace Example:
Workspaces Could be Created Using Following Commands:
```go
go work init ./module1 ./module2
go work use ./module1
```

Usecases:
- Local Module Development
- Interdependent Module Management
- Simplified Dependency resolution

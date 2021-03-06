
## Implant Core 

The `core/` directory contains two main things: 

----
### 1) Implant entrypoints 

The `core/` directory contains, for each major supported OS, entrypoints files (main functions). Depending on the attacked platform,
several implants are available for compilation. Also, several types of entrypoints are available: Windows DLLs, for instance.
Usually, these entrypoint files are named `ghost.go`, `ghost_dll.go`, etc. They will always either declare a `main()` function, or 
a wrapper function such as:

```go
func EntrypointWrapperFunction() { main() }

func DllInstall() { main() }
```

----
### 2) Core functionality 

The `core/` directory also contains all core functionality provided by Wiregost implants; that is, all post-exploitation capacity.
The code is separated between the major operating systems (Linux, MacOS, Windows), completed by a `generic/` directory containing
code working cross-platform.

This `generic` directory contains subpackages categorized by topic/domain, not OS. Usually, inside these, files are named like `proc_windows.go`
or `proc_linux.go`, for conditional compilation to occur based on target OS.

The other three following directories (`darwin/`, `linux/`, `windows/`) contain code and objects that are always specific to an OS.
This allows for easier/clearer package referencing, and for ensuring a Linux implant will never be compiled with Darwin/Windows code.

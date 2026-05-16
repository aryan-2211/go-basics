# go-basics

A structured, hands-on Go learning repository covering core language concepts through 25 self-contained modules. Each folder is a standalone program that demonstrates one concept, progressing from hello world to goroutines and channels.

## Module Index

| # | Topic | What it covers |
|---|-------|----------------|
| 01 | Hello World | `fmt.Print`, package main, entry point |
| 02 | Variables | Declaring strings and bools, `%T` format verb for type inspection |
| 03 | User Input | `bufio.NewReader`, reading stdin with `ReadString` |
| 04 | Conversions | `strconv.ParseFloat`, `strings.TrimSpace`, error handling pattern |
| 05 | Time | `time.Now()`, formatting and printing current time |
| 06 | Pointers | `&` (address-of), `*` (dereference), mutating values via pointer |
| 07 | Arrays | Fixed-length arrays, zero values, `len()` |
| 08 | Slices | Dynamic slices, `append`, `make`, `sort.Ints` |
| 09 | Maps | `make(map[K]V)`, insert, read, delete |
| 10 | Structs | Defining and using structs (Go's alternative to classes) |
| 11 | Loops | `for i`, `for range`, while-style loops |
| 12 | Functions | Single return, multiple return, variadic functions |
| 13 | Methods | Receiver functions on structs (`func (u User) Method()`) |
| 14 | Defer | `defer` execution order (LIFO), practical use in cleanup |
| 15 | File Handling | `os.Create`, `io.WriteString`, `ioutil.ReadFile` |
| 16 | Web Requests | `http.Get`, reading response body, panic on error |
| 17 | Handling URLs | `url.Parse`, extracting scheme/host/path/query params |
| 18 | GET Request | Making a GET request to a local server, reading JSON response |
| 19 | POST Request (JSON) | `http.Post` with a JSON body via `strings.NewReader` |
| 20 | POST Form Data | `url.Values`, `http.PostForm` for form-encoded requests |
| 21 | JSON | `json.MarshalIndent`, struct tags (`json:"field,omitempty"`) |
| 22 | Modules | `go.mod`, external dependency (Gorilla Mux), simple HTTP server |
| 23 | Build API | Course/Author structs for a REST API skeleton |
| 24 | Interfaces | Defining and implementing interfaces, embedded interfaces, custom errors |
| 25 | Goroutines | `go` keyword, channels (`chan`), `close`, ranging over a channel |

## Running Any Module

Each module is independent. Navigate into its folder and run:

```bash
cd go-basics/<module-folder>
go run main.go
```

For modules with external dependencies (`22mymodules`, `23buildAPI`), install deps first:

```bash
go mod tidy
go run main.go
```

## Key Concepts Covered

**Language fundamentals** — variables, types, zero values, type inference, error handling with multiple return values.

**Data structures** — arrays (fixed), slices (dynamic), maps, structs.

**Concurrency** — goroutines, channels, the `close` function, ranging over channels. Module 25 demonstrates a producer goroutine sending 1000 integers over a channel.

**HTTP & Networking** — making GET/POST requests, parsing URLs, form data, JSON marshaling/unmarshaling, and building a simple HTTP server with Gorilla Mux.

**Interfaces** — defining contracts (`Shape`, `Measurable`), composing interfaces, and implementing `error` with a custom struct.

## Tech Stack

- **Language:** Go 1.19+
- **External packages:** `github.com/gorilla/mux` (modules 22 and 23 only)
- **Everything else:** Go standard library only

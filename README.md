# JSON Parser

A lightweight, zero-dependency JSON parser and interactive REPL written in pure Go, built in accordance with the [RFC 8259](https://datatracker.ietf.org/doc/html/rfc8259) and [json.org](https://www.json.org/) specifications.

---

## 🚀 Features

- **Spec-Compliant Recursive Descent Parsing**: Parses valid JSON objects and arrays into native Go data structures (`map[string]any`, `[]any`, `float64`, `string`, `bool`, `nil`).
- **Full Type Support**:
  - **Objects**: Key-value mappings enclosed in `{}`.
  - **Arrays**: Ordered lists of heterogeneous values enclosed in `[]`.
  - **Strings**: Handles full escape sequences (`\"`, `\\`, `\/`, `\b`, `\f`, `\n`, `\r`, `\t`) and 4-digit hexadecimal Unicode escapes (`\uXXXX`).
  - **Numbers**: Integers, floating-point decimals, negative numbers, and scientific exponent notation (`e`/`E`).
  - **Booleans & Null**: `true`, `false`, and `null`.
- **Interactive REPL**: Read-Eval-Print Loop for testing and validating JSON snippets directly in the terminal.
- **CLI File Parser**: Command-line interface with dedicated flag handling (`-f`, `-file`) and proper exit status codes.
- **Zero External Dependencies**: Pure Go using only standard library packages (`bufio`, `flag`, `fmt`, `io`, `os`, `regexp`, `strconv`, `strings`, `unicode`).
- **Comprehensive Test Suite**: Includes a corpus of positive (`pass*.json`) and negative (`fail*.json`) test cases validating grammar boundaries.

---

## 📁 Project Structure

```text
.
├── cmd/
│   └── main.go           # CLI entrypoint (flag parsing, file evaluation & REPL launcher)
├── internal/
│   ├── json/
│   │   └── json.go       # Core parser implementation (combinators, lexing & grammar rules)
│   └── repl/
│       └── repl.go       # Interactive terminal REPL loop
├── test/                 # Test corpus containing passing and failing JSON test cases
│   ├── pass1.json ... pass5.json
│   ├── fail1.json ... fail33.json
│   └── test.json
├── go.mod                # Go module definition
└── README.md             # Project documentation
```

---

## 🛠️ Installation & Building

### Prerequisites
- [Go](https://go.dev/dl/) (version 1.20 or newer recommended)

### Clone and Build
```bash
# Clone the repository
git clone https://github.com/yashbaddi/jsonparser.git
cd jsonparser

# Build binary
go build -o jsonparser cmd/main.go
```

---

## 💻 Usage

### 1. CLI File Mode
Pass a JSON file using `-f` or `-file` to parse and validate its contents:

```bash
# Using the compiled binary
./jsonparser -f test/test.json

# Or using `go run`
go run cmd/main.go -f test/test.json
```

**Example Output:**
```text
Parsed Success
map[Phone:iPhone Skills:[servNow js] age:26 isFemale:true isMale:false laptop:Dell name:Mansi]
```

#### Exit Codes
| Exit Code | Description |
|---|---|
| `0` | Successful parsing |
| `1` | JSON syntax / parsing error |
| `2` | File I/O or CLI argument error |

---

### 2. Interactive REPL Mode
Running without arguments starts the interactive prompt:

```bash
./jsonparser
```

```text
JSON Parser > {"name": "Alice", "active": true, "scores": [98.5, 100]}
JSON Output: 
 map[active:true name:Alice scores:[98.5 100]]
JSON Parsed Successfully!! 
JSON Parser > [1, 2, "three", null]
JSON Output: 
 [1 2 three <nil>]
JSON Parsed Successfully!! 
JSON Parser > 
```

---

## 📦 Using as an Internal Library

You can import and use the parser directly in Go code:

```go
package main

import (
	"fmt"
	"github.com/yashbaddi/jsonparser/internal/json"
)

func main() {
	input := []byte(`{"name": "Gopher", "version": 1.26}`)

	parsed, ok := json.JSONParser(input)
	if !ok {
		fmt.Println("Invalid JSON")
		return
	}

	fmt.Printf("Parsed: %#v\n", parsed)
}
```

### Type Mapping Reference

| JSON Data Type | Go Representation | Example Input | Go Value |
|---|---|---|---|
| Object | `map[string]any` | `{"key": "val"}` | `map[string]any{"key": "val"}` |
| Array | `[]any` | `[1, "two"]` | `[]any{1.0, "two"}` |
| String | `string` | `"hello \u0041"` | `"hello A"` |
| Number | `float64` | `-42.5e+2` | `-4250.0` |
| Boolean | `bool` | `true` / `false` | `true` / `false` |
| Null | `nil` | `null` | `nil` |

---

## 🔍 How It Works

The parser utilizes a recursive descent parser design:
1. **`JSONParser` (Root Parser)**: Ensures top-level JSON conforms to an Object (`{...}`) or Array (`[...]`) and verifies all input bytes are consumed without trailing non-whitespace characters.
2. **`valueParser`**: Acts as a dispatcher, sequentially attempting to parse `null`, `bool`, `number`, `string`, `object`, or `array`.
3. **`stringParser` & `parseEscape`**: Reads rune by rune, decoding JSON escape characters such as `\n`, `\t`, `\"`, and `\uXXXX` unicode code points.
4. **`numberParser`**: Uses regular expression matching conforming to JSON specification for numeric formats before parsing to standard `float64`.
5. **`spaceParser`**: Consumes surrounding whitespace (`\x20`, `\t`, `\n`, `\r`) between tokens.

---

## 🧪 Testing

The repository comes with standard JSON test fixtures inside [`test/`](file:///Users/yashbaddi/Documents/PersonalProjects/JSONParser/test):
- **Positive test cases (`pass*.json`, `test.json`)**: Valid JSON payloads covering nested structures, escapes, and varied types.
- **Negative test cases (`fail*.json`)**: Malformed inputs testing unclosed brackets, invalid escapes, trailing commas, and unquoted keys.

You can verify any test file using:
```bash
./jsonparser -f test/pass1.json
```

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

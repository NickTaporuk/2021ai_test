# SCALC

## Summary
    Package was developed as the test task for 2021ai comapny
### Package directories structure
    ├── cmd         - all scripts of running cli
    ├── operators   - initialize new operator and add logic of operation for this operator
    ├── parser      - represents a syntax parser.
    ├── scanner     - represents a lexical scanner.
    ├── set         - realization of the set data structure
    ├── stack       - realization of the stack data structure
    ├── testdata    - test data of the solution
    ├── token       - represents a lexical token
    └── utils       - all helpers for the app

# Download

## Precompiled Binaries

You can download the precompiled release binary from [releases](https://github.com/NickTaporuk/2021ai_test/releases/) via web
or via

```bash
wget https://github.com/NickTaporuk/2021ai_test/releases/<version>/2021ai_test_<version>_<os>_<arch>
```

#### Go get

You can also use Go 1.12 or later to build the latest stable version from source:

```bash
GO111MODULE=on go get github.com/NickTaporuk/2021ai_test
```

#### Homebrew Tap

```bash
brew install nicktaporuk/tap/scalc
# After initial install you can upgrade the version via:
brew upgrade scalc
```
## Compilation

```bash
git clone git@github.com/NickTaporuk/2021ai_test.git
cd 2021ai_test
go build -o scalc cdm/main.go
```
## Usage
    The Solution can work in 2 mode: a onetime running or as a interactive mode.
    
    To reach out first type of the mode you should run the app in cli with some parameters
    
    Example :
        ./scalc [ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]
        
    To reach out second type of the mode you should run the app in cli without any parameters.
    After that you can work with the app with interactive mode
    Example :
        ./scalc
        
        You should see :
        scalc >

## Complexity
    Complexity is O(n2)
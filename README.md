# conv

A simple command line unit converter.

I created this project to learn [dynamic library loading](https://medium.com/learning-the-go-programming-language/writing-modular-go-programs-with-plugins-ec46381ee1a9) in Go. It probably has no practical use.

### Installation
1. Install [Go](https://go.dev/doc/install)
2. Make sure you have GOPATH in your PATH:  
    ```
    export PATH=$HOME/go/bin:$PATH
    ```
3. Fetch this repo  
    ```
    git clone https://github.com/againagainst/conv.git
    ```
4. Build and install    
    ```
    cd conv/cmd/conv
    go install
    ```    

### Usage
```
conv 149.76.12.4
conv 149.76.12.0 to ipv6
conv 9 inch to cm 
conv 150 nok to usd
```

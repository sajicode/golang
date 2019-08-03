# Important aspects of Golang

1. Multiple return values.
2. A modern standard library.
  a. Networking & HTTP
  b. HTML
  c. Cryptography
  d. Data Encoding
3. Concurrency with goroutines & channels.
4. Go toolchain.
  a. Package Management
  b. Testing
  c. Code Coverage
  d. Formatting

## Go Compiles faster than C

## Go does not need a runtime to run on the OS unlike Java

## Python and PHP sit behind a web server e.g. Nginx, Apache while Go has an inbuilt server

  . One of the reasons it's useful to put Python & PHP behind a web server has to do with threads & concurrent connections.
  . Python has a global interpreter lock that allows only one thread to execute at a time.
  . PHP applications tend to run from start to end in a process.
  . To enable multiple connections to an application concurrently, a web server e.g. Nginx can sit in front of the application
    and handle running the concurrent connections in separate processes.
  . The built-in Go web server takes advantage of goroutines to run connections concurrently.
  . Performance wise, PHP and Python take performance critical code and rewrite it in C, Go applications on the other hand
    compile into binary. Software in the standard library and aplication are both compiled into machine code. There is no distinction.

## JavaScript has a single-threaded model, although async I/O may use separate threads, the main program executes in a single thread.
  
  . When code in the main thread takes a significant amount of time, it blocks other code from executing.
  . Go uses a multithreaded model in which the runtime manages goroutines running conncurrently

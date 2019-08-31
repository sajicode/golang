# Goroutines

1. Concurrency is a program's ability to do multiple things at the same time.<br/>

2. A concurrent programs are programs that have two or more tasks that run independently of each other, at about the same time, but remain part of the same program.<br/>

3. Go uses the concurrency model called _Communicating Sequential Processes_ (CSP).<br/>

4. Two crucial concepts make Go's concurrency model work:

  . Goroutines - A function that run independently of the function that started it. It could also be explained as a function that runs as if it were on its own thread.

  . Channels - A channel is a pipeline for sending and receiving data. It is like a socket <socket.io> that runs inside your program. Channels provide a way for one goroutine to send structured data to another.

5. Goroutines run concurrently, but not necessarily in parallel. When you schedule a goroutine to run by calling <go func>, you're asking the Go runtime to execute that function for you as soon as it can, but that's not likely.

6. Although calling the Go scheduler may guarantee that the scheduler has a chance to check for other goroutines, we shouldn't rely on it as a tool for ensuring that other goroutines have a chance to complete.<br/>


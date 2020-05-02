Problem statement

This problem is the typical CR access. There is  a variable that will be incremented by a Go routine. The goal is not to miss any update.

One potential solution is to use WaitGroup with size 1 on each update. Instead it is required to use Mutex.
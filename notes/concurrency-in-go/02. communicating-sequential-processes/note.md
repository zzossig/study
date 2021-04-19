# Communicating Sequential Processes

## The Difference Between Concurrency and Parallelism

`Concurrency`는 하나의 cpu에서 여러 프로세스를 짧은 시간동안 번갈아가며 실행하는데, 이 때 사람이 느끼기에 프로세스가 동시에 진행된다고 느끼기 때문에 concurrency인 것이다.
실제로는 하나의 프로세스가 실행되는 동안에는 다른 프로세스는 중단되어 있다. 다만 인간이 그 차이를 느끼지 못하는 것일 뿐. 프로세스에서 다른 프로세스로 넘어 갈 때 context switcing이 일어나는데, 이 때 컴퓨팅 자원이 소모된다.

`Parallelism`은 여러 cpu코어에서 여러 프로세스가 진짜로 동시에 실행되는 것을 말한다. 짧은 순간의 context switcing도 하지 않고 말 그대로 병렬로 프로세스가 진행된다.

- Concurrency is property of the code; parallelism is a property of the running program.
- We do not write parallel code, only concurrent code that we hope will be run in parallel.

## What is CSP?

CSP stands for `Communicating Sequential Processes`. This concept impacted to Go's philosophy.

## Go's Philosophy on Concurrency

Do not communicate by sharing memory. Instead, share memory by communicating.
Go’s philosophy on concurrency can be summed up like this: aim for simplicity, use channels when possible, and treat goroutines like a free resource.
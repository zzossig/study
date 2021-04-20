# Go's Concurrency Building Blocks

## Goroutines

Goroutines는 실행 중간에 들여다 볼 수 없다(nonpreemptive). 대신 멈추거나 다시 시작할 수는 있다. Go runtime에서 Goroutines의 스케쥴을 관리한다.
Go에서 고루틴을 호스팅하는 방법은 `M:N scheduler`라 부르는 방식을 쓴다.
고루틴은 green thread로 생성되다가 한도가 초과되면 사용 가능한 threads로 생성된다. 
이 때 스케쥴러는 고루틴이 실행되고 멈추는 모든 행동이 concurreny할 수 있게 관리하게 된다.
Go는 concurrency 모델로 `fork-join`모델을 사용한다. `Go`키워드로 가지가 갈라져 나오면 어느 지점에서 가지가 다시 붙어야 하는 것이다.

다음은 메인 고루틴과 메인에서 생성된 고루틴의 `join-point`를 어떻게 만드는지 보여준다.
```go
var wg sync.WaitGroup
sayHello := func() {
  defer wg.Done()
  fmt.Println("Hello")
}
wg.Add(1)
go sayHello()
wg.Wait() // join-point
```

## WaitGroup

WaitGroup은 Concurrent 작업을 완료할 수 있게 도와준다. Concurrent 작업의 결과물이 뭐가되든 상관 없을 때 쓰는것이 좋다. 결과물을 제대로 control하려면 channel이나 select문을 써야한다.

## Mutex and RWMutex

Mutex stands for `mutual exclusion`. Mutex로 데이터의 원자성을 보장할 수 있게 된다.
Mutex는 Go의 concurrency 철학인 `share memory by communicating`이 아닌, 기존 프로그래밍에서 concurrency safe하기 위해 쓰던 관습적인 방법이다.
Mutex와 RWMutex의 성능 차이는 얼마 나지 않기 때문에 로직상 가능하다면 RWMutex를 더 선호하는게 좋다.
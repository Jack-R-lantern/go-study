# The Go Memory Model
_2022,06 version_
* [Introduction](#introduction)
    * [Advice](#advice)
    * [Informal Overview](#informal-overview)
* [Memory Model](#memory-model)
* [Synchronizaiton](#synchronization)
## Introduction
Go의 메모리 모델은 하나의 고루틴에서 변수 읽기가 다른 고루틴에서 동일한 변수에 대한 쓰기로 생성된 값을 관찰할 수 있는 조건을 지정.

### Advice
여러 고루틴에 의해 동시에 액세스되는 데이터를 수정하는 프로그램은 이러한 액세스를 직렬화해야 함. \
액세스를 직렬화하기 위해, `channel` 작업 또는 기본 동기화 요소인 `sync`, `sync/atomic`를 이용해 데이터를 보호.

### Informal Overview
Go는 타 언어들과 거의 같은 방식으로 메모리 모델에 접근하며, 의미론을 단순하고 이해하기 쉬우며 유용하게 유지하는것을 목표. \
`data race`는 `sync/atomic` 패키지에 의해 제공된 원자적 데이터 접근을 제외한, 특정 메모리 위치에 쓰기가 발생했을때, 동시에 같은 위치에 읽기 또는 쓰기가 발생한것으로 정의. 

## Memory Model
### Requirement 1:
### Requirement 2:
### Requirement 3:
## Implementation Restrictions for Programs Containing Data Races

## Synchronization

### Initialization
프로그램 초기화는 단일 고루틴으로 실행되지만, 그 고루틴은 동시에 실행되는 다른 고루틴을 생성할 수 있음.\
만일 `package p`가 `package q`를 참조한다면 `q`의 `init`함수의 완료는 `p`의 시작 전에 일어남.\
모든 `init` 함수의 완료는 `main.main`함수의 시작전에 동기화 됨.

### Goroutine creation
새로운 고루틴을 시작하는 `go`구문은 고루틴 수행의 시작전에 동기화 됨.

**example**
>```go
>var a string
>
>func f() {
>    print(a)
>}
>
>func hello()  {
>    a = "hello, world"
>    go f()
>}
>```
>`hello` 함수 호출은 `"hello, world"`을 특정 미래 시점에 출력(`hello`함수가 종료된 이후일지라도)

### Goroutine destruction
고루틴의 종료는 프로그램의 어떤 이벤트 전이라 동기화를 보장하지 않음.

**example**
>```go
>var a string
>
>func hello() {
>       go func() { a = "hello" } ()
>       print(a)
>}
>```
> `a` 할당에 대한 동기화 이벤트가 없음, 다른 고루틴에서 관찰되어짐을 보장 할 수 없음. \
> 진보적인 컴파일러의 경우 아마 `go`구문 전체를 삭제할 수 있음.

만일 고루틴의 영향이 반드시 다른 고루틴에 의해 관측되어야 한다면, `lock`, `channel communication` 같은 동기화 메커니즘을 사용하여 상대적 순서를 설정.

### Channel communication
채널 통신은 고루틴 간의 주요 동기화 방법. \
특정 채널의 각 송신부는 일반적으로 다른 고루틴에서 해당 채널로부터의 해당 수신부와 일치. \
채널의 송신부는 채널의 수신부의 대응이 완료되기전에 동기화됨.

**example**
>```go
>var c = make(chan int, 10)
>var a string
>
>func f() {
>	a = "hello, world"
>	c <- 0
>}
>
>func main() {
>	go f()
>	<-c
>	print(a)
>}
>```
> 예제는 `hello, world` 출력을 보장. \
> a에 대한 쓰기는 c에 대한 전송 전에 수행되고, c에 대한 해당 수신이 왼료되기 전에 동기화되며, 출력전에 수행됨.

채널 닫는것은 동기화 되어진다 

### Locks
**example**
>```go
>var l sync.Mutex
>var a string
>
>func f() {
>	a = "hello, world"
>	l.Unlock()
>}
>
>func main() {
>	l.Lock()
>	go f()
>	l.Lock()
>	print(a)
>}
>```
> 

### Once
**example**
```go
var a string
var once sync.Once

func setup() {
	a = "hello, world!"
}

func doprint()  {
	once.Do(setup)
	print(a)
}
```
### Atomic Values
### Finalizers
### Additional Mechanisms
## Incorrect synchronization
## Incorrect compilation
## Conclusion
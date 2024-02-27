# Go 레이스 컨디션 예제

이 프로젝트는 Go 언어에서 멀티스레딩 환경에서 발생할 수 있는 레이스 컨디션(Race Condition) 문제와 이를 해결하기 위한 방법을 보여줍니다.

레이스 컨디션은 여러 고루틴이 데이터에 동시에 접근하려 할 때 발생하는 문제로, 예측하지 못한 결과를 초래할 수 있습니다.

## 프로젝트 구성

이 프로젝트는 두 개의 테스트 케이스로 구성되어 있습니다:

1. `TestCounterRaceCondition`: 레이스 컨디션을 발생시키는 테스트 케이스입니다. 여러 고루틴이 공유 변수 `counter`를 동시에 수정하려고 시도하며, 이로 인해 최종 결과가 예상과 다를 수
   있습니다.

2. `TestCounterWithMutex`: `sync.Mutex`를 사용하여 레이스 컨디션 문제를 해결하는 테스트 케이스 입니다. 뮤텍스는 공유 변수에 대한 동시 접근을 제어하여 안전한 데이터 수정을
   보장합니다.

## 테스트 실행

```sh
go test -v

```

## 기대 결과

- `TestCounterRaceCondition` 테스트는 레이스 컨디션으로 인해 실패할 가능성이 높습니다. 이는 고루틴 간의 데이터 접근이 동기화되지 않기 때문입니다.

- `TestCounterWithMutex` 테스트는 `sync.Mutex`를 사용하여 데이터 접근을 동기화하기 때문에 성공할 것입니다. 최종 `counter` 값은 1000이 됩니다.

## 테스트 시 레이스 상태 감지

```sh
go test -race
```

```
==================
WARNING: DATA RACE
Read at 0x00c00009c158 by goroutine 11:
go-concurrency.TestCounterRaceCondition.func1()
/Users/yoojehwan/GolandProjects/go-concurrency/counter_test.go:15 +0x34

Previous write at 0x00c00009c158 by goroutine 7:
go-concurrency.TestCounterRaceCondition.func1()
/Users/yoojehwan/GolandProjects/go-concurrency/counter_test.go:15 +0x44

Goroutine 11 (running) created at:
go-concurrency.TestCounterRaceCondition()
/Users/yoojehwan/GolandProjects/go-concurrency/counter_test.go:14 +0x70
testing.tRunner()
/opt/homebrew/opt/go/libexec/src/testing/testing.go:1595 +0x1b0
testing.(*T).Run.func1()
/opt/homebrew/opt/go/libexec/src/testing/testing.go:1648 +0x40

Goroutine 7 (finished) created at:
go-concurrency.TestCounterRaceCondition()
/Users/yoojehwan/GolandProjects/go-concurrency/counter_test.go:14 +0x70
testing.tRunner()
/opt/homebrew/opt/go/libexec/src/testing/testing.go:1595 +0x1b0
testing.(*T).Run.func1()
/opt/homebrew/opt/go/libexec/src/testing/testing.go:1648 +0x40
==================
```

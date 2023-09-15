# fmt
Println
Printf



# net/http




# Mutex,Mux 객체

Mutex: Mutex는 "mutual exclusion"의 줄임말로, 여러 고루틴(goroutine)이 동시에 데이터를 수정하는 것을 방지하기 위해 사용되는 동기화 기법. 이를 통해 공유 자원에 대한 동시 접근을 제어할 수 있다. sync.Mutex라는 타입으로 제공.
Mux: Mux는 "multiplexer"의 줄임말로, HTTP 요청 라우터나 디스패처 역할. 이것은 들어오는 요청을 그 요청의 속성(예: URL 경로, HTTP 메소드 등)에 따라 적절한 핸들러 함수에 연결. http.ServeMux라는 타입으로 제공.

Mutex: sync 패키지에서 제공. sync.Mutex와 같은 형태로 사용.
Mux: net/http 패키지에서 제공. http.ServeMux와 같은 형태로 사용.

# Testing

https://pkg.go.dev/testing@go1.21.1

Test + '함수이름'(ex.)Xxx) 은 대문자로 시작해야하고, 식별의 용도. 

Error, Fail 또는 관련 메서드를 사용하여 실패하는 경우를 만듦.

새 테스트 만드는 법 : TestXxx 함수를 포함하는 파일 생성 및 "_test.go"로 끝나도록 이름 지정. 해당 파일은 일반 패키지 빌드에서는 제외되지만 "go test" 명령이 실행되면 포함됨.

테스트 파일은 테스트 중인 패키지와 동일한 패키지에 있거나 접미사 "_test"가 붙은 해당 패키지에 있을 수 있음.

테스트 역시 다른 Public 테스트 패키지를 import 해서 사용할 수 있고, "black box"testing이라고 함.

함수 이름을 BenchmarkXxx로 시작하게 하고, 매개변수로 *testing.B를 받아야 하는 문서상 조건을 만족하게 작성하면 benchmark, examples, fuzzing, skipping 등등의 내장 테스팅 로직을 자동으로 제공해줌.

### wrtn이 알려준 사용예시

Go의 테스팅 프레임워크는 성능 벤치마킹 기능을 제공합니다. 이를 사용하려면 함수 이름을 `BenchmarkXxx`로 시작하게 하고, 매개변수로 `*testing.B`를 받아야 합니다. 그리고 이 함수 내부에서 원하는 코드의 성능을 측정할 수 있습니다.

예를 들어, 다음과 같이 벤치마크 테스트 함수를 작성할 수 있습니다:

```go
func BenchmarkMyFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MyFunction()
    }
}
```

위의 예제에서 `MyFunction()`은 성능을 측정하고자 하는 대상 함수입니다.

`b.N`은 Go 테스팅 프레임워크가 자동으로 설정하는 루프 카운트 값입니다. 이 값은 동적으로 조절되며, 벤치마크 시간이 충분히 길어질 때까지 실행 횟수를 늘립니다.

벤치마크 테스트는 일반적인 `go test` 명령에 `-bench` 옵션을 추가하여 실행합니다:

```bash
go test -bench .
```

위 명령은 현재 패키지 내의 모든 벤치마크 테스트를 실행합니다. 결과는 각 함수당 실행 시간과 반복 횟수 등이 출력됩니다.
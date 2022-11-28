# useb-api-sdk-go

useB API 는 고객의 신원인증을 API 호출을 통해 간단하게 조회 할 수 있도록 제공하는 서비스입니다.

- 이 모듈은 useB에서 제공하는 REST API를 [Golang](https://go.dev/)으로 구현한 http client SDK 입니다.
- Golang 환경에서 별도의 API 호출 클라이언트를 구현하지 않고 간편하게 useB API를 호출할 수 있습니다.
- 관련 API 문서는 [useB API 문서](https://docs.useb.co.kr/)를 참고하세요.

## 설치하기

아래 명령어를 통해 `useb-api-sdk-go`를 설치할 수 있습니다.

```bash
$ go get github.com/useb-inc/useb-api-sdk-go
```

## API 호출 예제

- `github.com/useb-inc/useb-api-sdk-go` 모듈을 설치하면 `usebapi` 패키지를 사용할 수 있습니다.
- 모든 API 요청에 대한 요청 파라미터는 각자 `github.com/useb-inc/useb-api-sdk-go/interfaces/{API_리소스}` 패키지 내에 구조체로 정의되어 있습니다.
- 예를들어, 주민등록증 진위확인 API 호출에 필요한 요청 파라미터는 `github.com/useb-inc/useb-api-sdk-go/interfaces/status` 패키지 내에 `IdcardRequest` 구조체로 정의되어 있습니다.
- API 응답이 `200 OK` 가 아닐 경우, err 값으로 응답 코드와 응답 Body 메시지를 반환합니다.

```go
package main

import (
	"fmt"
	"github.com/useb-inc/useb-api-sdk-go"
	"github.com/useb-inc/useb-api-sdk-go/interfaces/status"
)

func main() {
	// useB API 호출을 위한 클라이언트 생성
	api, err := usebapi.NewUsebAPI("YOUR_API_KEY", "YOUR_API_SECRET")
	if err != nil {
		panic(err)
	}

	req := status.IdcardRequest{
		Identity:  "YOUR_IDENTITY",
		IssueDate: "YOUR_ISSUE_DATE",
		UserName:  "YOUR_USER_NAME",
	}
	res, err := api.StatusAPI.Idcard(&req)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res.Message))
}
```

## API 리소스 명칭

- `openbank` - 1원계좌인증(오픈뱅킹)
- `firmbank` - 1원계좌인증(펌뱅킹)
- `firmbankpremium` - 1원계좌인증(펌뱅킹-프리미엄)
- `ocr` - OCR 진위확인
- `masking` - 마스킹
- `face` - 안면인증
- `status` - 진위확인

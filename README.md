# go_google_images
get google images using go

## 기능
- cobra로 param을 입력 받는다.
- 입력받은 검색어를 아래 URL로 google에 요청한다.
  > "https://www.google.co.in/search?q="+QUERY+"&source=lnms&tbm=isch"
- html을 파싱해서 검색 결과 이미지 목록을 추출한다.
- 이미지를 다운 받아 경로에 저정한다.
- 파일명은 "검색어_0000." + 이미지 형식으로 저장한다.

## Usage
```
Usage:
  google-image-downloader [flags]
  google-image-downloader [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Download google images.
  help        Help about any command
  version     Display version number.

Flags:
  -h, --help   help for google-image-downloader
```

## Code Review

### Command
- [jeager](https://github.com/jaegertracing/jaeger/blob/4fc291568d8ac59a1c67cc47ee1d91ab20dd06c4/cmd/docs/command.go#L36)의 소스 코드를 보다보니, cobra라는 package를 사용하고 있었고 검색해보니, 복잡한 파라미터를 사용하는 cli app 개발에 유용하며 kubeneties, jeager, promethtus등 go로 만들어진 많은 app 에서 사용하고 있다고 한다. cobra를 활용한 간단한 프로젝트를 해보려고한다.

#### 사용법
 
  1. cmd라는 하위 디렉토리를 만들고, 추가한 command 마다 소스 파일을 구분해준다. 여기서는 ```get```, ```version``` 두개의 command를 만들었다.

  2. ```root```는 ```get```, ```version``` 두개의 하위 command를 가진다.
  ```golang:cmd/root.go [27-30]
  func init() {
    rootCmd.AddCommand(version)
    rootCmd.AddCommand(get)
  }
  ``` 
 
  3. main 에서 ```root```의 ```Execute()```를 호출한다.
  ```go:main.go [7-11]
  cmd.Execute()
  ```
 
  4. command 구조체에서 정의한 내용들을 필요에 따라 호출된다.
    - Use: 한줄짜리 Usage 메시지
    - Short: help에서 출력되는 짧은 메시지
    - Long: help에서 출려되는 긴 메시지 'help <현재-command>'와 같이 출력된다.
    - Example: 사용예 출력메시지 
    - Args: Run을 실행하기 전에 파라미터 유효성 검사를 하기 위한 함수. nil을 리턴해야 Run이 실행된다.
    - Run: 현재-command가 입력되었을때 실행한다.

### Crawling

  - html 파싱은 [goquery](https://pkg.go.dev/github.com/PuerkitoBio/goquery@v1.8.0)를 사용했다.
  - 'Search Result'라는 text 이후에 나오는 img 태그를 찾고, src 또는 data-src 속성의 값을 saveUrlImage 함수로 넘긴다
  - 이미지 저장시간이 오래 걸리 수 있으므로 go 루틴을 사용해서 병렬화하고
  - 전체 이미지가 모두 다운받은 후 프로그램을 종료해야하므로 sync.WaitGroup을 만들어 대기한다.

### Image Download

  - html body에 데이터를 이미지 파일로 저장한다. 
  - html 해더의 `Content-Type`으로 이미지 파일 확장자를 결정한다.

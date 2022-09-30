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
# blog-program with Go
프론트엔드 실습을 위한 블로그 백엔드 바이너리 프로그램

프론트엔드에서 빠른 api 호출 실습을 하기 위해 간단한 블로그 백엔드를 바이너리 프로그램으로 제작.


## 사용방법
1. `bin/` 폴더에서 운영체제에 맞는 프로그램 선택 후 다운로드
2. 클릭해서 실행


## API
### Endpoint : `http://127.0.0.1:8080`


### `/posts/`
![](https://img.shields.io/static/v1?label=method&message=GET&color=3688ff)


### `/post/create`
![](https://img.shields.io/static/v1?label=method&message=POST&color=3688ff)
`application/json` `charset=utf-8`
<table>
<tr>
  <td>data</td>
  <td>name</td>
  <td>type</td>
  <td>required</td>
  <td>description</td>
</tr>
<tr>
  <td>json</td>
  <td>title</td>
  <td>string</td>
  <td>o</td>
  <td>블로그 포스트의 제목에 해당</td>
</tr>
<tr>
  <td>json</td>
  <td>content</td>
  <td>string</td>
  <td>o</td>
  <td>블로그 포스트의 내용에 해당</td>
</tr>
</table>


### `/post/:id`
![](https://img.shields.io/static/v1?label=method&message=GET&color=3688ff)

### `/post/:id`
![](https://img.shields.io/static/v1?label=method&message=DELETE&color=3688ff)


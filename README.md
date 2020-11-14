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
#### return  
```json
{
  "posts": [
    {
      "ID": 1,
      "CreatedAt": "2020-11-13T21:51:29.620665+09:00",
      "UpdatedAt": "2020-11-13T21:51:29.620665+09:00",
      "DeletedAt": "0001-01-01T00:00:00Z",
      "title": "블로그 시작하겠습니다.",
      "content": "저의 블로그는 개발, 독서, 축구에 관련된 이야기들이 올라올 예정입니다."
    },
    {
      "ID": 2,
      "CreatedAt": "2020-11-13T21:56:04.523552+09:00",
      "UpdatedAt": "2020-11-13T21:56:04.523552+09:00",
      "DeletedAt": "0001-01-01T00:00:00Z",
      "title": "첫 번째 글은 바로 \"규칙 없음\" 리뷰입니다.",
      "content": "규칙 없음은 넷플릭스의 공동 창립자이자 현재까지 넷플릭스를 이끌고 있는 리드 헤이스팅스의 첫 글로..."
    },
  ]
}
```  

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

#### result
```json
{
  "post": {
    "ID": 4,
    "CreatedAt": "2020-11-14T09:51:40.640264+09:00",
    "UpdatedAt": "2020-11-14T09:51:40.640264+09:00",
    "DeletedAt": "0001-01-01T00:00:00Z",
    "title": "제목",
    "content": "내용"
  },
  "result": "success"
}
```  


### `/post/:id`
![](https://img.shields.io/static/v1?label=method&message=GET&color=3688ff)


#### result
```json
{
  "post": {
    "ID": 4,
    "CreatedAt": "2020-11-14T09:51:40.640264+09:00",
    "UpdatedAt": "2020-11-14T09:51:40.640264+09:00",
    "DeletedAt": "0001-01-01T00:00:00Z",
    "title": "제목",
    "content": "내용"
  }
}
```  
### `/post/:id`
![](https://img.shields.io/static/v1?label=method&message=DELETE&color=3688ff)


#### result
```json
{
  "result": "success"
}
```  

# query-example
- mongodb atlas에 연결하고 쿼리를 보내는 방법에 대한 예시
- 데이터베이스 조회시 자유롭게 변형해서 사용

## Preparation
- .bashrc에 DBID, DBPW, DBADDR 환경변수를 정의합니다. (따로 문의주세요)
- 몽고디비를 위한 라이브러리를 설치합니다
```bash
go get go.mongodb.org/mongo-driver
```
- [mongodb-with-go-tutorial](https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial)
- DNS error 발생시 /etc/resolve.conf 수정 [참조](https://stackoverflow.com/questions/55660134/cant-connect-to-mongo-cloud-mongodb-database-in-golang-on-ubuntu)
- github.com/go-some/crawler도 설치합니다

## Installation 
- [workspace-path] 밑에서 go-some crawler를 설치합니다
```bash
go get -u github.com/go-some/query-example
```

## Run
- go build 후에 생기는 바이너리 실행

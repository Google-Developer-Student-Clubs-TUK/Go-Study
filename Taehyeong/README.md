### GO-Study (이태형)

## 초기 설정
- study의 주제에 따라 폴더로 분리
- package 와 폴더 이름은 통일
- package 내 print 파일을 제외한 모든 파일은 private function 으로 구성

## 1. Go By Example
- link: https://gobyexample.com
- package: gobyexample
- exec: gobyexample.PrintExample("파일명")

## 2. Web_Server(gin) - album
- link: https://go.dev/doc/tutorial/web-service-gin
- package: album
- exec: album.Run()
- gin 을 이용해 album 웹서버 구축

## 3. DB_Server(gorm v2) - todo
- link: https://thedevsaddam.medium.com/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3
- package: todo
- exec: todo.Run()
- gorm_v2(mysql) 과 gin을 이용해 todo 웹서버 구현
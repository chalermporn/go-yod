---
marp: true
---
# <!--fit--> Golang by Yod
---
# GOLANG [go-yod]

Training golang by P'Yod
15 Feb 2020

## การติดตั้ง ภาษา Go

## วิธี run and compile
- การ compile ข้าม platform

## การเขียน Unit testing
- TDD concept

---
# <!--fit--> TDD concept
---
# วิธีการรัน test แบบ need tags
```sh
├── fizzbuzz
    ├── fizzbuzz.go
    ├── fizzbuzz_integration_test.go
    └── fizzbuzz_test.go
```
use diractory `fizzbuzz`
```sh
cd fizzbuzz
```
run test
```sh
go test -tags integration -run TestFizzBuzz1To10
```

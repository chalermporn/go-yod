---
marp: true
---
# <!--fit--> Golang by P'Yod
Training golang by P'Yod
15 - 16 Feb 2020
at THINK SOciety: Co-working Space & Cafe

---
## การติดตั้ง ภาษา Go

## วิธี run and compile
- การ compile ข้าม platform

## การเขียน Unit testing
- TDD concept

---
# <!--fit--> TDD concept
---
# วิธีการ Run test 
```sh
go test -v
```
```sh
=== RUN   TestFizzBuzzSayOrigin
=== RUN   TestFizzBuzzSayOrigin/given_1_say_1
--- PASS: TestFizzBuzzSayOrigin (0.00s)
    --- PASS: TestFizzBuzzSayOrigin/given_1_say_1 (0.00s)
=== RUN   TestFizzBuzzSayFizz
=== RUN   TestFizzBuzzSayFizz/given_3_say_Fizz
--- PASS: TestFizzBuzzSayFizz (0.00s)
    --- PASS: TestFizzBuzzSayFizz/given_3_say_Fizz (0.00s)
=== RUN   TestFizzBuzzSayBuzz
=== RUN   TestFizzBuzzSayBuzz/given_5_say_Buzz
--- PASS: TestFizzBuzzSayBuzz (0.00s)
    --- PASS: TestFizzBuzzSayBuzz/given_5_say_Buzz (0.00s)
PASS
ok      hello/fizzbuzz  0.192s
```
---
# วิธีการ Run test แบบ Need tags
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
go test -tags integration -run TestFizzBuzz1To100
```

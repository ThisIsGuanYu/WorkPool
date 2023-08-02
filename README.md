# WorkPool

***

### Introduction
In a hypothetical scenario, let's say you have a meat factory that deals with three types of meats: beef, pork, and chicken. Today, a batch of meats has arrived with 10 portions of beef, 7 portions of pork, and 5 portions of chicken.

You've got a team of five employees, labeled A, B, C, D, and E. Each employee works at the same speed: 1 second to handle beef, 2 seconds for pork, and 3 seconds for chicken. These employees work independently, and they can handle the meats one by one. Each person can only handle one piece of meat at a time, and each piece of meat can only be worked on by one person. Once a person takes a piece of meat, they can't give it back.

Your goal here is to have all 10 portions of beef, 7 portions of pork, and 5 portions of chicken processed by your team. Each employee will randomly choose and process pieces of meat.

### Simple Commend 

Run the program

```
go run main.go
```

Test the program 

```
go test
```

Build the program as binary

```
go build -o release/workpool
```

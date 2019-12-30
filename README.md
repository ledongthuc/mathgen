# Math generator

[![GoDoc](https://godoc.org/github.com/ledongthuc/mathgen?status.svg)](https://godoc.org/github.com/ledongthuc/mathgen)

As a dad, I frequently to play math quiz with my daughter.
Bu I'm too lazy to frequently to create math questions. This' reason why I create this lib/api/web, that I can help me to generate questions.
This repository will evolue frequently, I think, to support more topic.

## Use as library

Install library

```
  go get -u github.com/ledongthuc/mathgen
```

Usage:

```
  questionAndAnswers, _ := mathgen.AddIntegerN(3, 20)
  fmt.Println("Addends:", questionAndAnswers.Addends)
  fmt.Println("Sum:", questionAndAnswers.Sum)
  fmt.Println("Format:", questionAndAnswers.String())
```

More detail supported functions is cover at: https://godoc.org/github.com/ledongthuc/mathgen

## Run demo application

### Startup

Start with docker:

```
docker run -p 8080:8080 ledongthuc/mathgen-web:latest
```

### API

POST `http://localhost:8080/api/addition/generate`

Request:
```
{
  "number_of_addend": 2,
  "max_sum": 10
}
```

Response:
```
{
    "addends": [
        1,
        1
    ],
    "sum": 2
}
```

### Web

Local: http://localhost:8080

![Demo](https://user-images.githubusercontent.com/1828895/71580301-55574b00-2b00-11ea-9b2f-ace31dc5cbb7.gif)

# lets-go

## Hello World!

```go
package main

import "fmt"

func main(){
	fmt.Println("Hello World!") //Hello World!
}
```

### Uyarı!

```go
package main

func main(){
	var name string = "mustafa" //<-- Küçük harf private
	var Surname string = "dikyar" //<-- Büyük harf public
}
```

## Değişkenler

```go
package main

import "fmt"

var global1 = "örnek"
//global2 := "örnek" => hata : .\variables.go:6:1: syntax error: non-declaration statement outside function body

func main(){
	fmt.Println(global1, /*global2*/)

	//Usage
	var message1 string
	message1 = "Hello 1!"
	fmt.Println(message1) //Hello 1!

	//Usage 2
	var message2 string = "Hello 2!"
	fmt.Println(message2) //Hello 2!

	//Usage 3
	var message3 = "Hello 3!"
	fmt.Println(message3) //Hello 3!

	//Usage 4
	var number1, number2, number3 int = 1,2,3
	fmt.Println(number1, number2, number3) //*1 *2 *3

	//Usage 5
	var number, name, gender = 2021, "Mustafa", false
	fmt.Println(number, name, gender) //*2021 *Mustafa *false

	//Usage 6
	surname := "Dikyar"
	fmt.Println(surname) //Dikyar

	//Usage 7
	age, language, money := 40, "golang", 100.5
	fmt.Println(age, language, money) //*40 *golang *100.5

	//Usage 8
	letter := 'M'
	word := "kelime"
	sentence := `Bu bir cümledir.`
	fmt.Println(letter, word, sentence) //*77 *kelime *Bu bir cümledir.
}
```

```go
package main

import "fmt"

var (
	name = "Mustafa"
	surname = "Dikyar"
)

func main(){
	fmt.Println(name, surname) //Mustafa Dikyar
}
```

## Sabitler

```go
package main

import "fmt"

const name = "Mustafa"
const pi = 3.14
func main(){
	fmt.Println(name, pi) //Mustafa 3.14
}
```

## Operatörler

### Aritmatik Operatörler

| Operatör      |  Açıklama   |
| ------------- | ----------- |
| +             | Toplama     |
| -             | Çıkarma     |
| /             | Bölme       |
| *             | Çarpma      |
| %             | Mod         |
| ++            | Değeri bir artırır. |
| --            | Değeri bir azaltır. |


### İlişkisel Operatörler

| Operatör      |  Açıklama   |
| ------------- | ----------- |
| ==            | Eşit        |
| !=            | Eşit değil  |
| >             | Büyüktür    |
| <             | Küçüktür    |
| >=            | Büyük eşit  |
| <=            | Küçük eşit  |


### Mantıksal Operatörler

| Operatör      |  Açıklama   |
| ------------- | ----------- |
| &&            | ve          |
| ||            | veya        |
| !             | değil       |


### Temel Atama Operatörler

| Operatör      |  Açıklama   |
| ------------- | ----------- |
| =             | Sağdaki değeri sol tarafa atar.                                         |
| +=            | Sağdaki değeri sol taraftaki değere ekler ve soldaki değişkene atar.    |
| -=            | Sağdaki değeri sol taraftaki değerden çıkarır ve soldaki değişkene atar.|
| *=            | Sağdaki değeri soldaki değer ile çarpar ve soldaki değişkene atar.      |
| /=            | Sağdaki değeri soldaki değere böler ve soldaki değişkene atar.          |



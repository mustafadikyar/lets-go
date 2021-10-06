# lets-go

Herkese selam, **GoLang** öğrenmeye başladığım bu günlerde hem *Türkçe kaynak* oluşturmak hem de kendi adıma öğrenmenin 
en etkili yolunun birilerine anlatmak olduğunu düşündüğümden şuan incelemekte olduğunuz kaynak ortaya çıkmış bulunuyor. 

<details>
<summary>Hello World!</summary>
<br>
	
<a name="hello-world" />
	
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
</details>




<details>
<summary>Değişkenler</summary>
<br>
	
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
</details>

<details>
<summary>Sabitler</summary>
<br>

```go
package main

import "fmt"

const name = "Mustafa"
const pi = 3.14
func main(){
	fmt.Println(name, pi) //Mustafa 3.14
}
```
</details>

<details>
<summary>Operatörler</summary>
<br>
	

Aritmatik Operatörler

| Operatör      |  Açıklama   |
| ------------- | ----------- |
| +             | Toplama     |
| -             | Çıkarma     |
| /             | Bölme       |
| *             | Çarpma      |
| %             | Mod         |
| ++            | Değeri bir artırır. |
| --            | Değeri bir azaltır. |


İlişkisel Operatörler

| Operatör      |  Açıklama   |
| ------------- | ----------- |
| ==            | Eşit        |
| !=            | Eşit değil  |
| >             | Büyüktür    |
| <             | Küçüktür    |
| >=            | Büyük eşit  |
| <=            | Küçük eşit  |


Mantıksal Operatörler

| Operatör      |  Açıklama   |
| ------------- | ----------- |
| &&            | ve          |
| ||            | veya        |
| !             | değil       |


Temel Atama Operatörler

| Operatör      |  Açıklama   |
| ------------- | ----------- |
| =             | Sağdaki değeri sol tarafa atar.                                         |
| +=            | Sağdaki değeri sol taraftaki değere ekler ve soldaki değişkene atar.    |
| -=            | Sağdaki değeri sol taraftaki değerden çıkarır ve soldaki değişkene atar.|
| *=            | Sağdaki değeri soldaki değer ile çarpar ve soldaki değişkene atar.      |
| /=            | Sağdaki değeri soldaki değere böler ve soldaki değişkene atar.          |

</details>
	
<details>
<summary>Break Identifier</summary>
<br>
     
```go
package main

import "fmt"

func main(){
	
	/*GetData metotu ya veri dönecek ya hata dönecek olduğunu varsayarsak.
	İki değerden birisi null gelecek hata alacağız.
	Sadece data var mı yok mu ile ilgilenmek istediğimiz zaman _ (break identifier) ile boş bir tanımlayıcı kullanabiliriz.
	*/

	//errorCode, data := GetData()
	//_, data := GetData()

	var _, number, _ = 0, 15, ""
	fmt.Println(number) //15
}
```
	
</details>

<details>
<summary>Enums</summary>
<br>
	
Doğrudan bir enum kullanımı yok.

```go
package main

import "fmt"

type Language string
const(
	Turkish Language = "tr"
	English Language = "en"
	Russian Language = "ru"
)

func GetLanguage(language Language){
	fmt.Println(language)
}

func main(){
	GetLanguage(Turkish) //tr
	GetLanguage(Russian) //ru
}
```
	
</details>
	
<details>
<summary>Tür Dönüşümleri</summary>
<br>

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber = "10"

	//Convert

	//string to int
	number, _ := strconv.Atoi(firstNumber)
	result := number + 3
	fmt.Println(result) //13

	//int to string
	fmt.Println("Sonuç : " + strconv.Itoa(result)) //Sonuç : 13

	//Casting
	var secondNumber float64 = 11.1
	var intNumber = int(secondNumber)
	fmt.Println(intNumber) //11
}

```
	
</details>
	
<details>
<summary>Tarih-Saat İşlemleri</summary>
<br>

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	birthDate := time.Date(2021, time.April, 27, 20, 0, 0, 0, time.UTC) //custom tarih
	fmt.Println(birthDate)

	currentDate := time.Now()
	fmt.Println(currentDate)         //o anki tarih
	fmt.Println(currentDate.Year())  //o anki yıl
	fmt.Println(currentDate.Day())   //o anki gün
	fmt.Println(currentDate.Month()) //o anki ay

	newDay := birthDate.AddDate(0, 0, 1).Day()     //+1 gün ekler
	newMonth := birthDate.AddDate(0, 1, 0).Month() //+1 ay ekler
	newYear := birthDate.AddDate(1, 0, 0).Year()   //+1 yıl ekler
	fmt.Println(newDay, newMonth, newYear)
}

```
	
</details>
	
<details>
<summary>İf İfadesi</summary>
<br>

```go
package main

import "fmt"

func main() {
	number := 10

	if number > 0 {
		fmt.Println("Sayı pozitif")
	} else if number < 0 {
		fmt.Println("Sayı negatif")
	} else {
		fmt.Println("Sayı sıfıra eşit")
	}
}

```
</details>
	
<details>
<summary>Switch İfadesi</summary>
<br>
	
```go	
package main

import "fmt"

func main() {
	language := "de"

	switch language {
	case "tr":
		fmt.Println("Türkçe")
	case "en":
		fmt.Println("İngilizce")
	case "ru":
		fmt.Println("Rusça")
	default:
		fmt.Println("Diğer")
	}
}

```
</details>
	
<details>
<summary>For Döngüsü</summary>
<br>
	
```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```
</details>

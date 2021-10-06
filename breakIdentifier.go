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
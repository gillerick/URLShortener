package main


import (
	base62 "URLShortener/utils"
	"log"
)


func main()  {
	x := 1000
	base62String := base62.ToBase62(x)
	log.Println(base62String)
	normalNumber := base62.ToBase10(base62String)
	log.Println(normalNumber)

}

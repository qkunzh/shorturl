package main

import (
	"fmt"
	"shorturl/impl"
	"shorturl/util"
)

func main() {
	redisMap := util.NewRedisMap("localhost",6379)
	longURL := "https://www.google.com/search?q=go+error&oq=go+error&aqs=chrome..69i57j35i39j0l8.7363j0j7&sourceid=chrome&ie=UTF-8"
	conventor := impl.NewURLConventor(redisMap)
	shortURL,err:= conventor.Generator(longURL)
	if err == nil {
		fmt.Println(shortURL)
	}
	revertLong,_:= conventor.Revert(shortURL)
	fmt.Println(revertLong==longURL)
}

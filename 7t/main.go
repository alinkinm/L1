package main

import (
	"fmt"
	"github.com/goombaio/namegenerator"
	"strconv"
	"sync"
	"time"
)

type mapObj struct {
	name  string
	hello string
}

func main() {
	var mp = map[string]string{}
	names := make(chan mapObj)
	wg := &sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				seed := time.Now().UTC().UnixNano()
				nameGenerator := namegenerator.NewNameGenerator(seed)
				name := nameGenerator.Generate()
				str := strconv.Itoa(i)
				obj := mapObj{name: name, hello: "hello from " + str + " goroutine"}
				names <- obj
			}

		}(i, wg)
	}
	go func() {
		wg.Wait()
		//предупреждаем дедлок
		close(names)
	}()

	//почему обязательно делать это в мейне и не получается вынести в отдельную горутину
	//как конкурретно читать...
	for msg := range names {
		key := msg.name
		value := msg.hello
		mp[key] = value
		fmt.Println(mp[key])
		fmt.Println("\n")
	}

}

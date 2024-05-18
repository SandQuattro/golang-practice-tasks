package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type SiteData struct {
	URL     string
	Content string
}

func main() {
	site := make(chan SiteData)

	// некий процесс обрабатывает ссылки и возвращает контент
	// по завершении закрываем канал и выходим
	go func() {
		for i := 0; i < 10; i++ {
			site <- SiteData{fmt.Sprintf("url%d", i), fmt.Sprintf("content%d", i)}
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
		close(site)
	}()

	// Запуск worker-ов
	data := processSiteData(site, 3)

	// Выводим данные с сайтов
	fmt.Println(data)
}

// выгребаем данные сайтов из канала
func processSiteData(c <-chan SiteData, numWorkers int) map[string]SiteData {
	sites := make(map[string]SiteData)

	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		i := i
		go func() {
			log.Print(fmt.Sprintf("worker %d online", i+1))
			defer wg.Done()
			for {
				val, ok := <-c
				if !ok {
					// канал закрыт, выходим
					break
				}
				log.Print(fmt.Sprintf(">> worker %d did the job:%s", i+1, val))
				sites[val.URL] = val
			}

		}()
	}
	wg.Wait()
	return sites
}

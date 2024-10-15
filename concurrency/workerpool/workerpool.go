package main

import (
	"fmt"
	"log"
	"runtime"
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
			time.Sleep(time.Duration(i) * time.Millisecond * 10)
		}
		close(site)
	}()

	// Запуск worker-ов
	for data := range processSiteData(site, runtime.NumCPU()) {
		// Выводим данные с сайтов
		log.Println(data)
	}

}

// выгребаем данные сайтов из канала
func processSiteData(c <-chan SiteData, numWorkers int) <-chan string {
	sites := make(chan string, numWorkers)

	// оптимизация, мы сразу асинхронно вернем канал в main и начнем из него читать
	go func() {
		log.Println("Starting workers:", numWorkers)
		wg := sync.WaitGroup{}
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			i := i
			go func() {
				log.Printf("worker %d online", i+1)
				defer wg.Done()
				for {
					val, ok := <-c
					if !ok {
						// канал закрыт, выходим
						break
					}
					log.Printf(">> worker %d did the job:%s", i+1, val)
					sites <- val.Content
				}
			}()
		}

		wg.Wait()
		close(sites)
	}()

	return sites
}

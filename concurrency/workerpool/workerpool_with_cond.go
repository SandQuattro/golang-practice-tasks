package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type CondSiteData struct {
	URL     string
	Content string
}

var (
	processingDone = false
	wrks           = runtime.NumCPU()
)

// RULE: Wait should be ALWAYS before Broadcast, otherwise workers will wait forever - deadlock
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	dataSlice := make([]CondSiteData, 0)

	site := make(chan CondSiteData)
	go func() {
		// Запуск worker-ов
		for data := range processCondSiteData(cond, site, wrks) {
			// Выводим данные с сайтов
			log.Println(data)
		}
	}()

	// некий процесс обрабатывает ссылки и возвращает контент, процесс долгий
	// Зпускаем горутины и ожидаем получения данных
	// по завершении закрываем канал и выходим
	for i := 0; i < 10; i++ {
		dataSlice = append(dataSlice, CondSiteData{fmt.Sprintf("url%d", i), fmt.Sprintf("content%d", i)})
		time.Sleep(time.Duration(i) * time.Millisecond * 100)
	}

	// Данные обработаны, надо разбудить горутины
	cond.L.Lock()
	processingDone = true
	cond.L.Unlock()
	cond.Broadcast()

	// Пишем наши данные в канал и закрываем его
	for i := 0; i < 100; i++ {
		for _, data := range dataSlice {
			site <- data
		}
	}
	close(site)
}

// выгребаем данные сайтов из канала
func processCondSiteData(cond *sync.Cond, c <-chan CondSiteData, numWorkers int) <-chan string {
	wg := sync.WaitGroup{}

	log.Println("Starting workers:", numWorkers)

	// создаем канал и тут же его возвращаем, он сразу доступен для чтения снаружи
	sites := make(chan string, numWorkers)

	go func() {
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			i := i
			go func() {
				defer wg.Done()
				log.Printf("worker %d online, waiting...", i+1)

				// waiting for broadcast signal, and we ready to go...
				cond.L.Lock()
				for !processingDone {
					cond.Wait()
				}
				cond.L.Unlock()

				for {
					val, ok := <-c
					if !ok {
						// канал закрыт, выходим
						break
					}
					log.Printf(">> worker %d did the job:%s", i+1, val)
					sites <- val.Content
				}
				log.Printf("!! worker %d done", i+1)
			}()
		}
		wg.Wait()
		close(sites)
	}()

	return sites
}

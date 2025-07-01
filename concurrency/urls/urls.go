package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

/*
http://ёёёё – not ok (Get "http://%D1%91%D1%91%D1%91%D1%91": dial tcp: lookup xn--61aaaa: no such host)
http://somesite.com – ok
http://non-existent.domain.tld – not ok (Get "http://non-existent.domain.tld": dial tcp: lookup non-existent.domain.tld: no such host)
https://ya.ru – ok
https://ozon.ru – not ok (Get "https://ozon.ru/?__rr=2": context canceled)
http://ozon.ru – not ok (Get "https://ozon.ru/?__rr=2": context canceled)
http://google.com – not ok (Get "http://www.google.com/": context canceled)
*/

var urls = []string{
	"http://ozon.ru",
	"https://ozon.ru",
	"http://google.com",
	"http://somesite.com",
	"http://non-existent.domain.tld",
	"https://ya.ru",
	"http://ya.ru",
	"http://ёёёё",
}

func main() {
	// Максимум одновременных запросов (настраивается).
	const maxParallel = 4

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := &http.Client{Timeout: 5 * time.Second}

	var okCount int32
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxParallel) // лимитер

	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			sem <- struct{}{}        // занимаем слот
			defer func() { <-sem }() // освобождаем

			// Если уже найдены два успешных ответа — выходим сразу.
			if atomic.LoadInt32(&okCount) >= 2 {
				return
			}

			// моментальная проверка отмены
			select {
			case <-ctx.Done():
				return // работы уже не требуется
			default:
			}

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				fmt.Printf("%s – not ok (bad url: %v)\n", url, err)
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("%s – not ok (%v)\n", url, err)
				return
			}
			resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				if atomic.AddInt32(&okCount, 1) == 2 {
					// получили второй 200 – отменяем оставшиеся запросы
					cancel()
				}
				fmt.Printf("%s – ok\n", url)
			} else {
				fmt.Printf("%s – not ok (status %d)\n", url, resp.StatusCode)
			}
		}(u)
	}

	wg.Wait()
}

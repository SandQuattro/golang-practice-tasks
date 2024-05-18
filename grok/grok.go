package main

import "log"
import "github.com/vjeantet/grok"

func main() {
	g, err := grok.New()
	if err != nil {
		log.Panic("Error initializing patterns processor, ", err)
	}

	txt := "2023/06/25 17:21:24 [error] 2097633#2097633: *20397 open() \"/var/www/jwebhelp.ru/www.jwebhelp.ru/htdocs/xmls/owa/auth/logon.aspx.xml\" failed (2: No such file or directory), client: 162.243.134.10, server: , request: \"GET /owa/auth/logon.aspx?url=https%3a%2f%2f1%2fecp%2f HTTP/1.1\", host: \"213.239.217.186\""
	pattern := "(?P<timestamp>%{YEAR}/%{MONTHNUM}/%{MONTHDAY} %{TIME}) \\[%{WORD:log_level}\\] %{NUMBER:nginx_pid}#%{NUMBER:request_id}: %{GREEDYDATA:error_message}, client: %{IPV4:client_ip}, server: %{GREEDYDATA:server_name}, request: \"%{WORD:http_method} %{URIPATH:path} HTTP/%{NUMBER:http_version}\", host: \"%{HOSTNAME:host}\""

	data, err := g.Parse(pattern, txt)
	if err != nil {
		log.Panic("error ", err)
	}

	if len(data) == 0 {
		log.Println("construct default ")

		pattern := "(?P<timestamp>%{YEAR}/%{MONTHNUM}/%{MONTHDAY} %{TIME}) %{GREEDYDATA:message}"
		data, err = g.Parse(pattern, txt)

		//data["message"] = txt
	}
	log.Println(data)
}

package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func Titulo(urls ...string) <-chan string {
	ch := make(chan string)
	for _, url := range urls {
		//semelhante ao length
		go func(url string){
			resp, _ := http.Get(url)
			//chama a estrutura da url
			html, _ := ioutil.ReadAll(resp.Body)
			//ioutil - leitura (manipulacao de buffer)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			//expressao regular - autômato
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return ch
}
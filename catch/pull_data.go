package catch

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

//获取网页文本对象
func PullData(url string) (*goquery.Document, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil,err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11")
	client := http.DefaultClient

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil,err
	}
	defer response.Body.Close()

	return goquery.NewDocumentFromReader(response.Body)
}

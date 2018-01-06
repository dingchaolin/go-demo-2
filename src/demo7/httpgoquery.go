package main

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"errors"
)
func cleanUrls( u string, urls []string)[]string{

}
func fetch( url string) ([]string, error){
	// http://www.baidu.com/index/name=dcl/#1
	// schema :// host / path / queryString # 锚点
	//url := "https://baike.baidu.com/item/鲁迅/36231?fr=aladdin"
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()//此处需要关闭 否则会造成资源泄露
	// 2XX 正常  3XX 重定向 4XX 客户端错误 5XX 服务器错误
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New( resp.Status )
	}
	doc, err := goquery.NewDocumentFromResponse( resp )
	if err != nil {
		return nil, err
	}
	urls := []string{}
	doc.Find("img").Each(func( i int, s *goquery.Selection){
		link, ok := s.Attr("src")
		if ok && link[:4] == "http"{
			urls = append( urls, link )
		}else{
			fmt.Println( "src not found")
		}

	})

	return urls, nil

}

// 链接的形式
// http://xxx.com.jpg
// //xx.com/a.jpg
// /static/a.jpg
// a.jpg
func main(){
	url := "https://baike.baidu.com/item/鲁迅/36231?fr=aladdin"
	urls, err := fetch(url)
	if err != nil {
		log.Fatal( err )
	}
	for _, u := range urls{
		fmt.Println( u )
	}
}
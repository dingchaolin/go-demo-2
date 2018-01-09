package main

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"errors"
	"net/url"
	"strings"
	"io/ioutil"
	"os"
	"path/filepath"
	"path"
	"io"
	"time"
	"archive/tar"
	"compress/gzip"
	"flag"
)

var (
	label = flag.String("label", "img", "label to download")
)

var labelAttrMap = map[string]string{
	"img":"src",
	"script" : "src",
	"a":"href",
}

// 链接的形式
// http://xxx.com.jpg
// //xx.com/a.jpg
// /static/a.jpg
// a.jpg
func CleanUrl(uri *url.URL, link string) string {
	switch {
	case strings.HasPrefix(link, "https") || strings.HasPrefix(link, "http"):
		return link
	case strings.HasPrefix(link, "//"):
		return uri.Scheme + ":" + link
	case strings.HasPrefix(link, "/"):
		return fmt.Sprintf("%s://%s%s", uri.Scheme, uri.Host, link)
	default:
		p := strings.SplitAfter(uri.Path, "/")
		path := strings.Join(p[:2], "")
		return fmt.Sprintf("%s://%s%s%s", uri.Scheme, uri.Host, path, link)
	}
}

func cleanUrls(u string, urls []string) []string {
	var ret [] string
	uri, _ := url.Parse(u)
	for i := range urls {
		ret = append(ret, CleanUrl(uri, urls[i]))
	}
	return ret

}
func fetch(url string) ([]string, error) {
	// http://www.baidu.com/index/name=dcl/#1
	// schema :// host / path / queryString # 锚点
	//url := "https://baike.baidu.com/item/鲁迅/36231?fr=aladdin"
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() //此处需要关闭 否则会造成资源泄露
	// 2XX 正常  3XX 重定向 4XX 客户端错误 5XX 服务器错误
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	urls := []string{}
	doc.Find(/*"img"*/*label).Each(func(i int, s *goquery.Selection) {
		//doc.Find("javascript").Each(func(i int, s *goquery.Selection) {//可以获取所有的js文件
		/*
		以此类推 所有的同类型的都可以下载 视频也可以
		 */
		link, ok := s.Attr(/*"src"*/ labelAttrMap[*label])
		if ok {
			urls = append(urls, link)
		} else {
			fmt.Println("src not found")
		}

	})

	return urls, nil

}

func downloadImgs(urls []string, dir string) error {
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			continue
			//return errors.New(resp.Status)
		}
		fullname := filepath.Join(dir, path.Base(u))
		f, err := os.Create(fullname)
		if err != nil {
			return err
		}
		defer f.Close()
		io.Copy(f, resp.Body)
	}
	return nil
}

func makeTar(dir string, w io.Writer) error {
	compress := gzip.NewWriter(w)
	defer compress.Close()
	//tr := tar.NewWriter(w) //创建一个可写的对象
	tr := tar.NewWriter( compress)//压缩
	defer tr.Close()
	filepath.Walk(dir, func(name string /*有路径信息*/ , info os.FileInfo, err error) error { //遍历文件
		// 写入tar的fileHeader
		// 以读取的方法打开文件
		// 判断目录和文件 如果是文件写入 其他跳过
		//把文件的内容写入到body
		header, err := tar.FileInfoHeader(info, "")//软连接为空
		if err != nil {
			return err
		}
		// name 中有路径信息 info.Name header.Name中没有路径信息 两者是一样的
		//fmt.Println(header.Name)
		/*
		name有几层目录 tar中也会有几层目录
		 */
		 basedir := filepath.Base(dir)

		//header.Name = name //将有路径信息的文件信息写入
		//info.Name 获取的是全路径
		p, _ := filepath.Rel(dir, name )//获取 name 到 dir的相对路径
		header.Name = filepath.Join(basedir, p )//以临时目录路径为最后一层目录
		tr.WriteHeader(header)
		//如果跳过 tar包中就会缺少一个目录
		//if info.IsDir() {
		//	return nil
		//}

		f, err := os.Open(name)//只读方式
		if err != nil{
			return err
		}
		defer f.Close()

		io.Copy( tr, f )
		return nil
	})
	return nil
}

func main() {
	flag.Parse()
	start := time.Now()
	url := "http://daily.zhihu.com"
	urls, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	urls = cleanUrls(url, urls)
	//for _, u := range urls{
	//	fmt.Println( u )
	//}

	tmpdir, err := ioutil.TempDir(".", "spider") //保证生成的目录是唯一的
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("临时文件路径为:", tmpdir)
	defer os.RemoveAll( tmpdir )//移除临时文件

	err = downloadImgs(urls, tmpdir)
	if err != nil {
		log.Fatal(err)
	}

	//f, err := os.Create("img.tar")
	f, err := os.Create("img.tar.gz")
	if err != nil {
		log.Fatal(err)
	}
	makeTar(tmpdir, f )//查看文件内容 tar tf img.tar

	useTime := time.Since(start)
	fmt.Println("耗时为:=====", useTime)
}
/*
命令
go run -label=img httpgoquery.go
 */
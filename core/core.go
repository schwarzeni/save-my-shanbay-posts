package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/schwarzeni/save-my-shanbay-posts/model"
	"github.com/schwarzeni/save-my-shanbay-posts/parser"
	"github.com/schwarzeni/save-my-shanbay-posts/parser/hugo"
	"github.com/schwarzeni/save-my-shanbay-posts/request/image"
	"github.com/schwarzeni/save-my-shanbay-posts/request/post"
	"github.com/schwarzeni/save-my-shanbay-posts/saver"
)

var config model.AppConfig

var tickTime = 1 * time.Second // 每过2秒请求一次

var wg sync.WaitGroup

func Run(fPath string) {
	readConfigFile(fPath)
	wg.Add(4)
	saveFile(
		processPost(
			fetchPost(
				fetchPostsList()),
			&hugo.HugoParser{BlogRoot: config.BlogPathRoot}))
	wg.Wait()
}

// 读取配置文件
func readConfigFile(f string) {
	var (
		e     error
		bytes []byte
	)
	log.Println("Read config file ....")
	if bytes, e = ioutil.ReadFile(f); e != nil {
		log.Fatal(e)
	}

	if e = json.Unmarshal(bytes, &config); e != nil {
		log.Fatal(e)
	}
}

// 获取文章列表
func fetchPostsList() chan []model.PostItem {
	var (
		err          error
		postList     model.ReqPostList
		postListChan = make(chan []model.PostItem, 2)
	)
	go func() {
		defer func() {
			log.Println("[fetchPostsList] 请求文章列表结束")
			close(postListChan)
			wg.Done()
		}()
		page := 1
		for {
			<-time.Tick(tickTime)
			log.Println("[fetchPostsList] 请求文章列表第", page, "页")
			if postList, err = post.GetPostList(config, page); err != nil {
				// TODO: handle error
				log.Fatal(err)
			}
			if len(postList.Data.Objects) == 0 {
				return
			}
			page++
			postListChan <- postList.Data.Objects
		}
	}()
	return postListChan
}

// 获取单个post的信息
func fetchPost(postListChan chan []model.PostItem) chan model.PostItem {
	var (
		postItemChan = make(chan model.PostItem, 2)
		err          error
		postItem     model.ReqPostItem
	)

	go func() {
		defer func() {
			log.Println("[fetchPost] 请求文章全部信息结束")
			close(postItemChan)
			wg.Done()
		}()
		f := hugo.HugoParser{}
		for items := range postListChan {
			for _, item := range items {
				<-time.Tick(tickTime)
				// TODO 考虑是否当文件不存在时才进行下载
				fp := f.GetSavePath(item)
				if _, err = os.Stat(fp); os.IsNotExist(err) {
					log.Println("[fetchPost] 开始请求文章，ID为 ", item.Id)
					if postItem, err = post.GetPostContent(item.Id); err != nil {
						// TODO: handle error
						log.Fatal(err)
					}
					postItemChan <- postItem.Data
				}
				// TODO: >
			}
		}
	}()

	return postItemChan
}

func processPost(postchan chan model.PostItem, parser parser.Parser) chan model.SaveFileInfo {
	var fileChan = make(chan model.SaveFileInfo, 1)
	go func() {
		defer func() {
			close(fileChan)
			log.Println("[processPost] porcess posts finish!")
			wg.Done()
		}()
		for p := range postchan {
			log.Println("[processPost] 处理文章，代号为：", p.Id, " 标题为：", p.Title)
			blogPostInfo, imageUrl := parser.Parse(p)
			for _, img := range imageUrl {
				log.Println("[processPost] 获取图片：", img.FromUrl)
				// 额外获取图片
				// 此处为sync获取
				if result, err := image.GetImage(img.FromUrl); err != nil {
					// TODO: handle error here
					log.Fatal(err)
				} else {
					// 将文件交由saver保存
					fileChan <- model.SaveFileInfo{
						Data:        result,
						FilePathStr: img.SavePath,
					}
				}
			}
			// 将文件交由saver保存
			fileChan <- model.SaveFileInfo{
				Data:        []byte(blogPostInfo.BasicInfo.Content),
				FilePathStr: blogPostInfo.SavePath,
			}

		}
	}()
	return fileChan
}

// 保存文件
func saveFile(fileChan chan model.SaveFileInfo) {
	go func() {
		defer func() {
			log.Println("[saveFile] finish")
			wg.Done()
		}()
		for f := range fileChan {
			if err := saver.Save(f); err != nil {
				// TODO: handle error here
				log.Fatal(err)
			}
			log.Println("[saveFile] save file ", filepath.Base(f.FilePathStr), " success!")
		}
	}()
}

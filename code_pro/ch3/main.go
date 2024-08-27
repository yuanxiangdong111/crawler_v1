package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "go_code/code_pro/bili"
    "go_code/code_pro/crawlerFactor"
)

type Crawler interface {
    crawlerFactor.AbstractCrawler
}

type CrawlerFactory struct {
    Crawlers map[string]Crawler
}

func (f *CrawlerFactory) CreateCrawler(platform string) (Crawler, error) {
    crawler, ok := f.Crawlers[platform]
    if !ok {
        return nil, fmt.Errorf("Invalid Media Platform Currently only supported xhs or dy or ks or bili ...")
    }
    return crawler, nil
}

func main() {
    // platform := flag.String("platform", "xhs", "Media platform select (xhs | dy | ks | bili | wb)")
    // loginType := flag.String("lt", "qrcode", "Login type (qrcode | phone | cookie)")
    // crawlerType := flag.String("type", "search", "crawler type (search | detail)")
    // flag.Parse()

    factory := &CrawlerFactory{
        Crawlers: map[string]Crawler{
            // "xhs":   &XiaoHongShuCrawler{},
            // "dy":    &DouYinCrawler{},
            // "ks":    &KuaishouCrawler{},
            "bili": &bili.BilibiliCrawler{},
            // "wb":    &WeiboCrawler{},
        },
    }

    createCrawler, err := factory.CreateCrawler("bili")
    if err != nil {
        fmt.Println(err)
        return
    }
    createCrawler.InitConfig("bili", "qrcode", "search")

    err = createCrawler.StartCrawler()
    if err != nil {
        fmt.Println("err = ", err)
        return
    }

    // handle Ctrl+C
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    <-c
}

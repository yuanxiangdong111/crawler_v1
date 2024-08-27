package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/PuerkitoBio/goquery"
    "github.com/gocolly/colly/v2"
)

func main() {

    // 存储文件名
    fName := "douban_movie_top250.csv"
    file, err := os.Create(fName)
    if err != nil {
        log.Fatalf("创建文件失败 %q: %s\n", fName, err)
        return
    }
    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()
    // 写CSV头部
    writer.Write([]string{"片名", "导演", "编剧", "主演", "类型", "官方网站", "制片国家/地区", "语言", "上映日期", "片长", "又名", "IMDb"})

    // var movies []model.Movie
    c := colly.NewCollector(
        // 设置用户代理
        colly.Async(true),
        // colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
    )

    // c.Limit(&colly.LimitRule{
    //     DomainGlob:  "*",
    //     RandomDelay: 1 * time.Second, // 随机延迟
    // })

    c.OnRequest(func(r *colly.Request) {
        // r.Headers.Set("cookie", "bid=_NDUeNU0K9s; _pk_id.100001.4cf6=dd862285e99ecf33.1704551498.; __utmc=30149280; __utmc=223695111; __yadk_uid=51UZ80RGKei0BxzYvIWCtOu6kfv1yNSv; ll=\"108288\"; _vwo_uuid_v2=DB67E2BBCD060A47DA5959803812BFCB7|365d659f82480719bd34ec471eca6195; dbcl2=\"234905083:PmndCKQ6tjo\"; ck=33f0; __utma=30149280.2009793875.1704551498.1704622350.1704626738.5; __utmb=30149280.0.10.1704626738; __utmz=30149280.1704626738.5.3.utmcsr=accounts.douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utma=223695111.791405552.1704551498.1704622350.1704626738.5; __utmb=223695111.0.10.1704626738; __utmz=223695111.1704626738.5.3.utmcsr=accounts.douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1704626738%2C%22https%3A%2F%2Faccounts.douban.com%2F%22%5D; _pk_ses.100001.4cf6=1; push_doumail_num=0; push_noty_num=0")
        log.Println("start visiting:", r.URL.String())
    })

    // newC := c.Clone()
    // 每一页的处理
    c.OnHTML("ol.grid_view", func(e *colly.HTMLElement) {
        e.DOM.Find("li").Each(func(_ int, selection *goquery.Selection) {
            url, ok := selection.Find("div.hd > a").Attr("href")
            fmt.Println("url = ", url)
            // 找到 a 元素里面的 href 内容了
            if ok {
                // newC.Visit(url)
                c.Visit(e.Request.AbsoluteURL(url))
            }
        })
    })
    // c.OnHTML("div#content div#info", func(e *colly.HTMLElement) {
    c.OnHTML("div#content", func(e *colly.HTMLElement) {
        // 定义结构体
        var movieMap = map[string]string{
            "片名":          "title",
            "导演":          "director",
            "编剧":          "screenwriter",
            "主演":          "starring",
            "类型":          "type",
            "官方网站":      "officialUrl",
            "制片国家/地区": "country",
            "语言":          "language",
            "上映日期":      "releaseDate",
            "片长":          "duration",
            "又名":          "alias",
            "IMDb":          "IMDb",
        }

        // 查找电影名 并且清洗整理数据
        title := e.DOM.Find("h1>span").First().Text()
        title = strings.TrimSpace(title)
        titles := strings.Split(title, " ")

        e1 := e.DOM.Find("div#info")
        if len(e1.Text()) == 0 {
            return
        }
        Text := strings.TrimSpace(e1.Text())
        Text = strings.ReplaceAll(Text, " ", "")

        moviesInfo := strings.Split(Text, "\n")

        // 开头添加电影名
        moviesInfo = append([]string{"片名:" + titles[0]}, moviesInfo...)

        // 官方网站:https://www.facebook.com/StarWars/
        var order = []string{"片名", "导演", "编剧", "主演", "类型", "官方网站", "制片国家/地区", "语言", "上映日期", "片长", "又名", "IMDb"}
        for _, v := range moviesInfo {
            processStr := strings.SplitN(v, ":", 2)
            if len(processStr) < 2 {
                continue
            }
            mapKey := processStr[0]
            mapVal := processStr[1]
            if _, ok := movieMap[mapKey]; ok {
                movieMap[mapKey] = mapVal
            }
        }

        var resultStr []string
        for _, v := range order {
            resultStr = append(resultStr, movieMap[v])
            // fmt.Printf("%s: %s\n", v, movieMap[v])
        }
        writer.Write([]string{
            resultStr[0],
            resultStr[1],
            resultStr[2],
            resultStr[3],
            resultStr[4],
            resultStr[5],
            resultStr[6],
            resultStr[7],
            resultStr[8],
            resultStr[9],
            resultStr[10],
            resultStr[11],
        })

    })

    // // 查找下一页
    c.OnHTML("div.paginator > span.next", func(element *colly.HTMLElement) {
        href, found := element.DOM.Find("a").Attr("href")
        // 如果有下一页，则继续访问
        if found {
            element.Request.Visit(element.Request.AbsoluteURL(href))
        }
    })

    c.Visit("https://movie.douban.com/top250")
    c.Wait()
}

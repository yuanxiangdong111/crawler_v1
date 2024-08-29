package main

import (
    "fmt"
    "time"

    "github.com/playwright-community/playwright-go"
)

func main() {
    err := startCrawler()
    if err != nil {
        fmt.Println("startCrawler error = ", err)
    }
}

func startCrawler() error {
    // 启动浏览器
    pw, err := playwright.Run()
    if err != nil {
        return err
    }

    chromium := pw.Chromium

    browser, err := chromium.Launch(playwright.BrowserTypeLaunchOptions{
        Args:                 nil,
        Channel:              nil,
        ChromiumSandbox:      nil,
        Devtools:             nil,
        DownloadsPath:        nil,
        Env:                  nil,
        ExecutablePath:       nil,
        FirefoxUserPrefs:     nil,
        HandleSIGHUP:         nil,
        HandleSIGINT:         nil,
        HandleSIGTERM:        nil,
        Headless:             playwright.Bool(false),
        IgnoreAllDefaultArgs: nil,
        IgnoreDefaultArgs:    nil,
        Proxy:                nil,
        SlowMo:               nil,
        Timeout:              playwright.Float(3000000),
        TracesDir:            nil,
    })

    if err != nil {
        fmt.Println("1111")
        return err
    }

    // 添加User-Agent
    page, err := browser.NewPage(playwright.BrowserNewPageOptions{
        UserAgent: playwright.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"),
    })
    if err != nil {
        fmt.Println("2222")
        return err
    }

    url := "https://www.baidu.com"
    _, err = page.Goto(url, playwright.PageGotoOptions{
        Referer:   nil,
        Timeout:   playwright.Float(3000000),
        WaitUntil: nil,
    })
    if err != nil {
        fmt.Println("33333")
        return err
    }

    err = page.Locator("span > #kw.s_ipt").Fill("今天是周三?")
    if err != nil {
        fmt.Println("44444")
        return err
    }

    for i := 1; i < 10; i++ {
        time.Sleep(1 * time.Second)
        err = page.Mouse().Wheel(0, 200*float64(i))
    }

    // url := "http://bang.dangdang.com/books/bestsellers/01.00.00.00.00.00-24hours-0-0-2-1"
    // response, err := page.Request().Get(url, playwright.APIRequestContextGetOptions{
    //     Data:              nil,
    //     FailOnStatusCode:  nil,
    //     Form:              nil,
    //     Headers:           nil,
    //     IgnoreHttpsErrors: nil,
    //     MaxRedirects:      nil,
    //     Multipart:         nil,
    //     Params:            nil,
    //     Timeout:           nil,
    // })
    //
    // if err != nil {
    //     return err
    // }

    // body, err := response.Body()
    // if err != nil {
    //     return err
    // }

    // fmt.Println("body = ", string(body))

    time.Sleep(10000 * time.Second)

    return nil
}

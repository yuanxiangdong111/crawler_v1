package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"

    "github.com/playwright-community/playwright-go"
)

type Response struct {
    Code    int                    `json:"code"`
    Message string                 `json:"message"`
    Data    map[string]interface{} `json:"data"`
}

func main() {
    pw, err := playwright.Run()
    if err != nil {
        log.Fatalf("could not start playwright: %v", err)
    }

    chromium := pw.Chromium

    // 持久化浏览器
    context, err := chromium.LaunchPersistentContext("/Users/xd_yuan/go_code/saveDir1", playwright.BrowserTypeLaunchPersistentContextOptions{
        AcceptDownloads: playwright.Bool(true),
        Headless:        playwright.Bool(false),
        UserAgent:       playwright.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
    })
    if err != nil {
        log.Fatalf("could not launch browser: %v", err)
    }

    // 打开一个新的页面
    page := context.Pages()
    // if len(page) > 0 {
    //     page[0].Close()
    // }

    response, err := page[0].Goto("https://www.baidu.com")
    if err != nil {
        log.Fatalf("could not create page: %v", err)
    }

    fmt.Println("response = ", response)
    time.Sleep(10 * time.Second)

}

// func Pong1(page playwright.Page) (login_flag bool) {
//     // log_flag := false
//     // 检查登录状态
//     check_login_url := "https://api.bilibili.com/x/web-interface/nav"
//     pageRespons, _ := page.Request().Get(check_login_url, playwright.APIRequestContextGetOptions{
//         Data: nil,
//     })
//     if pageRespons.Status() != 0 {
//         log.Println("未登录,请登录")
//         return false
//     }
//     // 登录了,获取登录网站的的返回信息
//     pageRespons.JSON()
// }

func request(method string, url string, headers map[string]string) (string, string) {
    client := &http.Client{}
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        log.Fatalf("could not NewRequest: %v", err)
    }

    for k, v := range headers {
        req.Header.Add(k, v)
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalf("could not Do: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("could not ReadAll: %v", err)
    }

    var response Response
    err = json.Unmarshal(body, &response)
    if err != nil {
        log.Fatalf("could not Unmarshal: %v", err)
    }

    if response.Code != 0 {
        log.Fatalf("response.Code: %v", response.Code)
    }

    dataMap := response.Data
    wbi_img := dataMap["wbi_img"].(map[string]interface{})["img_url"].(string)
    wbi_sub_url := dataMap["wbi_sub_url"].(map[string]interface{})["wbi_sub_url"].(string)

    fmt.Println("wbi_img = ", wbi_img)
    fmt.Println("wbi_sub_url = ", wbi_sub_url)

    return wbi_img, wbi_sub_url

}

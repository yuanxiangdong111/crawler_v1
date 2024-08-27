package main

import (
    "bytes"
    "encoding/base64"
    "encoding/json"
    "errors"
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "io"
    "log"
    "net/http"
    "os"
    "os/exec"
    "strings"
    "time"

    "go_code/utils/files"

    "github.com/playwright-community/playwright-go"
)

type Response struct {
    Code    int                    `json:"code"`
    Message string                 `json:"message"`
    Data    map[string]interface{} `json:"data"`
}

type AutoGenerated struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    TTL     int    `json:"ttl"`
    Data    struct {
        Replies []struct {
            Content struct {
                Message string `json:"message"`
            } `json:"content"`
            Replies      []any `json:"replies"`
            ReplyControl struct {
                MaxLine  int    `json:"max_line"`
                TimeDesc string `json:"time_desc"`
                Location string `json:"location"`
            } `json:"reply_control"`
            Folder struct {
                HasFolded bool   `json:"has_folded"`
                IsFolded  bool   `json:"is_folded"`
                Rule      string `json:"rule"`
            } `json:"folder"`
            TrackInfo string `json:"track_info"`
        } `json:"replies"`
    } `json:"data"`
}

func main() {
    pw, err := playwright.Run()
    if err != nil {
        log.Fatalf("could not start playwright: %v", err)
    }

    userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"

    // 指定浏览器类型
    chromium := pw.Chromium

    // 设置浏览器参数
    // browser, err := chromium.Launch(playwright.BrowserTypeLaunchOptions{
    //     Headless: playwright.Bool(false),
    //     Timeout:  playwright.Float(1000),
    // })

    // 持久化浏览器
    context, err := chromium.LaunchPersistentContext("/Users/xd_yuan/go_code/saveDir1", playwright.BrowserTypeLaunchPersistentContextOptions{
        AcceptDownloads: playwright.Bool(true),
        UserAgent:       &userAgent,
        Headless:        playwright.Bool(false),
    })

    defer func() {
        if err = context.Close(); err != nil {
            log.Fatalf("could not close browser: %v", err)
        }

        if err = pw.Stop(); err != nil {
            log.Fatalf("could not stop Playwright: %v", err)
        }
    }()

    // 创建新页面
    pages := context.Pages()
    // if err != nil {
    //     log.Fatalf("could not create page: %v", err)
    // }

    // 获取第一个页面
    page := pages[0]

    // 防止被检测脚本
    ScriptPath := "/Users/xd_yuan/go_code/chromepd/Test/stealth.min.js"
    context.AddInitScript(playwright.Script{Path: playwright.String(ScriptPath)})

    // 设置页面参数 如：User-Agent
    page.SetExtraHTTPHeaders(map[string]string{
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
    })

    // cookies, _ := page.Context().Cookies()

    // fmt.Println("page.Context().Cookies() = ", cookies)

    // 从本地查看是否有cookie文件，如果有则直接读取cookie文件中的cookie
    // 检查是否登录
    flag := Pong1(page)

    if !flag {
        // 设置页面视图大小 如：1920*1080 访问移动端页面时需要设置
        if _, err = page.Goto("https://www.bilibili.com/video/BV1S5411i7Me"); err != nil {
            log.Fatalf("could not goto: %v", err)
        }
        // 如果没有登录，则进行登录操作
        // 点击登录按钮
        err = page.Locator(".right-entry__outside.go-login-btn div").Click()
        if err != nil {
            log.Fatalf("could not click: %v", err)
        }
        qrcode_img_selector := ".login-scan-box img"
        qrcodeStr := findLoginQrcode(page, qrcode_img_selector)
        showQRCode(qrcodeStr)

        // 等待上一步登录完成,将获取登录后的cookie存在page.Context().Cookies()中

        // 检查登录状态
        err = check_login_state(page)
        if err != nil {
            log.Fatalf("could not check login state: %v", err)
        }
    }

    // 访问目标页面 并且等待页面加载完成
    if _, err = page.Goto("https://www.bilibili.com/video/BV1S5411i7Me", playwright.PageGotoOptions{
        WaitUntil: playwright.WaitUntilStateLoad,
    }); err != nil {
        log.Fatalf("could not goto: %v", err)
    }

    // // 将//div[@class='right-entry__outside go-login-btn']//div 改成css
    // // 定位到登录的位置
    // err = page.Locator(".right-entry__outside.go-login-btn div").Click()
    // if err != nil {
    //     log.Fatalf("could not click: %v", err)
    // }
    // qrcode_img_selector := ".login-scan-box img"
    // qrcodeStr := findLoginQrcode(page, qrcode_img_selector)
    // showQRCode(qrcodeStr)
    //
    // // 等待上一步登录完成,将获取登录后的cookie存在page.Context().Cookies()中
    //
    // // 检查登录状态
    // err = check_login_state(page)
    // if err != nil {
    //     log.Fatalf("could not check login state: %v", err)
    // }

    page.OnResponse(func(response playwright.Response) {
        if strings.Contains(response.URL(), "https://api.bilibili.com/x/v2/reply/wbi/main?") {
            fmt.Println("response.URL() = ", response.URL())
            fmt.Println("response.Status() = ", response.Status())
            body, err := response.Body()

            if err != nil {
                log.Fatalf("could not create page: %v", err)
            }

            fmt.Println("response.Body() = ", string(body))
        }
    })

    // fmt.Println("登录成功,开始监听")
    // go func() {
    //     err = page.Route("**/x/v2/reply/wbi/main?*", func(route playwright.Route) {
    //         go func(gRoute playwright.Route) {
    //
    //             // 使用fetch 基于当前页面的上下文去访问
    //             fetch, _ := gRoute.Fetch()
    //
    //             var response1 AutoGenerated
    //             err2 := fetch.JSON(&response1)
    //             if err2 != nil {
    //                 log.Fatalf("could not create page: %v", err2)
    //             }
    //
    //             for _, v := range response1.Data.Replies {
    //                 fmt.Println(v.Content.Message)
    //             }
    //
    //             gRoute.Continue()
    //         }(route)
    //     })
    //     if err != nil {
    //         log.Fatalf("could not set route handler: %v", err)
    //     }
    // }()

    // time.Sleep(5 * time.Second)

    for {
        // fmt.Println("1111")
        _, _ = page.Evaluate(`window.scrollBy(0, 500);`)
        time.Sleep(1 * time.Second)

        // scrollTop, _ := page.Evaluate("document.documentElement.scrollTop")
        // scrollHeight, _ := page.Evaluate("document.documentElement.scrollHeight")
        // clientHeight, _ := page.Evaluate("document.documentElement.clientHeight")
        //
        // scrollTopFloat, ok := scrollTop.(float64)
        // if !ok {
        //     scrollTopFloat = float64(scrollTop.(int))
        // }
        //
        // scrollHeightFloat, ok := scrollHeight.(float64)
        // if !ok {
        //     scrollHeightFloat = float64(scrollHeight.(int))
        // }
        //
        // clientHeightFloat, ok := clientHeight.(float64)
        // if !ok {
        //     clientHeightFloat = float64(clientHeight.(int))
        // }
        //
        // if scrollTopFloat+clientHeightFloat >= scrollHeightFloat {
        //     break
        // }
    }

    // if err = browser.Close(); err != nil {
    //     log.Fatalf("could not close browser: %v", err)
    // }
    //
    // if err = pw.Stop(); err != nil {
    //     log.Fatalf("could not stop Playwright: %v", err)
    // }
}

func findLoginQrcode(page playwright.Page, selector string) string {
    elements, err := page.WaitForSelector(selector)
    if err != nil {
        log.Fatalf("WaitForSelector : %v", err)
    }
    // 获取
    property, _ := elements.GetProperty("src")

    return property.String()
}

// 处理二维码
func showQRCode(qrCode string) {
    if strings.Contains(qrCode, ",") {
        qrCode = strings.Split(qrCode, ",")[1]
    }

    qrCodeBytes, _ := base64.StdEncoding.DecodeString(qrCode)
    qrCodeImage, _ := png.Decode(bytes.NewReader(qrCodeBytes))

    width, height := qrCodeImage.Bounds().Dx(), qrCodeImage.Bounds().Dy()
    newImage := image.NewRGBA(image.Rect(0, 0, width+20, height+20))

    white := color.RGBA{255, 255, 255, 255}
    draw.Draw(newImage, newImage.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)
    draw.Draw(newImage, image.Rect(10, 10, width+10, height+10), qrCodeImage, image.Point{}, draw.Src)

    black := color.RGBA{0, 0, 0, 255}
    rect := image.Rect(0, 0, width+19, height+19)
    for x := rect.Min.X; x < rect.Max.X; x++ {
        newImage.Set(x, rect.Min.Y, black)
        newImage.Set(x, rect.Max.Y, black)
    }
    for y := rect.Min.Y; y < rect.Max.Y; y++ {
        newImage.Set(rect.Min.X, y, black)
        newImage.Set(rect.Max.X, y, black)
    }

    dir, _ := os.Getwd()
    filePath := fmt.Sprintf("%s/output.png", dir)

    outputFile, _ := os.Create(filePath)
    defer outputFile.Close()
    png.Encode(outputFile, newImage)

    // filePath := "/Users/xd_yuan/go_code/output.png" // 替换为你的图像文件路径
    cmd := exec.Command("open", filePath)
    err := cmd.Run()
    if err != nil {
        panic(err)
    }
}

func Pong1(page playwright.Page) (login_flag bool) {
    // log_flag := false
    // 检查登录状态
    check_login_url := "https://api.bilibili.com/x/web-interface/nav"
    pageRespons, _ := page.Request().Get(check_login_url, playwright.APIRequestContextGetOptions{
        Data: nil,
        // 10秒超时
        Timeout: playwright.Float(10000),
    })
    // fmt.Println("pageRespons.URL() = ", pageRespons.URL())
    // fmt.Println("pageRespons.Status() = ", pageRespons.Status())
    // text, _ := pageRespons.Text()
    // fmt.Println("pageRespons.Text() = ", text)
    // fmt.Println("pageRespons = ", pageRespons)
    if pageRespons.Status() != 200 {
        log.Println("网页出错,未返回200")
        return false
    }
    // 查看返回信息
    var response Response

    err := pageRespons.JSON(&response)
    if err != nil {
        log.Fatalf("could not create page: %v", err)
    }

    if response.Code != 0 {
        log.Println("未登录,请登录")
        return false
    }

    // 登录了,获取登录网站的的返回信息
    return response.Data["isLogin"].(bool)

}

func check_login_state(page playwright.Page) error {

    // Pong(page, "https://api.bilibili.com/x/web-interface/nav", false)

    // 如果cookies文件存在，则直接读取文件中的cookie
    // if _, err := os.Stat("/Users/xd_yuan/go_code/cookies.csv"); err == nil {
    //     log.Println("cookies文件存在，直接读取文件中的cookie")
    //     cookies, allCookies := files.ReadCSV()
    //     if len(cookies) > 0 {
    //         log.Println("cookies文件存在，直接读取文件中的cookie")
    //         var playwrightCookies []playwright.OptionalCookie
    //         // playwrightCookies := make([]playwright.OptionalCookie, 1)
    //         domain := ".bilibili.com"
    //         path := "/"
    //         allCookiesStr := strings.Split(allCookies, "; ")
    //         // 将cookie转换为playwright的cookie
    //         for _, cookieStr := range allCookiesStr {
    //             cookie := strings.Split(cookieStr, "=")
    //             playwrightCookies = append(playwrightCookies, playwright.OptionalCookie{
    //                 Name:   cookie[0],
    //                 Value:  cookie[1],
    //                 Domain: &domain,
    //                 Path:   &path,
    //             })
    //         }
    //
    //         // 将cookie设置到浏览器中
    //         err := page.Context().AddCookies(playwrightCookies)
    //         if err != nil {
    //             return err
    //         }
    //         log.Println("启用存储的cookie")
    //         return nil
    //     }
    // }

    // 如果cookies文件不存在，则等待用户扫码登录

    select {
    case <-time.After(10 * time.Second):
        log.Println("超时退出")
        return errors.New("超时退出")
    default:
        for {
            time.Sleep(500 * time.Millisecond)
            cookies, err := page.Context().Cookies()
            if err != nil {
                log.Fatalf("could not get cookies: %v", err)
            }
            dict, cookieStrs := convert_cookies_to_dict(cookies)
            _, ok1 := dict["SESSDATA"]
            _, ok2 := dict["DedeUserID"]
            if ok1 || ok2 {
                log.Println("success")
                // 将cookie保存到csv文件中
                files.WriteCSV(dict, cookieStrs)
                return nil
            }

        }
    }
}

func Pong(page playwright.Page, url string, enableSign bool) (string, string) {

    localStorage, err := page.Evaluate(`() => {
    let result = {};
    for (let i = 0; i < localStorage.length; i++) {
        let key = localStorage.key(i);
        let value = localStorage.getItem(key);
        result[key] = value;
    }
    return result;
}`)
    if err != nil {
        log.Fatalf("could not evaluate script: %v", err)
    }

    // Check if localStorage is empty
    localStorageMap, ok := localStorage.(map[string]interface{})
    if !ok {
        log.Fatalf("could not convert localStorage to map: %v", localStorage)
    }
    if len(localStorageMap) == 0 {
        log.Println("localStorage is empty")
    }

    wbi_img, ok := localStorageMap["wbi_img"].(string)
    if !ok {
        log.Fatalf("could not convert wbi_img to string: %v", localStorage)
    }

    wbi_sub_url, ok := localStorageMap["wbi_sub_url"].(string)
    if !ok {
        log.Fatalf("could not convert wbi_img to string: %v", localStorage)
    }

    wbi_img_urls, ok := localStorageMap["wbi_img_urls"].(string)
    if !ok {
        log.Fatalf("could not convert wbi_img to string: %v", localStorage)
    }

    old_wbi_img_urls := wbi_img + "-" + wbi_sub_url

    if strings.Contains(old_wbi_img_urls, wbi_img_urls) {
        return wbi_img, wbi_sub_url
    } else {
        img_url, sur_url := request("GET", url, map[string]string{
            "Referer":    "https://www.bilibili.com",
            "Origin":     "https://www.bilibili.com",
            "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
        })
        return img_url, sur_url
    }
    return "", ""
}

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

func convert_cookies_to_dict(cookies []playwright.Cookie) (map[string]string, string) {
    var cookiesStr []string
    cookieDict := make(map[string]string)
    // 拿到cookie 键值对的字符串
    // 例如： buvid3=F546933D-E914-C21C-D7FA-8BCE47BE301B23789infoc; b_nut=1697074923; i-wanna-go-back=-1; b_ut=7;
    for _, cookie := range cookies {
        cookieStr := fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
        cookiesStr = append(cookiesStr, cookieStr)
        cookieDict[cookie.Name] = cookie.Value
    }

    return cookieDict, strings.Join(cookiesStr, "; ")
}

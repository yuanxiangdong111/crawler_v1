package bili

import (
    "bytes"
    "encoding/base64"
    "errors"
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "log"
    "os"
    "os/exec"
    "strings"
    "time"

    "github.com/playwright-community/playwright-go"
    "go_code/code_pro/dataconfig"
)

// 定义通用返回体
type Response struct {
    Code    int                    `json:"code"`
    Message string                 `json:"message"`
    Data    map[string]interface{} `json:"data"`
}

func Pong(page playwright.Page) (login_flag bool) {
    // log_flag := false
    log.Println("检查登录状态")
    // 检查登录状态
    check_login_url := "https://api.bilibili.com/x/web-interface/nav"
    pageRespons, _ := page.Request().Get(check_login_url, playwright.APIRequestContextGetOptions{
        Data: nil,
        // 10秒超时
        Timeout: playwright.Float(10000),
    })

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

func PongAndRespKeys(page playwright.Page) (imgKeys, subKeys string, err error) {
    // log_flag := false
    log.Println("检查登录状态")
    // 检查登录状态
    check_login_url := "https://api.bilibili.com/x/web-interface/nav"
    pageRespons, err := page.Request().Get(check_login_url, playwright.APIRequestContextGetOptions{
        // 10秒超时
        Timeout: playwright.Float(10000),
    })

    if pageRespons.Status() != 200 {
        log.Println("网页出错,未返回200")
        return "", "", err
    }
    // 查看返回信息
    var response Response

    err = pageRespons.JSON(&response)
    if err != nil {
        log.Fatalf("could not create page: %v", err)
    }

    // 如论是否成功登录,都去获取 img_key 和 sub_key
    wbiMap := response.Data["wbi_img"].(map[string]interface{})

    imgKeys = wbiMap["img_url"].(string)
    imgKeysList := strings.Split(imgKeys, "/")
    imgKeys = strings.Split(imgKeysList[len(imgKeysList)-1], ".")[0]

    subKeys = wbiMap["sub_url"].(string)
    subKeysList := strings.Split(subKeys, "/")
    subKeys = strings.Split(subKeysList[len(subKeysList)-1], ".")[0]

    return imgKeys, subKeys, nil

    // 登录了,获取登录网站的的返回信息

}

func ReadyLogin(page playwright.Page) error {

    log.Println("准备登录")
    time.Sleep(3 * time.Second)
    // 设置页面视图大小 如：1920*1080 访问移动端页面时需要设置
    if _, err := page.Goto(dataconfig.LOGIN_URL, playwright.PageGotoOptions{
        Timeout:   playwright.Float(90000),
        WaitUntil: playwright.WaitUntilStateLoad,
    }); err != nil {
        log.Println("could not goto: %v", err)
        return fmt.Errorf("could not goto: %v", err)
    }
    // 如果没有登录，则进行登录操作
    // 点击登录按钮
    err := page.Locator(".right-entry__outside.go-login-btn div").Click()
    if err != nil {
        // log.Fatalf("could not click: %v", err)
        log.Println("could not click: %v", err)
        return fmt.Errorf("could not click: %v", err)
    }

    // 点击二维码之后的登录图像获取
    qrcode_img_selector := ".login-scan-box img"

    qrcodeStr := findLoginQrcode(page, qrcode_img_selector)
    showQRCode(qrcodeStr)

    // 等待上一步登录完成,将获取登录后的cookie存在page.Context().Cookies()中

    // 检查登录状态
    err = check_login_state(page)
    if err != nil {
        log.Println("could not check login state: %v", err)
        return fmt.Errorf("could not check login state: %v", err)
    }

    return nil

}

func findLoginQrcode(page playwright.Page, selector string) string {

    // page.Locator(selector).WaitFor(playwright.LocatorWaitForOptions{
    //     State:   nil,
    //     Timeout: nil,
    // })

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

func check_login_state(page playwright.Page) error {
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
            dict, _ := convert_cookies_to_dict(cookies)
            _, ok1 := dict["SESSDATA"]
            _, ok2 := dict["DedeUserID"]
            if ok1 || ok2 {
                log.Println("login success!")
                // 将cookie保存到csv文件中
                // files.WriteCSV(dict, cookieStrs)
                return nil
            }

        }
    }
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

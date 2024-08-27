package Login

import (
    "log"

    "github.com/playwright-community/playwright-go"
)

var BilibiliLogin = bilibiliLogin{}

type bilibiliLogin struct {
    LoginType  string
    LoginPhone string
    CookieStr  string
    Browser    *playwright.Browser
    Page       *playwright.Page
}

func init() {
    pw, err := playwright.Run()
    if err != nil {
        log.Fatalf("could not launch playwright: %v", err)
    }

    browser, err := pw.Chromium.Launch()
    if err != nil {
        log.Fatalf("could not launch browser: %v", err)
    }
    b.Browser = &browser

    page, err := browser.NewPage()
    if err != nil {
        log.Fatalf("could not create page: %v", err)
    }
    b.Page = &page
}

func NewBilibiliLogin(loginType string, loginPhone string, cookies string) *bilibiliLogin {
    return &bilibiliLogin{
        LoginType:  loginType,
        LoginPhone: loginPhone,
        CookieStr:  cookies,
    }
}

func (b *bilibiliLogin) Begin() {

    switch b.LoginType {
    case "phone":
        b.phoneLogin()
    case "cookies":
        b.cookiesLogin()
    case "qrcode":
        b.qrcodeLogin()
    default:
        log.Fatalln("Invalid Login Type. Currently only supported qrcode or phone or cookie ...")
    }
}

func (b *bilibiliLogin) phoneLogin() {
    // TODO
}

func (b *bilibiliLogin) cookiesLogin() {
    // TODO
}

func (b *bilibiliLogin) qrcodeLogin() {
    log.Println("Qrcode Login")
    b.Page
    var a playwright.Page
    a.Locator("xpath=//div[@class='right-entry__outside go-login-btn']//div")

    // if err := b.ContextPage("xpath=//div[@class='right-entry__outside go-login-btn']//div"); err != nil {
    //     log.Fatalf("could not click: %+v", err)
    // }

    // TODO
}

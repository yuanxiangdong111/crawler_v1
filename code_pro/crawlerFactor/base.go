package crawlerFactor

import (
    "context"

    "github.com/playwright-community/playwright-go"
)

type AbstractCrawler interface {
    InitConfig(platform string, loginType string, crawlerType string)
    StartCrawler() error
    Search(ctx context.Context) error
    LaunchBrowser(chromium playwright.BrowserType, playwrightProxy map[string]interface{}, userAgent string, headless bool) (playwright.BrowserContext, error)
}

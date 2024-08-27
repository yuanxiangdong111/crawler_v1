package main

import (
    "fmt"
    "io/ioutil"
    "os"

    jsoniter "github.com/json-iterator/go"
)

type CountryInfo struct {
    NameCN         string `json:"name_cn"`
    NameEN         string `json:"name_en"`
    CountryCode    string `json:"country_code"`
    CountryCode3   string `json:"country_code3"`
    ContinentCode  string `json:"continent_code"`  // 大洲码
    ContinentCode2 string `json:"continent_code2"` // 大洲码，包括 亚太地区
    EuropeanUnion  bool   `json:"european_union"`
    Timezone       string `json:"timezone"`
    UtcOffset      string `json:"utc_offset"`
    IsMutiTimezone string `json:"is_muti_time_zone"`
    CurrencyCode   string `json:"currency_code"`
    CurrencyName   string `json:"currency_name"`
}

func ReadCountryInfos(pth string) (map[string]*CountryInfo, error) {
    f, err := os.OpenFile(pth, os.O_RDONLY, 0600)
    if err != nil {
        return nil, err
    }
    contentByte, err := ioutil.ReadAll(f)
    fmt.Println(string(contentByte))
    var ms map[string]*CountryInfo

    // json.Unmarshal(contentByte, &ms)

    err = jsoniter.Unmarshal(contentByte, &ms)
    return ms, err
}

func main() {
    pathStr := "/Users/xd_yuan/go_code/code_pro/ipdbTest/ch1/china-ipv4-demo.ipdb"
    // city, err := ipdb.NewCity(pathStr)
    // if err != nil {
    //     panic(err)
    // }

    infos, err := ReadCountryInfos(pathStr)
    if err != nil {
        panic(err)
    }
    for k, v := range infos {
        println(k, v)
    }

}

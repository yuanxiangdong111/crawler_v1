package bili

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "net/url"
    "sort"
    "strconv"
    "time"
)

type BilibiliHelp struct {
    imgKey   string
    subKey   string
    mapTable []int
}

func NewBilibiliSign(imgKey, subKey string) *BilibiliHelp {
    return &BilibiliHelp{
        imgKey: imgKey,
        subKey: subKey,
        mapTable: []int{
            46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49,
            33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40,
            61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11,
            36, 20, 34, 44, 52,
        },
    }
}

func (b *BilibiliHelp) getSalt() string {
    mixinKey := b.imgKey + b.subKey
    salt := ""
    for _, mt := range b.mapTable {
        salt += string(mixinKey[mt])
    }
    return salt[:32]
}

func (b *BilibiliHelp) Sign(urlStr string) (string, error) {
    // currentTs := strconv.FormatInt(time.Now().Unix(), 10)
    // reqData["wts"] = "1705742897"
    // /x/v2/reply/wbi/main?mode=0&next=8&oid=239052422&ps=20&type=1&wts=1705916844&w_rid=26e214ff3fb1f086e6c068b34175e94c
    // https://api.bilibili.com/x/v2/reply/wbi/main?mode=0&next=8&oid=239052422&ps=20&type=1&wts=1705916844&w_rid=26e214ff3fb1f086e6c068b34175e94c

    // 解析url 拿到name和value
    urlObj, err := url.Parse(urlStr)
    if err != nil {
        // 返回
        return "", err
    }

    fmt.Println("1 urlObj.String() = ", urlObj.String())

    values := urlObj.Query()
    params := make(map[string]string, len(values))
    for name, value := range values {
        params[name] = value[0]
    }

    // 根据img_url 和 sub_url 生成salt
    // saltKey := b.getSalt()

    currentTime := strconv.FormatInt(time.Now().Unix(), 10)

    // 生成时间戳map
    params["wts"] = currentTime

    keys := make([]string, 0, len(params))
    for k := range params {
        keys = append(keys, k)
    }

    sort.Strings(keys)

    // for k, v := range params {
    //     v = sanitizeString(v)
    //     params[k] = v
    // }

    query := url.Values{}
    for _, k := range keys {
        query.Set(k, params[k])
    }

    queryStr := query.Encode()

    // 根据img_url 和 sub_url 生成salt
    salt := b.getSalt()

    hash := md5.Sum([]byte(queryStr + salt))
    params["w_rid"] = hex.EncodeToString(hash[:])

    baseUrl := "https://api.bilibili.com/x/v2/reply/wbi/main?"
    resUrl := fmt.Sprintf(baseUrl+"mode=%s&next=%s&oid=%s&ps=%s&type=%s&wts=%s&w_rid=%s", params["mode"], params["next"], params["oid"], params["ps"], params["type"], params["wts"], params["w_rid"])
    fmt.Println("resUrl = ", resUrl)
    // https://api.bilibili.com/x/v2/reply/wbi/main?mode=0&next=0&oid=239052422&ps=20&type=1&wts=1705916837&w_rid=4e218d24e7250c9935ac0e99263b127b
    return resUrl, nil
}

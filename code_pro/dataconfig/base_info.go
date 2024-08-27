package dataconfig

const (
    SAVE_LOGIN_STATUS = 1
    LOGIN_URL         = "http://www.bilibili.com"
)

// 视频配置
// https://api.bilibili.com/x/v2/reply/wbi/main?mode=0&next=0&oid=239052422&ps=20&type=1&wts=1705916837&w_rid=4e218d24e7250c9935ac0e99263b127b
var (
    VIDEO_LIST = []string{
        "239052422",
        // "http://vjs.zencdn.net/v/oceans.mp4",
    }
    VIDEO_MODE = "0"
    VIDEO_TYPE = "1"
    VIDEO_PS   = "20"
    // VIDEO_NEXT = "0"
)

/*
   requestParam := map[string]string{
       "oid":  "BV1S5411i7Me",
       "mode": "3",
       "type": "1",
       "ps":   "20",
       "next": "0",
   }
*/

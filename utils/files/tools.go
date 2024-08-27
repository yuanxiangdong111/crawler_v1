package files

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

func WriteCSV(cookieMap map[string]string, allCookies string) {
    // 创建CSV文件
    file, err := os.Create("cookies.csv")
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
    defer file.Close()

    // 创建CSV writer
    writer := csv.NewWriter(file)
    defer writer.Flush()

    // 写入文件头
    headers := []string{"1", allCookies}
    if err := writer.Write(headers); err != nil {
        log.Fatalf("error writing record to csv: %s", err)
    }

    // 写入数据
    for name, value := range cookieMap {
        record := []string{name, value}
        if err := writer.Write(record); err != nil {
            log.Fatalf("error writing record to csv: %s", err)
        }
    }
}

func ReadCSV() (map[string]string, string) {
    // 打开文件
    f, err := os.Open("cookies.csv")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // 读取文件内容
    r := csv.NewReader(f)
    records, err := r.ReadAll()
    if err != nil {
        fmt.Println("111212121")
        panic(err)
    }
    allCookies := records[0][1]
    // 将数据转换为map
    cookieMap := make(map[string]string)
    for k, record := range records {
        if k == 0 {
            continue
        }
        cookieMap[record[0]] = record[1]
    }
    return cookieMap, allCookies
}

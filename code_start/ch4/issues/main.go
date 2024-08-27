package main

import (
    "fmt"
    "log"
    "time"

    "go_code/code_start/ch4/github"
)

type TimeItem struct {
    AfterMonth  []*github.Issue
    BeforeMonth []*github.Issue
    BeforeYear  []*github.Issue
    AfterYear   []*github.Issue
}

func main() {
    strs := []string{"repo:golang/go", "is:open", "json", "decoder"}
    result, err := github.SearchIssues(strs)
    var timeType TimeItem

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%d issues:\n", result.TotalCount)
    nowTime := time.Now()
    for _, item := range result.Items {

        // 一个月之内
        if item.CreatedAt.After(nowTime.AddDate(0, -1, 0)) {
            timeType.BeforeMonth = append(timeType.BeforeMonth, item)
        } else if item.CreatedAt.After(nowTime.AddDate(-1, 0, 0)) { // 一个月之外和一年之内
            timeType.BeforeYear = append(timeType.BeforeYear, item)
        } else { // 一年之外
            timeType.AfterYear = append(timeType.AfterYear, item)
        }

        // fmt.Printf("#%-5d %9.9s %.55s\n",
        //     item.Number, item.User.Login, item.Title)
    }

    for _, v := range timeType.BeforeMonth {
        fmt.Printf("#%-5d %9.9s %.55s\n", v.Number, v.User.Login, v.Title)
    }
    fmt.Println("-----------------------------------------------------------------------")
    for _, v := range timeType.AfterMonth {
        fmt.Printf("#%-5d %9.9s %.55s\n", v.Number, v.User.Login, v.Title)
    }
    fmt.Println("-----------------------------------------------------------------------")

    for _, v := range timeType.BeforeYear {
        fmt.Printf("#%-5d %9.9s %.55s\n", v.Number, v.User.Login, v.Title)
    }
    fmt.Println("-----------------------------------------------------------------------")

    for _, v := range timeType.AfterYear {
        fmt.Printf("#%-5d %9.9s %.55s\n", v.Number, v.User.Login, v.Title)
    }

}

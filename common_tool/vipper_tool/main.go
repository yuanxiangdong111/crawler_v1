package main

import (
    "fmt"
    "log"

    "github.com/spf13/viper"
    "go_code/utils"
)

func main() {

    config, err := utils.LoadConfig("./common_tool")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

    viper.SetDefault("ContentDir", "content")
    viper.SetDefault("LayoutDir", "layouts")
    viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

    fmt.Println(viper.Get("ContentDir"))
    fmt.Println(viper.Get("LayoutDir"))
    fmt.Println(viper.Get("Taxonomies"))

    fmt.Println(config.DBSource)
    fmt.Println(config.DBDriver)
    fmt.Println(config.ServerAddress)

}

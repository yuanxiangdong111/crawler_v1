package main

import (
    "fmt"
    "os"
    "sort"
    "text/tabwriter"
    "time"
)

// 排序输出结构体类型并且格式化

type Track struct {
    Title  string
    Artist string
    Album  string
    Year   int
    Length time.Duration
}

var tracks = []*Track{
    {"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
    {"Go", "Moby", "Moby", 1992, length("3m37s")},
    {"Go Ahead", "Alicia Keys", "As I Am ds wd qq qds", 2007, length("4m36s")},
    {"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
    d, err := time.ParseDuration(s)
    if err != nil {
        panic(s)
    }
    return d
}

// 根据作者姓名排序
type byArtist []*Track

func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// 按照年份排序
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func PrintTracks(tracks []*Track) {
    const format = "%v\t%v\t%v\t%v\t%v\t\n"
    tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
    fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
    fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
    for _, t := range tracks {
        fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
    }
    tw.Flush()
}

// 对结构体排序
// 自定义排序
type customSort struct {
    t    []*Track
    less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
    // 正序
    // sort.Sort(byArtist(tracks))

    // 逆序
    // sort.Sort(sort.Reverse(byArtist(tracks)))

    // 年份正序
    // sort.Sort(byYear(tracks))
    // PrintTracks(tracks)

    // 结构体自定义排序
    // 先 Title
    // 后 Year
    // 然后 Length
    sort.Sort(customSort{tracks, func(x, y *Track) bool {
        if x.Title != y.Title {
            return x.Title < y.Title
        }
        if x.Year != y.Year {
            return x.Year < y.Year
        }
        if x.Length != y.Length {
            return x.Length < y.Length
        }
        return false
    }})
    PrintTracks(tracks)
}

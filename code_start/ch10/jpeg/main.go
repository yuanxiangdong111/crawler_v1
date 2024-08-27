package main

import (
    "fmt"
    "image"
    "image/gif"
    "image/jpeg"
    "image/png"
    _ "image/png"
    "io"
    "log"
    "os"
)

func main() {
    // /Users/xd_yuan/go_code/code_start/ch10/jpeg/
    inFilePath, _ := os.Getwd()
    inFileName := "/test1.png"
    // /Users/xd_yuan/go_code/code_start/ch10/jpeg/output
    infile, err := os.Open(fmt.Sprintf("%s%s", inFilePath, inFileName))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to open input file: %v\n", err)
        os.Exit(1)
    }
    defer infile.Close()

    fmt.Println("inFilePath = ", inFilePath)

    outFileName := "/output.png"
    need2Tpte := fmt.Sprintf("%s%s", inFilePath, outFileName)
    outfile, err := os.Create(need2Tpte)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to create output file: %v\n", err)
        os.Exit(1)
    }
    defer outfile.Close()

    format := "png"

    err = switch2Ph(infile, format, outfile)
    if err != nil {
        log.Fatal(err)
    }
    // 实现快速排序

}

func switch2Ph(in io.Reader, format string, out io.Writer) error {

    img, _, err := image.Decode(in)
    if err != nil {
        log.Fatal(err)
        return err
    }

    switch format {
    case "jpeg":
        // encode
        err := jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
        if err != nil {
            log.Fatal(err)
            return err
        }
    case "png":
        err := png.Encode(out, img)
        if err != nil {
            log.Fatal(err)
            return err
        }
    case "gif":
        err := gif.Encode(out, img, &gif.Options{})
        if err != nil {
            log.Fatal(err)
            return err
        }
    }

    return nil
}

func toJPEG(in io.Reader, out io.Writer) error {
    img, kind, err := image.Decode(in)
    if err != nil {
        return err
    }
    fmt.Fprintln(os.Stderr, "Input format =", kind)
    return jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
}

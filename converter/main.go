package main

import (
    "os"
    "fmt"
    "bytes"
    "strings"
    "io/ioutil"
    "path/filepath"
    "html/template"

    "github.com/russross/blackfriday"
)

type Menu map[string]Field

type Field map[string]PostInfo

type PostInfo struct {
    Field string
    PostName string
    Url string
}

type FileInfo struct {
    Name string
    IsDir bool
    Path string
}

func main() {
    searchDir := "..\\md"

    menu := Menu{}
    fileList := []FileInfo{}
    err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        fileList = append(fileList, FileInfo{
            Name: f.Name(),
            IsDir: f.IsDir(),
            Path: path,
        })
        return nil
    })
    if err != nil {
        panic(err.Error())
    }

    snakeFieldName := ""
    camelFieldName := ""
    for _, info := range fileList {
        if info.IsDir {
            newDir := strings.Replace(info.Path, "..\\md", "..\\post", 1)
            fmt.Println("Directory: ", info.Path, newDir)
            os.MkdirAll(newDir, 0644)

            if newDir == "..\\post" {
                continue
            }

            snakeFieldName = info.Name
            camelFieldName = SnakeToCamel(snakeFieldName)
            menu[camelFieldName] = Field{}
            continue
        }
        if strings.HasSuffix(info.Path, ".md") {
            newPath := strings.Replace(info.Path, "..\\md", "..\\post", 1)
            newPath = strings.Replace(newPath, ".md", ".html", 1)
            convert(info.Path, newPath)
            fmt.Println("File: ", info.Path)

            field := SnakeToCamel(snakeFieldName)
            snakePostName := strings.Replace(info.Name, ".md", "", 1)
            fmt.Println(info.Name, menu)
            menu[camelFieldName][info.Name] = PostInfo{
                Field: field,
                PostName: SnakeToCamel(snakePostName),
                Url: "/post/" + snakeFieldName + "/" + strings.Replace(info.Name, ".md", "", 1),
            }
        }
    }

    fmt.Println(menu)

    makeMenu(menu)
}

func convert(path, newPath string) {
    b, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err.Error())
    }
    output := blackfriday.MarkdownCommon(b)
    //fmt.Println(string(output))

    err = ioutil.WriteFile(newPath, output, 0644)
    if err != nil {
        panic(err.Error())
    }
}

func makeMenu(menu Menu) {
    tmplName := "menu.html"

    t, err := template.ParseFiles(tmplName)
    if err != nil {
        panic(err.Error())
    }

    result := bytes.Buffer{}

    t.ExecuteTemplate(&result, tmplName, menu)
    fmt.Println(result.String())
    err = ioutil.WriteFile("../src/ila-menu.html", result.Bytes(), 0644)
    if err != nil {
        panic(err.Error())
    }
}

func SnakeToCamel(snake string) string {
    b := []byte(snake)
    l := len(b)
    for i, s := range b {
        if i == 0 {
            b[i] = []byte(strings.ToUpper(string(s)))[0]
        }

        if string(s) == "-" {
            b[i] = byte(' ')
            if i + 1 > l - 1 {
                continue
            }

            b[i + 1] = []byte(strings.ToUpper(string(b[i + 1])))[0]
        }
    }
    return string(b)
}
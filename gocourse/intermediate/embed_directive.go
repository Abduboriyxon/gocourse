package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example_.txt
var content string

//go:embed basic
var basicFolder embed.FS

func main() {

	fmt.Println("Embedded Content:", content)
	content, err := basicFolder.ReadFile("basic/hello.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Embedded File Content:", string(content))

	err = fs.WalkDir(basicFolder, "basic", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error walking directory:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

}
package astraljs

import (
	"errors"
	"log"
	"os"
)

type WebApp struct {
	Title  string
	Path   string
	Source string
}

func ResolveWebApp() (webApp WebApp) {
	path, err := getWebAppPath()
	if err != nil {
		log.Fatal(err)
	}

	webApp, err = getWebApp(path)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func TryResolveWebApp() (webApp WebApp) {
	path, err := getWebAppPath()
	if err != nil {
		log.Println(err)
		return
	}

	webApp, err = getWebApp(path)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func getWebAppPath() (path string, err error) {
	args := os.Args[1:]
	if len(args) < 1 {
		err = errors.New("path to js project required")
		return
	}
	path = args[0]
	return
}

func getWebApp(path string) (app WebApp, err error) {
	app = WebApp{
		Title: path,
		Path:  path,
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return
	}
	app.Source = string(bytes)
	return
}

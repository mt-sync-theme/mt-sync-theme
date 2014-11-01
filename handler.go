package main

import (
	"os"
	"os/exec"
	"path"
	"regexp"
)

func previewNameRemapper(name string) string {
	dir, base := path.Split(name)
	templateDir := "templates" + string(os.PathSeparator)
	if dir == templateDir {
		return path.Join("templates_preview", base)
	} else {
		return name
	}
}

func newUrlHandler(opts *cmdOptions) URLHandler {
	handlerRegexp, _ := regexp.Compile("'([^']*)'|\"([^\"]*)\"|(\\S+)")
	envRegexp, _ := regexp.Compile("\\$\\w+")

	return func(url string) error {
		if opts.OptUrlHandler == "" {
			return nil
		}

		matches := handlerRegexp.FindAllStringSubmatch(opts.OptUrlHandler, -1)
		parts := []string{}
		for _, ms := range matches {
			arg := ""
			for _, s := range ms[1:] {
				if s != "" {
					arg = s
				}
			}

			arg = envRegexp.ReplaceAllStringFunc(arg, func(str string) string {
				str = str[1:]
				if str == "URL" {
					return url
				} else {
					return os.Getenv(str)
				}
			})

			parts = append(parts, arg)
		}

		head := parts[0]
		if len(parts) > 1 {
			parts = parts[1:]
		} else {
			parts = []string{}
		}

		_, err := exec.Command(head, parts...).Output()
		return err
	}
}

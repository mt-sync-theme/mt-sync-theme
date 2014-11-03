package main

import (
	"gopkg.in/fsnotify.v1"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getPrefixLength(theme Theme) int {
	themeDir := strings.TrimRight(theme.Directory, "/")
	if themeDir == "." {
		return 0
	} else {
		return len(themeDir) + 1
	}
}

func doWatchCommand(actions []string, theme Theme, client MTSyncThemeClient, opts *cmdOptions, remapper NameRemapper, errorWriter io.Writer, done chan bool) error {
	events := make(chan fsnotify.Event)
	prefixLength := getPrefixLength(theme)
	go func() {
		for {
			event := <-events
			name := event.Name[prefixLength:]

			if event.Op&(fsnotify.Create|fsnotify.Write) != 0 {
				err := client.PutFiles(theme, []string{name}, actions, newUrlHandler(opts), remapper)
				if err != nil {
					log.Println(err)
				}
			} else if event.Op&(fsnotify.Remove|fsnotify.Rename) != 0 {
				err := client.DeleteFiles(theme, []string{name}, actions, remapper)
				if err != nil {
					log.Println(errorWriter, err)
				}
			}
		}
	}()

	Watch(opts.OptThemeDirectory, events, done)

	return nil
}

func doPreview(theme Theme, client MTSyncThemeClient, opts *cmdOptions, errorWriter io.Writer, done chan bool) error {
	return doWatchCommand([]string{"preview"}, theme, client, opts, previewNameRemapper, errorWriter, done)
}

func doOnTheFly(theme Theme, client MTSyncThemeClient, opts *cmdOptions, errorWriter io.Writer, done chan bool) error {
	return doWatchCommand([]string{"on-the-fly"}, theme, client, opts, nil, errorWriter, done)
}

func doSyncDirectory(directory string, theme Theme, client MTSyncThemeClient, opts *cmdOptions, remapper NameRemapper) error {
	paths := []string{}
	prefixLength := getPrefixLength(theme)
	err := filepath.Walk(directory, func(path string, stat os.FileInfo, err error) error {
		if err != nil || stat.IsDir() {
			return nil
		}

		path = path[prefixLength:]

		p := path
		for {
			if filepath.Base(p)[:1] == "." {
				return nil
			}
			np := filepath.Dir(p)
			if p == np || np == "." || (len(np) == 0 && np[0] == filepath.Separator) {
				break
			}
			p = np
		}

		paths = append(paths, path)

		return nil
	})
	if err != nil {
		return err
	}

	return client.PutFiles(theme, paths, []string{}, nil, remapper)
}

func doSync(theme Theme, client MTSyncThemeClient, opts *cmdOptions) error {
	return doSyncDirectory(theme.Directory, theme, client, opts, nil)
}

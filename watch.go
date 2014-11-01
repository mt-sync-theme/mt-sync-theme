package main

import (
	"gopkg.in/fsnotify.v1"
	"log"
	"os"
	"path/filepath"
)

func Watch(themeDir string, events chan fsnotify.Event, done chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if filepath.Base(event.Name)[:1] == "." {
					continue
				}

				stat, err := os.Stat(event.Name)
				if stat != nil && stat.IsDir() {
					if event.Op&fsnotify.Create == fsnotify.Create {
						err = watcher.Add(event.Name)
						if err != nil {
							log.Fatal(err)
						}
					} else if event.Op&(fsnotify.Remove|fsnotify.Rename) != 0 {
						_ = watcher.Remove(event.Name)
					}
				} else {
					events <- event
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = filepath.Walk(themeDir, func(path string, stat os.FileInfo, err error) error {
		if filepath.Base(path)[:1] == "." {
			return nil
		}

		if err == nil && stat.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	})
	<-done
}

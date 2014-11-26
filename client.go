package main

import (
	"github.com/usualoma/mt-data-api-sdk-go"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type MTSyncThemeClient struct {
	dataapi.Client
}

type FileData struct {
	Path    string `json:"path"`
	Content string `json:"content"`
	Action  string `json:"action"`
}

type ActionResult struct {
	Urls []string
}

type FilesResult struct {
	dataapi.Result
	Actions []ActionResult
}

type URLHandler func(string) error
type NameRemapper func(string) string

func NewMTSyncThemeClient(opts dataapi.ClientOptions) MTSyncThemeClient {
	return MTSyncThemeClient{
		dataapi.Client{
			Opts: opts,
		},
	}
}

func (c *MTSyncThemeClient) PutFiles(theme Theme, names []string, actions []string, handler URLHandler, remapper NameRemapper) error {
	var err error

	targetNames := make([]string, len(names))
	for i, name := range names {
		if remapper != nil {
			targetNames[i] = remapper(name)
		} else {
			targetNames[i] = name
		}
	}
	if len(actions) != 0 {
		log.Println("Actions:", actions)
	}
	if len(targetNames) != 0 {
		log.Println("Upload:", targetNames)
	}

	i := 0
	for {
		files := []FileData{}
		contentSize := 0
		for ; i < len(names); i++ {
			if contentSize > 1024*1024 {
				break
			}

			name := names[i]
			file := path.Join(theme.Directory, name)
			content, err := ioutil.ReadFile(file)
			if err != nil {
				if _, e := os.Stat(file); os.IsNotExist(e) {
					continue
				} else {
					return err
				}
			}

			if remapper != nil {
				name = remapper(name)
			}
			// Windows
			name = strings.Replace(name, string(os.PathSeparator), "/", -1)

			files = append(files, FileData{
				Path:    name,
				Content: string(content),
				Action:  "put",
			})

			contentSize += len(content)
		}

		result := FilesResult{}
		err = c.SendRequest(
			"POST",
			"/synced-theme/"+theme.Id+"/files",
			&dataapi.RequestParameters{
				"files":   files,
				"actions": actions,
			},
			&result)
		if err != nil {
			return err
		}
		if result.Result.Error != nil {
			return result.Result.Error
		}

		for _, action := range result.Actions {
			for _, url := range action.Urls {
				log.Println("URL:", url)
				err = handler(url)
				if err != nil {
					return err
				}
			}
		}

		if i >= len(names) {
			break
		}
	}

	return nil
}

func (c *MTSyncThemeClient) DeleteFiles(theme Theme, names []string, actions []string, remapper NameRemapper) error {
	var err error

	targetNames := []string{}
	for _, name := range names {
		targetNames = append(targetNames, name)
		if remapper != nil {
			altName := remapper(name)
			if altName != name {
				targetNames = append(targetNames, altName)
			}
		}
	}
	names = targetNames

	if len(actions) != 0 {
		log.Println("Actions:", actions)
	}
	if len(targetNames) != 0 {
		log.Println("Delete:", targetNames)
	}

	files := []FileData{}
	for _, name := range names {
		// Windows
		name = strings.Replace(name, string(os.PathSeparator), "/", -1)

		files = append(files, FileData{
			Path:   name,
			Action: "delete",
		})
	}

	result := FilesResult{}
	err = c.SendRequest(
		"POST",
		"/synced-theme/"+theme.Id+"/files",
		&dataapi.RequestParameters{
			"files":   files,
			"actions": actions,
		},
		&result)
	if err != nil {
		return err
	}
	if result.Result.Error != nil {
		return result.Result.Error
	}

	return nil
}

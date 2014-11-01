package main

import (
	"errors"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var version = "v0.1.0"

type cmdOptions struct {
	OptHelp           bool   `short:"h" long:"help" description:"Show this help message and exit"`
	OptVersion        bool   `long:"version" description:"Print the version and exit"`
	OptVerbose        bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	OptConfigFile     string `short:"c" long:"config-file" description:"Config file"`
	OptEndpoint       string `long:"endpoint" description:"Endpoint" yaml:"endpoint"`
	OptApiVersion     string `long:"api-version" default:"1" description:"API version" yaml:"api_version"`
	OptClientId       string `long:"client-id" description:"Client ID" default:"mt-sync-theme" yaml:"client_id"`
	OptUsername       string `long:"username" description:"Username" yaml:"username"`
	OptThemeDirectory string `long:"theme-directory" default:"." description:"Theme directory"`
	OptUrlHandler     string `long:"url-handler" description:"URL handler" yaml:"url_handler"`
}

type clientOptions struct {
	*cmdOptions
	PasswordData string
}

func (o clientOptions) Endpoint() string {
	return o.OptEndpoint
}

func (o clientOptions) ApiVersion() string {
	return o.OptApiVersion
}

func (o clientOptions) ClientId() string {
	return o.OptClientId
}

func (o clientOptions) Username() string {
	return o.OptUsername
}

func (o clientOptions) Password() string {
	pass := o.PasswordData
	// If we need to clear password.
	// o.PasswordData = ""
	return pass
}

func getPassword() (string, error) {
	var err error
	var password []byte

	passwordFile := os.Getenv("MT_SYNCED_THEME_PASSWORD_FILE")
	if passwordFile != "" {
		password, err = ioutil.ReadFile(passwordFile)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Can not read password from password file: %s\n", passwordFile))
		}

		if os.Getenv("MT_SYNCED_THEME_PASSWORD_FILE_REMOVE") != "" {
			err = os.Remove(passwordFile)
			if err != nil {
				return "", errors.New(fmt.Sprintf("Can not remove password file: %s\n", passwordFile))
			}
		}
	} else {
		stat, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if stat.Mode()&os.ModeNamedPipe == 0 {
			fmt.Printf("Password: ")
			password = gopass.GetPasswd()
		} else {
			return "", errors.New(fmt.Sprintf("Can not read password from other than a terminal\n"))
		}
	}

	return strings.Trim(string(password), "\n"), nil
}

func Run(cmdArgs []string, errorWriter io.Writer) int {
	var err error

	opts := &cmdOptions{}
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.ParseArgs(cmdArgs)
	if len(args) > 1 || err != nil {
		p.WriteHelp(errorWriter)
		return 1
	}

	if opts.OptHelp {
		p.WriteHelp(errorWriter)
		return 0
	}

	loadConfigFile := func(file string) error {
		yamlData, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal([]byte(yamlData), opts)
		if err != nil {
			return err
		}

		_, err = p.ParseArgs(cmdArgs)
		return err
	}

	var configFile string
	if opts.OptConfigFile != "" {
		configFile = opts.OptConfigFile
	} else {
		for _, f := range []string{path.Join(opts.OptThemeDirectory, "mt-sync-theme.yaml"), "mt-sync-theme.yaml"} {
			if _, err := os.Stat(f); !os.IsNotExist(err) {
				configFile = f
				break
			}
		}
	}

	if configFile != "" {
		err = loadConfigFile(configFile)
		if err != nil {
			fmt.Fprint(errorWriter, err)
			p.WriteHelp(errorWriter)
			return 1
		}
	}

	if opts.OptEndpoint == "" || opts.OptUsername == "" {
		fmt.Fprintln(errorWriter, "Both '--endpoint' and '--username' are required.\n")
		p.WriteHelp(errorWriter)
		return 1
	}

	if opts.OptVersion {
		fmt.Fprintf(errorWriter, "mt-sync-theme: %s\n", version)
		return 0
	}

	command := "preview"
	if len(args) == 1 {
		command = args[0]
	}

	theme, err := NewTheme(opts.OptThemeDirectory)
	if err != nil {
		fmt.Fprintf(errorWriter, "Can not load theme: %s\n", opts.OptThemeDirectory)
		return 0
	}

	password, err := getPassword()
	if err != nil {
		fmt.Fprint(errorWriter, err)
		return 0
	}
	if password == "" {
		fmt.Fprint(errorWriter, "The password is required")
		return 0
	}

	client := NewMTSyncThemeClient(clientOptions{
		cmdOptions:   opts,
		PasswordData: password,
	})

	log.Println("Command: " + command)
	switch command {
	case "preview":
		doSyncDirectory(path.Join(theme.Directory, "templates"), theme, client, opts, previewNameRemapper)
		done := make(chan bool)
		err = doPreview(theme, client, opts, errorWriter, done)
	case "on-the-fly":
		done := make(chan bool)
		err = doOnTheFly(theme, client, opts, errorWriter, done)
	case "sync":
		err = doSync(theme, client, opts)
	case "apply":
		err = client.PutFiles(theme, []string{}, []string{"apply-template-set", "apply-static-files", "apply-custom-fields"}, nil, nil)
	case "rebuild":
		err = client.PutFiles(theme, []string{}, []string{"rebuild"}, nil, nil)
	default:
		fmt.Fprintln(errorWriter, "Available commands: preview, on-the-fly, sync, apply, rebuild")
		p.WriteHelp(errorWriter)
		return 1
	}

	if err != nil {
		log.Println(err)
	}

	return 0
}

func main() {
	os.Exit(Run(os.Args[1:], os.Stderr))
}

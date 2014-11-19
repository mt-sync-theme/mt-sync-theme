# mt-sync-theme

This is a command-line tool for the Movable Type. This program can synchronize files, and generate preview, and apply theme.


## Features

* Can synchronize local theme files to the remote MT.
* Can generate preview without any side-effect to production environment.
* Can apply theme to the blogs, in bulk.
* Can rebuild blogs, in bulk.


## Installation

### Mac OS X / Homebrew

```sh
brew tap mt-sync-theme/mt-sync-theme
brew install mt-sync-theme
```

### Binary (Please choose this option if you use Windows)

Binary packages are available in the [releases page](https://github.com/mt-sync-theme/mt-sync-theme/releases).

### go get

```sh
go get github.com/mt-sync-theme/mt-sync-theme
```

## Setup

1. Install the plug-in [SyncedTheme](https://github.com/mt-sync-theme/mt-plugin-SyncedTheme/releases) to your Movable Type.
1. (Optinal) Link templates to the files of theme directory for a blog via [SyncedTheme](https://github.com/mt-sync-theme/mt-plugin-SyncedTheme/releases).
1. (Optinal) Export the theme from a blog.
1. Download a (exported or existing) theme to your local environment.
1. (Optinal) Create the `mt-sync-theme.yaml` at your local theme directory (a place with theme.yaml). examples: [for Mac](https://github.com/mt-sync-theme/mt-sync-theme/blob/master/example/mt-sync-theme.yaml), [for Windows](https://github.com/mt-sync-theme/mt-sync-theme/blob/master/example/windows/mt-sync-theme.yaml)
    * These configuration variable can be specified via command-line options.
1. Ready to run `mt-sync-theme`.


## Usage

### Overview

The `mt-sync-theme` takes a command, like this.
```
mt-sync-theme preview
```

#### Available commands

* preview
    * Generate a preview page when a file is modified, and open generated preview page via specified handler.
    * In this command, `mt-sync-theme` does not make change to production environment.
* on-the-fly
    * Rebuild a published page when a file is modified, and open updated page via specified handler.
    * In this command, `mt-sync-theme` makes change to production environment. This command should be used in developing stage of the site.
* sync
    * Synchronize local theme files to the remote MT.
* apply
    * Re-apply current theme to the blogs with which this theme is related.
* rebuild
    * Rebuild blogs with which current theme is related.

### preview

* Watch filesystem, and generate a preview page when a file is modified.
    * This command enters to the loop of watching filesystem. You can get quit of this loop by Ctrl-C.
* If modified file is a template (index, or archive), open preview URL via "url_handler".
    * You can preview a module template through a template that is specified by "preview_via".

```
mt-sync-theme preview
```

### on-the-fly

* Watch filesystem, and rebuild a page for preview when a file is modified.
    * This command enters to the loop of watching filesystem. You can get quit of this loop by Ctrl-C.
* You should rebuilt with current templates at least once, before running this command.
* If modified file is a template (index, or archive), open preview URL via "url_handler".
    * You can handle a module template through a template that is specified by "preview_via".

```
mt-sync-theme on-the-fly
```

### sync

* Synchronize local theme files to the remote MT.

```
mt-sync-theme sync
```

### apply

* Re-apply current theme to the blogs with which this theme is related.
* Only these importer will be applied.
    * template_set
    * static_files
    * custom_fields
        * However, application of the custom-field goes wrong in many cases.

```
mt-sync-theme apply
```

### rebuild

* Rebuild blogs with which current theme is related.

```
mt-sync-theme rebuild
```


## Command-Line Options

```
Usage:
  mt-sync-theme

Application Options:
  -h, --help             Show this help message and exit
      --version          Print the version and exit
  -v, --verbose          Show verbose debug information
  -c, --config-file=     Config file
      --endpoint=        Endpoint
      --api-version=     API version (1)
      --client-id=       Client ID (mt-sync-theme)
      --username=        Username
      --theme-directory= Theme directory (.)
      --url-handler=     URL handler
```

## LICENSE

Copyright (c) 2014 Taku AMANO

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

# mt-sync-theme

This is a command-line tool for the Movable Type. This program can synchronize files, and generate preview, and apply theme.


## Features

* Can synchronize local theme files to the remote MT.
* Can generate preview without any side-effect to production environment.
* Can apply theme to the blogs, in bulk.
* Can rebuild blogs, in bulk.


## Installation

### Binary

Binary packages are available in the [releases page](https://github.com/mt-sync-theme/mt-sync-theme/releases).

### go get

```
go get github.com/mt-sync-theme/mt-sync-theme
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

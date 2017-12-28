# gem-search

> Imitate [go-search](https://github.com/tj/go-search/)

 [rubygems.org](rubygems.org) via the command-line.

## Installation

```
$ go get github.com/yunshang/gem-search
```

## Usage

- Help text:
    ```sh
    $ gem-search --help
    Usage:
        gem-search <query>... [--top] [--count n] [--open]
        gem-search -h | --help
        gem-search --version

      Options:
        -n, --count n    number of results [default: 5]
        -o, --open       open rubygems.org search results in default browser
        -h, --help       output help information
        -v, --version    output version
    ```
- Examples:
    ```sh
    $ gme-search rails

    rails
    http://github.com/rails/rails
    Ruby on Rails is a full-stack web framework optimized for
    programmer happiness and sustainable productivity. It
    encourages beautiful code by favoring convention over
    configuration.
    
    jquery-rails
    https://github.com/rails/jquery-rails
    This gem provides jQuery and the jQuery-ujs driver for your
    Rails 4+ application.
    ```
    ```sh
    $ gem-search -o rails
    # opens rubygems.org search results in default browser
    ```
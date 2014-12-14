
# pick

  `pick` is a command line program that allows you to scrape and "pick" stuff from html.

  ![](https://cldup.com/QbqVqqL8l4.gif)

## Installation

  ```bash
  $ go get github.com/yields/pick
  ```

## Examples

  ```bash
  # pick <script>s and output sources.
  curl --silent https://github.com | pick script :src
  
  # pick <title> and output it's text.
  echo "<title>foo</title>" | pick title --text

  # pick all <span>'s in <a>.
  curl --silent https://github.com | pick a | pick span
  curl --silent https://github.com | pick "a span"

  # pick all [src=*].
  curl --silent https://github.com | pick :src
  ```

## Usage

  ```bash
  Usage: 
    pick [options] <selector>...
    pick -v | --version
    pick -h | --help

  Options:
    -t, --text          output inner text for each element
    -v, --version       show version information
    -h, --help          show help information
  ```

## License

  (MIT) 2014 Amir Abu Shareb
# TWIG (top website image gatherer)

TWIG (top website image gatherer) is a cli tool that downloads screenshots of top websites

## Usage

Download the binary that fits your OS and architecture from the [releases](https://github.com/OliverMKing/twig/releases) page. Then either run the binary or add it to your path.

Alternatively you can use `go install`.

```bash
go install github.com/OliverMKing/twig@latest
```

Note: You must have Google Chrome installed (or an alternative driver compatable with [chromedp](https://github.com/chromedp/chromedp)).

### Example usage

Walking through a scenario of how to install and use this project on Ubuntu.

```bash
wget -O twig https://github.com/OliverMKing/twig/releases/download/v0.0.1/twig-linux-amd64
chmod +x ./twig
sudo cp twig /usr/local/bin/twig
```

Installs the binary, makes it executable, and adds it to the path. Now the `twig` is available and can be run.

Running `twig gather -o ./screenshots -n 5` will install the top 5 website's screenshots to the `./screenshots` directory.
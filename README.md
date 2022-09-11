# TWIG (top website image gatherer)

TWIG (top website image gatherer) is a cli tool that downloads screenshots of top websites

## Install

Download the binary that fits your OS and architecture from the [releases](https://github.com/OliverMKing/twig/releases) page. Then either run the binary or add it to your path.

Alternatively you can use `go install`.

```bash
go install github.com/OliverMKing/twig@latest
```

Note: You must have Google Chrome installed (or an alternative driver compatable with [chromedp](https://github.com/chromedp/chromedp)).

## Usage

Walking through a scenario of how to install and use this project on Ubuntu.


```bash
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo apt install ./google-chrome-stable_current_amd64.deb
```

Installs Google Chrome and is only needed if it's not already installed.


```bash
wget -O twig https://github.com/OliverMKing/twig/releases/download/v0.0.1/twig-linux-amd64
chmod +x ./twig
sudo cp twig /usr/local/bin/twig
```

Installs the binary, makes it executable, and adds it to the path. Now the `twig` is available and can be run.

Running `twig gather -o ./screenshots -n 5` will install the top 5 website's screenshots to the `./screenshots` directory.

This tool will commonly be used to download an extremely large number of screenshots. However, due to storage constraints pagination could be needed. You can write a script to assist.

We created a [starting bash script](./example.sh) for you to build off.

```bash
wget https://raw.githubusercontent.com/OliverMKing/twig/main/example.sh
chmod +x ./example.sh
./example.sh
```

Installs the script, makes it executable, then runs it. 

The example script downloads 1000 top website images in batches of 100. There's a comment marking where you could add functionality that does something with the screenshots
# Bip Portfolio
Portfolio watcher for Minter Network


### How to Run

1. Install [Go](https://golang.org) and [dep](https://github.com/golang/dep)

2. Download and compile source code
```bash
mkdir -p $GOPATH/src/github.com/danil-lashin
cd $GOPATH/src/github.com/danil-lashin
git clone https://github.com/danil-lashin/bip-portfolio.git
cd bip-portfolio
cp config.json.example config.json
dep ensure
go build
```

3. Edit `config.json` and add your addresses

4. Run `./bip-portfolio`

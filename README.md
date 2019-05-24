# Bip Portfolio
Portfolio watcher for Minter Network

```
+-----------+-----------+-----------+---------------------------+
|   COIN    |   VALUE   | BIP VALUE | USD VALUE (1 BIP = $0.32) |
+-----------+-----------+-----------+---------------------------+
| BIP       | 30 118.95 | 30 118.95 |                 $9 638.06 |
| BTCSECURE |    100.00 |      2.80 |                     $0.90 |
+-----------+-----------+-----------+---------------------------+
|                 TOTAL | 30 121.75 |                 $9 638.96 |
+-----------+-----------+-----------+---------------------------+
```

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

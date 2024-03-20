# typing-test-server

## Run

```[go]
go run main.go
```

↓ でも可能．一度ビルドすれば 2 回目以降は`./main`で OK．

```[go]
go build main.go
./main
```

## Require

Go のインストールが必要．

### インストールされているか確認

コマンドの先頭の「$」や「>」は入力しない．

-   Linux

```
$ go version （←これを実行）
go version go1.18.1 linux/amd64 （←こんな感じでバージョン情報が結果として表示されればOK）
```

-   Windows

```
> go version（←これを実行）
go version go1.21.3 windows/amd64 （←こんな感じでバージョン情報が結果として表示されればOK）
```

Go のインストールは[こちら](https://go.dev/doc/install)を参照．

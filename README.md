# Go Training

某社でのGoの研修の資料

1日でGoの概要を理解する。。。

## 参考

* [https://golang.org/](https://golang.org/) 公式ページ
* [A Tour of Go](https://tour.golang.org/welcome/1) とりあえずやってみるが途中で飽きる
* [Effective Go](https://golang.org/doc/effective_go.html) どう書くのが良いんだっけという時に見る
* [The Go Blog](http://blog.golang.org/) 最新情報とTips
* [The Go Programming Language](http://www.gopl.io/) 現状一番良い本で6月に日本語版出るらしい

## 準備

### インストール

* [公式ドキュメント](https://golang.org/doc/install)に従って入れれば良い
* 実際の研修ではいろいろインストール済のVMが配られるはず

### エディタ

* Vimだったら[vim-go](https://github.com/fatih/vim-go)入れておくと良いです
* 他のエディタについては知りませんがググればいろいろ出てくると思います
* とにかく保存時に[goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)とか[golint](https://github.com/golang/lint)とかでチェックするようにしておくと捗ります

### このリポジトリを`go get`する

```sh
go get github.com/higebu/gotraining
```

エラーになったらエラーメッセージをよく読んで対処してください

### みんな大好き Hello, World!

* [ソース](./helloworld/helloworld.go)
* [Playground](http://play.golang.org/p/992fMmkkxr)
    * [The Go Playground](https://play.golang.org)、ブラウザ上でGolangのコードを実行できます

`go run` で実行できる

```sh
cd $GOPATH/src/github.com/higebu/gotraining/helloworld/
go run helloworld.go
```

`Hello, World!` と表示されるはず

`go build` でコンパイルできる

```sh
go build helloworld.go
./helloworld
```

`Hello, World!` と表示されるはず

# Anoside

実験的にGo / ClojureScriptで開発中。
http://anoside.com で実際に稼働しているシステムは [anoside/anoside](https://github.com/anoside/anoside) になります。


## 動かし方

### 必要なソフトウェアのインストールする

#### Google App Engine SDK

https://developers.google.com/appengine/downloads?hl=ja からSDKをダウンロードしてください。
そして解凍して生まれたディレクトリ `go_appengine` を任意の場所に配置すれば完了です。


#### Go

```
$ brew install go
$ export GOPATH=$HOME/golang
$ export PATH=$GOPATH/bin:$PATH
```


#### ClojureScript

[Leiningen](https://github.com/technomancy/leiningen) をインストールします。

```
$ brew install leiningen
```


### 各種コマンド

#### 開発用サーバを起動する

```
$ cd path/to/anoside/web
$ goapp serve .
```


#### ClojureScriptのソースコードをコンパイルする

```
$ lein cljsbuild once
```


#### ClojureScriptのREPLを起動する

```
$ lein trampoline cljsbuild repl-rhino
```

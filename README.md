# Monkey Language
[Go言語でつくるインタプリタ](https://www.oreilly.co.jp/books/9784873118222/)で作成したプログラミング言語 ”Monkey” です。

ほとんどは書籍の写経ですが、一部機能の追加や改善をしています。

# 使用方法

以下のコマンドでインタプリタが起動します。

```
go run main.go
```

# 環境
Go:1.21
Docker:24.0.6

# 追加機能

書籍のままでは最低限の機能しか提供されていないため、機能の追加を行いました。
## ファイル操作
テキストファイルの読み込み、書き込みを行う機能を作成しました。
### 使用方法
```
# ファイルの読み込み
>> read('text.txt')

# ファイルの書き込み
>> write('string','path')
```
## monkeyコマンド
rubyやpythonといったインタプリタのようにコマンドラインからインタプリタを起動するコマンドを作成しました。
コマンドラインで`monkey`を実行するとインタプリタが起動します。

## 使用方法
```
# インタプリタの起動
$ monkey 

# monkeyスクリプトが記述されたテキストファイルを読み込み、実行する。
# monkey script.txt
```

## オプション
コマンドのオプションです。
- v:バージョン情報
- h:ヘルプ



その他の使い方や基本構文、理論的なことのまとめは[こちら]()のページにまとめています

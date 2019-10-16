# 実装中のメモ
gorillaライブラリを使用してwebsocketを実装するとUnity経由でも動作
golang は先頭が大文字か小文字かでprivateとpublicを使い分けている。
パッケージ化するときは注意
dbrからパラメータを受け取る際にも変数が大文字で始まっていないと値が入らない
下記は例
```
type User struct {
    Id int
}
```

# echoフレームワーク
公式サイト：https://echo.labstack.com/  
FWのGinより早いらしい。日本語ドキュメントは皆無なので、頑張るしか無さそう

# 起動手順
下記コマンドを実行
```
$ docker-compose build
$ docker-compose up
```
バックグラウンドで実行してしまうとログが見れなくなるので、フォアグラウンドで実行すること  
その後、ブラウザを起動し、下記サイトへアクセス  
http://localhost:8000  
勝手にwebsocket通信を行う。  
なお、直接websocketに繋げる場合は、 http://localhost:8000/ws へアクセスすれば良い。

# テストコマンド
追々は単体テストできるようにする
```
$ curl -X POST -d '{"id":"1", "name":"hoge", "room_id":"1"}' localhost:8080/user

$ curl localhost:8080/user/1

$ curl -X POST -d '{"id":"1", "name":"hoge", "room_id":"2"}' localhost:8080/user/1
```
mysql関連
user: root
pass: root
```
create database mydb character set utf8mb4;
create table user (id int, name varchar(20), room_id int);
alter table user default character set utf8mb4;
create table chat (id int, room_id int, user_id int, message text);
alter table chat default character set utf8mb4;
```

# ビルド手順
```
$ cd models
$ go build -o models user.go
```

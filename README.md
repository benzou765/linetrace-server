# ライントレース、サーバ側コード
Arduino、ESP8266、groveカメラモジュール、モータを使用して、ライントレースカーを作成。
TODO: どこかのタイミングでコードを整理

## 実装中のメモ
カメラで取得したデータは、フォームデータに変換できる余裕があるかわからないので、
バイナリで送信。使うヘッダーは下記
```
application/octet-stream
```
リクエストヘッダーの取得
```
c.Request().Header
```
型確認
reflectのインポートが必要
reflect.TypeOf(hoge)

画像は下記のように送信してあげれば良さそう
```
Content-Type: image/jpeg
（データ）
```

## 機械学習のライブラリ
tensorflowを利用しようとしたが、
* 多くのミドルウェアが必要なこと
* ハードウェア構成を意識した設計が必要なこと
* 環境構築するのに手間がかかること
から断念。今回は機械学習のテストをするだけなので、gobrainを使用
うまく行かない場合は、python＋scikit-learnで実装。

## docker便利コマンド
```
docker run -it --rm (image) bash
```
上記コマンドで終了時にdockerイメージも一緒に削除してくれる。
イメージ作成中は便利

## テストデータ
```
curl -v http://localhost:8080/id1aefup8oozahlo6etai4gei2aew5ee -H "Content-Type: image/jpeg" --data-binary @test_image.jpg
```

## バイトデータの取扱
参考：https://qiita.com/mztnnrt/items/ddf6920a484e74f0ee1a
```
body := c.Request().Body
octet, err := ioutil.ReadAll(body)
if err != nil {
    return err
}
c.Echo().Logger.Debug(octet)
```

```
body := c.Request().Body
buffer := new(bytes.Buffer)
io.Copy(buffer, body)
octet := buffer.Bytes()
c.Echo().Logger.Debug(octet)
```
今回は、ioutil.ReadAllを使用

## MNISTの中身
TRAINING SET LABEL FILE (train-labels-idx1-ubyte):
[offset] [type]          [value]          [description]
0000     32 bit integer  0x00000801(2049) magic number (MSB first)
0004     32 bit integer  60000            number of items
0008     unsigned byte   ??               label
0009     unsigned byte   ??               label
........
xxxx     unsigned byte   ??               label
The labels values are 0 to 9.

TRAINING SET IMAGE FILE (train-images-idx3-ubyte):
[offset] [type]          [value]          [description]
0000     32 bit integer  0x00000803(2051) magic number
0004     32 bit integer  60000            number of images
0008     32 bit integer  28               number of rows
0012     32 bit integer  28               number of columns
0016     unsigned byte   ??               pixel
0017     unsigned byte   ??               pixel
........
xxxx     unsigned byte   ??               pixel
Pixels are organized row-wise. Pixel values are 0 to 255. 0 means background (white), 255 means foreground (black).

TEST SET LABEL FILE (t10k-labels-idx1-ubyte):
[offset] [type]          [value]          [description]
0000     32 bit integer  0x00000801(2049) magic number (MSB first)
0004     32 bit integer  10000            number of items
0008     unsigned byte   ??               label
0009     unsigned byte   ??               label
........
xxxx     unsigned byte   ??               label
The labels values are 0 to 9.

TEST SET IMAGE FILE (t10k-images-idx3-ubyte):
[offset] [type]          [value]          [description]
0000     32 bit integer  0x00000803(2051) magic number
0004     32 bit integer  10000            number of images
0008     32 bit integer  28               number of rows
0012     32 bit integer  28               number of columns
0016     unsigned byte   ??               pixel
0017     unsigned byte   ??               pixel
........
xxxx     unsigned byte   ??               pixel
Pixels are organized row-wise. Pixel values are 0 to 255. 0 means background (white), 255 means foreground (black).

## MNISTを使用して実験する場合
コピーをこのgitに保存するとファイルサイズが大きいので、gitignoreで除去
実験する場合は、ml_mnist配下に下記ファイルを保存すること。
公式サイトと圧縮形式が異なる場合は、再圧縮が必要
THE MNIST DATABASE：http://yann.lecun.com/exdb/mnist/

* t10k-images-idx3-ubyte.gz
* t10k-labels-idx1-ubyte.gz
* train-images-idx3-ubyte.gz
* train-labels-idx1-ubyte.gz

## 起動手順
下記コマンドを実行
```
$ docker-compose build
$ docker-compose up
```

# 参考資料
echo公式サイト：https://echo.labstack.com/  
httpプロトコルのデータ送信：https://qiita.com/ts-3156/items/93af082d0479c0eb9646
低レベルアクセスの方法：https://ascii.jp/elem/000/001/252/1252961/
io.Readerから[]byteへの変換のベンチマーク：https://qiita.com/imishinist/items/be9073a03ae2e903d913
Go でバイナリ処理：https://qiita.com/Jxck_/items/c64d9ae0e910762eab37
image/jpegの扱い方：https://developer.mozilla.org/ja/docs/Web/HTTP/Basics_of_HTTP/MIME_types#JPEG
Awesome Go : 素晴らしい Go のフレームワーク・ライブラリ・ソフトウェアの数々：https://qiita.com/hatai/items/f31914f37dc6c53b2bce
Golangだけでやる機械学習と画像分析：https://mattn.kaoriya.net/software/lang/go/20181108123756.htm
THE MNIST DATABASE：http://yann.lecun.com/exdb/mnist/
Alpine Linuxにnumpy, scipy, scikit-learn, pandasを入れた：https://qiita.com/ricesho/items/e56bf08f51ea406674eb

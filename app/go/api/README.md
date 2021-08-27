# コードについての説明
## ./main.go
---
```go
e.a(b, c)
```
aにはHTTPのメソッド名(GET, POST, PUT, DELETEなど)を指定し、 bにはパスを指定し、cには受け取ったデータを処理してどのようにして返すかを記述した関数を指定します。

例
```go
e.POST("/coin", apifunc.CoinPost)
```
localhost:8080/**coin**
にきた**POST**リクエストを処理する関数(apifunc.CoinPost)を指定しています。

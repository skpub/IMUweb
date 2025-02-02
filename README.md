# インモラル大学Webサイト

インモラル大学Webサイトのフロントエンドとバックエンド。

## 基本的な構成

インモラル大学のサイトは学部情報や学長挨拶のような固定のコンテンツと、
それから学長が適宜更新する記事とで出来ている。
記事はMarkdownで書かれており、これはサーバが握っている。

構成としては、フロントはSSGで成り立っていて、
記事だけサーバから引っ張ってきて適当にDOMに変換して表示するものである。

- バックエンド
    - HTTPサーバ(Goa) (/IMUbackend)
    - オブジェクトストレージ(MinIO)
- フロントエンド(Svelte SSG) (/)

コイツらの挙動を .env で管理する。環境変数は以下の通り。

```
MINIO_ROOT_USER=
MINIO_ROOT_PASSWORD=

MINIO_PORT= (デフォ9000)
MINIO_CONSOLE_PORT= (デフォ 9001)
MINIO_SERVER_URL=localhost:${MINIO_PORT}

POOL_PATH= (/var/mdstore などで良いと思う。ホストのどこにMarkdown記事を永続化するか)

MDFILESYSTEM= (Markdown記事を突っ込むファイルシステムの名前(MinIOのやつ))
MDBUCKET= (Markdown記事を突っ込むバケットの名前(MinIOのやつ))

# ↑ここまではオブジェクトストレージMinIOの設定
# ↓ここからはバックエンド(Goa)の設定

BACKEND_ADDR= (http://localhost など)
BACKEND_PORT= (8080でいいと思う)
```

## 動かし方

基本的には一つのサーバで動かすことを考えているが、
やりたいのであれば、まぁ.env の構成を頑張ればオブジェクトストレージを他で動かすことも可能そうである。
更に頑張れば並列で動かしてLBを設置する構成にもまぁできそうである。

とりあえず、以下では一つのOriginですべてを動かす小規模構成で説明する。

### 手順

1. 環境変数の設定
1. dockerでMinIOを動かす
1. goa gen
1. バックエンドを動かす
1. フロントをビルドして/var/wwwなど(適宜)配置
1. nginx などでリバースプロキシの設定をしたりsystemdでサービスにしたり(適宜)

### 1. 環境変数の設定

上に書いてあるものを適宜設定

### 2. dockerでMinIOを動かす

カレントディレクトリはプロジェクトルートで

```bash
~$ docker compose up -d
```

→ MinIOが立ち上がる

### 3. goa gen

Goa のインストールは適宜頑張ってください。
カレントディレクトリは/IMUbackendで

```bash
~$ goa gen IMUbackend/design
```

→ API が生成される

### 4. バックエンドを動かす

カレントディレクトリは/IMUbackendで
```bash
~$ go run ./main
```

 当然一つセッションを奪われるのでtmuxかnohupかscreenか何かでやると良いと思います。
 もちろんsystemdでサービスにしても良いし。

### 5. フロントをビルドして/var/wwwなど(適宜)配置
カレントディレクトリはプロジェクトルートで

```bash
~$ npm install
~$ npm run build
```

これでSSGされるので生成物を適宜移動してください。

### 6.
お好きにどうぞ。
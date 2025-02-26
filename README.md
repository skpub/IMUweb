# インモラル大学Webサイト

インモラル大学Webサイトのフロントエンドとバックエンド。

## 基本的な構成

インモラル大学のサイトは学部情報や学長挨拶のような固定のコンテンツと、
それから学生・教員が適宜更新する記事とで出来ている。
記事はMarkdownで書かれており、これはサーバが握っている。

構成としては、フロントはSSRで成り立っていて、
記事だけサーバから引っ張ってきて適当にDOMに変換して表示するものである。
フロント側のサーバは bun で動かしている。

- バックエンド
    - HTTPサーバ(Goa) (/IMUbackend)
    - オブジェクトストレージ(MinIO: docker)
    - DB(postgres: docker)
- フロントエンド(Svelte SSR(bun)) (/)

コイツらの挙動を .env で管理する。環境変数は以下の通り。

```
# オブジェクトストレージMinIOの設定
MINIO_ROOT_USER=
MINIO_ROOT_PASSWORD=

MINIO_PORT= (デフォ9000)
MINIO_CONSOLE_PORT= (デフォ 9001)
MINIO_SERVER_URL=localhost:${MINIO_PORT}

POOL_PATH= (/var/mdstore などで良いと思う。ホストのどこにMarkdown記事を永続化するか)

MDFILESYSTEM= (Markdown記事を突っ込むファイルシステムの名前(MinIOのやつ))

# DB の設定
PG_HOST=
PG_USER=
PG_PASSWORD=
PG_DBNAME

# バックエンド(Goa)の設定
PRIVATE_BACKEND_ADDR= (SSRのサーバ側から見えるアドレス)
PRIVATE_BACKEND_PORT= 

PUBLIC_BACKEND_PORT= (クライアントから見えるアドレス)
PUBLIC_BACKEND_PORT=

JWT_SECRET= (JWT署名の共通鍵)
SALT=       (パスワードをハッシュ化するときにお塩も掛ける)

# フロントエンド(SvelteKit SSR)の設定
PUBLIC_FRONTEND_ORIGIN=

PRIVATE_MC_JE_ADDR= (コイツら3つはマイクラサーバの情報。
PRIVATE_MC_BE_ADDR= ゆくゆくはバックエンドからマイクラサーバと連携して
PRIVATE_MC_VERSION= これらの情報を持ってくるやつを作る)
```

## 動かし方

基本的にはフロントエンドもバックエンドも一つのサーバで動かすことを考えているが、
やりたいのであれば、まぁ.env の構成を頑張ればオブジェクトストレージやDBやその他を他で動かすことも可能そうである。
更に頑張れば並列で動かしてLBを設置する構成にもまぁできそうである。

とりあえず、以下では一つのOriginですべてを動かす小規模構成で説明する。

### 手順

1. 環境変数の設定
1. dockerでMinIOとpostgresを動かす
1. DBのマイグレーション
1. goa gen
1. バックエンドを動かす
1. フロントをビルドして/var/wwwなど(適宜)配置
1. nginx などでリバースプロキシの設定をしたりsystemdでサービスにしたり(適宜)

### 1. 環境変数の設定

上に書いてあるものを適宜設定

### 2. dockerでMinIOとpostgresを動かす

カレントディレクトリはプロジェクトルートで

```bash
~$ docker compose up -d
```

→ MinIOとpostgresが立ち上がる

### 3. マイグレーション

マイグレーション用のスクリプトを用意しています。

```bash
~$ ./migrate.sh
```


### 4. goa gen

Goa のインストールは適宜頑張ってください。
カレントディレクトリは/IMUbackendで

```bash
~$ goa gen IMUbackend/design
```

→ API が生成される

### 5. バックエンドを動かす

カレントディレクトリは/IMUbackendで
```bash
~$ go run ./main
```

 当然一つセッションを奪われるのでtmuxかnohupかscreenか何かでやると良いと思います。
 もちろんsystemdでサービスにしても良いし。

### 6. フロントをビルドして/var/wwwなど(適宜)配置
カレントディレクトリはプロジェクトルートで

```bash
~$ npm install
~$ bun --bun run build
```

bun の指示通りに動かしてください。
あと、そもそもbunをインストールしておいてください。

### 7.
お好きにどうぞ。

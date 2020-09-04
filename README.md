## astquery のフロントエンド(WASM)

#### 概要
- goファイルに対して、astqueryをかける事ができる(パッケージにたいしてもできるようにしたかったが、packages.LoadがWASMではできなかった。)

#### public/
- wasmの実行用ファイル

#### 立ち上げ
```
cd public && \
./run.sh


open http://localhost:8080

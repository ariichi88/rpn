# rpn

逆ポーランド記法で四則演算する、簡単なコマンドライン計算機  

## インストール
GitHubからクローン 
```
git clone git@github.com:ariichi88/rpn.git
```
ビルド 
'''
go build 
```
できた実行ファイルをパスの通ったディレクトリにコピーしてください  

## 使い方  
コマンドライン引数に計算式を指定してください。  
```
$ rpn 200 300 + 100 * 10 /
5000
```
or
引数なしで起動すると対話的に計算をすることができます。 
```
$ rpn
exit with q
200
300
+
500
q -> ｑでコマンドを終了します
```
> コマンドラインオプションとヘルプオプションはありません  


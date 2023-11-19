# usi-engine-delayer

usi-engine-delayer は USI (Universal Shogi Interface) プロトコルのコマンド送信を意図的に遅延させるツールです。

将棋の GUI アプリで USI プロトコル関連の動作検証を行う際、必ずしも CPU に負荷をかけてより優れた指し手を追及する必要はありません。
かといって瞬時に応答するプレイヤーでは、思考時間が長いケースの動作検証として十分ではありません。
そこでこのツールではコマンドの送信を意図的に遅延させることで、マシンに負荷をかけずに思考時間を延長します。

## 使用方法

### 必要なもの

Go 言語実行環境

### インストール

```
go install github.com/sunfish-shogi/usi-engine-delayer@main
```

### 設定ファイルを使用する場合

`-config` で設定ファイルのパスを指定します。省略するとカレントディレクトリの `config.json` を参照します。
設定ファイルは以下のように記述します。

```json
{
  "exePath": "C:\\Users\\GuestUser\\MyEngines\\Lesserkai.exe",
  "delaySeconds": 5
}
```

なお、 `-exe` オプションが指定された場合は設定ファイルを参照しません。

### 引数で設定値を渡す場合

- `-exe` - エンジンのパスを指定します。
- `-delay` - bestmove コマンドの遅延秒数を指定します。

エンジンと同じディレクトリに bat ファイルを設置する場合は、例えば次のように記述します。

```bat
usi-engine-delayer -exe %~dp0Lesserkai.exe -delay 5
```

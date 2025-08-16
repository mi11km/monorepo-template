# Logging package

下記ログの出力をすることができるパッケージです

- アプリケーションログ（JSON、コンソール）

## 使い方

### アプリケーションログ

基本的に、公開関数 (Debug, Info, Warn, Error, Fatal) を使ってログを出力してください。
デフォルトでは、標準エラー出力に一般的なJSON形式で出力されます。

```go
package main

import (
	"context"

	"github.com/mi11km/monorepo-template/go/pkg/logger"
)

func main() {
	ctx := context.Background()
	logger.Info(ctx, "Hello, World!", logger.FieldString("key", "value"))
}

// Output: {"level":"info","time":"2024-07-11T17:58:09.873934+09:00","message":"Hello, World!","key":"value"}
```

ログレベルとログフォーマットに関しては、環境変数 `LOG_LEVEL` と `LOG_FORMAT` で設定できます。

- `LOG_LEVEL` は `debug`, `info`, `warn`, `error`, `fatal` のいずれかを指定できます(デフォルトは `info`)
- `LOG_FORMAT` は `json`, `console` を指定できます(デフォルトは`json`)

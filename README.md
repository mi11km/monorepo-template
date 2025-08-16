# Monorepo Template

一般的な Web サービス、アプリ開発に使えるモノレポのテンプレートを提供

### はじめに

開発者間で複数のランタイムバージョンを共通化するために [mise](https://mise.jdx.dev/) の利用を推奨します。

### コミットメッセージのガイドライン

[commitlint](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional) でコミットメッセージの形式を強制します。
以下に頻繁に使うであろうタイプの一覧を転載しておきます。その他スコープなどの詳細についてはリンク先を参照してください。

| Type     | Description                                                          |
| -------- | -------------------------------------------------------------------- |
| build    | ビルドシステムまたは外部依存関係に影響する変更（例: npm, webpack）   |
| ci       | CI構成ファイルとスクリプトの変更（例: GitHub Actions）               |
| docs     | ドキュメントのみの変更                                               |
| feat     | 新しい機能                                                           |
| fix      | バグ修正                                                             |
| perf     | パフォーマンスを向上させるコード変更                                 |
| refactor | バグを修正したり機能を追加したりしないコード変更                     |
| style    | コードの意味に影響を与えない変更（空白、書式、セミコロンの欠落など） |
| test     | 不足しているテストの追加または既存のテストの修正                     |

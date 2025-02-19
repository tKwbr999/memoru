# AI駆動開発を自分なりにやってみる
## 前提やきっかけ
まず、僕自身はgoでのbackend開発は得意ですが、frontend・UIデザイン・supabaseのようなbaasの利用はやってできなくはないといった感じで専門性はほとんどない状態です。

AI agentによるコーディングの波に乗れ！ということで以前から欲しかったけどあまり聞いたことのないサービスを自分なりに作ってみようという企画です。

## 計画を作成
まず自分が言語化できる計画内容をマークダウンにまとめる。

まとめたものがこちら。

やりたいことは自分の独り言をまとめ、ある程度溜まったらAIによる解析で何をしたいのかを提案するサービスの作成。
```
# 企画内容
## サービス
### 名称
memoru(メモる)
### 概要
誰にも伝えない自分だけのメモ
メモとはメモリーの略
### 価値
メモをとる→メモが溜まる→やりたいことが生まれる→実現の道筋を提案する
やりたいことの解析をAIで行う。
ある程度解析がまとまってくると実現したい本音を見つけだす
本音の実現に向けてアドバイスをする。
アドバイスを受けたユーザーの評価を解析し提案の質を高める
## アーキテクチャ
### 前提
supabaseにdeployできる形式でアーキテクチャをプランニングする
### backend開発言語
supabaseへのデプロイ可能な言語を優先度の高い順から選択する
第一候補: golang
第二候補: python
### backendフレームワーク
言語別に指定のフレームワークを利用する
- goの場合
第一候補: atreugo
第二候補: fiber
- pythonの場合
市場に知識が出回っているおすすめのフレームワーク
### データベース
supabaseへのデプロイ可能なDBを選択する
実装に手間のかからず安定する方法を優先する
### 認証
supabaseへのデプロイ可能な認証方式を選択する
実装に手間のかからず安定する方法を優先する
### frontend開発言語
supabaseへのデプロイ可能なfrontend実装を選択する
優先度の高い順から選択する
第一候補: typescript
第二候補: javascript
## UIデザイン
### テーマカラー
白やグレー
### デザインイメージ
任天堂のゲーム「motherシリーズ」のような見た目
ドットを基調として柔らかな雰囲気が選定理由
```

そしてここまでまとまったらAIが読みやすい内容に改善するためclaudeによるチェックを依頼する。

先ほど自分がまとめたマークダウンをそのまま提出する。

チェックの依頼はこのようなプロンプトで行う。
```
プロンプトとしてわかりやすい内容に変更。またこの計画において過不足や新たな提案がある場合はその内容を盛り込む。
```

そしてチェック後のマークダウンがこちら。
```
# プロジェクト概要: memoru（メモる）

## 1. サービスコンセプト
### 製品ビジョン
「あなたの想いを育てる、AIパーソナルメモ」

### 主要機能
1. パーソナルメモ機能
   - プライバシー重視の完全非公開メモ
   - 直感的な入力インターフェース
   
2. AI分析・提案機能
   - メモの集積からユーザーの本質的な願望を分析
   - 実現可能な行動計画の自動生成
   - ユーザーフィードバックによる継続的な精度向上

## 2. 技術スタック
### バックエンド
- 言語: Go（推奨）
  - フレームワーク: Fiber
  - 理由: Supabaseとの親和性、パフォーマンス、開発効率

### データベース
- PostgreSQL（Supabase提供）
  - リアルタイム機能活用
  - 組み込みの行レベルセキュリティ

### 認証
- Supabase Auth
  - メール/パスワード認証
  - ソーシャルログイン（Google, GitHub等）

### フロントエンド
- 言語: TypeScript
- フレームワーク: Next.js
  - SSR対応
  - 優れたDX（Developer Experience）

## 3. UIデザイン指針
### ビジュアルテーマ
- ベースカラー
  - プライマリ: #FFFFFF（白）
  - セカンダリ: #F5F5F5（ライトグレー）
  - アクセント: #E0E0E0（ミディアムグレー）

### デザインコンセプト
- MOTHERシリーズインスピレーション
  - ドットグラフィックス要素
  - 優しく親しみやすいUI
  - ミニマルでクリーンなレイアウト

## 4. 追加提案
1. プログレッシブWeb App（PWA）対応
   - オフライン機能
   - プッシュ通知サポート

2. AIフィードバックループ
   - ユーザー行動分析
   - パーソナライズされた提案最適化

3. プライバシー強化機能
   - エンドツーエンド暗号化
   - データエクスポート機能
```
内容が味気なくなってしまったが、AIが理解できれば問題なし。

ただ、対人間の資料としては自分で考えた文章がいいこともあると思うので、自分の考えたマークダウンは `seed-plan.md`, claudeによるチェックを受けたマークダウンは `optimized-plan.md` として保存しておく

## roo-clineによるアーキテクトモードによるプランニング
`optimized-plan.md` をさらに技術的に掘り下げるためvscodeのプラグイン `roo code(roo cline)` でアーキテクチャを考えてもらう。

roo codeで利用するAI modelは `gemini-2.0-pro-exp-02-05` 。

この構成で利用する理由についてはこちらのポストをみてから興味があったのでやってみたから。

https://x.com/exceedsystem/status/1889036769339085015

geminiのAPIキーの発行方法とかは割愛します。

Architectモードにしてclaudeにチェックしてもらったマークダウンを確認してもらいます。

その確認するためのプロンプトはこんな感じにしました。
```
マークダウン記載内容を実現するための道筋をステップバイステップで検討・レポートして。
一個一個のステップがデプロイ可能で利用可能となるようにする。
一番重要なのはメモを投稿する機能。
```

で、これもclaundeによってチェックしてもらったのがこちら。
```
memoruアプリケーションの開発ロードマップを作成してください。以下の要件に従って、段階的な実装計画を詳細に説明してください：

1. 実装の優先順位
- コアフィーチャーであるメモ投稿機能を最優先
- 各開発フェーズは独立してデプロイ可能であること
- 各フェーズで動作確認可能な状態まで実装すること

2. 技術要件
- Supabaseをインフラとして使用
- バックエンド：Go/Fiber
- フロントエンド：TypeScript/Next.js
- データベース：PostgreSQL（Supabase提供）

3. 開発フェーズごとに必要な情報
- 実装する機能の詳細仕様
- 必要な技術コンポーネント
- 想定される技術的課題
- テスト項目
- デプロイ手順

4. 成功基準
- 各フェーズでの具体的な達成目標
- 品質基準
- パフォーマンス要件

最終的な成果物として期待すること：
1. フェーズ分けされた実装計画
2. 各フェーズの詳細な技術仕様
3. 具体的なタイムライン
4. リスク管理計画
5. スケーラビリティを考慮した設計方針

補足：実装の優先順位は以下の順で検討してください：
1. メモの作成・保存機能（コア機能）
2. ユーザー認証システム
3. AIによるメモ分析機能
4. 行動提案機能
5. フィードバックシステム
```

こうしてみるとAIに伝わりやすい用語の選定もキーになってくるような気がしますが、とりあえずAIのことはAIが一番わかると思うので質問内容のチェックもしたほうがいいかもって思ってます。

そして作成できたものがこちら。roadmap.mdとしてファイル作成されていました。
```
# memoru開発ロードマップ

## 全体概要

本ロードマップは、memoruアプリケーションの開発を段階的に進めるための計画です。各フェーズは独立してデプロイ可能であり、動作確認可能な状態まで実装することを目標とします。

### 実装優先順位

1.  メモの作成・保存機能（コア機能）
2.  ユーザー認証システム
3.  AIによるメモ分析機能
4.  行動提案機能
5.  フィードバックシステム

### 技術要件

-   インフラ: Supabase
-   バックエンド: Go/Fiber
-   フロントエンド: TypeScript/Next.js
-   データベース: PostgreSQL (Supabase提供)

## フェーズ1: メモの作成・保存機能

### 詳細仕様

-   ユーザーはテキスト形式でメモを作成できる。
-   作成したメモはデータベースに保存される。
-   保存されたメモは一覧表示できる。
-   メモは編集・削除できる。
-   メモは完全に非公開で、作成者のみがアクセスできる。

### 必要な技術コンポーネント

-   **フロントエンド:**
    -   メモ入力フォーム (Next.js)
    -   メモ一覧表示コンポーネント (Next.js)
-   **バックエンド:**
    -   メモ作成API (Go/Fiber)
    -   メモ取得API (Go/Fiber)
    -   メモ更新API (Go/Fiber)
    -   メモ削除API (Go/Fiber)
-   **データベース:**
    -   メモを保存するテーブル (PostgreSQL)

### 想定される技術的課題

-   FiberとSupabaseの連携方法の調査
-   PostgreSQLのスキーマ設計
-   Next.jsでのフォーム処理と状態管理

### テスト項目

-   メモの作成、保存、読み込み、更新、削除が正常に行えること。
-   空のメモが作成できないこと。
-   非常に長いメモが問題なく保存・表示できること。
-   特殊文字を含むメモが問題なく保存・表示できること。
-   複数のメモを作成し、正しく一覧表示されること。

### デプロイ手順

1.  Supabaseプロジェクトを作成し、データベースと認証機能を有効にする。
2.  バックエンドAPIをデプロイする（例：Vercel, Heroku, Google Cloud Run）。
3.  フロントエンドアプリケーションをデプロイする（例：Vercel）。
4.  環境変数を設定する（SupabaseのURL、APIキーなど）。

### 成功基準

-   **達成目標:** ユーザーが基本的なメモの作成・保存・編集・削除機能を問題なく利用できる。
-   **品質基準:**
    -   コードは可読性が高く、保守しやすいこと。
    -   エラーハンドリングが適切に行われていること。
    -   基本的なセキュリティ対策が施されていること（例：入力値の検証）。
-   **パフォーマンス要件:**
    -   メモの作成・取得APIのレスポンスタイムが1秒以内であること。

## フェーズ2: ユーザー認証システム

### 詳細仕様

- ユーザーはメールアドレスとパスワードでサインアップ/ログインできる。
- Supabase Authを利用して認証を実装する。
- 認証状態に応じて、メモの作成・閲覧・編集・削除を制限する。
- ソーシャルログイン（Google, GitHub）を実装する（オプション）。

### 必要な技術コンポーネント

- **フロントエンド:**
  - サインアップ/ログインフォーム (Next.js)
  - 認証状態管理 (Next.js, Supabase Auth)
- **バックエンド:**
  - 認証関連のAPI (Go/Fiber, Supabase Auth)
- **データベース:**
  - ユーザー情報を保存するテーブル (PostgreSQL, Supabase Auth)

### 想定される技術的課題

- Supabase AuthのNext.jsへの統合
- セッション管理
- セキュリティ対策（CSRF、XSSなど）

### テスト項目

- メールアドレスとパスワードでサインアップ/ログインできること。
- 無効なメールアドレスやパスワードでログインできないこと。
- ログアウトできること。
- 認証状態に応じて、メモへのアクセスが制限されること。
- (オプション) ソーシャルログインが正常に機能すること。

### デプロイ手順

1. Supabase Authの設定を行う。
2. フロントエンド/バックエンドのコードを更新し、デプロイする。

### 成功基準

- **達成目標:** ユーザーが安全にサインアップ/ログインでき、認証状態に基づいてメモを管理できる。
- **品質基準:**
  - 認証情報は安全に保管されること。
  - 認証フローがスムーズで、ユーザーフレンドリーであること。
- **パフォーマンス要件:**
  - ログイン/サインアップAPIのレスポンスタイムが1秒以内であること。

## フェーズ3: AIによるメモ分析機能

### 詳細仕様

- 収集されたメモをAI（自然言語処理）で分析する。
- メモの内容から、ユーザーの興味関心、感情、思考パターンなどを抽出する。
- 分析結果は、ユーザーにわかりやすく可視化する（オプション）。

### 必要な技術コンポーネント

- **AI/ML:**
  - 自然言語処理ライブラリ/API (例: Go-nlp, spaCy, Transformers.js, OpenAI API)
- **バックエンド:**
  - メモ分析API (Go/Fiber)
- **フロントエンド:**
  - 分析結果表示コンポーネント (Next.js) (オプション)

### 想定される技術的課題

- 適切な自然言語処理ライブラリ/APIの選定
- 分析結果の精度向上
- プライバシーへの配慮

### テスト項目

- メモの内容が正しく分析されること。
- 分析結果がユーザーに理解しやすい形で提供されること（オプション）。
- 分析処理がパフォーマンスに影響を与えないこと。

### デプロイ手順

1. 選択したAI/MLライブラリ/APIをバックエンドに統合する。
2. バックエンド/フロントエンドのコードを更新し、デプロイする。

### 成功基準

- **達成目標:** メモの内容から、ユーザーの関心事や思考パターンに関する有用な洞察が得られる。
- **品質基準:**
  - 分析結果は、ユーザーのプライバシーを尊重した形で扱われること。
  - 分析結果の精度は、継続的に改善されること。
- **パフォーマンス要件:**
  - メモ分析APIのレスポンスタイムが5秒以内であること。

## フェーズ4: 行動提案機能

### 詳細仕様

- AIによるメモ分析結果に基づいて、ユーザーに具体的な行動を提案する。
- 提案は、ユーザーの目標達成を支援するものでなければならない。
- 提案は、ユーザーが実行可能な範囲のものでなければならない。

### 必要な技術コンポーネント

- **AI/ML:**
  - 提案生成アルゴリズム
- **バックエンド:**
  - 行動提案API (Go/Fiber)
- **フロントエンド:**
  - 提案表示コンポーネント (Next.js)

### 想定される技術的課題

- 提案の質と多様性の確保
- ユーザーのコンテキストを考慮した提案
- 提案のパーソナライズ

### テスト項目

- ユーザーのメモの内容に基づいた適切な行動提案が行われること。
- 提案が具体的で実行可能であること。
- 提案がユーザーの目標達成に役立つものであること。

### デプロイ手順

1. 提案生成アルゴリズムをバックエンドに統合する。
2. バックエンド/フロントエンドのコードを更新し、デプロイする。

### 成功基準

- **達成目標:** ユーザーが、自分のメモに基づいた有用な行動提案を受け取れる。
- **品質基準:**
  - 提案は、ユーザーの目標や状況に合わせてパーソナライズされること。
  - 提案は、ユーザーが受け入れやすい形で提示されること。
- **パフォーマンス要件:**
  - 行動提案APIのレスポンスタイムが3秒以内であること。

## フェーズ5: フィードバックシステム

### 詳細仕様

- ユーザーは、AIによる分析結果や行動提案に対してフィードバックを提供できる。
- フィードバックは、AIモデルの改善に利用される。
- フィードバックは、ユーザーインターフェースを通じて簡単に提供できる。

### 必要な技術コンポーネント

- **フロントエンド:**
  - フィードバック入力コンポーネント (Next.js)
- **バックエンド:**
  - フィードバック収集API (Go/Fiber)
- **データベース:**
  - フィードバックを保存するテーブル (PostgreSQL)

### 想定される技術的課題

- フィードバックの収集と活用のためのデータモデル設計
- フィードバックデータのプライバシー保護

### テスト項目

- ユーザーがフィードバックを簡単に提供できること。
- フィードバックが正しく保存されること。
- フィードバックがAIモデルの改善に活用されること（長期的なテスト）。

### デプロイ手順

1. フィードバック関連のコンポーネントをフロントエンド/バックエンドに追加する。
2. データベースにフィードバックテーブルを作成する。
3. コードを更新し、デプロイする。

### 成功基準

- **達成目標:** ユーザーがAIの分析結果や提案に対してフィードバックを提供し、AIの精度向上に貢献できる。
-   **品質基準:**
    -   フィードバックは、ユーザーのプライバシーを尊重した形で扱われること。
    -   フィードバックシステムは、使いやすく、ユーザーの負担にならないこと。
-   **パフォーマンス要件:**
    -   フィードバック送信APIのレスポンスタイムが1秒以内であること。

## リスク管理計画

-   **技術的リスク:**
    -   AIモデルの精度が低い: 継続的な学習とチューニング、ユーザーフィードバックの活用。
    -   Supabaseの制限: 代替サービスの検討（例：Firebase）。
    -   Go/Fiberの学習コスト: 早期の技術検証、ドキュメントの整備。
-   **スケジュールリスク:**
    -   開発の遅延: 定期的な進捗確認、タスクの優先順位付け、リソースの調整。
-   **その他リスク:**
    -   ユーザーからのフィードバックが少ない: フィードバックシステムの改善、ユーザーへの働きかけ。

## スケーラビリティを考慮した設計方針

-   **データベース:** Supabaseのスケールアップオプションを活用。
-   **バックエンド:** Go/Fiberの水平スケーリング（コンテナ化、ロードバランシング）。
-   **フロントエンド:** Next.jsの静的サイト生成（SSG）とサーバーサイドレンダリング（SSR）の適切な使い分け。
-   **AI/ML:** モデルの軽量化、推論処理の最適化、APIリクエストのキャッシュ。

## タイムライン

| フェーズ                     | 開始     | 終了     | 期間     |
| ---------------------------- | -------- | -------- | -------- |
| フェーズ1: メモの作成・保存機能 | 2025/2/24 | 2025/3/10 | 2週間    |
| フェーズ2: ユーザー認証システム | 2025/3/11 | 2025/3/24 | 2週間    |
| フェーズ3: AIによるメモ分析機能 | 2025/3/25 | 2025/4/21 | 4週間    |
| フェーズ4: 行動提案機能         | 2025/4/22 | 2025/5/19 | 4週間    |
| フェーズ5: フィードバックシステム | 2025/5/20 | 2025/6/2  | 2週間    |
```

なかなかにごつい。。

しかしながらなんかいけそうなステップに分かれているのがいいです。

こうゆう資料作成は僕の能力だとふわっと「まぁなんとかこれでいけるでしょう」くらいのアピールしかできないので非常に助かります。

## とりあえずstep1だけ実現してみる
記事作成の都合上、とりあえずコアとなるメモのCRUDを実現できるstep1をAI駆動開発していきます。

`roadmap.md` のstep1をコード生成してもらいましょう。

ちなみにですが、cursorのcomposer機能もかなり捨て難い選択肢だと思っています。別の機会でAIエージェントサービスを変えてみたて何がどう違うか知見を貯めてみようかなと思ってます。

一モードはarchitectのまま `作成したroadmap.mdを理解し、step1を実現する` といった内容で引き続き指示していきます。

するとcodeモードに変更を促されるので変更を承認(approve)します。

このコード生成の時点でREAD/WRITEはauto approveにしています。理由は壊されたらまずいファイルなど一切ないからで、もしなにかしらのコードがすでにあるのでしたら念の為にWRITE権限は外すことをかなりお勧めします。結構AIは既存コードを破壊してきますので。。。

また、コマンドに関しては承認のボタンを押しながら進めます。これはコマンドが他のプロジェクトやマシンの設定に影響のある場合もあるからで、念の為そんなことさせたくないので目で確認してOKそうなら承認、やばそうなら止めるということをします。

コマンドをオートにして失敗した例を上げると、AIの考えでは既存実装で利用しているライブラリを推奨せす、そのライブラリを廃止した上で実装を書き換え既存コードを書き直すということをしてきたことがあります。

ライブラリの変更に関しては変更するためのコマンド（この例だと `go mod tidy` の実行）をやられたので、ある意味こうゆうコードの破壊方法があるのかぁ・・・みたいな気持ちになってました。

またコード生成時にレートリミットなどのエラーが出ることもありますが、トークンやデポジットの利用限界を迎えなければretryのようなことをすれば大体動き直します。

途中でsupabaseのDatabaseの設定のためにいくつか手動で色々やってくれという内容が表示されてきます。
```
Supabaseのウェブサイト（https://supabase.com/）にアクセスし、アカウントを作成またはログインします。
新しいプロジェクトを作成します。
プロジェクトのダッシュボードから、「Database」を選択し、「New table」をクリックしてmemosテーブルを作成します。
```
ウェブサイトにアクセスし、必要に応じてアカウント登録・組織作成・プロジェクト作成を行なっていきます。

僕の場合はこれら全ての作成を実施した上で、databaseのパスワードは.envファイルを作成し、そこに保存しました。そしてdatabaseを言われている内容で設定・作成します

さらに認証についてもこのようなことを言われているので後々のステップで利用してきます。
```
「Authentication」を選択し、「Enable Email & Password」を有効にします（今回は認証機能はまだ使用しませんが、後のフェーズで必要になります）。
プロジェクトのAPIキーとURLを取得します。「Settings」>「API」から確認できます。
```

また、プロジェクトURLやAPI KEYも問われましたので.envにそれらも追記し保存、それを読み込むように指示します。
```
求められている情報はプロジェクトルートにある.envファイルに保管しました。
そこから適宜読み取り利用してください。
```

そして、依頼と承認をこなし続けていくとbackendとfrontendが出来上がりそれぞれ起動し動作内容を確認できるところまで進みました。

## 動作確認
ここからは生成したものを確認していくので、もしこの記事を参考になにかしら手を動かしている方がいる場合、こちらの状況とは全く違うことになっているかもしれません。

あくまで僕の場合だと思ってみてもらえるとありがたいです。

### とりあえずPOSTメソッドが動かない
メモの保存ボタンを押すと404エラーが発生しています。

ブラウザの開発モードの画面から該当するネットワークやコンソール上のエラーを拾って修正依頼をログとともにします。

この時の修正依頼時に思ったことなのですが、一旦動作確認を求められた時点でコミットを残したほうがいいと思いました。

というのも、動作確認から先はbugfixのターンになるのと、次回以降同じようなコード生成をするのだったらプロンプトに注意点として書けば事前に防げると思ったからです。

気づいた時点でコミットして修正したあとの動作を逐一確認してbugfixを進めていきます。

### バグフィックスの精度とスピードでモデルを決める
ここまでは0->1でコードもある程度は自由でよかったのですが、bugfixになると話は別です。

AI agentと使っているとバグAに対して１という手順で修正、解決していないので２という手順で修正、解決していないので１という手順で修正・・・・

というようにループし始めることがあります。

そのため、そのようなことが起きたらモデルを変えて実行することにしています。今回の場合はgeminiで難しいようならclaudeに変えるという感じです。

また生成のスピードについてもmodelによってまばらです。自分に合うモデルを探すというのも必要になると思いますが、とりあえずはgeminiとclaudeが選択肢の二強な気がします。

### 細かな修正は手動の方が早い
bugfixを進めるとちょっとしたimportができてないとか、変数が違うとか、細かすぎる部分がコンパイルエラーになっていたりします。

またAIに修正をしてもらう以前に自分でやった方が早いものも出てきますので、そのようなときは手動で直してしまうのが早いし正確だし、なんといっても自分の手を動かすだけなら無料！（と言えなくもない）

あまりAIにこだわることなくAIor自分を取捨選択できるといいかなと思いました。



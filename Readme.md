# go_react_template
## 概要
Dockerを使ったGo(Gin)+Typescript(React)でアプリケーションを作るときの雛形

## 使い方
### ビルド

```
cd src \
&& go mod init gin_react_template \
&& go mod tidy \
&& cd ..
```

```
docker-compose build
```

```
docker-compose run --rm react sh -c 'npx create-react-app react_app --template typescript'
```

### 起動
```
docker-compose up
```

# tune-bot core
Go module for interacting with tune-bot's various services.

Downloads songs as they are added using [yt-dlp](https://github.com/yt-dlp/yt-dlp).


```
git commit -am "my update"
git tag vx.x.x
git push origin vx.x.x
```

## Docker
```
docker build -t tune-bot-database .
docker run -p 3306:3306 -td tune-bot-database
```

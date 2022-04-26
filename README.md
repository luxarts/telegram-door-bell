# telegram-door-bell
Telegram Bot that sends a message when your door bell rings.

### Run Redis
```
docker run -d \
    --name redis-stack \
    -v tdbredisdata:/var/lib/redis/data \
    -p 6379:6379 \
    -e REDIS_ARGS="--requirepass redis" \
    arm64v8/redis
```
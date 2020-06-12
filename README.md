# Makaba Engine Posting
Говнокод на golang для отправки сообщений на Два.ч используя пасскод.

Используется их публичный Makaba Engine API https://2ch.hk/api/index.html. Отправка сообщений только по пасскоду, потому что капчу я вертел 20 раз и это не вайпалка, уважайте ребят из Olaneta. Ola-la~. Купить пасскод можно тут https://2ch.hk/2ch/ 

## Как использовать?

У тебя уже должен быть установлен golang.org. Забираешь себе сорцы.
```
git clone https://github.com/wmw9/makaba-engine-posting
cd ./makaba-engine-posting
```

Потом в environment ОС сетишь свой пасскод

```
export PASSCODE=2cHKDgdfzlYRieBlcrQ7wMPU4R91gB21
```
Внутри кода в post.go укажи доску и номер треда. Теперь запускаем:

```
go run *.go
```
Готово!

# P.S.

**Автор**: https://instagram.com/wmw


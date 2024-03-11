# Week 1 - Marketpalce App

### Link notion

[https://openidea.notion.site/Brief-ProjectSprint-1-Marketplace-app-e95b110f434142a5a144de1e7fc98cef?pvs=4](Ling disini)

## Development

### env

Jangan lupa populate file `.env` dengan file `.env.example` yang sudah ada.

```bash
cp .env.example .env
```

### Database

Database bisa di run dengan menggunakan `docker-compose` dengan perintah dibawah ini.

```bash
make docker-up
```

### Migration

Tools untuk migrasi database, dengan menggunakan [migrate](https://github.com/golang-migrate/migrate/tree/master) dan `makefile`. Untuk migrasi database dengan perintah dibawah ini, nanti file sql akan ada di folder `db/migrations`.

```bash
make migrate-create filename.sql
```

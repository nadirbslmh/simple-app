# simple-app

Simple REST API for demonstration purpose. This application is deployed with [fly.io](https://fly.io/).

## How to deploy?

1. Buat akun [fly.io](https://fly.io/app/sign-in).

2. Install `flyctl` [link](https://fly.io/docs/flyctl/installing/).

3. Login dari CLI.

```sh
flyctl auth login
```

4. Buat folder baru untuk membuat database MySQL.

```sh
mkdir my-mysql
cd my-mysql
```

6. Buat aplikasi untuk database MySQL.

```sh
fly launch
```

7. Isi nama aplikasi, nama bebas seperti `coba-mysql`.

8. Membuat volume agar data di dalam MySQL bisa disimpan.

```sh
fly volumes create mysqldata --size 10
```

9. Membuat MySQL password.

```sh
fly secrets set MYSQL_PASSWORD=password MYSQL_ROOT_PASSWORD=password
```

10. Ubah file `fly.toml` di dalam folder `my-mysql`.

```
app = "coba-mysql"
kill_signal = "SIGINT"
kill_timeout = 5

[mounts]
  source="mysqldata"
  destination="/data"

[env]
  MYSQL_DATABASE = "dbku"
  MYSQL_USER = "coolkid"

[build]
  image = "mysql:8"

[experimental]
  cmd = [
    "--default-authentication-plugin", "mysql_native_password",
    "--datadir", "/data/mysql",
    "--performance-schema=OFF",
    "--innodb-buffer-pool-size", "64M"
  ]

```

11. Deploy aplikasi database MySQL.

```sh
fly deploy
```

12. Lakukan perubahan pada bagian koneksi database. Contohnya dapat dilihat di file `database.go`.

13. Lakukan perubahan juga di file `main.go`. Contohnya dapat dilihat di file `main.go`.

14. Persiapkan file untuk Docker yaitu `Dockerfile`.

15. Lakukan deployment pada aplikasi Golang. Pastikan berada di directory aplikasi Golang yang akan di deploy.

```sh
cd golangapp

flyctl launch
```

16. Isi nama aplikasi, nama aplikasi bebas.

17. Ketika ada pertanyaan mengenai database Postgre, ketik `N` karena database MySQL sudah digunakan.

18. Ketik `y` karena aplikasi siap untuk di deploy.

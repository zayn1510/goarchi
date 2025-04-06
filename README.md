**goarchi** adalah sebuah mini-framework open-source berbasis Golang yang menggunakan **Gin**, **GORM**, dan struktur clean architecture. Framework ini memudahkan kamu dalam membangun REST API dengan arsitektur yang rapi, scalable, dan cepat dalam pengembangan.

---

## 🚀 Fitur Utama

- 🛠️ generator built-in
- 🔁 Arsitektur Clean (MVC + Service + Request + Resource)
- 🔐 Middleware JWT siap pakai
- 🌐 Dukungan CORS (bisa diaktif/nonaktif)
- 🧪 Struktur routing yang mudah dan terpisah
- 📁 Konfigurasi lewat file `.env`

---

## 📦 Instalasi

### 1. Clone repo

```bash
git clone https://github.com/zayn1510/goarchi.git my-app
cd my-app
2. Salin file environment


Copy
cp .env.example .env
Lalu atur konfigurasi database kamu:

env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
3. Install dependency Go
go mod tidy

🛠️ CLI Generator
Framework ini menyediakan generator bawaan yang menggunakan makefile. Kamu bisa membuat berbagai komponen hanya dengan satu baris perintah.

make controller name=User	Generate controller User
make service name=User	Generate service User
make request name=User	Generate request User
make resource name=User	Generate resource User
make models name=User	Generate model dan migration User
📂 Struktur Folder
.
├── app/
│   ├── controllers/
│   ├── services/
│   ├── requests/
│   ├── resources/
│   └── models/
│   └── seeders/
│   └── migarations/
│   └── migrate/
├── config/
│   └── database.go/
├── middleware/
├── routes/
│   └── web.go
├── .env.example
├── main.go
└── go.mod
🧬 Menjalankan Project
bash
Always show details

Copy
go run main.go
Project akan berjalan di http://localhost:8080

🧱 Middleware
✅ Mengaktifkan / Menonaktifkan CORS
Di dalam main.go, kamu bisa mengatur CORS dengan baris berikut:

router := gin.Default()
middleware.SetCors(router) // aktifkan CORS
routers.RegisterRoutes(router)
router.Run(":8080")
Jika ingin menonaktifkan, cukup hapus atau komen middleware.SetCors(router).

🔐 Middleware JWT
Untuk mengamankan grup route dengan JWT, kamu tinggal gunakan:
users := api.Group("users")
users.Use(middleware.JWTMiddleware())
🔀 Routing
Semua definisi routing API dilakukan di file routes/web.go. Kamu bisa mengatur grouping dan handler di sana.

api := r.Group("/api/v1")
UserRouter(api) // memanggil router khusus user
🤝 Kontribusi
Pull request dan issue sangat terbuka untuk siapa saja yang ingin ikut berkontribusi. Yuk bantu kembangkan bareng!

📄 Lisensi
Framework ini dirilis di bawah lisensi MIT.

🙌 Terima Kasih
Framework ini dibangun dengan semangat open-source dan kolaborasi. Semoga bermanfaat buat proyek-proyek kamu! """

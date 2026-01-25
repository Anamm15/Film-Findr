# 🎬 FilmFindr – Temukan, Simpan, dan Review Film Favoritmu

**FilmFindr** adalah platform film interaktif yang dirancang untuk para pencinta film sejati. Temukan film menarik, berikan ulasan pribadi, dan simpan film ke dalam daftar tontonmu hanya dengan satu akun!

---

## 🚀 Fitur Unggulan

✨ **Autentikasi & Manajemen User (Secure Authentication)**  
Sistem autentikasi aman menggunakan mekanisme token (JWT / session-based)  
untuk memastikan akses pengguna terkontrol dan data tetap terlindungi.

📝 **Review & Rating Film**  
Pengguna dapat menulis ulasan dan memberikan rating pada film.  
Data review disimpan secara terstruktur untuk keperluan analitik dan rekomendasi.

🔥 **Weekly Trending Film**  
Menampilkan daftar film yang sedang trending setiap minggu berdasarkan:
- Frekuensi review & rating
- Aktivitas pengguna
- Popularitas dalam periode waktu tertentu  
Menggunakan agregasi data dan time-based analysis.

📊 **Sentiment Analysis Review**  
Ulasan pengguna dianalisis untuk mengklasifikasikan sentimen:
- Positif
- Negatif  

Hasil analisis digunakan untuk:
- Insight kualitas film
- Ringkasan opini publik
- Pendukung sistem rekomendasi

🤖 **AI Chatbot Film Assistant**  
- Memberikan rekomendasi film berdasarkan preferensi pengguna
- Menjawab pertanyaan seputar film, genre, aktor, atau review
- Membantu eksplorasi film secara interaktif dan kontekstual

🎯 **Watch List Pribadi**  
Fitur manajemen watch list untuk menyimpan film yang ingin ditonton,  
lengkap dengan status (planned, watching, completed).

🔎 **Eksplorasi & Pencarian Film**  
Fitur pencarian dan eksplorasi film berdasarkan:
- Judul
- Sinopsis
- Tahun Rilis
- Genre
Dirancang untuk skalabilitas dan performa query yang optimal.


---

## 🛠️ Teknologi yang Digunakan

| Layer       | Teknologi                  |
|-------------|----------------------------|
| Frontend    | ⚛️ React.js                |
| Backend     | 🧠 Golang (Gin + GORM)     |
| Database    | 🐘 PostgreSQL              |

Arsitektur modern yang memisahkan frontend dan backend demi skalabilitas dan fleksibilitas pengembangan.


---

## 🧪 Cara Menjalankan di Lokal

### 1. Clone Proyek
```bash
git clone https://github.com/Anamm15/film-findr.git
cd film-findr
```

### 2. Atur Backend
#### a. Masuk ke direktori backend 
```bash
cd backend
```

#### b. Buat file .env dan isi dengan
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=film-findr
PORT=8080
```

#### c. Jalankan backend
```bash
go mod tidy
go run main.go
```

### 2. Atur Frontend
#### a. Masuk ke direktori frontend 
```bash
cd frontend
```

#### b. Install dependency
```bash
npm install
```

#### c. Jalankan React dev server
```bash
npm run dev
```



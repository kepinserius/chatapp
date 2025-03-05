# Aplikasi Chat Real-Time Mobile

## Deskripsi

Aplikasi Chat Real-Time Mobile adalah aplikasi perpesanan instan yang memungkinkan pengguna untuk berkomunikasi secara real-time. Aplikasi ini dibangun dengan teknologi WebSocket untuk memastikan pengiriman dan penerimaan pesan secara instan.

## Fitur Utama

- **Autentikasi Pengguna**: Sistem login dan registrasi yang aman
- **Ruang Chat**: Pengguna dapat membuat dan bergabung ke berbagai ruang chat
- **Pesan Real-Time**: Pesan dikirim dan diterima secara instan tanpa perlu refresh
- **Daftar Pengguna Online**: Melihat siapa saja yang sedang aktif di ruang chat
- **Antarmuka Responsif**: Desain yang nyaman digunakan di berbagai ukuran layar

## Teknologi yang Digunakan

### Backend
- **Golang**: Bahasa pemrograman utama untuk backend
- **Gin**: Framework web untuk Golang
- **WebSocket**: Untuk komunikasi real-time
- **PostgreSQL**: Database untuk menyimpan pesan dan data pengguna

### Frontend
- **React Native**: Framework untuk pengembangan aplikasi mobile
- **TypeScript**: Untuk type-safety dan pengembangan yang lebih baik
- **React Navigation**: Untuk navigasi antar halaman
- **Axios**: Untuk komunikasi HTTP dengan backend
- **AsyncStorage**: Untuk penyimpanan lokal

## Struktur Aplikasi

### Backend
```
backend/
├── api/
│   ├── auth.go
│   ├── room.go
│   ├── router.go
│   └── websocket.go
├── database/
│   └── db.go
├── models/
│   ├── message.go
│   ├── room.go
│   └── user.go
├── websocket/
│   ├── client.go
│   └── hub.go
├── .env
├── go.mod
├── go.sum
└── main.go
```

### Frontend
```
frontend/
├── src/
│   ├── components/
│   │   ├── MessageItem.tsx
│   │   ├── RoomItem.tsx
│   │   └── UserList.tsx
│   ├── context/
│   │   └── AuthContext.tsx
│   ├── screens/
│   │   ├── LoginScreen.tsx
│   │   ├── RegisterScreen.tsx
│   │   ├── RoomListScreen.tsx
│   │   └── ChatScreen.tsx
│   ├── services/
│   │   ├── api.ts
│   │   └── websocket.ts
│   └── types/
│       └── index.ts
├── App.tsx
├── package.json
└── tsconfig.json
```

## Alur Kerja Aplikasi

1. **Autentikasi**: Pengguna mendaftar atau masuk ke aplikasi
2. **Daftar Ruang Chat**: Pengguna melihat daftar ruang chat yang tersedia
3. **Membuat Ruang Chat**: Pengguna dapat membuat ruang chat baru
4. **Bergabung ke Ruang Chat**: Pengguna memilih ruang chat untuk bergabung
5. **Komunikasi Real-Time**: Pengguna dapat mengirim dan menerima pesan secara instan
6. **Melihat Pengguna Online**: Pengguna dapat melihat siapa saja yang sedang aktif di ruang chat

## Cara Menjalankan Aplikasi

### Prasyarat
- Go (versi 1.16 atau lebih baru)
- Node.js (versi 14 atau lebih baru)
- PostgreSQL
- React Native CLI atau Expo CLI

### Langkah-langkah Menjalankan Backend

1. **Siapkan Database PostgreSQL**
   ```
   createdb chatapp
   ```

2. **Konfigurasi Environment Variables**
   Buat file `.env` di folder `backend` dengan isi:
   ```
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=password_anda
   DB_NAME=chatapp
   DB_PORT=5432
   SERVER_PORT=8080
   ```

3. **Instal Dependensi Backend**
   ```
   cd backend
   go mod tidy
   ```

4. **Jalankan Server Backend**
   ```
   go run main.go
   ```
   Server akan berjalan di `http://localhost:8080`

### Langkah-langkah Menjalankan Frontend

1. **Instal Dependensi Frontend**
   ```
   cd frontend
   npm install
   ```

2. **Konfigurasi API URL**
   Buka file `src/services/api.ts` dan sesuaikan URL API:
   - Untuk emulator Android: `http://10.0.2.2:8080/api`
   - Untuk perangkat fisik: Gunakan IP komputer Anda, misalnya `http://192.168.1.100:8080/api`
   - Untuk iOS simulator: `http://localhost:8080/api`

3. **Jalankan Aplikasi**
   
   Untuk React Native CLI:
   ```
   npx react-native run-android
   # atau
   npx react-native run-ios
   ```
   
   Untuk Expo:
   ```
   npx expo start
   ```

## Troubleshooting

- **Masalah Koneksi Backend**: Pastikan server backend berjalan dan dapat diakses dari perangkat/emulator
- **Masalah WebSocket**: Periksa konfigurasi CORS di backend
- **Masalah Database**: Pastikan PostgreSQL berjalan dan kredensial di file `.env` benar

## Pengembangan Lebih Lanjut

Beberapa fitur yang dapat ditambahkan untuk pengembangan lebih lanjut:
- Notifikasi push untuk pesan baru
- Fitur berbagi media (gambar, video, file)
- Status "sedang mengetik"
- Fitur grup chat dengan admin
- End-to-end encryption untuk keamanan pesan

---

Dengan mengikuti panduan ini, Anda dapat menjalankan dan mengembangkan aplikasi Chat Real-Time Mobile dengan sukses. Jika ada pertanyaan atau masalah, silakan buat issue di repositori ini.

# Tugas Pemrograman Web 3

## Deskripsi
Proyek ini adalah aplikasi web sederhana yang menampilkan informasi tentang buku menggunakan bahasa pemrograman Go. Aplikasi ini menggunakan `net/http` sebagai web server dan `html/template` untuk rendering halaman.

## Struktur Direktori
```
TugasPemogreamanWeb3/
├── .git/                # Direktori Git
├── controller/          # Controller untuk menangani request
│   ├── Index_book.go    # Controller utama untuk halaman buku
├── entity/              # Struktur data utama
│   ├── book.go          # Struct Book
│   ├── publisher.go     # Struct Publisher
├── routes/              # Routing aplikasi
│   ├── routes.go        # Konfigurasi rute aplikasi
├── views/               # Folder untuk file HTML
│   ├── bookInfo.html    # Template tampilan informasi buku
├── go.mod               # Modul Go
├── main.go              # Entry point aplikasi
└── README.md            # Dokumentasi proyek ini
```

## Cara Menjalankan
### 1. Pastikan Go telah terinstal
Jika belum, unduh dan instal Go dari [golang.org](https://golang.org/).

### 2. Clone repository (jika diperlukan)
```sh
 git clone https://github.com/nabilulilalbab/assigmentpeomgramanweb3.git
 cd assigmentpeomgramanweb3
```

### 3. Jalankan Aplikasi
```sh
go run main.go
```
Aplikasi akan berjalan di `http://localhost:8080`.

## Teknologi yang Digunakan
- **Go**: Backend utama
- **net/http**: Untuk menjalankan server HTTP
- **html/template**: Untuk rendering halaman HTML
- **embed**: Untuk menyematkan file HTML ke dalam binary

## Fungsi Utama
### 1. **Menampilkan Informasi Buku**
Ketika pengguna mengakses `http://localhost:8080`, mereka akan melihat informasi buku berikut:
- **Judul Buku**
- **Penulis**
- **Harga**
- **Nama dan Kota Penerbit**

## Kontribusi
Jika ingin berkontribusi, silakan fork proyek ini dan buat pull request. Pastikan kode yang ditambahkan telah diuji dengan baik.

## Lisensi
Proyek ini dibuat untuk keperluan tugas kuliah dan dapat digunakan dengan bebas.

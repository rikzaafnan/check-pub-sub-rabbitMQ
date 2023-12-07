# Gunakan image node terbaru sebagai dasar
FROM node:14

# Setel direktori kerja di dalam kontainer
WORKDIR /app

# Salin package.json dan package-lock.json (jika ada)
COPY package*.json ./

# Install dependensi
RUN npm install

# Salin seluruh kode sumber proyek ke dalam kontainer
COPY . .

# Ekspose port yang digunakan oleh aplikasi
EXPOSE 3000

# Perintah untuk menjalankan aplikasi ketika kontainer dimulai
CMD ["node", "index.js"]

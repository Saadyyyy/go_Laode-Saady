# 17 . ORM & Code Structure (MVC)

ORM (Object-Relational Mapping) adalah sebuah konsep dalam pengembangan perangkat lunak yang digunakan untuk menghubungkan antara model objek dalam aplikasi dengan struktur data dalam database relasional. ORM memungkinkan pengembang untuk berinteraksi dengan database menggunakan objek dan bahasa pemrograman yang lebih familiar, daripada menulis query SQL secara langsung. Di sisi lain, MVC (Model-View-Controller) adalah pola desain arsitektur yang digunakan untuk memisahkan tiga komponen utama dalam aplikasi: model, tampilan, dan pengendali. Berikut adalah penjelasan dan poin-poin utama dari keduanya:

### ORM (Object-Relational Mapping):

1. **Definisi**: ORM adalah sebuah teknik atau kerangka kerja yang memungkinkan pengembang untuk mengakses dan memanipulasi data dalam database relasional menggunakan objek dan konsep pemrograman berorientasi objek.
2. **Tujuan**: ORM bertujuan untuk mengurangi kompleksitas dalam pengembangan perangkat lunak dengan menyediakan cara yang lebih mudah dan efisien untuk berinteraksi dengan database.
3. **Keuntungan**:
   - **Abstraksi Database**: Mengabstraksi detail teknis dari database, seperti tabel, kolom, dan koneksi, sehingga pengembang tidak perlu terlalu memahami detail-detail ini.
   - **Kode yang Lebih Bersih**: Memungkinkan pengembang untuk menulis kode yang lebih bersih dan lebih mudah dipahami dengan menggunakan objek dan metode pemrograman berorientasi objek.
   - **Portabilitas**: Memungkinkan aplikasi untuk lebih mudah berpindah dari satu database ke database lainnya tanpa mengubah kode aplikasi.
4. **Contoh ORM**: Beberapa kerangka kerja ORM yang populer termasuk Hibernate untuk Java, SQLAlchemy untuk Python, dan Entity Framework untuk .NET.

### MVC (Model-View-Controller):

1. **Definisi**: MVC adalah pola desain arsitektur yang memisahkan aplikasi menjadi tiga komponen utama: Model, View, dan Controller.
2. **Tujuan**: Tujuan dari MVC adalah untuk memisahkan tugas-tugas dalam aplikasi sehingga kode lebih terstruktur, lebih mudah dipelihara, dan lebih fleksibel.
3. **Komponen-komponen utama**:
   - **Model**: Bertanggung jawab untuk mengelola data dan bisnis logic. Ini adalah bagian yang berhubungan dengan database atau sumber data lainnya.
   - **View**: Bertanggung jawab untuk menampilkan informasi kepada pengguna. Ini berfokus pada tampilan antarmuka pengguna.
   - **Controller**: Bertanggung jawab untuk mengatur aliran aplikasi dan mengendalikan interaksi antara Model dan View. Controller menerima input dari pengguna dan memutuskan bagaimana data harus diambil atau diperbarui dari Model dan bagaimana tampilan harus diperbarui.
4. **Keuntungan**:
   - **Pisahkan Tugas**: Memisahkan tugas-tugas aplikasi menjadi komponen-komponen yang terpisah sehingga mempermudah pemeliharaan dan pengembangan.
   - **Fleksibilitas**: Memungkinkan perubahan pada satu komponen tanpa memengaruhi komponen lainnya, asalkan antarmuka tetap konsisten.
   - **Skalabilitas**: Memungkinkan aplikasi untuk tumbuh dengan lebih baik karena komponen-komponennya dapat digantikan atau diperluas secara independen.

Kombinasi penggunaan ORM dengan struktur MVC adalah praktik umum dalam pengembangan aplikasi web dan desktop modern untuk mencapai kode yang lebih bersih, terstruktur, dan mudah dipelihara. Dengan ORM, Model sering kali merupakan representasi objek dari data dalam database, sementara Controller mengatur interaksi dengan Model dan View digunakan untuk menampilkan data kepada pengguna.

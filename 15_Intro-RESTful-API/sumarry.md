# 15. RESTful API

RESTful API adalah singkatan dari Representational State Transfer API. Ini adalah gaya arsitektur yang digunakan dalam pengembangan aplikasi web untuk mengelola dan mentransfer data antara sistem yang berbeda. RESTful API didasarkan pada sejumlah prinsip dan konsep yang membantu dalam merancang sistem yang efisien, scalable, dan mudah diakses. Berikut adalah beberapa poin utama yang menjelaskan RESTful API:

1. **Resource (Sumberdaya):** Dalam REST, segala sesuatu dianggap sebagai sumber daya, seperti objek, data, atau layanan. Setiap sumber daya diberikan URI (Uniform Resource Identifier) yang unik sebagai alamatnya.

2. **HTTP Methods (Metode HTTP):** RESTful API menggunakan metode HTTP untuk berinteraksi dengan sumber daya. Metode HTTP yang paling umum digunakan adalah:
   - **GET:** Mengambil data dari sumber daya.
   - **POST:** Membuat sumber daya baru.
   - **PUT:** Mengubah atau memperbarui sumber daya.
   - **DELETE:** Menghapus sumber daya.
   - **PATCH:** Mengubah sebagian dari sumber daya.
3. **Representations (Representasi):** Sumber daya dalam RESTful API dapat memiliki beberapa representasi, seperti JSON, XML, atau HTML. Klien dapat memilih representasi yang sesuai dengan kebutuhan mereka dengan menggunakan header HTTP yang benar.

4. **Stateless (Tanpa Status):** Setiap permintaan klien ke server harus mengandung semua informasi yang diperlukan untuk memahami dan memproses permintaan tersebut. Server tidak menyimpan status klien antara permintaan. Ini meningkatkan skalabilitas sistem.

5. **Client-Server Architecture (Arsitektur Klien-Server):** REST memisahkan klien (pengguna antarmuka) dari server (penyedia sumber daya). Ini memungkinkan perkembangan dan evolusi independen dari keduanya.

6. **Layered System (Sistem Bertingkat):** REST memungkinkan penggunaan sistem berlapis (layered system), di mana setiap lapisan memiliki tugas dan tanggung jawab yang terdefinisi dengan baik. Ini memungkinkan fleksibilitas dan peningkatan sistem yang lebih mudah.

7. **Uniform Interface (Antarmuka Seragam):** REST memiliki antarmuka seragam yang didefinisikan dengan baik melalui metode HTTP standar seperti GET, POST, PUT, DELETE, dan lainnya. Ini mempermudah penggunaan dan pemahaman API.

8. **Statelessness (Tanpa Kehandalan):** Setiap permintaan dari klien ke server harus berisi semua informasi yang diperlukan, tanpa mengandalkan permintaan sebelumnya. Ini membuat sistem lebih sederhana dan dapat di-cache.

9. **HATEOAS (Hypermedia As The Engine Of Application State):** Ini adalah konsep di mana API memberikan tautan ke sumber daya terkait dalam responsnya. Dengan kata lain, klien dapat menavigasi sistem secara dinamis melalui tautan yang disediakan oleh server.

10. **Resource Identification Through URI (Identifikasi Sumber Daya Melalui URI):** Setiap sumber daya dalam RESTful API diidentifikasi oleh URI yang unik. URI ini digunakan untuk mengakses, mengelola, dan berinteraksi dengan sumber daya tersebut.

RESTful API adalah pendekatan yang populer dalam pengembangan web karena kemudahan penggunaan, skabilitas, dan interoperabilitasnya. Ini telah menjadi standar de facto dalam pengembangan layanan web yang efisien.

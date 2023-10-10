## 21. DOCKER

Docker adalah platform perangkat lunak yang digunakan untuk mengembangkan, mengemas, dan menjalankan aplikasi dalam wadah (container). Container adalah lingkungan terisolasi yang berisi semua yang diperlukan untuk menjalankan aplikasi, termasuk kode, perpustakaan, dan dependensi. Docker memungkinkan aplikasi untuk dijalankan secara konsisten di berbagai lingkungan, baik di mesin pengembangan lokal maupun di lingkungan produksi yang kompleks.

Berikut adalah beberapa poin penting tentang Docker:

1. **Containerization**: Docker menggunakan teknologi container untuk mengisolasi aplikasi. Ini memungkinkan aplikasi dan dependensinya untuk dijalankan dalam lingkungan yang terisolasi, yang membuat aplikasi lebih mudah dipindahkan dan dikelola.

2. **Image**: Image Docker adalah paket yang berisi semua yang diperlukan untuk menjalankan aplikasi, termasuk kode, runtime, perpustakaan, dan pengaturan. Image ini digunakan sebagai dasar untuk membuat wadah.

3. **Container**: Container adalah instance yang berjalan dari image Docker. Container dapat dimulai, dihentikan, dimodifikasi, dan dihapus sesuai kebutuhan. Mereka berjalan di atas sistem operasi host, tetapi terisolasi dari host dan kontainer lainnya.

4. **Dockerfile**: Dockerfile adalah file teks yang digunakan untuk mendefinisikan bagaimana image Docker harus dibangun. Ini berisi langkah-langkah untuk menginstal perangkat lunak, menyalin file, mengatur pengaturan, dan lainnya.

5. **Registry**: Docker menyediakan registri pusat tempat image Docker dapat disimpan dan dibagikan. Docker Hub adalah salah satu registry publik yang paling populer, tetapi Anda juga dapat membuat registry pribadi.

6. **Orkestrasi**: Docker memiliki alat-alat seperti Docker Compose, Kubernetes, dan Swarm yang memungkinkan Anda untuk mengelola dan mengeksekusi wadah dalam lingkungan yang lebih kompleks, termasuk clustering, penjadwalan, dan otomatisasi.

7. **Portability**: Docker container dapat dijalankan di berbagai lingkungan, termasuk mesin fisik, mesin virtual, cloud, dan pusat data. Ini membuat aplikasi lebih mudah dipindahkan dari satu lingkungan ke lingkungan lainnya.

8. **Efisien**: Docker menghemat sumber daya dengan memungkinkan beberapa wadah berbagi kernel sistem operasi yang sama. Ini membuatnya lebih efisien daripada virtualisasi tradisional.

9. **Ekosistem Luas**: Docker memiliki ekosistem yang luas dan aktif, dengan ribuan image dan proyek terbuka yang tersedia di Docker Hub dan GitHub.

10. **Keamanan**: Docker memiliki fitur keamanan seperti pengelolaan hak akses, izin, dan pengontrolan sumber daya untuk menjaga aplikasi yang berjalan dalam wadah tetap aman.

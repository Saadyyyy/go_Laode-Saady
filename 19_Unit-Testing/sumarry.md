# 19 . Unit Testing

Unit Testing adalah sebuah teknik pengujian perangkat lunak yang digunakan untuk menguji komponen-komponen perangkat lunak secara terisolasi, seperti fungsi atau metode, untuk memastikan bahwa setiap komponen berfungsi sesuai dengan yang diharapkan. Tujuan utama dari unit testing adalah untuk memastikan bahwa setiap unit kode berperilaku dengan benar dan menghasilkan hasil yang sesuai dengan spesifikasi.

Berikut adalah beberapa poin penting tentang unit testing:

1. **Terisolasi**: Unit testing dilakukan pada tingkat yang sangat terisolasi, yaitu pada tingkat komponen individu. Hal ini memungkinkan pengujian untuk fokus secara khusus pada satu fungsi atau metode tanpa dipengaruhi oleh komponen lain dalam perangkat lunak.

2. **Automatisasi**: Unit testing harus dapat diotomatisasikan sepenuhnya. Ini berarti pengujian dapat dijalankan secara otomatis tanpa intervensi manusia. Alat-alat pengujian otomatis seperti JUnit (untuk Java) atau pytest (untuk Python) sering digunakan untuk melakukan unit testing.

3. **Menggunakan Data Pengujian yang Terkontrol**: Unit testing melibatkan penggunaan data pengujian yang sudah diketahui dan terkontrol dengan baik. Ini termasuk data masukan yang mencakup kondisi ekstrem dan batasan-batasan yang mungkin timbul.

4. **Menetapkan Ekspektasi**: Sebelum menjalankan unit test, Anda harus menetapkan ekspektasi tentang hasil yang diharapkan dari unit yang diuji. Ini bisa berupa asertasi bahwa hasil fungsi atau metode harus sama dengan nilai yang sudah diketahui.

5. **Dilakukan Secara Berulang**: Unit testing harus dilakukan secara berulang setiap kali ada perubahan dalam kode sumber. Ini membantu untuk mendeteksi perubahan yang tidak diinginkan atau kerusakan yang mungkin muncul akibat perubahan tersebut.

6. **Menggunakan Kerangka Pengujian**: Biasanya, unit testing melibatkan penggunaan kerangka pengujian yang menyediakan fungsi-fungsi dan alat-alat yang diperlukan untuk melakukan pengujian dengan mudah.

7. **Membantu dalam Refaktorisasi**: Unit testing memungkinkan pengembang untuk melakukan refaktorisasi kode dengan lebih percaya diri. Saat kode diubah, unit test akan memastikan bahwa perilaku yang diharapkan tetap konsisten.

8. **Mengurangi Debugging**: Dengan melakukan unit testing secara intensif, banyak bug dapat terdeteksi dan diperbaiki pada tahap awal pengembangan, yang dapat mengurangi waktu dan biaya debugging di tahap yang lebih lanjut.

9. **Memungkinkan Integrasi Terus-Menerus**: Unit testing adalah bagian penting dari praktik integrasi berkelanjutan (CI/CD) di mana setiap perubahan kode diuji secara otomatis sebelum diintegrasikan ke dalam kode basis utama.

10. **Meningkatkan Kualitas Perangkat Lunak**: Kesalahan dapat dideteksi lebih awal, dan perangkat lunak memiliki tingkat keandalan yang lebih tinggi karena unit testing membantu meminimalkan kesalahan yang dapat muncul.

Unit testing adalah praktik yang sangat penting dalam pengembangan perangkat lunak karena membantu memastikan bahwa setiap komponen perangkat lunak berfungsi sebagaimana mestinya dan berkontribusi pada kualitas keseluruhan perangkat lunak.

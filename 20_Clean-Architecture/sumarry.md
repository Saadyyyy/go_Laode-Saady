# 20 . Clean Architecture

Clean Architecture adalah sebuah konsep arsitektur perangkat lunak yang dirancang untuk memisahkan berbagai komponen dalam suatu aplikasi dengan tujuan agar aplikasi menjadi lebih terstruktur, mudah dipahami, mudah diuji, dan mudah untuk dipelihara. Konsep ini diperkenalkan oleh Robert C. Martin, seorang pakar dalam dunia pengembangan perangkat lunak, dan merupakan evolusi dari konsep arsitektur sebelumnya seperti Hexagonal Architecture dan Ports and Adapters.

Clean Architecture mengusung beberapa prinsip utama, yang dapat dijelaskan dengan beberapa poin penting berikut:

1. **Independence of Frameworks:** Komponen-komponen inti dari aplikasi harus bebas dari ketergantungan pada framework tertentu. Artinya, kode bisnis atau domain perusahaan tidak boleh terkait secara langsung dengan framework seperti UI framework, database framework, atau framework lainnya.

2. **Testability:** Arsitektur ini didesain untuk memudahkan pengujian (testing). Kode bisnis yang independen dari framework akan lebih mudah diuji dengan baik, baik itu pengujian unit, integrasi, atau pengujian fungsional.

3. **Separation of Concerns:** Aplikasi dibagi menjadi beberapa lapisan atau komponen yang berbeda, dengan setiap lapisan memiliki tanggung jawab tertentu. Lapisan-lapisan ini biasanya mencakup Domain, Use Cases, Interface, dan Infrastructure, dan masing-masing memiliki peran dan tanggung jawab yang jelas.

4. **Single Responsibility Principle (SRP):** Prinsip SRP dari SOLID digunakan secara ekstensif dalam Clean Architecture. Setiap komponen atau kelas harus memiliki satu alasan untuk berubah. Ini membantu menjaga kode agar tidak menjadi monolitik dan sulit dipelihara.

5. **Dependency Rule:** Lapisan dalam Clean Architecture harus mengikuti aturan dependensi yang mengalir dari dalam ke luar. Ini berarti bahwa lapisan dalam (misalnya, Domain) tidak boleh bergantung pada lapisan luar (misalnya, UI atau Database).

6. **Entities:** Komponen Domain mengandung entitas yang mewakili objek bisnis inti dalam aplikasi. Ini mencakup aturan bisnis dan logika bisnis yang paling penting.

7. **Use Cases:** Lapisan Use Cases (atau Application Layer) mengandung aturan bisnis yang lebih tingkat dan interaksi antara entitas dan objek domain. Ini adalah tempat di mana operasi-operasi bisnis utama didefinisikan.

8. **Interfaces:** Lapisan Interface menghubungkan aplikasi dengan dunia luar, seperti UI, database, dan sumber daya eksternal lainnya. Ini menyediakan cara bagi aplikasi untuk berkomunikasi dengan komponen eksternal tanpa mencampuradukkan logika bisnis.

9. **Frameworks and Drivers:** Lapisan ini adalah yang terluar dan merupakan tempat di mana implementasi spesifik dari framework dan driver (seperti UI, database, atau API) berada. Lapisan ini harus bergantung pada lapisan-lapisan internal, bukan sebaliknya.

10. **Dependency Inversion Principle (DIP):** Prinsip ini menekankan penggunaan abstraksi dan kontrak (interfaces) untuk mengurangi ketergantungan langsung antara komponen-komponen berbeda, sehingga mengikuti prinsip DIP.

Clean Architecture adalah konsep yang sangat membantu dalam pengembangan perangkat lunak yang bersih, terstruktur, dan mudah dipelihara. Dengan mengikuti prinsip-prinsip ini, pengembang dapat menciptakan aplikasi yang lebih mudah diuji, dioptimalkan, dan diperluas.

# IoTGo-Services

Bu repoda, **IoTGo** uygulamasında kullanılan IoT cihazlarının simülasyonunu sağlayacak servisler bulunmaktadır.  
Gerçek IoT cihazlar yerine, bu servisler test ortamında istekleri işleyerek uygulamanın cihazlarla olan iletişimini doğrulamak için kullanılır.

---

## Amaç

- **IoT Cihaz Simülasyonu:** Gerçek bir IoT cihazı gibi davranarak uygulamanın cihazlara gönderdiği talepleri kabul eder ve yanıtlar.
- **Test Ortamı Desteği:** IoT cihazların henüz mevcut olmadığı veya erişilemediği durumlarda geliştirme ve test süreçlerini kolaylaştırır.
- **Gerçekçi İşlevler:** IoT cihazlarının beklenen davranışlarını simüle ederek gerçek ortamı mümkün olduğunca yansıtır.

---

## Özellikler

1. **Simüle Edilmiş Cihaz İşlevleri:**
    - Cihaz durumunu sorgulama (örneğin açık/kapalı).
    - Cihaz durumunu güncelleme (örneğin açma/kapama).
    - Cihaz ayarlarını değiştirme (örneğin sıcaklık, mod).

2. **RESTful API Desteği:**
    - Simüle edilmiş cihazlarla HTTP üzerinden iletişim kurulur.
    - JSON formatında veri alışverişi sağlanır.

3. **Geliştirici Dostu Yapı:**
    - Loglama ve hata ayıklama araçları ile geliştirme sürecini destekler.
    - Gerçek cihazlarla iletişim protokollerini taklit eder.

4. **Hafif ve Performanslı:**
    - Fiber framework kullanılarak hızlı bir backend sağlanır.
    - Redis ile önbellekleme yapılarak sık erişilen verilere hızlı ulaşım sağlanır.

---

## Kullanılan Teknolojiler

### Fiber
- **HTTP Sunucusu:** RESTful API servislerinin hızlı bir şekilde geliştirilmesi.
- **Middleware:** İsteklerin doğrulanması ve loglanması.

### Redis
- **Önbellekleme:** Cihaz bilgileri gibi sık erişilen verilerin saklanması.

### SQLC ve PostgreSQL
- **Veritabanı İşlemleri:** Cihaz bilgileri ve ayarlarının güvenli bir şekilde saklanması ve sorgulanması.

---

## Proje Yapısı

```plaintext
IoTGo-Services/
│
├── controllers/         # API kontrolörleri
│   ├── device.go        # Cihaz işlevlerini yöneten kontrolör
│   └── user.go          # Kullanıcı işlemlerini yöneten kontrolör
│
├── models/              # Veritabanı modelleri
│   ├── device.go        # Cihaz modeli
│   └── user.go          # Kullanıcı modeli
│
├── services/            # İş mantığı ve servis katmanı
│   ├── device_service.go # Cihazla ilgili iş mantığı
│   └── auth_service.go  # Kimlik doğrulama servisi
│
├── routes/              # API rotaları
│   ├── device_routes.go # Cihazlara ait rotalar
│   └── user_routes.go   # Kullanıcıya ait rotalar
│
├── main.go              # Uygulamanın giriş noktası
└── config/              # Konfigürasyon dosyaları
    └── config.go        # Uygulama ayarları

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
```

# Simüle Edilebilecek Akıllı Cihazlar ve İşlevleri

## Akıllı Kamera
- Hareket algılama (ör. eve birisi yaklaştığında bildirim gönderme).
- Canlı yayın ve kayıt (ör. güvenlik amacıyla).
- Gece görüş modu (ör. düşük ışıkta görüntü sağlama).
- Yüz tanıma (ör. izinli kişilerle yabancıları ayırt etme).

## Akıllı Işık
- Açma/Kapama (ör. uzaktan kontrol veya belirli saatlerde).
- Renk değiştirme (ör. ortam ışığını ayarlama).
- Parlaklık ayarı (ör. enerji tasarrufu veya ambiyans oluşturma).
- Hareket sensörü (ör. birisi odaya girince ışığı otomatik açma).

## Akıllı Termostat
- Sıcaklık kontrolü (ör. odanın sıcaklığını sabit tutma).
- Zamanlama (ör. belirli saatlerde ısıtma/soğutma).
- Enerji verimliliği raporları (ör. ne kadar enerji harcandığını gösterme).
- Uzaktan kontrol (ör. dışarıdayken sıcaklık ayarı yapma).

## Akıllı Priz
- Açma/Kapama kontrolü (ör. bağlı cihazları uzaktan açma/kapatma).
- Enerji tüketimi ölçümü (ör. hangi cihaz ne kadar enerji harcıyor).
- Zamanlayıcı (ör. belirli cihazların belirli saatlerde çalışması).

## Akıllı Kapı Kilidi
- Uzaktan kilitleme/açma (ör. kapıyı telefonla kontrol etme).
- Şifre veya biyometrik doğrulama (ör. parmak iziyle açma).
- Geçmiş kayıtları (ör. kim ne zaman kapıyı açmış).
- Ziyaretçi yönetimi (ör. geçici erişim izni verme).

## Akıllı Perdeler/Jaluziler
- Açma/Kapama (ör. uzaktan kontrol veya ışık durumuna göre otomasyon).
- Zamanlama (ör. sabah otomatik açılma).
- Enerji verimliliği (ör. sıcak havalarda güneş ışığını engelleme).

## Akıllı Buzdolabı
- İçerik takibi (ör. hangi ürünlerin tükenmek üzere olduğunu bildirme).
- Enerji tüketimi optimizasyonu.
- Sıcaklık ve nem kontrolü.


---
# IoT Cihaz Fabrikası - Simülasyon Sunucusu

Bu sistem, simüle edilecek cihazların bir fabrika gibi üretildiği ve yönetildiği temel bir altyapı sunar. Üretilen cihazlar, benzersiz ID'ler ile tanımlanır ve bu ID'ler üzerinden IoTGO uygulamasıyla entegre şekilde çalışır.

IoTGO uygulaması bu sisteme bağlanarak simüle edilen cihazların kontrolünü sağlar. Örneğin, ışık sisteminin açma, kapama ve renk değiştirme özellikleri bu sunucuda tanımlıdır. IoTGO üzerinden bu işlevlere zamanlama eklemek gerekiyorsa, sistem mevcut zaman servisi ile entegre şekilde çalışarak işlemleri gerçekleştirir.

Bu yapı, cihazların temel işlevlerini merkezi bir noktada toplayarak uygulama ile daha kapsamlı bir entegrasyon sağlar.

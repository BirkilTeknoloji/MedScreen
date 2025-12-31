import NfcManager, { NfcEvents } from 'react-native-nfc-manager';
import AsyncStorage from '@react-native-async-storage/async-storage';

export async function sendRfidToBackend(tagId) {
  // IP adresini apiService.js ile aynı tutmaya özen göster
  const url = `http://192.168.1.101:8080/api/v1/nfc-kart/authenticate/${tagId}`;

  try {
    console.log('📡 NFC authenticate isteği:', url);

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
    });

    if (!response.ok) return null;

    const resJson = await response.json();
    const data = resJson.data;
    const token = data?.token;
    const personel = data?.personel;

    if (token && personel) {
      // ÖNEMLİ: apiService.js 'auth_token' beklediği için burayı güncelledik
      await AsyncStorage.setItem('auth_token', token);
      await AsyncStorage.setItem('userInfo', JSON.stringify(personel));

      if (personel.kod) {
        await AsyncStorage.setItem('personel_kod', personel.kod);
      }

      console.log('✅ Giriş başarılı, token kaydedildi.');

      return {
        success: true,
        token: token,
        personel: personel,
      };
    }
    return null;
  } catch (error) {
    console.error('NFC Network hatası:', error.message);
    return null;
  }
}

// BAŞLATMA FONKSİYONU - DÜZELTİLMİŞ
// nfcHandler.js

export async function startNfcReading(handleTagDiscovered, setIsReading) {
  try {
    // 1. Durumu güncelle
    setIsReading(true);

    // 2. Önce donanımı başlatmayı dene (Eğer App.js'de başarısız olduysa burada tekrar dener)
    try {
      await NfcManager.start();
    } catch (e) {
      // "Already started" veya "Activity" hatası gelirse burada yutuyoruz çünkü
      // bazen donanım arka planda hazır olsa da hata dönebilir.
      console.log('NFC Start bypass:', e.message);
    }

    // 3. Mevcut dinleyicileri temizle (Önemli: Çakışmaları önler)
    await NfcManager.unregisterTagEvent().catch(() => {});
    NfcManager.setEventListener(NfcEvents.DiscoverTag, null);

    // 4. Yeni dinleyiciyi bağla
    NfcManager.setEventListener(NfcEvents.DiscoverTag, handleTagDiscovered);

    // 5. Okumayı başlat
    await NfcManager.registerTagEvent();
    console.log('📡 NFC Okuma moduna girildi.');
  } catch (error) {
    setIsReading(false);
    console.error('NFC Kayıt Hatası:', error);
    // Kullanıcıya activity hatası hakkında bilgi verebilirsin
    if (error.toString().includes('current activity')) {
      console.warn(
        'Uygulama henüz hazır değil, lütfen bir saniye sonra tekrar deneyin.',
      );
    }
  }
}
export async function stopNfcReading(setIsReading, isProcessingRef) {
  try {
    setIsReading(false);
    if (isProcessingRef) isProcessingRef.current = false;

    // Event listener'ı kaldır
    NfcManager.setEventListener(NfcEvents.DiscoverTag, null);

    // Kaydı iptal et
    await NfcManager.unregisterTagEvent();
    console.log('🛑 NFC Durduruldu.');
  } catch (error) {
    // Genellikle zaten durmuşsa hata verir, sessizce geçebiliriz
  }
}

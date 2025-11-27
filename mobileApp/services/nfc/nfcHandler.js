import NfcManager, { NfcEvents } from 'react-native-nfc-manager';
import { BASE_API_URL } from '@env';
import AsyncStorage from '@react-native-async-storage/async-storage';

// nfcHandler.js (veya ilgili dosya)

export async function sendRfidToBackend(tagId) {
  const url = `${BASE_API_URL}/nfc-cards/authenticate`;

  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
      body: JSON.stringify({ card_uid: tagId }),
    });

    if (!response.ok) {
      console.error('Backend error:', response.status);
      return null;
    }

    const resJson = await response.json();

    // 🛠️ DÜZELTME BURADA:
    // Token bazen direkt { token: ... } olarak, bazen { data: { token: ... } } olarak gelebilir.
    // İkisini de kontrol ediyoruz:
    const token = resJson.token || resJson.data?.token;
    const user = resJson.user || resJson.data?.user;

    if (token) {
      try {
        console.log('✅ Token bulundu, hafızaya kaydediliyor:', token.substring(0, 10) + '...');
        await AsyncStorage.setItem('userToken', token);
        
        if (user) {
          await AsyncStorage.setItem('userInfo', JSON.stringify(user));
        }
        
        // Fonksiyonun döndürdüğü veriyi standartlaştıralım ki HomeScreen şaşırmasın
        // HomeScreen'e her zaman { token: "...", success: true } dönelim
        return { success: true, token: token, user: user, originalData: resJson };

      } catch (e) {
        console.error('Token kaydetme hatası:', e);
      }
    } else {
        console.warn("⚠️ Yanıtta Token bulunamadı!", resJson);
    }

    return resJson;
  } catch (error) {
    console.error('Network request failed:', error.message);
    return null;
  }
}

// Aşağıdaki fonksiyonlarda bir değişiklik gerekmiyor, standart NFC yönetimi
export async function startNfcReading(handleTagDiscovered, setIsReading) {
  try {
    setIsReading(true);

    await NfcManager.start();
    NfcManager.setEventListener(NfcEvents.DiscoverTag, handleTagDiscovered);
    await NfcManager.registerTagEvent();
  } catch (error) {
    setIsReading(false);
    console.error('NFC başlatma hatası:', error);
  }
}

export async function stopNfcReading(setIsReading, isProcessingRef) {
  try {
    setIsReading(false);
    if (isProcessingRef) isProcessingRef.current = false;
    NfcManager.setEventListener(NfcEvents.DiscoverTag, null);
    await NfcManager.unregisterTagEvent().catch(() => {});
  } catch (error) {
    console.error('NFC durdurma hatası:', error);
  }
}

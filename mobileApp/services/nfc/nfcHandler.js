import NfcManager, { NfcEvents } from 'react-native-nfc-manager';
import { BASE_API_URL } from '@env';
import AsyncStorage from '@react-native-async-storage/async-storage';

// nfcHandler.js (veya ilgili dosya)

export async function sendRfidToBackend(tagId) {
  const url = `${BASE_API_URL}/nfc-cards/authenticate`;

  try {
    console.log('🔍 NFC tag ID received:', tagId);
    console.log('🔍 Tag type:', typeof tagId);
    if (typeof tagId === 'object') {
      console.log('🔍 Tag object:', JSON.stringify(tagId, null, 2));
    }
    console.log('📡 NFC authenticate isteği gönderiliyor:', url);
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
      body: JSON.stringify({ card_uid: tagId }),
    });

    console.log('NFC response status:', response.status);

    if (!response.ok) {
      console.error('NFC Backend error:', response.status);
      const errorText = await response.text();
      console.error('Error response:', errorText);
      return null;
    }

    const resJson = await response.json();
    console.log('NFC response data:', resJson);

    // backendV2 response format: { success: true, code: "SUCCESS_NFC_AUTHENTICATION", message: "...", data: { user: {...}, token: "..." } }
    const data = resJson.data || resJson;
    const token = data.token || resJson.token;
    const user = data.user || resJson.user;

    if (token && user) {
      try {
        console.log('✅ Token bulundu, hafızaya kaydediliyor:', token.substring(0, 10) + '...');
        await AsyncStorage.setItem('userToken', token);
        await AsyncStorage.setItem('userInfo', JSON.stringify(user));
        
        // Device MAC'i AsyncStorage'a kaydet (QR token atama için)
        if (user.id) {
          const deviceMac = await getDeviceMacFromBackend(user.id, token);
          if (deviceMac) {
            await AsyncStorage.setItem('device_mac', deviceMac);
            console.log('Device MAC kaydedildi:', deviceMac);
          }
        }
        
        return { 
          success: true, 
          token: token, 
          user: user,
          originalData: resJson 
        };

      } catch (e) {
        console.error('Token/User kaydetme hatası:', e);
        return null;
      }
    } else {
      console.warn("⚠️ Yanıtta Token veya User bulunamadı!", resJson);
      return null;
    }
  } catch (error) {
    console.error('NFC Network request failed:', error.message);
    return null;
  }
}

// Helper: Device MAC'i backend'den al
async function getDeviceMacFromBackend(userId, token) {
  try {
    // Bu endpoint'i backend'e eklemek gerekebilir
    // Şu an için cihazların genel listesinden ilkini al
    const url = `${BASE_API_URL}/devices`;
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.warn('Device listesi alınamadı');
      return null;
    }

    const resJson = await response.json();
    const devices = resJson.data || resJson;
    
    if (Array.isArray(devices) && devices.length > 0) {
      return devices[0].mac_address;
    }

    return null;
  } catch (error) {
    console.error('Device MAC alınamadı:', error);
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

import NfcManager, { NfcEvents } from 'react-native-nfc-manager';
import { BASE_API_URL } from '@env';

export async function sendRfidToBackend(tagId) {
  const url = `${BASE_API_URL}/nfc-cards/authenticate`;
  
  console.log('BASE_API_URL:', BASE_API_URL);
  console.log('Full URL:', url);
  console.log('Card UID:', tagId);
  console.log('Request headers:', {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  });

  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
      body: JSON.stringify({
        card_uid: tagId,
      }),
    });

    if (!response.ok) {
      console.error('Backend error:', response.status, response.statusText);
      return null;
    }

    const data = await response.json();
    return data;
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

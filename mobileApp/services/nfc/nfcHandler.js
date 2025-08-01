import NfcManager, { NfcEvents } from 'react-native-nfc-manager';
import { BASE_API_URL } from '@env';

export async function sendRfidToBackend(tagId) {
    const url = `${BASE_API_URL}/users/card/${tagId}`;
    console.log('🌐 İstek gönderiliyor:', url);

    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            },
        });

        if (!response.ok) {
            const errorText = await response.text();
            console.log('❌ Backend hatası:', response.status, errorText);
            return null;
        }

        const data = await response.json();
        console.log('✅ Backend yanıtı:', data);
        return data;
    } catch (error) {
        console.log('❌ Fetch hatası:', error);
        return null;
    }
}

export async function startNfcReading(handleTagDiscovered, setIsReading) {
    try {
        setIsReading(true);
        console.log('NFC okuma başlatılıyor...');

        await NfcManager.start();
        NfcManager.setEventListener(NfcEvents.DiscoverTag, handleTagDiscovered);
        await NfcManager.registerTagEvent();

        console.log('NFC okuma aktif, kart bekleniyor...');
    } catch (error) {
        setIsReading(false);
        console.log('NFC başlatma hatası:', error);
    }
}

export async function stopNfcReading(setIsReading, isProcessingRef) {
    try {
        setIsReading(false);
        isProcessingRef.current = false;
        NfcManager.setEventListener(NfcEvents.DiscoverTag, null);
        await NfcManager.unregisterTagEvent().catch(() => {});
        console.log('NFC okuma durduruldu');
    } catch (error) {
        console.log('NFC durdurma hatası:', error);
    }
}

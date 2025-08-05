import NfcManager, { NfcEvents } from 'react-native-nfc-manager';
import { BASE_API_URL } from '@env';

export async function sendRfidToBackend(tagId) {
    const url = `${BASE_API_URL}/users/card/${tagId}`;
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
            console.error('❌ Backend hatası:', response.status, errorText);
            return null;
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('❌ Fetch hatası:', error);
        return null;
    }
}

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
        isProcessingRef.current = false;
        NfcManager.setEventListener(NfcEvents.DiscoverTag, null);
        await NfcManager.unregisterTagEvent().catch(() => {});
    } catch (error) {
        console.error('NFC durdurma hatası:', error);
    }
}

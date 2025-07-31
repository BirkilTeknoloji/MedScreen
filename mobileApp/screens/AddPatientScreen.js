import { useState, useRef, useCallback } from 'react';
import { View, Text, Image, Button } from 'react-native';
import { useNavigation, useFocusEffect } from '@react-navigation/native';
import DeviceInfo from 'react-native-device-info';
import Toast from 'react-native-root-toast';
import { startNfcReading, stopNfcReading, sendRfidToBackend } from '../services/nfc/nfcHandler';
import { addPatient } from '../services/api';
import styles from './styles/AddPatientScreenStyle';

const showToast = (message, backgroundColor) => {
    Toast.show(message, {
        duration: 3000,
        position: 100,
        backgroundColor,
        textColor: '#fff',
    });
};

export default function AddPatientScreen() {
    const navigation = useNavigation();
    const [isReading, setIsReading] = useState(false);
    const isProcessingRef = useRef(false);
    const deviceId = DeviceInfo.getUniqueIdSync();
    const [userData, setUserData] = useState(null);

    const handleTagDiscovered = async (tag) => {
        if (isProcessingRef.current) return;
        isProcessingRef.current = true;

        try {
            const user = await sendRfidToBackend(tag.id);
            if (!user?.ID || !user?.PatientInfo) throw new Error('Kullanıcı bulunamadı');
            setUserData(user);
        } catch (error) {
            console.log('❌ Hata:', error);
        } finally {
            setTimeout(() => {
                isProcessingRef.current = false;
            }, 3000);
        }
    };

    const handleAddPatient = async () => {
        if (!userData?.ID || !userData?.PatientInfo) {
            showToast('Lütfen geçerli bir kart okutun.', '#f57c00');
            return;
        }

        try {
            const result = await addPatient(deviceId, userData.ID);
            showToast('✅ Hasta başarıyla kaydedildi.', '#4caf50');
            navigation.navigate('PatientScreen', { userData: result });
        } catch {
            showToast('❌ Hasta kaydı başarısız oldu.', '#b00020');
        }
    };

    useFocusEffect(
        useCallback(() => {
            startNfcReading(handleTagDiscovered, setIsReading);
            return () => stopNfcReading(setIsReading, isProcessingRef);
        }, [])
    );

    return (
        <View style={styles.container}>
            <Image source={require('../assets/nfc.png')} style={styles.nfcImage} />
            <Text style={styles.infoText}>Kartı okutun ve hasta bilgilerini görüntüleyin</Text>

            <View style={styles.infoBox}>
                <Text style={styles.label}>İsim:</Text>
                <Text style={styles.value}>{userData?.PatientInfo?.Name || '-'}</Text>

                <Text style={styles.label}>Rol:</Text>
                <Text style={styles.value}>{userData?.Role || '-'}</Text>

                <Text style={styles.label}>Cihaz ID:</Text>
                <Text style={styles.value}>{deviceId}</Text>
            </View>

            <Button
                title="➕ Hastayı Kaydet"
                onPress={handleAddPatient}
                disabled={!userData}
            />

            <Text style={styles.statusText}>
                {isReading ? '📱 NFC okuma aktif...' : '❌ NFC okuma durdu'}
            </Text>
        </View>
    );
}
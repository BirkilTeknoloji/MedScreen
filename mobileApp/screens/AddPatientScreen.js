import React, { useState, useRef } from 'react';
import { View, Text, StyleSheet, Image, Button } from 'react-native';
import { useNavigation, useFocusEffect } from '@react-navigation/native';
import { startNfcReading, stopNfcReading, sendRfidToBackend } from '../services/nfc/nfcHandler';
import Toast from 'react-native-root-toast';
import { addPatient } from '../services/api';
import DeviceInfo from 'react-native-device-info';

export default function AddPatientScreen({ route }) {
    const navigation = useNavigation();
    const [isReading, setIsReading] = useState(false);
    const isProcessingRef = useRef(false);
    const [deviceId] = useState(DeviceInfo.getUniqueIdSync());
    const [userData, setUserData] = useState(null);

    const handleTagDiscovered = async (tag) => {
        if (isProcessingRef.current) return;
        isProcessingRef.current = true;

        try {
            const user = await sendRfidToBackend(tag.id);
            if (!user || !user.ID || !user.PatientInfo) throw new Error('Kullanıcı bulunamadı');

            setUserData(user); // name, role, userId gibi bilgileri al
        } catch (error) {
            console.log('❌ Hata:', error);
        } finally {
            setTimeout(() => {
                isProcessingRef.current = false;
            }, 3000);
        }
    };

    const handleAddPatient = async () => {
        if (!userData || !userData.ID || !userData.PatientInfo) {
            Toast.show('Lütfen geçerli bir kart okutun.', {
                duration: 3000,
                position: 100,
                backgroundColor: '#f57c00',
                textColor: '#fff',
            });
            return;
        }

        try {
            const result = await addPatient(deviceId, userData.ID);
            Toast.show('✅ Hasta başarıyla kaydedildi.', {
                duration: 3000,
                position: 100,
                backgroundColor: '#4caf50',
                textColor: '#fff',
            });
            navigation.navigate('PatientScreen', { userData: result });
        } catch (err) {
            Toast.show('❌ Hasta kaydı başarısız oldu.', {
                duration: 3000,
                position: 100,
                backgroundColor: '#b00020',
                textColor: '#fff',
            });
        }
    };

    useFocusEffect(
        React.useCallback(() => {
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

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#f0f4f7',
        padding: 20,
        justifyContent: 'center',
    },
    nfcImage: {
        width: 140,
        height: 140,
        alignSelf: 'center',
        marginBottom: 30,
    },
    infoText: {
        fontSize: 20,
        fontWeight: 'bold',
        textAlign: 'center',
        marginBottom: 20,
    },
    infoBox: {
        backgroundColor: '#fff',
        padding: 15,
        borderRadius: 10,
        marginBottom: 20,
        borderWidth: 1,
        borderColor: '#ccc',
    },
    label: {
        fontWeight: 'bold',
        fontSize: 16,
        marginTop: 10,
    },
    value: {
        fontSize: 16,
        color: '#333',
    },
    statusText: {
        fontSize: 14,
        marginTop: 15,
        textAlign: 'center',
        color: '#666',
    },
});

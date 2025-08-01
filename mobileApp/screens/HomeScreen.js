import React, { useRef, useState } from 'react';
import { View, Text, Image, TouchableOpacity } from 'react-native';
import { useNavigation, useFocusEffect } from '@react-navigation/native';
import { sendRfidToBackend, startNfcReading, stopNfcReading } from '../services/nfc/nfcHandler';
import Toast from 'react-native-root-toast';
import styles from './styles/HomeScreenStyle';
import { addPatient } from '../services/api';

const showSuccessToast = (message) => {
    return Toast.show(message, {
        duration: Toast.durations.LONG,
        position: 100,
        shadow: true,
        animation: true,
        hideOnPress: true,
        delay: 0,
        backgroundColor: '#333',
        textColor: '#fff',
        containerStyle: {
            paddingHorizontal: 24,
            paddingVertical: 18,
            borderRadius: 15,
        },
        textStyle: {
            fontSize: 24,
            fontWeight: 'bold',
            textAlign: 'center',
        },
    });
};

const showErrorToast = (message) => {
    Toast.show(message, {
        duration: Toast.durations.LONG,
        position: 100,
        backgroundColor: '#b00020',
        textColor: '#fff',
        containerStyle: {
            paddingHorizontal: 24,
            paddingVertical: 16,
            borderRadius: 12,
        },
        textStyle: {
            fontSize: 18,
            fontWeight: 'bold',
            textAlign: 'center',
        },
    });
};

export default function HomeScreen() {
    const navigation = useNavigation();
    const [isReading, setIsReading] = useState(false);
    const isProcessingRef = useRef(false);
    const [userData, setUserData] = useState(null);

    const handleTagDiscovered = async (tag) => {
        if (isProcessingRef.current) {
            console.log('Tag işlemi zaten devam ediyor, yeni işlem engellendi');
            return;
        }
        isProcessingRef.current = true;
        try {
            console.log('📦 NFC Tag Okundu:', JSON.stringify(tag));
            const backendResponse = await sendRfidToBackend(tag.id || JSON.stringify(tag.id));

            if (backendResponse?.Role) {
                const toast = showSuccessToast('✅ Giriş başarılı, yönlendiriliyorsunuz...');

                setTimeout(() => {
                    Toast.hide(toast);

                    if (backendResponse.Role === 'patient') {
                        addPatient(backendResponse.ID)
                            .then(() => {
                                navigation.navigate('PatientScreen', { isPatientLogin: false });
                            })
                            .catch(() => {
                                showErrorToast('❌ Hasta otomatik kaydı başarısız.');
                            })
                            .finally(() => {
                                isProcessingRef.current = false;
                            });
                    } else {
                        navigation.navigate('PatientScreen', { isPatientLogin: false });
                        isProcessingRef.current = false;
                    }
                }, 2000);
            } else {
                showErrorToast('❌ Giriş başarısız: Kart tanımlı değil.');
                setTimeout(() => {
                    isProcessingRef.current = false;
                }, 3000);
            }
        } catch (error) {
            console.log('Tag işleme hatası:', error);
            showErrorToast('❌ NFC işleminde hata oluştu.');
            setTimeout(() => {
                isProcessingRef.current = false;
            }, 3000);
        }
    };

    useFocusEffect(
        React.useCallback(() => {
            console.log('HomeScreen focused - NFC başlatılıyor');
            startNfcReading(handleTagDiscovered, setIsReading);
            setUserData(null);
            return () => {
                console.log('HomeScreen unfocused - NFC durduruluyor');
                stopNfcReading(setIsReading, isProcessingRef);
            };
        }, [])
    );

    return (
        <View style={styles.container}>
            <Image source={require('../assets/nfc.png')} style={styles.nfcImage} />
            <Text style={styles.infoText}>
                Giriş için lütfen kartınızı okutunuz <Text style={styles.arrow}>⤴</Text>
            </Text>
            <Text style={styles.statusText}>
                {isReading ? '📱 NFC okuma aktif...' : '❌ NFC okuma durdu'}
            </Text>
        </View>
    );
}
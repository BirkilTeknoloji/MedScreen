import React, { useRef, useState } from 'react';
import { View, Text, Image, TouchableOpacity } from 'react-native';
import { useNavigation, useFocusEffect } from '@react-navigation/native';
import { sendRfidToBackend, startNfcReading, stopNfcReading } from '../services/nfc/nfcHandler';
import Toast from 'react-native-root-toast';
import styles from './styles/HomeScreenStyle';
import { addPatient, getPatientByUserId, getPatientById } from '../services/api';

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
            console.warn('Tag işlemi zaten devam ediyor, yeni işlem engellendi');
            return;
        }
        isProcessingRef.current = true;
        try {
            const backendResponse = await sendRfidToBackend(tag.id || JSON.stringify(tag.id));

            console.log('Backend Response Keys:', Object.keys(backendResponse || {}));
            console.log('Full Backend Response:', JSON.stringify(backendResponse, null, 2));
            
            if (backendResponse?.success && backendResponse?.data) {
                const toast = showSuccessToast('✅ Giriş başarılı, yönlendiriliyorsunuz...');

                setTimeout(async () => {
                    Toast.hide(toast);
                    console.log('Hasta verileri çekiliyor...');
                    
                    try {
                        const userData = backendResponse.data;
                        const userId = userData.ID;
                        
                        navigation.navigate('PatientScreen', { isPatientLogin: false });
                        
                        console.log('Navigation tamamlandı');
                    } catch (navError) {
                        console.error('Navigation hatası:', navError);
                        showErrorToast('❌ Hasta verileri alınamadı.');
                    }
                    
                    isProcessingRef.current = false;
                }, 1000);
            } else {
                showErrorToast('❌ Giriş başarısız: Kart tanımlı değil.');
                setTimeout(() => {
                    isProcessingRef.current = false;
                }, 3000);
            }
        } catch (error) {
            console.error('Tag işleme hatası:', error);
            showErrorToast('❌ NFC işleminde hata oluştu.');
            setTimeout(() => {
                isProcessingRef.current = false;
            }, 3000);
        }
    };

    useFocusEffect(
        React.useCallback(() => {
            startNfcReading(handleTagDiscovered, setIsReading);
            setUserData(null);
            return () => {
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
import React, { useEffect, useRef, useState } from 'react';
import { View, Text, StyleSheet, Image, TouchableOpacity } from 'react-native';
import { useNavigation, useFocusEffect } from '@react-navigation/native';
import { sendRfidToBackend, startNfcReading, stopNfcReading } from '../services/nfc/nfcHandler';
import Toast from 'react-native-root-toast';

export default function HomeScreen() {
    const navigation = useNavigation();
    const [isReading, setIsReading] = useState(false);
    const isProcessingRef = useRef(false);
    const [showOptions, setShowOptions] = useState(false);
    const [userData, setUserData] = useState(null);

    const handleTagDiscovered = async (tag) => {
        if (isProcessingRef.current) {
            console.log('Tag işlemi zaten devam ediyor, yeni işlem engellendi');
            return;
        }

        try {
            isProcessingRef.current = true;
            console.log('📦 NFC Tag Okundu:', JSON.stringify(tag));

            const backendResponse = await sendRfidToBackend(tag.id || JSON.stringify(tag.id));

            if (backendResponse && backendResponse.Role) {
                const toast = Toast.show('✅ Giriş başarılı, yönlendiriliyorsunuz...', {
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

                setTimeout(() => {
                    Toast.hide(toast);
                    if (backendResponse.Role === 'patient') {
                        // Patient ise direkt PatientScreen'e git ama kendi bilgilerini göster
                        navigation.navigate('PatientScreen', { userData: backendResponse, isPatientLogin: true });
                    } else {
                        // Doctor/Nurse ise seçenekleri göster
                        setUserData(backendResponse);
                        setShowOptions(true);
                    }
                }, 2000);

            } else {
                Toast.show('❌ Giriş başarısız: Kart tanımlı değil.', {
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

                // 3 saniye sonra tekrar işleme izin ver
                setTimeout(() => {
                    isProcessingRef.current = false;
                }, 3000);
            }

        } catch (error) {
            console.log('Tag işleme hatası:', error);
            setTimeout(() => {
                isProcessingRef.current = false;
            }, 3000);
        }
    };

    // Screen focus yönetimi
    useFocusEffect(
        React.useCallback(() => {
            console.log('HomeScreen focused - NFC başlatılıyor');
            startNfcReading(handleTagDiscovered, setIsReading);
            
            // Seçenekler ekranından geri dönüldüğünde ana ekrana dön
            setShowOptions(false);
            setUserData(null);

            return () => {
                console.log('HomeScreen unfocused - NFC durduruluyor');
                stopNfcReading(setIsReading, isProcessingRef);
            };
        }, [])
    );

    return (
        <View style={styles.container}>
            {showOptions ? (
                <>
                    <Text style={styles.optionTitle}>Lütfen bir seçenek seçin:</Text>
                    <Text style={styles.welcomeText}>Hoş geldiniz, {userData?.PatientInfo?.Name || userData?.Name}</Text>
                    <View style={styles.cardContainer}>
                        <TouchableOpacity 
                            style={styles.card} 
                            onPress={() => {
                                // Cihaza kayıtlı hasta bilgilerini göster
                                navigation.navigate('PatientScreen');
                            }}
                        >
                            <Text style={styles.cardText}>📋 Hasta Bilgileri</Text>
                            <Text style={styles.cardSubText}>(Bu cihaza kayıtlı hasta)</Text>
                        </TouchableOpacity>
                        <TouchableOpacity 
                            style={styles.card} 
                            onPress={() => {
                                // Hasta ekleme ekranına git
                                navigation.navigate('AddPatientScreen', { userData });
                            }}
                        >
                            <Text style={styles.cardText}>➕ Hasta Ekle</Text>
                            <Text style={styles.cardSubText}>(Bu cihaza hasta kaydet)</Text>
                        </TouchableOpacity>
                    </View>
                </>
            ) : (
                <>
                    <Image source={require('../assets/nfc.png')} style={styles.nfcImage} />
                    <Text style={styles.infoText}>
                        Giriş için lütfen kartınızı okutunuz <Text style={styles.arrow}>⤴</Text>
                    </Text>
                    <Text style={styles.statusText}>
                        {isReading ? '📱 NFC okuma aktif...' : '❌ NFC okuma durdu'}
                    </Text>
                </>
            )}
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#f0f4f7',
        justifyContent: 'center',
        alignItems: 'center',
        paddingHorizontal: 20,
    },
    nfcImage: {
        width: 180,
        height: 180,
        marginBottom: 60,
        resizeMode: 'contain',
    },
    infoText: {
        fontSize: 40,
        fontWeight: 'bold',
        color: '#3370b0ff',
        textAlign: 'center',
        backgroundColor: 'rgba(15, 88, 165, 0.07)',
        paddingHorizontal: 25,
        paddingVertical: 15,
        borderRadius: 15,
        textShadowColor: 'rgba(0,0,0,0.25)',
        textShadowOffset: { width: 1, height: 1 },
        textShadowRadius: 3,
    },
    arrow: {
        fontSize: 45,
        marginLeft: 5,
    },
    statusText: {
        fontSize: 16,
        color: '#666',
        marginTop: 20,
        textAlign: 'center',
    },
    optionTitle: {
        fontSize: 24,
        fontWeight: 'bold',
        marginBottom: 10,
        color: '#333',
        textAlign: 'center',
    },
    welcomeText: {
        fontSize: 18,
        color: '#666',
        marginBottom: 30,
        textAlign: 'center',
    },
    cardContainer: {
        width: '100%',
        justifyContent: 'center',
        alignItems: 'center',
    },
    card: {
        backgroundColor: '#e0ecf8',
        paddingVertical: 20,
        paddingHorizontal: 30,
        borderRadius: 12,
        marginVertical: 10,
        width: '90%',
        alignItems: 'center',
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
        elevation: 3,
    },
    cardText: {
        fontSize: 20,
        color: '#1c4a7e',
        fontWeight: '600',
        marginBottom: 5,
    },
    cardSubText: {
        fontSize: 14,
        color: '#666',
        textAlign: 'center',
    },
});
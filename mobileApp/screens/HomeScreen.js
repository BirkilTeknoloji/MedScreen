import React, { useEffect } from 'react';
import { View, Text, StyleSheet, Image, ToastAndroid, Platform, Alert } from 'react-native';
import NfcManager, { NfcTech } from 'react-native-nfc-manager';
import { useNavigation } from '@react-navigation/native';
import Toast from 'react-native-root-toast';

// const API_URL = "http://192.168.1.110:8080/api/v1/users/card/:card_id";
const API_URL = "http://192.168.1.104:8080/api/v1/users/card/:card_id";

async function sendRfidToBackend(tagId) {
    try {
        const url = API_URL.replace(':card_id', tagId);
        const response = await fetch(url, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });
        if (!response.ok) {
            console.log('Backend bağlantı hatası:', response.status);
            return null;
        }
        const data = await response.json();
        console.log('Backend yanıtı:', data);
        return data;
    } catch (error) {
        console.log('Backend gönderme hatası:', error);
        return null;
    }
}

function startNfc() {
    NfcManager.start();
    console.log('NFC Başlatıldı');
}

async function readNfcTag(navigation) {
    try {
        await NfcManager.requestTechnology([NfcTech.Ndef, NfcTech.NfcA, NfcTech.NfcB, NfcTech.NfcF, NfcTech.NfcV, NfcTech.NdefFormatable, NfcTech.MifareClassic]);
        const tag = await NfcManager.getTag();
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
                    navigation.navigate('PatientScreen', { userData: backendResponse });
                } else {
                    navigation.navigate('PersonnelScreen', { userData: backendResponse });
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

            // Kart tanımsızsa NFC tekrar aktifleştir
            await NfcManager.cancelTechnologyRequest();
            setTimeout(() => readNfcTag(navigation), 3000);
        }

    } catch (ex) {
        console.log('❌ NFC Hatası:', ex?.message || ex?.toString());

        // Hata durumunda NFC teknolojisini bırak
        try {
            await NfcManager.cancelTechnologyRequest();
        } catch (cancelError) {
            console.warn('NFC cancel error:', cancelError);
        }

        // Belirli süre sonra tekrar dene (isteğe bağlı)
        setTimeout(() => readNfcTag(navigation), 3000);
    }
}



export default function HomeScreen() {
    const navigation = useNavigation();

    useEffect(() => {
        startNfc();
        async function initNfcRead() {
            await readNfcTag(navigation);
        }
        initNfcRead();
    }, []);

    return (
        <View style={styles.container}>
            <Image source={require('../assets/nfc.png')} style={styles.nfcImage} />
            <Text style={styles.infoText}>
                Giriş için lütfen kartınızı okutunuz <Text style={styles.arrow}>⤴</Text>
            </Text>
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
});

import React, { useEffect } from 'react';
import { View, Text, StyleSheet, Image, Alert } from 'react-native';
import NfcManager, { NfcTech } from 'react-native-nfc-manager';
import { useNavigation } from '@react-navigation/native';

const API_URL = "http://192.168.1.106:8080/api/v1/users/card/:card_id";

async function sendRfidToBackend(tagId) {
    try {
        const url = API_URL.replace(':card_id', tagId);
        console.log('Backend url:', url);
        const response = await fetch(url, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });
        if (!response.ok) throw new Error('Backend bağlantı hatası: ' + response.status);
        const data = await response.json();
        console.log('Backend yanıtı:', data);
        return data;
    } catch (error) {
        console.error('Backend gönderme hatası:', error);
        Alert.alert('Hata', 'Veri gönderilirken hata oluştu.');
        return null;
    }
}


// NFC başlatma fonksiyonu
function startNfc() {
    NfcManager.start();
    console.log('NFC Başlatıldı');
}

// NFC okuma fonksiyonu
async function readNfcTag(navigation) {
    console.log("NFC dinleniyor... "); 
    try {
        await NfcManager.requestTechnology(NfcTech.Ndef);
        console.log("✅ NFC teknolojisi seçildi"); 
        const tag = await NfcManager.getTag();
        console.log('📦 NFC Tag Okundu:', JSON.stringify(tag));
        console.log('🆔 NFC Tag ID:', tag.id);

        const backendResponse = await sendRfidToBackend(tag.id || JSON.stringify(tag.id));
        if (backendResponse) {
            Alert.alert('Başarılı', 'Giriş başarılı!');
        }

        if (backendResponse.Role === "patient") {
            navigation.navigate('PatientScreen', { userData: backendResponse });
        } else {
            navigation.navigate('PersonnelScreen', { userData: backendResponse });
        }

    } catch (ex) {
        console.error('❌ NFC Hatası:', ex);
        Alert.alert('Hata', ex.toString());
    } finally {
        NfcManager.cancelTechnologyRequest();
    }
}


export default function HomeScreen() {
    const navigation = useNavigation();

    useEffect(() => {
        startNfc();

        // Async fonksiyonu useEffect içinde tanımla ve çağır
        async function initNfcRead() {
            console.log('NFC Okuma Başlatılıyor...');
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

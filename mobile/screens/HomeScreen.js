import React, { useEffect } from 'react';
import { View, Text, StyleSheet, Image, Alert } from 'react-native';
import NfcManager, { NfcTech } from 'react-native-nfc-manager';

// NFC başlatma fonksiyonu
function startNfc() {
  NfcManager.start();
  console.log('NFC Başlatıldı');
}

// NFC okuma fonksiyonu
async function readNfcTag() {
  try {
    await NfcManager.requestTechnology(NfcTech.Ndef);
    const tag = await NfcManager.getTag();
    console.log('NFC Tag Okundu:', JSON.stringify(tag));
    Alert.alert('Kart Okundu!', JSON.stringify(tag));
  } catch (ex) {
    Alert.alert('Hata', ex.toString());
  } finally {
    NfcManager.cancelTechnologyRequest();
  }
}

export default function HomeScreen() {
  useEffect(() => {
    startNfc();

    // Async fonksiyonu useEffect içinde tanımla ve çağır
    async function initNfcRead() {
      await readNfcTag();
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

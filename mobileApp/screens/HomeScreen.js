import React, { useRef, useState } from 'react';
import {
  View,
  Text,
  Image,
  TouchableOpacity,
  Modal,
  Alert,
} from 'react-native';
import { useNavigation, useFocusEffect } from '@react-navigation/native';
import {
  Camera,
  useCameraDevice,
  useCodeScanner,
} from 'react-native-vision-camera';
import {
  sendRfidToBackend,
  startNfcReading,
  stopNfcReading,
} from '../services/nfc/nfcHandler';
import { parseQrCode } from '../services/api';
import Toast from 'react-native-root-toast';
import styles from './styles/HomeScreenStyle';

// --- Toast Yardımcı Fonksiyonları ---
const showSuccessToast = message => {
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

const showErrorToast = message => {
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
  const [qrModalVisible, setQrModalVisible] = useState(false);
  const [qrScanning, setQrScanning] = useState(false);
  const [qrCameraPermission, setQrCameraPermission] = useState(false);
  const isProcessingRef = useRef(false);
  const [userData, setUserData] = useState(null);

  // Try back camera first, fallback to front if not available
  let device = useCameraDevice('back');
  if (!device) {
    device = useCameraDevice('front');
  }

  const handleTagDiscovered = async tag => {
    if (isProcessingRef.current) return;
    isProcessingRef.current = true;

    try {
      // tag.id bazen obje (Android) bazen string dönebilir, garantiye alıyoruz
      const tagId =
        typeof tag.id === 'object'
          ? tag.id.id || JSON.stringify(tag.id)
          : tag.id;

      const backendResponse = await sendRfidToBackend(tagId);

      if (backendResponse && backendResponse.success) {
        const personel = backendResponse.personel || backendResponse.user;

        // API'den gelen alan adlarını normalize ediyoruz
        const firstName = personel.AD || personel.ad || personel.first_name;
        const lastName =
          personel.SOYADI || personel.soyadi || personel.last_name;
        const role = personel.GOREV || personel.gorev || personel.role;
        const pKod = personel.KODU || personel.kod || personel.personel_kod;

        const toast = showSuccessToast(
          `✅ Hoş geldiniz ${firstName} ${lastName}!`,
        );

        setTimeout(() => {
          Toast.hide(toast);

          // PatientScreen'e yönlendirme
          navigation.navigate('PatientScreen', {
            isPatientLogin: false,
            doctorData: {
              ...personel,
              AD: firstName,
              SOYADI: lastName,
              GOREV: role,
              KODU: pKod,
            },
            // Eğer QR/NFC ile bir hasta bilgisi de geldiyse (NFC karta hasta atanmışsa)
            isQrNavigation: true,
            qrTokenType: 'NFC_CARD',
            qrTokenData: { patient_id: 'H0001' }, // Burada normalde result'dan gelen hasta kodu olmalı
          });
        }, 1500);
      } else {
        showErrorToast('❌ Giriş başarısız: Kart tanımlı değil.');
        isProcessingRef.current = false;
      }
    } catch (error) {
      showErrorToast('❌ NFC Hatası: ' + error.message);
      isProcessingRef.current = false;
    }
  };

  // --- NFC Başlatma / Durdurma ---
  useFocusEffect(
    React.useCallback(() => {
      startNfcReading(handleTagDiscovered, setIsReading);
      setUserData(null);

      return () => {
        stopNfcReading(setIsReading, isProcessingRef);
      };
    }, []),
  );

  // --- QR Scanner Handler ---
  const handleQrScan = async codes => {
    if (codes.length > 0 && qrScanning) {
      const qrValue = codes[0].value;
      setQrScanning(false);
      console.log('QR scanned from HomeScreen:', qrValue);

      try {
        const AsyncStorage =
          require('@react-native-async-storage/async-storage').default;
        const userToken = await AsyncStorage.getItem('userToken');

        // YENİ: QR parse işlemi - backend'inizde varsa kullanın
        // Yoksa bu kısmı backend'inize göre düzenlemeniz gerekebilir
        const parseResult = await parseQrCode(qrValue, userToken);

        if (!parseResult) {
          Alert.alert('Hata', 'QR token doğrulanamadı.');
          setQrScanning(true);
          return;
        }

        // Handle token_used
        if (parseResult.type === 'token_used') {
          Alert.alert('Geçersiz QR', 'Bu QR daha önce kullanılmış.');
          setQrScanning(true);
          return;
        }

        // Handle prescription_info token
        if (
          parseResult.tokenType === 'prescription_info' &&
          parseResult.data?.patient_id
        ) {
          console.log(
            'Prescription info token detected, navigating to PatientScreen with QR data',
          );
          setQrModalVisible(false);
          navigation.navigate('PatientScreen', {
            qrTokenData: parseResult.data,
            qrTokenType: 'prescription_info',
            isQrNavigation: true,
          });
          return;
        }

        // Handle patient_assignment token
        if (
          parseResult.tokenType === 'patient_assignment' &&
          parseResult.data?.patient_id
        ) {
          console.log(
            'Patient assignment token detected, navigating to PatientScreen with QR data',
          );
          setQrModalVisible(false);
          navigation.navigate('PatientScreen', {
            qrTokenData: parseResult.data,
            qrTokenType: 'patient_assignment',
            isQrNavigation: true,
          });
          return;
        }

        // Generic token_validated
        if (
          parseResult.type === 'token_validated' &&
          parseResult.data?.patient_id
        ) {
          setQrModalVisible(false);
          navigation.navigate('PatientScreen', {
            qrTokenData: parseResult.data,
            qrTokenType: parseResult.tokenType,
            isQrNavigation: true,
          });
          return;
        }

        Alert.alert(
          'Bilgi',
          'QR token başarıyla doğrulandı ancak hasta verisi bulunamadı.',
        );
        setQrScanning(true);
      } catch (error) {
        console.error('QR scan error:', error);
        Alert.alert('Hata', 'QR okuma hatası: ' + error.message);
        setQrScanning(true);
      }
    }
  };

  const codeScanner = useCodeScanner({
    codeTypes: ['qr', 'ean-13'],
    onCodeScanned: handleQrScan,
  });

  // Request camera permission when QR modal opens
  const openQrScanner = async () => {
    try {
      const permission = await Camera.requestCameraPermission();
      if (permission === 'granted') {
        setQrCameraPermission(true);
        setQrScanning(true);
        setQrModalVisible(true);
      } else {
        Alert.alert(
          'Hata',
          'Kamera izni gerekli. Lütfen uygulama ayarlarından izin verin.',
        );
      }
    } catch (error) {
      console.error('Camera permission error:', error);
      Alert.alert('Hata', 'Kamera izni alınamadı.');
    }
  };

  return (
    <View style={styles.container}>
      <Image source={require('../assets/nfc.png')} style={styles.nfcImage} />
      <Text style={styles.infoText}>
        Giriş için lütfen kartınızı okutunuz <Text style={styles.arrow}>⤴</Text>
      </Text>
      <Text style={styles.statusText}>
        {isReading ? '📱 NFC okuma aktif...' : '❌ NFC okuma durdu'}
      </Text>

      {/* QR Scanner Button */}
      <TouchableOpacity
        style={{
          marginTop: 20,
          paddingHorizontal: 16,
          paddingVertical: 12,
          backgroundColor: '#2196F3',
          borderRadius: 8,
          alignSelf: 'center',
        }}
        onPress={openQrScanner}
      >
        <Text
          style={{
            color: 'white',
            fontSize: 16,
            fontWeight: 'bold',
            textAlign: 'center',
          }}
        >
          📱 QR Kod Tara
        </Text>
      </TouchableOpacity>

      {/* QR Scanner Modal */}
      <Modal
        visible={qrModalVisible}
        transparent={true}
        animationType="slide"
        onRequestClose={() => {
          setQrModalVisible(false);
          setQrScanning(false);
        }}
      >
        <View style={{ flex: 1, backgroundColor: '#000' }}>
          {device && qrCameraPermission ? (
            <Camera
              style={{ flex: 1 }}
              device={device}
              isActive={qrScanning}
              codeScanner={qrScanning ? codeScanner : undefined}
            />
          ) : (
            <View
              style={{
                flex: 1,
                justifyContent: 'center',
                alignItems: 'center',
              }}
            >
              <Text
                style={{ color: '#fff', fontSize: 16, textAlign: 'center' }}
              >
                {!device ? 'Kamera bulunamadı' : 'İzin bekleniyor...'}
              </Text>
            </View>
          )}

          {/* Close Button */}
          <TouchableOpacity
            style={{
              position: 'absolute',
              bottom: 30,
              left: 20,
              right: 20,
              paddingHorizontal: 16,
              paddingVertical: 12,
              backgroundColor: '#f44336',
              borderRadius: 8,
            }}
            onPress={() => {
              setQrModalVisible(false);
              setQrScanning(false);
            }}
          >
            <Text
              style={{
                color: 'white',
                fontSize: 16,
                fontWeight: 'bold',
                textAlign: 'center',
              }}
            >
              Kapat
            </Text>
          </TouchableOpacity>
        </View>
      </Modal>
    </View>
  );
}

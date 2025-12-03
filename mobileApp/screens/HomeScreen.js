import React, { useRef, useState } from 'react';
import { View, Text, Image, TouchableOpacity, Modal, Alert } from 'react-native';
import { useNavigation, useFocusEffect } from '@react-navigation/native';
import { Camera, useCameraDevice, useCodeScanner } from 'react-native-vision-camera';
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

  const handleTagDiscovered = async (tag) => {
    if (isProcessingRef.current) {
      console.warn('Tag işlemi zaten devam ediyor...');
      return;
    }
    isProcessingRef.current = true;

    try {
      console.log("🔐 Kart okundu, backend'e soruluyor...");
      const backendResponse = await sendRfidToBackend(tag.id || JSON.stringify(tag.id));

      console.log('HomeScreen NFC Cevabı:', backendResponse);

      if (backendResponse && backendResponse.success && backendResponse.token && backendResponse.user) {
        const toast = showSuccessToast('✅ Giriş başarılı, yönlendiriliyorsunuz...');

        setTimeout(() => {
          Toast.hide(toast);
          
          try {
            const user = backendResponse.user;
            console.log('Kullanıcı login başarılı:', user.first_name, user.last_name);

            navigation.navigate('PatientScreen', { 
              isPatientLogin: false,
              doctorData: user 
            });
            console.log('✅ Navigation başarılı');

          } catch (navError) {
            console.error('❌ Navigation hatası:', navError);
            showErrorToast('❌ Sayfa yönlendirme hatası.');
            isProcessingRef.current = false;
          }
        }, 1000);

      } else {
        console.warn("❌ Giriş Başarısız. Backend Cevabı:", backendResponse);
        showErrorToast('❌ Giriş başarısız: Kart tanımlı değil veya sistem hatası.');
        
        setTimeout(() => {
          isProcessingRef.current = false;
        }, 3000);
      }

    } catch (error) {
      console.error('❌ Tag işleme hatası:', error);
      showErrorToast('❌ NFC işleminde hata oluştu: ' + error.message);
      setTimeout(() => {
        isProcessingRef.current = false;
      }, 3000);
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
  const handleQrScan = async (codes) => {
    if (codes.length > 0 && qrScanning) {
      const qrValue = codes[0].value;
      setQrScanning(false);
      console.log('QR scanned from HomeScreen:', qrValue);

      try {
        const userToken = require('@react-native-async-storage/async-storage').default.getItem('userToken');
        const parseResult = await parseQrCode(qrValue);

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
        if (parseResult.tokenType === 'prescription_info' && parseResult.data?.patient_id) {
          console.log('Prescription info token detected, navigating to PatientScreen with QR data');
          setQrModalVisible(false);
          navigation.navigate('PatientScreen', {
            qrTokenData: parseResult.data,
            qrTokenType: 'prescription_info',
            isQrNavigation: true,
          });
          return;
        }

        // Handle patient_assignment token
        if (parseResult.tokenType === 'patient_assignment' && parseResult.data?.patient_id) {
          console.log('Patient assignment token detected, navigating to PatientScreen with QR data');
          setQrModalVisible(false);
          navigation.navigate('PatientScreen', {
            qrTokenData: parseResult.data,
            qrTokenType: 'patient_assignment',
            isQrNavigation: true,
          });
          return;
        }

        // Generic token_validated
        if (parseResult.type === 'token_validated' && parseResult.data?.patient_id) {
          setQrModalVisible(false);
          navigation.navigate('PatientScreen', {
            qrTokenData: parseResult.data,
            qrTokenType: parseResult.tokenType,
            isQrNavigation: true,
          });
          return;
        }

        Alert.alert('Bilgi', 'QR token başarıyla doğrulandı ancak hasta verisi bulunamadı.');
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
        Alert.alert('Hata', 'Kamera izni gerekli. Lütfen uygulama ayarlarından izin verin.');
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
          alignSelf: 'center'
        }}
        onPress={openQrScanner}
      >
        <Text style={{ color: 'white', fontSize: 16, fontWeight: 'bold', textAlign: 'center' }}>
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
            <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
              <Text style={{ color: '#fff', fontSize: 16, textAlign: 'center' }}>
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
            <Text style={{ color: 'white', fontSize: 16, fontWeight: 'bold', textAlign: 'center' }}>
              Kapat
            </Text>
          </TouchableOpacity>
        </View>
      </Modal>
    </View>
  );
}
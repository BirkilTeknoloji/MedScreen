import { useState, useEffect } from 'react';
import { View, Text, Alert, ActivityIndicator, TouchableOpacity } from 'react-native';
import { Camera, useCameraDevice, useCodeScanner } from 'react-native-vision-camera';
import styles from './styles/QrScannerScreenStyle';
import { BASE_API_URL } from '@env';

export default function QrScannerScreen({ navigation }) {
  const [hasPermission, setHasPermission] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [isScanning, setIsScanning] = useState(true);
  const [scannedData, setScannedData] = useState(null);
  const [patientInfo, setPatientInfo] = useState(null);
  const [fetchingInfo, setFetchingInfo] = useState(false);
  const device = useCameraDevice('front');

  useEffect(() => {
    (async () => {
      try {
        setIsLoading(true);
        const permission = await Camera.requestCameraPermission();
        setHasPermission(permission === 'granted');
      } catch (error) {
        console.error('Permission request error:', error);
        setHasPermission(false);
      } finally {
        setIsLoading(false);
      }
    })();
  }, []);

  // API çağrısı
  const fetchPatientInfo = async ({ id, field, itemId }) => {
    try {
      setFetchingInfo(true);
      const response = await fetch(`${BASE_API_URL}/users/${id}/patientinfo/${field}/${itemId}`);
      console.log('Hasta bilgisi alınıyor:', response.url);
      if (!response.ok) {
        throw new Error(`API error: ${response.status}`);
      }
      const data = await response.json();
      setPatientInfo(data);
    } catch (error) {
      console.error('Hasta bilgisi alınamadı:', error);
      Alert.alert('Hata', 'Hasta bilgisi alınamadı. Lütfen tekrar deneyin.');
      setPatientInfo(null);
      // İstersen yeniden taramaya izin ver:
      setIsScanning(true);
      setScannedData(null);
    } finally {
      setFetchingInfo(false);
    }
  };

  // QR kod okunduğunda otomatik çalışan fonksiyon
  const handleScan = async (codes) => {
    if (codes.length > 0 && isScanning) {
      const value = codes[0].value;
      setIsScanning(false);
      setScannedData({
        content: value,
        type: typeof value,
        length: value.length,
      });
      setPatientInfo(null);

      try {
        const parsed = JSON.parse(value.trim());
        if (parsed.id && parsed.field && parsed.itemId) {
          await fetchPatientInfo(parsed);
        } else {
          Alert.alert(
            'Eksik Veri',
            'QR kod içeriği gerekli alanları içermiyor: id, field, itemId'
          );
          setIsScanning(true);
          setScannedData(null);
        }
      } catch (err) {
        Alert.alert(
          'JSON Parse Hatası',
          'QR içeriği JSON formatında değil veya hatalı.'
        );
        setIsScanning(true);
        setScannedData(null);
      }
    }
  };

  const codeScanner = useCodeScanner({
    codeTypes: ['qr', 'ean-13'],
    onCodeScanned: handleScan,
  });

  const handleRescan = () => {
    setScannedData(null);
    setPatientInfo(null);
    setIsScanning(true);
  };

  if (isLoading) {
    return (
      <View style={styles.center}>
        <ActivityIndicator size="large" />
        <Text>İzinler kontrol ediliyor...</Text>
      </View>
    );
  }

  if (!hasPermission) {
    return (
      <View style={styles.center}>
        <Text style={styles.errorText}>Kamera izni gerekli</Text>
        <Text style={styles.infoText}>
          Lütfen uygulama ayarlarından kamera iznini etkinleştirin
        </Text>
      </View>
    );
  }

  if (!device) {
    return (
      <View style={styles.center}>
        <Text style={styles.errorText}>Kamera bulunamadı</Text>
        <Text style={styles.infoText}>Cihazınızda kullanılabilir kamera yok</Text>
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <View style={styles.background}>
        <Text style={styles.title}>QR Kod Tarayıcı</Text>
        <View style={styles.cameraContainer}>
          <Camera
            style={styles.camera}
            device={device}
            isActive={isScanning}
            codeScanner={isScanning ? codeScanner : undefined}
            photo={true}
            video={false}
          />
          <View style={styles.cameraFrame} />
        </View>
        {isScanning && (
          <View style={styles.infoContainer}>
            <Text style={styles.instructionText}>
              QR kodu kamera görüş alanına yerleştirin
            </Text>
            <Text style={styles.cameraInfoText}>
              📱 Kamera: {device.position === 'front' ? 'Ön' : 'Arka'}
            </Text>
          </View>
        )}
      </View>

      {fetchingInfo && (
        <View style={{ padding: 10 }}>
          <ActivityIndicator size="large" />
          <Text>Hasta bilgisi yükleniyor...</Text>
        </View>
      )}

      {patientInfo && (
        <View style={styles.resultOverlay}>
          <View style={styles.resultContainer}>
            <Text style={styles.resultTitle}>{patientInfo.title}</Text>

            <View style={{ marginTop: 16 }}>                
              <Text style={{ fontSize: 16, fontWeight: 'bold', color: '#333' }}>
                Sonuç:
              </Text>
              <Text style={{
                fontSize: 18,
                marginBottom: 12,
                color: patientInfo.result === 'normal' ? 'green' : 'red',
                fontWeight: '600'
              }}>
                {patientInfo.result}
              </Text>
            </View>

            <TouchableOpacity
              style={[styles.button, styles.rescanButton]}
              onPress={handleRescan}
              disabled={fetchingInfo}
            >
              <Text style={styles.buttonText}>Yeniden Tara</Text>
            </TouchableOpacity>
            <TouchableOpacity
              style={[styles.button, styles.homeButton]}
              onPress={() => navigation.navigate('PatientScreen')}
              disabled={fetchingInfo}
            >
              <Text style={styles.buttonText}>Ana Sayfaya Dön</Text>
            </TouchableOpacity>
          </View>
        </View>
      )}

      {/* Eğer istersen, hata veya eksik veri durumunda yeniden tarama otomatik aktif olacak */}
    </View>
  );
}

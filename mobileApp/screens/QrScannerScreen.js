import { useState, useEffect } from 'react';
import { View, Text, Alert, TouchableOpacity } from 'react-native';
import { Camera, useCameraDevice, useCodeScanner } from 'react-native-vision-camera';
import styles from './styles/QrScannerScreenStyle';

export default function QrScannerScreen({ navigation }) {
  const [hasPermission, setHasPermission] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [isScanning, setIsScanning] = useState(true);
  const [scannedData, setScannedData] = useState(null);

  const device = useCameraDevice('front');

  useEffect(() => {
    (async () => {
      try {
        const devices = await Camera.getAvailableCameraDevices();
        console.log('Available Devices:', JSON.stringify(devices, null, 2));
      } catch (error) {
        console.error('Error getting camera devices:', error);
      }
    })();
  }, []);

  useEffect(() => {
    (async () => {
      try {
        setIsLoading(true);
        const permission = await Camera.requestCameraPermission();
        setHasPermission(permission === 'granted');
        console.log('Camera permission:', permission);
      } catch (error) {
        console.error('Permission request error:', error);
        setHasPermission(false);
      } finally {
        setIsLoading(false);
      }
    })();
  }, []);

  const handleScan = (codes) => {
    if (codes.length > 0 && isScanning) {
      const value = codes[0].value;
      setIsScanning(false);
      setScannedData({
        content: value,
        type: typeof value,
        length: value.length
      });
      console.log('Scanned QR Code:', value);
    }
  };

  const codeScanner = useCodeScanner({
    codeTypes: ['qr', 'ean-13'],
    onCodeScanned: handleScan,
  });

  const handleRescan = () => {
    setScannedData(null);
    setIsScanning(true);
  };

  const tryParseAsJSON = () => {
    try {
      const trimmedValue = scannedData.content.trim();
      const parsed = JSON.parse(trimmedValue);
      if (parsed.role === 'patient') {
        navigation.replace('PatientScreen', { userData: parsed });
      } else {
        Alert.alert('Bilinmeyen Rol', `Rol: ${parsed.role || 'Tanımlanmadı'}`);
      }
    } catch (err) {
      Alert.alert(
        'JSON Parse Hatası',
        `Hata: ${err.message}\n\nQR içeriği JSON formatında değil.`
      );
      console.error('JSON Parse Error:', err);
    }
  };

  const handleDirectValue = () => {
    const value = scannedData.content;
    if (value.toLowerCase().includes('patient')) {
      navigation.replace('PatientScreen', {
        userData: {
          role: 'patient',
          rawData: value,
          id: Date.now().toString()
        }
      });
    } else {
      Alert.alert('Bilinmeyen Format', 'QR kod formatı tanınamadı.');
    }
  };

  if (isLoading) {
    return (
      <View style={styles.center}>
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
        <Text style={styles.infoText}>
          Cihazınızda kullanılabilir kamera yok
        </Text>
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
      {scannedData && (
        <View style={styles.resultOverlay}>
          <View style={styles.resultContainer}>
            <Text style={styles.resultTitle}>QR Kod Okundu</Text>
            <View style={styles.dataContainer}>
              <Text style={styles.dataLabel}>İçerik:</Text>
              <Text style={styles.dataContent} numberOfLines={5}>
                {scannedData.content}
              </Text>
              <Text style={styles.dataLabel}>Detaylar:</Text>
              <Text style={styles.dataDetails}>
                Tip: {scannedData.type} | Uzunluk: {scannedData.length}
              </Text>
            </View>
            <View style={styles.buttonContainer}>
              <Text style={styles.buttonLabel}>Bu veriyi nasıl işlemek istiyorsunuz?</Text>
              <View style={styles.buttonRow}>
                <TouchableOpacity
                  style={[styles.button, styles.jsonButton]}
                  onPress={tryParseAsJSON}
                >
                  <Text style={styles.buttonText}>JSON Olarak Dene</Text>
                </TouchableOpacity>
                <TouchableOpacity
                  style={[styles.button, styles.directButton]}
                  onPress={handleDirectValue}
                >
                  <Text style={styles.buttonText}>Doğrudan Kullan</Text>
                </TouchableOpacity>
              </View>
              <TouchableOpacity
                style={[styles.button, styles.rescanButton]}
                onPress={handleRescan}
              >
                <Text style={styles.buttonText}>Yeniden Tara</Text>
              </TouchableOpacity>
            </View>
          </View>
        </View>
      )}
    </View>
  );
}

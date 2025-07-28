// QrScannerScreen.js
import React, { useState, useEffect } from 'react';
import { View, Text, StyleSheet, Alert } from 'react-native';
import { Camera, useCameraDevice, useCodeScanner } from 'react-native-vision-camera';

export default function QrScannerScreen({ navigation }) {
  const [hasPermission, setHasPermission] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [isScanning, setIsScanning] = useState(true);
  const [scannedData, setScannedData] = useState(null);

  const device = useCameraDevice('front');

  useEffect(() => {
    const checkDevices = async () => {
      try {
        const devices = await Camera.getAvailableCameraDevices();
        console.log('Available Devices:', JSON.stringify(devices, null, 2));
      } catch (error) {
        console.error('Error getting camera devices:', error);
      }
    };
    checkDevices();
  }, []);

  const codeScanner = useCodeScanner({
    codeTypes: ['qr', 'ean-13'],
    onCodeScanned: (codes) => {
      if (codes.length > 0 && isScanning) {
        const value = codes[0].value;
        console.log('Scanned QR Code:', value);
        
        // Taramayı durdur
        setIsScanning(false);
        
        // Scanned data'yı kaydet
        setScannedData({
          content: value,
          type: typeof value,
          length: value.length
        });
      }
    },
  });

  const handleRescan = () => {
    setScannedData(null);
    setIsScanning(true);
  };

  const tryParseAsJSON = () => {
    try {
      const trimmedValue = scannedData.content.trim();
      console.log('Trimmed value:', trimmedValue);
      
      const parsed = JSON.parse(trimmedValue);
      console.log('Parsed JSON:', parsed);
      
      if (parsed.role === 'patient') {
        navigation.replace('PatientScreen', { userData: parsed });
      } else {
        Alert.alert('Bilinmeyen Rol', `Rol: ${parsed.role || 'Tanımlanmadı'}`);
      }
    } catch (err) {
      console.error('JSON Parse Error:', err);
      Alert.alert(
        'JSON Parse Hatası', 
        `Hata: ${err.message}\n\nQR içeriği JSON formatında değil.`
      );
    }
  };

  const handleDirectValue = () => {
    const value = scannedData.content;
    if (value.toLowerCase().includes('patient')) {
      const userData = {
        role: 'patient',
        rawData: value,
        id: Date.now().toString()
      };
      navigation.replace('PatientScreen', { userData });
    } else {
      Alert.alert('Bilinmeyen Format', 'QR kod formatı tanınamadı.');
    }
  };

  useEffect(() => {
    const requestPermission = async () => {
      try {
        setIsLoading(true);
        const permission = await Camera.requestCameraPermission();
        console.log('Camera permission:', permission);
        setHasPermission(permission === 'granted');
      } catch (error) {
        console.error('Permission request error:', error);
        setHasPermission(false);
      } finally {
        setIsLoading(false);
      }
    };

    requestPermission();
  }, []);

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

  if (device == null) {
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
      {/* Ana arka plan */}
      <View style={styles.background}>
        <Text style={styles.title}>QR Kod Tarayıcı</Text>
        
        {/* Kamera container */}
        <View style={styles.cameraContainer}>
          <Camera
            style={styles.camera}
            device={device}
            isActive={isScanning}
            codeScanner={isScanning ? codeScanner : undefined}
            photo={true}
            video={false}
          />
          
          {/* Kamera çerçevesi */}
          <View style={styles.cameraFrame} />
        </View>
        
        {/* Bilgi metni */}
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

      {/* Scanned Data Display */}
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
                <Text 
                  style={[styles.button, styles.jsonButton]}
                  onPress={tryParseAsJSON}
                >
                  JSON Olarak Dene
                </Text>
                
                <Text 
                  style={[styles.button, styles.directButton]}
                  onPress={handleDirectValue}
                >
                  Doğrudan Kullan
                </Text>
              </View>
              
              <Text 
                style={[styles.button, styles.rescanButton]}
                onPress={handleRescan}
              >
                Yeniden Tara
              </Text>
            </View>
          </View>
        </View>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: { 
    flex: 1,
  },
  background: {
    flex: 1,
    backgroundColor: '#f0f0f0',
    alignItems: 'center',
    paddingTop: 60,
    paddingHorizontal: 20,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#333',
    marginBottom: 30,
    textAlign: 'center',
  },
  cameraContainer: {
    width: 400,
    height: 400,
    borderRadius: 20,
    overflow: 'hidden',
    elevation: 8,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 4 },
    shadowOpacity: 0.3,
    shadowRadius: 8,
    backgroundColor: '#000',
    position: 'relative',
  },
  camera: {
    width: '100%',
    height: '100%',
  },
  cameraFrame: {
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    borderWidth: 3,
    borderColor: '#4CAF50',
    borderRadius: 20,
    backgroundColor: 'transparent',
  },
  infoContainer: {
    marginTop: 30,
    alignItems: 'center',
    paddingHorizontal: 20,
  },
  instructionText: {
    fontSize: 16,
    color: '#333',
    textAlign: 'center',
    marginBottom: 10,
    lineHeight: 22,
  },
  cameraInfoText: {
    fontSize: 14,
    color: '#666',
    textAlign: 'center',
  },
  center: { 
    flex: 1, 
    justifyContent: 'center', 
    alignItems: 'center',
    padding: 20 
  },
  errorText: {
    fontSize: 18,
    color: '#ff4444',
    marginBottom: 10,
    textAlign: 'center',
  },
  infoText: {
    fontSize: 14,
    color: '#666',
    textAlign: 'center',
    lineHeight: 20,
  },
  // Sonuç overlay stilleri
  resultOverlay: {
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    backgroundColor: 'rgba(0,0,0,0.8)',
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  resultContainer: {
    backgroundColor: 'white',
    borderRadius: 12,
    padding: 20,
    width: '100%',
    maxWidth: 400,
    maxHeight: '80%',
  },
  resultTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    textAlign: 'center',
    marginBottom: 20,
    color: '#333',
  },
  dataContainer: {
    marginBottom: 20,
  },
  dataLabel: {
    fontSize: 14,
    fontWeight: 'bold',
    color: '#555',
    marginTop: 10,
    marginBottom: 5,
  },
  dataContent: {
    fontSize: 13,
    color: '#333',
    backgroundColor: '#f5f5f5',
    padding: 10,
    borderRadius: 6,
    fontFamily: 'monospace',
  },
  dataDetails: {
    fontSize: 12,
    color: '#777',
    fontStyle: 'italic',
  },
  buttonContainer: {
    alignItems: 'center',
  },
  buttonLabel: {
    fontSize: 14,
    color: '#555',
    textAlign: 'center',
    marginBottom: 15,
  },
  buttonRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    width: '100%',
    marginBottom: 10,
  },
  button: {
    paddingVertical: 12,
    paddingHorizontal: 20,
    borderRadius: 8,
    textAlign: 'center',
    fontSize: 14,
    fontWeight: 'bold',
  },
  jsonButton: {
    backgroundColor: '#4CAF50',
    color: 'white',
    flex: 0.48,
  },
  directButton: {
    backgroundColor: '#2196F3',
    color: 'white',
    flex: 0.48,
  },
  rescanButton: {
    backgroundColor: '#FF9800',
    color: 'white',
    width: '100%',
    marginTop: 10,
  },
});
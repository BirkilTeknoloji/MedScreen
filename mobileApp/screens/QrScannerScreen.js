import { useState, useEffect } from 'react';
import {
  View,
  Text,
  Alert,
  ActivityIndicator,
  TouchableOpacity,
  Modal,
  Pressable,
  Image,
} from 'react-native';
import {
  Camera,
  useCameraDevice,
  useCodeScanner,
} from 'react-native-vision-camera';
import styles from './styles/QrScannerScreenStyle';
import { BASE_API_URL } from '@env';
import { parseQrCode, getPrescriptionsByPatientId } from '../services/api';
import PrescriptionsDetail from './components/modal/PrescriptionsDetail';

export default function QrScannerScreen({ navigation }) {
  const [hasPermission, setHasPermission] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [isScanning, setIsScanning] = useState(true);
  const [scannedData, setScannedData] = useState(null);
  const [patientInfo, setPatientInfo] = useState(null);
  const [fetchingInfo, setFetchingInfo] = useState(false);
  const [modalVisible, setModalVisible] = useState(false);
  const [modalImageUrl, setModalImageUrl] = useState(null);
  const [prescriptionModalVisible, setPrescriptionModalVisible] =
    useState(false);
  const [prescriptionData, setPrescriptionData] = useState(null);
  const [qrTokenType, setQrTokenType] = useState(null);
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

  const fetchPatientInfo = async ({ id, field, itemId }) => {
    try {
      setFetchingInfo(true);
      const response = await fetch(
        `${BASE_API_URL}/users/${id}/patientinfo/${field}/${itemId}`,
      );
      if (!response.ok) throw new Error(`API error: ${response.status}`);

      const data = await response.json();
      setPatientInfo(data);
    } catch (error) {
      console.error('Failed to retrieve patient information:', error);
      Alert.alert(
        'Error',
        'Failed to retrieve patient information. Please try again.',
      );
      setPatientInfo(null);
      setIsScanning(true);
      setScannedData(null);
    } finally {
      setFetchingInfo(false);
    }
  };

  const handleScan = async codes => {
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
        // First try new backend QR token validation
        const parseResult = await parseQrCode(value);

        if (parseResult) {
          // Handle token already used
          if (parseResult.type === 'token_used') {
            console.warn('QR token already used:', parseResult.token);
            Alert.alert(
              'Geçersiz QR',
              'Bu QR daha önce kullanılmış. Yeni bir QR oluşturun veya yetkilinize başvurun.',
            );
            setIsScanning(true);
            setScannedData(null);
            return;
          }

          // Handle prescription_info token type
          if (
            parseResult.tokenType === 'prescription_info' &&
            parseResult.data
          ) {
            console.log('📋 QR Token data from QR:', parseResult.data);
            setQrTokenType('prescription_info');
            setFetchingInfo(true);

            // Get patient ID from QR token
            const patientId =
              parseResult.data.patient_id || parseResult.data.patient?.id;
            console.log('👤 Patient ID from QR token:', patientId);

            if (patientId) {
              // Fetch prescriptions for this patient
              const prescriptions = await getPrescriptionsByPatientId(
                patientId,
              );

              if (prescriptions && prescriptions.length > 0) {
                console.log(
                  `📋 Retrieved ${prescriptions.length} prescriptions`,
                );
                // Show first prescription or create a list
                const firstPrescription = prescriptions[0];
                setPrescriptionData(firstPrescription);
                setPrescriptionModalVisible(true);
              } else {
                console.warn('No prescriptions found for patient');
                Alert.alert('Bilgi', 'Bu hasta için reçete kaydı bulunamadı.');
              }
            } else {
              console.error('No patient ID found in QR token');
              Alert.alert('Hata', "QR token'de hasta ID bulunamadı.");
            }

            setFetchingInfo(false);
            return;
          }

          // Handle different response types from parseQrCode
          if (parseResult.type === 'assignment_success') {
            Alert.alert('Başarı', 'Hasta cihaza başarıyla atanmıştır', [
              {
                text: 'Tamam',
                onPress: () => {
                  setPatientInfo({
                    title: 'Hasta Atama Başarılı',
                    result: {
                      message: 'Hasta cihaza başarıyla atanmıştır',
                      tokenType: parseResult.tokenType,
                      token: parseResult.token,
                    },
                  });
                },
              },
            ]);
            return;
          } else if (parseResult.type === 'assignment_failed') {
            Alert.alert(
              'Atama Başarısız',
              'Hasta cihaza atanırken hata oluştu. Token doğruysa, manuel olarak atayabilirsiniz.',
              [
                {
                  text: 'Bilgiyorum',
                  onPress: () => {
                    setPatientInfo({
                      title: 'Hasta Atama Başarısız',
                      result: parseResult.data,
                    });
                  },
                },
              ],
            );
            return;
          } else if (parseResult.type === 'token_validated') {
            // Token validated successfully
            setPatientInfo({
              title: `Token Doğrulandı (${parseResult.tokenType})`,
              result: parseResult.data,
            });
            return;
          }
        }

        // Fallback to old format: QR contains JSON with { id, field, itemId }
        const parsed = JSON.parse(value.trim());
        if (parsed.id && parsed.field && parsed.itemId) {
          await fetchPatientInfo(parsed);
        } else {
          Alert.alert(
            'Missing Data',
            'QR code content does not contain required fields',
          );
          setIsScanning(true);
          setScannedData(null);
        }
      } catch {
        Alert.alert(
          'JSON Parse Error',
          'QR code content is not in valid format',
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

  const capitalize = str => str.charAt(0).toUpperCase() + str.slice(1);

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

      {fetchingInfo && (
        <View style={{ padding: 10, alignItems: 'center' }}>
          <ActivityIndicator size="large" />
          <Text>Hasta bilgisi yükleniyor...</Text>
        </View>
      )}

      {patientInfo && (
        <View style={styles.resultOverlay}>
          <View style={styles.resultContainer}>
            <Text style={styles.resultTitle}>{patientInfo.title}</Text>

            <View style={{ marginTop: 16 }}>
              <Text
                style={{
                  fontSize: 16,
                  fontWeight: 'bold',
                  color: '#333',
                  marginBottom: 8,
                }}
              >
                Sonuç:
              </Text>

              {Array.isArray(patientInfo.result) ? (
                patientInfo.result.map((item, index) => (
                  <View
                    key={index}
                    style={{
                      marginBottom: 12,
                      padding: 12,
                      backgroundColor: '#f9f9f9',
                      borderRadius: 8,
                      borderWidth: 1,
                      borderColor: '#ddd',
                    }}
                  >
                    {Object.entries(item).map(([key, value]) => {
                      if (
                        key.toLowerCase() === 'imageurl' &&
                        typeof value === 'string'
                      ) {
                        return (
                          <Pressable
                            key={key}
                            onPress={() => {
                              setModalImageUrl(value);
                              setModalVisible(true);
                            }}
                          >
                            <Image
                              source={{ uri: value }}
                              style={{
                                width: 150,
                                height: 150,
                                marginBottom: 8,
                                borderRadius: 8,
                              }}
                              resizeMode="contain"
                            />
                          </Pressable>
                        );
                      }
                      return (
                        <Text
                          key={key}
                          style={{
                            fontSize: 14,
                            marginBottom: 4,
                            color: '#444',
                          }}
                        >
                          <Text style={{ fontWeight: 'bold' }}>
                            {capitalize(key)}:{' '}
                          </Text>
                          {String(value)}
                        </Text>
                      );
                    })}
                  </View>
                ))
              ) : (
                <View
                  style={{
                    marginBottom: 12,
                    padding: 12,
                    backgroundColor: '#f9f9f9',
                    borderRadius: 8,
                    borderWidth: 1,
                    borderColor: '#ddd',
                  }}
                >
                  {Object.entries(patientInfo.result).map(([key, value]) => {
                    if (
                      key.toLowerCase() === 'imageurl' &&
                      typeof value === 'string'
                    ) {
                      return (
                        <Pressable
                          key={key}
                          onPress={() => {
                            setModalImageUrl(value);
                            setModalVisible(true);
                          }}
                        >
                          <Image
                            source={{ uri: value }}
                            style={{
                              width: 150,
                              height: 150,
                              marginBottom: 8,
                              borderRadius: 8,
                            }}
                            resizeMode="contain"
                          />
                        </Pressable>
                      );
                    }
                    return (
                      <Text
                        key={key}
                        style={{ fontSize: 14, marginBottom: 4, color: '#444' }}
                      >
                        <Text style={{ fontWeight: 'bold' }}>
                          {capitalize(key)}:{' '}
                        </Text>
                        {String(value)}
                      </Text>
                    );
                  })}
                </View>
              )}
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

      {/* Modal for showing large image */}
      <Modal
        visible={modalVisible}
        transparent={true}
        animationType="fade"
        onRequestClose={() => setModalVisible(false)}
      >
        <Pressable
          style={{
            flex: 1,
            backgroundColor: 'rgba(0,0,0,0.8)',
            justifyContent: 'center',
            alignItems: 'center',
          }}
          onPress={() => setModalVisible(false)}
        >
          <Image
            source={{ uri: modalImageUrl }}
            style={{ width: '90%', height: '90%', borderRadius: 12 }}
            resizeMode="contain"
          />
        </Pressable>
      </Modal>

      {/* Prescription Detail Modal - for QR prescription_info tokens */}
      <PrescriptionsDetail
        visible={prescriptionModalVisible}
        prescription={prescriptionData}
        onClose={() => {
          setPrescriptionModalVisible(false);
          handleRescan();
        }}
      />
    </View>
  );
}

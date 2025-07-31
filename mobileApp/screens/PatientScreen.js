import React, { useState, useEffect } from 'react';
import { View, StyleSheet, Text } from 'react-native';
import PatientProfile from './components/PatientProfile';
import TabBar from './components/TabBar';
import AppointmentsTestsTab from './components/AppointmentsTestsTab';
import MedicalHistoryTab from './components/MedicalHistoryTab';
import ActionButtons from './components/ActionButtons';
import DeviceInfo from 'react-native-device-info';
import { getPatientByDeviceId } from '../services/api'; // API fonksiyonunu değiştiriyoruz

export default function PatientScreen({ route, navigation }) {
  const [userData, setUserData] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [activeTab, setActiveTab] = useState('randevularTetkikler');

  useEffect(() => {
    const deviceId = DeviceInfo.getUniqueIdSync();
    console.log('Cihaz ID:', deviceId);
    
    // Eğer patient olarak giriş yapıldıysa, kendi bilgilerini göster
    const isPatientLogin = route.params?.isPatientLogin;
    const patientData = route.params?.userData;
    
    if (isPatientLogin && patientData) {
      console.log('Patient girişi - kendi bilgileri gösteriliyor:', patientData);
      setUserData(patientData);
      setIsLoading(false);
      return;
    }
    
    // Doctor/Nurse ise cihaza kayıtlı hasta verilerini al
    getPatientByDeviceId(deviceId)
      .then(data => {
        if (data) {
          setUserData(data);
          console.log('Cihaza kayıtlı hasta bulundu:', data);
        } else {
          setError("Bu cihaza kayıtlı hasta bulunamadı.");
          console.log('Cihaza kayıtlı hasta yok');
        }
        setIsLoading(false);
      })
      .catch(err => {
        console.error('Hasta verilerini alma hatası:', err);
        setError("Hasta verilerini alırken hata oluştu.");
        setIsLoading(false);
      });
  }, [route.params]);

  const tabs = [
    { key: 'randevularTetkikler', label: 'Randevular & Tetkikler' },
    { key: 'saglikGecmisi', label: 'Sağlık Geçmişi' },
  ];

  return (
    <View style={styles.container}>
      {isLoading ? (
        <Text style={styles.loadingText}>Yükleniyor...</Text>
      ) : error ? (
        <Text style={styles.errorText}>{error}</Text>
      ) : !userData ? (
        <Text style={styles.noDataText}>Hasta verisi bulunamadı.</Text>
      ) : (
        <>
          <PatientProfile userData={userData} />
  
          <TabBar 
            tabs={tabs}
            activeTab={activeTab}
            onTabPress={setActiveTab}
          />
  
          {activeTab === 'randevularTetkikler' && (
            <AppointmentsTestsTab data={userData} />
          )}
  
          {activeTab === 'saglikGecmisi' && (
            <MedicalHistoryTab data={userData} />
          )}
  
          <ActionButtons navigation={navigation} />
        </>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: { 
    flex: 1, 
    padding: 20 
  },
  loadingText: {
    textAlign: 'center',
    fontSize: 16,
    marginTop: 50,
  },
  errorText: {
    color: 'red', 
    textAlign: 'center',
    fontSize: 16,
    marginTop: 50,
  },
  noDataText: {
    textAlign: 'center',
    fontSize: 16,
    marginTop: 50,
  },
});
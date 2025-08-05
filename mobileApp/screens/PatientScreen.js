import { useState, useEffect } from 'react';
import { View, Text } from 'react-native';
import DeviceInfo from 'react-native-device-info';
import ActionButtons from './components/ActionButtons';
import AppointmentsTestsTab from './components/AppointmentsTestsTab';
import MedicalHistoryTab from './components/MedicalHistoryTab';
import PatientProfile from './components/PatientProfile';
import TabBar from './components/TabBar';
import { getPatientByDeviceId } from '../services/api';
import styles from './styles/PatientScreenStyle';

export default function PatientScreen({ route, navigation }) {
  const [userData, setUserData] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [activeTab, setActiveTab] = useState('randevularTetkikler');

  useEffect(() => {
    const deviceId = DeviceInfo.getUniqueIdSync();

    // cihaz ID'ye bağlı hasta verisi çekiliyor
    fetchPatientData(deviceId);
  }, [route.params]);

  const fetchPatientData = async (deviceId) => {
    try {
      const data = await getPatientByDeviceId(deviceId);
      if (data) {
        setUserData(data);
      } else {
        setError("Bu cihaza kayıtlı hasta bulunamadı.");
        console.warn('Cihaza kayıtlı hasta yok');
      }
    } catch (err) {
      console.error('Hasta verilerini alma hatası:', err);
      setError("Hasta verilerini alırken hata oluştu.");
    } finally {
      setIsLoading(false);
    }
  };

  const tabs = [
    { key: 'randevularTetkikler', label: 'Randevular & Tetkikler' },
    { key: 'saglikGecmisi', label: 'Sağlık Geçmişi' },
  ];

  const renderContent = () => {
    if (isLoading) {
      return <Text style={styles.loadingText}>Yükleniyor...</Text>;
    }
    if (error) {
      return <Text style={styles.errorText}>{error}</Text>;
    }
    if (!userData) {
      return <Text style={styles.noDataText}>Hasta verisi bulunamadı.</Text>;
    }
    return (
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
    );
  };

  return (
    <View style={styles.container}>
      {renderContent()}
    </View>
  );
}
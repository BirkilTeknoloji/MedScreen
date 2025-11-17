import { useState, useEffect } from 'react';
import { View, Text, TouchableOpacity, Animated } from 'react-native';
import DeviceInfo from 'react-native-device-info';
import ActionButtons from './components/ActionButtons';
import AppointmentsTestsTab from './components/AppointmentsTestsTab';
import MedicalHistoryTab from './components/MedicalHistoryTab';
import PatientProfile from './components/PatientProfile';
import TabBar from './components/TabBar';
import { getPatientByDeviceId } from '../services/api';
import styles from './styles/PatientScreenStyle';
import Icon from 'react-native-vector-icons/Entypo';
import CustomDropdown from './components/CustomDropdown';
import AppointmentsContainer from './components/AppointmentsContainer';

export default function PatientScreen({ route, navigation }) {
  const [userData, setUserData] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [activeTab, setActiveTab] = useState('randevularTetkikler');
  // const [dropModal, setDropModal] = useState(false);
  // const [animation] = useState(new Animated.Value(0));

  useEffect(() => {
    const deviceId = DeviceInfo.getUniqueIdSync();
    console.log('deviceeffect:', deviceId);
    fetchPatientData(deviceId);
  }, [route.params]);

  // const toggleDropdown = () => {
  //   const toValue = dropModal ? 0 : 1;

  //   Animated.timing(animation, {
  //     toValue,
  //     duration: 300,
  //     useNativeDriver: false,
  //   }).start();

  //   setDropModal(!dropModal);
  // };

  // // ✅ maxHeight interpolation eklendi
  // const maxHeight = animation.interpolate({
  //   inputRange: [0, 1],
  //   outputRange: [0, 300], // 0'dan 300px'e kadar genişler
  // });

  const fetchPatientData = async deviceId => {
    try {
      const data = await getPatientByDeviceId(deviceId);
      console.log('deviceID:', deviceId);
      console.log('data:', data);
      if (data) {
        console.log('Cihaza kayıtlı hasta bulundu:', data);
        setUserData(data);
      } else {
        setError('Bu cihaza kayıtlı hasta bulunamadı.');
        console.warn('Cihaza kayıtlı hasta yok');
      }
    } catch (err) {
      console.error('Hasta verilerini alma hatası:', err);
      setError('Hasta verilerini alırken hata oluştu.');
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
        <View style={styles.contentRow}>
          <AppointmentsContainer userData={userData} />


        </View>
        <ActionButtons navigation={navigation} />
      </>
    );
  };

  return <View style={styles.container}>{renderContent()}</View>;
}

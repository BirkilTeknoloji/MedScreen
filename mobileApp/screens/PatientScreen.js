import { useState, useEffect } from 'react';
import { View, Text, TouchableOpacity, Animated } from 'react-native';
import DeviceInfo from 'react-native-device-info';
import ActionButtons from './components/ActionButtons';
import AppointmentsTestsTab from './components/AppointmentsTestsTab';
import MedicalHistoryTab from './components/MedicalHistoryTab';
import PatientProfile from './components/PatientProfile';
import TabBar from './components/TabBar';
import {
  getAppointmentsByPatientId,
  getDiagnosesByPatientId,
  getFirstPatient,
  getPatientById,
  getPatientByDeviceId,
  getPrescriptionsByPatientId,
  getMedicalTestsByPatientId,
  getMedicalHistoryByPatientId,
  getSurgeryHistoryByPatientId,
  getAllergiesByPatientId,
} from '../services/api';
import styles from './styles/PatientScreenStyle';
import Icon from 'react-native-vector-icons/Entypo';
import CustomDropdown from './components/CustomDropdown';
import AppointmentsContainer from './components/AppointmentsContainer';

export default function PatientScreen({ route, navigation }) {
  const [userData, setUserData] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [firstData, setFirstData] = useState(null);
  const [appointments, setAppointments] = useState([]);
  const [diagnoses, setDiagnoses] = useState([]);
  const [prescriptions, setPrescriptions] = useState([]);
  const [medicalTests, setMedicalTests] = useState([]);
  const [medicalHistory, setMedicalHistory] = useState([]);
  const [surgeryHistory, setSurgeryHistory] = useState([]);
  const [allergies, setAllergies] = useState([]);
  const { patientData, doctorData, isPatientLogin, qrTokenData, qrTokenType, isQrNavigation } = route.params || {};
  const [activeTab, setActiveTab] = useState('randevularTetkikler');

  useEffect(() => {
    if (!route.params?.patientData) {
      // Check if coming from QR navigation with patient_id
      if (isQrNavigation && qrTokenData?.patient_id) {
        console.log('Fetching patient data from QR token:', qrTokenData.patient_id);
        fetchPatientDataById(qrTokenData.patient_id);
      } else {
        fetchPatientFirstData();
      }
    }
  }, []);

  useEffect(() => {
    if (userData) {
      fetchAppointmentsOnly();
    }
  }, [userData]);

  const fetchPatientFirstData = async () => {
    try {
      setIsLoading(true);

      const data = await getFirstPatient();

      if (data) {
        setUserData(data);
      } else {
        setError('Görüntülenecek hasta bulunamadı.');
      }
    } catch (err) {
      console.error('Hata:', err);
      setError('Veri alınamadı.');
    } finally {
      setIsLoading(false);
    }
  };

  const fetchPatientDataById = async (patientId) => {
    try {
      setIsLoading(true);
      // Fetch specific patient by ID from QR token
      const patient = await getPatientById(patientId);
      
      if (patient) {
        setUserData(patient);
      } else {
        console.warn('Patient not found for ID:', patientId);
        setError('Hasta verisi alınamadı.');
      }
    } catch (err) {
      console.error('Patient fetch error:', err);
      setError('Veri alınamadı.');
    } finally {
      setIsLoading(false);
    }
  };
  const fetchAppointmentsOnly = async () => {
    if (!userData?.ID) return;

    try {
      // Randevuları çek
      const appts = await getAppointmentsByPatientId(userData.ID);
      setAppointments(appts);

      // 3. TANILARI VE İLAÇLARI ÇEK
      console.log('Tanılar isteniyor...');
      const diags = await getDiagnosesByPatientId(userData.ID);
      console.log('İlaçlar isteniyor...');
      const pres = await getPrescriptionsByPatientId(userData.ID);
      console.log('Tetkikler isteniyor...');
      const tests = await getMedicalTestsByPatientId(userData.ID);
      console.log('Tıbbi geçmiş isteniyor...');
      const history = await getMedicalHistoryByPatientId(userData.ID);
      console.log('Ameliyat geçmişi isteniyor...');
      const surgery = await getSurgeryHistoryByPatientId(userData.ID);
      console.log('Alerjiler isteniyor...');
      const allergiesData = await getAllergiesByPatientId(userData.ID);
      setDiagnoses(diags);
      setPrescriptions(pres);
      setMedicalTests(tests);
      setMedicalHistory(history);
      setSurgeryHistory(surgery);
      setAllergies(allergiesData);
    } catch (err) {
      console.error(err);
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
        {isQrNavigation && (
          <View style={{ backgroundColor: '#e3f2fd', padding: 12, marginBottom: 8, borderRadius: 4 }}>
            <Text style={{ fontSize: 12, color: '#1976d2', fontWeight: 'bold' }}>
              ✅ QR Token ile Yüklendi ({qrTokenType})
            </Text>
          </View>
        )}
        <PatientProfile userData={userData} />
        <View style={styles.contentRow}>
          <AppointmentsContainer
            userData={userData}
            appointments={appointments}
            diagnoses={diagnoses}
            prescriptions={prescriptions}
            medicalTests={medicalTests}
            medicalHistory={medicalHistory}
            surgeryHistory={surgeryHistory}
            allergies={allergies}
          />
        </View>
        <ActionButtons navigation={navigation} />
      </>
    );
  };

  return <View style={styles.container}>{renderContent()}</View>;
}

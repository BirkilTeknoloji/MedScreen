import { useState, useEffect, useCallback } from 'react';
import { View, Text, ActivityIndicator, ScrollView } from 'react-native';
import ActionButtons from './components/ActionButtons';
import PatientProfile from './components/PatientProfile';
import AppointmentsContainer from './components/AppointmentsContainer';
import * as Api from '../services/apiService';
import styles from './styles/PatientScreenStyle';

export default function PatientScreen({ route, navigation }) {
  const [userData, setUserData] = useState(null);
  const [patientId, setPatientId] = useState('H0001');

  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  // Veri State'leri
  const [appointments, setAppointments] = useState([]);
  const [diagnoses, setDiagnoses] = useState([]);
  const [prescriptions, setPrescriptions] = useState([]);
  const [medicalTests, setMedicalTests] = useState([]);
  const [medicalHistory, setMedicalHistory] = useState([]);
  const [surgeryHistory, setSurgeryHistory] = useState([]);
  const [allergies, setAllergies] = useState([]);

  const { isQrNavigation, qrTokenData } = route.params || {};

  const resetDetails = () => {
    setAppointments([]);
    setDiagnoses([]);
    setPrescriptions([]);
    setMedicalTests([]);
    setMedicalHistory([]);
    setSurgeryHistory([]);
    setAllergies([]);
  };

  // Normalize helper
  const norm = v =>
    (v ?? '').toString().trim().toLocaleUpperCase('tr-TR').replaceAll('İ', 'I');

  // 1) Base data
  useEffect(() => {
    const fetchBaseData = async () => {
      try {
        setIsLoading(true);
        setError(null);
        resetDetails();

        const pid =
          isQrNavigation && qrTokenData?.patient_id
            ? qrTokenData.patient_id
            : 'H0001';

        setPatientId(pid);

        console.log('Hasta yükleniyor:', pid);
        const data = await Api.getPatientById(pid);

        if (data) {
          setUserData(data);
          console.log('Hasta yüklendi:', data);
        } else {
          setUserData(null);
          setError('Hasta verisi boş döndü.');
        }
      } catch (err) {
        setUserData(null);
        setError('Sunucu bağlantı hatası.');
        console.error('Hasta yükleme hatası:', err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchBaseData();
  }, [isQrNavigation, qrTokenData]);

  // 2) Details
  const fetchAllDetails = useCallback(async () => {
    if (!userData || !patientId) return;

    try {
      const [appts, diags, pres, tests, historyRes] = await Promise.all([
        Api.getAppointmentsByPatientId(patientId),
        Api.getDiagnosesByPatientId(patientId),
        Api.getPrescriptionsByPatientId(patientId),
        Api.getMedicalTestsByPatientId(patientId),
        Api.getMedicalHistoryByPatientId(patientId),
      ]);

      const history = Array.isArray(historyRes)
        ? historyRes
        : historyRes?.data ?? [];

      // Debug
      console.log('history raw:', history);
      console.log(
        'history types:',
        (history || []).map(x => x.turu || x.TURU),
      );

      // PatientScreen.js içindeki filtreleme mantığı
      if (history && Array.isArray(history)) {
        // 1. Alerjiler
        const allergyData = history.filter(
          i => i.tibbi_bilgi_turu_kodu === 'ALERJI',
        );
        setAllergies(allergyData);

        // 2. Ameliyatlar
        const surgeryData = history.filter(
          i => i.tibbi_bilgi_turu_kodu === 'AMELIYAT',
        );
        setSurgeryHistory(surgeryData);

        // 3. Kronik Hastalıklar (Tıbbi Geçmiş)
        const chronicData = history.filter(
          i =>
            i.tibbi_bilgi_turu_kodu === 'KRONIK' ||
            i.aciklama?.toLowerCase().includes('kronik'),
        );
        setMedicalHistory(chronicData);
      }

      setAppointments(appts || []);
      setDiagnoses(diags || []);
      setPrescriptions(pres || []);
      setMedicalTests(tests || []);
    } catch (err) {
      console.error('Detay veri çekme/ayrıştırma hatası:', err);
    }
  }, [userData, patientId]);

  useEffect(() => {
    fetchAllDetails();
  }, [fetchAllDetails]);

  useEffect(() => {
    console.log("Dropdown'a gelen alerjiler:", allergies);
  }, [allergies]);

  // --- Render ---
  if (isLoading) {
    return (
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
        <ActivityIndicator size="large" color="#4A90E2" />
        <Text style={{ marginTop: 10, color: '#666' }}>
          Hasta bilgileri yükleniyor...
        </Text>
      </View>
    );
  }

  if (error) {
    return (
      <View
        style={{
          flex: 1,
          justifyContent: 'center',
          alignItems: 'center',
          padding: 20,
        }}
      >
        <Text style={{ color: 'red', fontSize: 16, textAlign: 'center' }}>
          {error}
        </Text>
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <View style={{ flex: 1 }}>
        {userData ? (
          <>
            <PatientProfile
              userData={userData}
              actionButtons={<ActionButtons navigation={navigation} />}
            />
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
          </>
        ) : (
          <View style={{ padding: 20 }}>
            <Text>Hasta bilgisi bulunamadı.</Text>
          </View>
        )}
      </View>
    </View>
  );
}

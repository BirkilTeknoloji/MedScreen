import { useState, useEffect, useCallback } from 'react';
import { View, Text, ActivityIndicator, ScrollView } from 'react-native';
import ActionButtons from './components/ActionButtons';
import PatientProfile from './components/PatientProfile';
import AppointmentsContainer from './components/AppointmentsContainer';
import * as Api from '../services/apiService';
import styles from './styles/PatientScreenStyle';

export default function PatientScreen({ route, navigation }) {
  const [userData, setUserData] = useState(null);
  const [patientInfo, setPatientInfo] = useState(null);
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

        const data = await Api.getPatientById(pid);
        const info = await Api.getPatientInfoByPatientId(pid);

        if (data) {
          // Merge patient data with latest application info
          const latestApp =
            info && Array.isArray(info) && info.length > 0 ? info[0] : {};
          const combinedData = {
            ...latestApp, // Application details (protokol no, birim, vb.)
            ...data, // Patient demographics (ad, soyadi, vb.)
            hasta: data, // Explicitly set for PatientProfile logic
            hekim: latestApp.hekim || {}, // Set doctor from application
          };

          setUserData(combinedData);
          setPatientInfo(info);
          console.log('Hasta yüklendi (Combined):', combinedData);
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
      // Her biri için ayrı try-catch kullanarak birinin hatasının diğerlerini bozmasını engelliyoruz
      const fetchData = async (apiFunc, setter, label) => {
        try {
          const res = await apiFunc(patientId);
          setter(res || []);
          console.log(`${label} yüklendi:`, res ? res.length : 0, 'kayıt');
        } catch (err) {
          console.error(`${label} çekme hatası:`, err);
          setter([]);
        }
      };

      await Promise.all([
        fetchData(
          Api.getAppointmentsByPatientId,
          setAppointments,
          'Randevular',
        ),
        fetchData(Api.getDiagnosesByPatientId, setDiagnoses, 'Tanılar'),
        fetchData(
          Api.getPrescriptionsByPatientId,
          setPrescriptions,
          'Reçeteler',
        ),
        fetchData(Api.getMedicalTestsByPatientId, setMedicalTests, 'Tetkikler'),
        (async () => {
          try {
            const historyRes = await Api.getMedicalHistoryByPatientId(
              patientId,
            );
            const history = Array.isArray(historyRes)
              ? historyRes
              : historyRes?.data ?? [];

            console.log(
              'Geçmiş verisi (history) yüklendi:',
              (history || []).length,
              'kayıt',
            );

            if (history && Array.isArray(history)) {
              setAllergies(
                history.filter(i => i.tibbi_bilgi_turu_kodu === 'ALERJI'),
              );
              setSurgeryHistory(
                history.filter(i => i.tibbi_bilgi_turu_kodu === 'AMELIYAT'),
              );
              setMedicalHistory(
                history.filter(
                  i =>
                    i.tibbi_bilgi_turu_kodu === 'KRONIK' ||
                    i.aciklama?.toLowerCase().includes('kronik'),
                ),
              );
            }
          } catch (err) {
            console.error('Geçmiş verisi çekme hatası:', err);
          }
        })(),
      ]);
    } catch (err) {
      console.error('Genel veri çekme hatası:', err);
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
  console.log('userdata: ', userData);
  return (
    <View style={styles.container}>
      <View style={{ flex: 1 }}>
        {userData ? (
          <>
            <PatientProfile
              userData={userData}
              patientInfo={patientInfo}
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

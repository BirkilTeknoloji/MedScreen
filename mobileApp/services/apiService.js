import axios from 'axios';
import AsyncStorage from '@react-native-async-storage/async-storage';

const BASE_URL = 'http://192.168.1.101:8080/api/v1';

const apiClient = axios.create({
  baseURL: BASE_URL,
  timeout: 10000,
});

apiClient.interceptors.request.use(async config => {
  try {
    const token = await AsyncStorage.getItem('auth_token');
    if (token) config.headers.Authorization = `Bearer ${token}`;
  } catch (e) {
    console.error('Token alınamadı:', e);
  }
  return config;
});

const handleResponse = res => res.data?.data ?? res.data;

// ==================== HASTA ====================
export const getPatientById = id =>
  apiClient.get(`/hasta/${id}`).then(handleResponse);
export const getPatientInfoByPatientId = id =>
  apiClient.get(`/hasta-basvuru/hasta/${id}`).then(handleResponse);

// ==================== RANDEVU ====================
export const getAppointmentsByPatientId = id =>
  apiClient.get(`/randevu/hasta/${id}`).then(handleResponse);

// ==================== TANILAR ====================
export const getDiagnosesByPatientId = id =>
  apiClient.get(`/basvuru-tani/hasta/${id}`).then(handleResponse);

// ==================== REÇETELER ====================
export const getPrescriptionsByPatientId = async patientId => {
  try {
    const applicationsResponse = await apiClient.get(
      `/hasta-basvuru/hasta/${patientId}`,
    );
    const applications = applicationsResponse.data?.data ?? [];

    if (!applications.length) return [];

    const allPrescriptions = [];
    for (const app of applications) {
      try {
        const prescResponse = await apiClient.get(
          `/recete/basvuru/${app.hasta_basvuru_kodu}`,
        );
        const prescData = prescResponse.data?.data ?? [];
        if (prescData.length) allPrescriptions.push(...prescData);
      } catch (error) {
        console.log(`Başvuru ${app.hasta_basvuru_kodu} için reçete bulunamadı`);
      }
    }

    return allPrescriptions;
  } catch (error) {
    console.error('Reçeteler çekilemedi:', error);
    return [];
  }
};

// ==================== TETKİKLER ====================
export const getMedicalTestsByPatientId = async patientId => {
  try {
    const applicationsResponse = await apiClient.get(
      `/hasta-basvuru/hasta/${patientId}`,
    );
    const applications = applicationsResponse.data?.data ?? [];

    if (!applications.length) return [];

    const allTests = [];
    for (const app of applications) {
      try {
        const testsResponse = await apiClient.get(
          `/tetkik-sonuc/basvuru/${app.hasta_basvuru_kodu}`,
        );
        const testsData = testsResponse.data?.data ?? [];
        if (testsData.length) allTests.push(...testsData);
      } catch (error) {
        console.log(`Başvuru ${app.hasta_basvuru_kodu} için tetkik bulunamadı`);
      }
    }

    return allTests;
  } catch (error) {
    console.error('Tetkikler çekilemedi:', error);
    return [];
  }
};

// ==================== TIBBİ BİLGİ ====================
export const getMedicalHistoryByPatientId = id =>
  apiClient.get(`/hasta-tibbi-bilgi/hasta/${id}`).then(handleResponse);

// ==================== VİTAL ====================
export const getVitalSignsByApplicationId = applicationId =>
  apiClient.get(`/vital-bulgu/basvuru/${applicationId}`).then(handleResponse);

export const getClinicalCourseByApplicationId = applicationId =>
  apiClient.get(`/klinik-seyir/basvuru/${applicationId}`).then(handleResponse);

export const getPatientAlertsByApplicationId = applicationId =>
  apiClient.get(`/hasta-uyari/basvuru/${applicationId}`).then(handleResponse);

export const getRiskScoresByApplicationId = applicationId =>
  apiClient.get(`/risk-skorlama/basvuru/${applicationId}`).then(handleResponse);

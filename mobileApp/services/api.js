import { BASE_API_URL } from '@env';
import DeviceInfo from 'react-native-device-info';
import AsyncStorage from '@react-native-async-storage/async-storage';

export const addPatient = async userId => {
  try {
    const deviceId = DeviceInfo.getUniqueIdSync();
    const response = await fetch(`${BASE_API_URL}/devices/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ device_id: deviceId, user_id: userId }),
    });
    console.log(response);
    if (!response.ok) {
      const errorData = await response.text();
      console.error('API response error:', errorData);
      throw new Error('Hasta kaydı başarısız');
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Fetch ile hasta kaydı hatası:', error);
    throw error;
  }
};

export const getPatientData = async id => {
  try {
    const url = `${BASE_API_URL}/patients`;
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });
    if (!response.ok) {
      console.error('Backend bağlantı hatası:', response.status);
      return null;
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Backend gönderme hatası:', error.message);
    return null;
  }
};
export const getFirstPatient = async () => {
  const token = await AsyncStorage.getItem('userToken');

  if (!token) {
    console.warn('Hata: Token bulunamadı, istek atılmıyor.');
    return null;
  }

  try {
    const url = `${BASE_API_URL}/patients`;
    console.log('İlk hasta isteniyor:', url);

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`, 
      },
    });

    if (!response.ok) {
      console.error('Bağlantı hatası:', response.status);
      return null;
    }

    const responseJson = await response.json();

    if (
      responseJson.success && 
      Array.isArray(responseJson.data) &&
      responseJson.data.length > 0
    ) {
      const firstPatient = responseJson.data[0];
      console.log('İlk hasta başarıyla çekildi:', firstPatient.first_name);
      return firstPatient;
    } else {
      console.warn('Liste boş veya beklenen formatta değil.', responseJson);
      return null;
    }
  } catch (error) {
    console.error('getFirstPatient Hatası:', error.message);
    throw error;
  }
};
export const getPatientByUserId = async userId => {
  try {
    const url = `${BASE_API_URL}/patients`;
    console.log('Patients URL: ', url);
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      timeout: 10000,
    });

    if (!response.ok) {
      console.error('HTTP Error:', response.status, response.statusText);
      throw new Error(`HTTP ${response.status}: ${response.statusText}`);
    }

    const data = await response.json();
    console.log('All patients data:', data);

    const patient = data.find(
      p => p.primary_doctor_id === userId || p.id === userId,
    );
    return patient;
  } catch (error) {
    console.error('Network Error Details:', {
      message: error.message,
      name: error.name,
      stack: error.stack,
    });
    throw new Error(`Bağlantı hatası: ${error.message}`);
  }
};

export const getPatientById = async patientId => {
  try {
    const url = `${BASE_API_URL}/patients/1`;
    console.log('Patient by ID URL: ', url);
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      timeout: 10000,
    });

    if (!response.ok) {
      console.error('HTTP Error:', response.status, response.statusText);
      throw new Error(`HTTP ${response.status}: ${response.statusText}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Network Error Details:', {
      message: error.message,
      name: error.name,
      stack: error.stack,
    });
    throw new Error(`Bağlantı hatası: ${error.message}`);
  }
};

export const getPatientMedicalHistory = async patientId => {
  try {
    const url = `${BASE_API_URL}/patients/${patientId}/medical-history`;
    console.log('Medical History URL: ', url);
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      timeout: 10000,
    });

    if (!response.ok) {
      console.error('HTTP Error:', response.status, response.statusText);
      throw new Error(`HTTP ${response.status}: ${response.statusText}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Network Error Details:', {
      message: error.message,
      name: error.name,
      stack: error.stack,
    });
    throw new Error(`Bağlantı hatası: ${error.message}`);
  }
};

export const getAppointmentsByPatientId = async patientId => {
  const token = await AsyncStorage.getItem('userToken');
  try {
    const url = `${BASE_API_URL}/appointments`;
    console.log(`Randevular çekiliyor (Hedef Hasta ID: ${patientId})...`);

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.error('Randevu isteği başarısız:', response.status);
      return [];
    }

    const responseJson = await response.json();

    if (responseJson.success && Array.isArray(responseJson.data)) {
      const allAppointments = responseJson.data;


      const actualPatientId = patientId?.ID || patientId?.id || patientId;
      const filteredAppointments = allAppointments.filter(
        appointment => appointment.patient_id === actualPatientId,
      );
      console.log(`Randevular filtreleniyor...`);
      console.log(`Hedef Hasta ID: ${patientId}`);

      console.log(
        `Toplam ${allAppointments.length} randevudan, bu hastaya ait ${filteredAppointments.length} randevu bulundu.`,
      );

      return filteredAppointments;
    } else {
      console.warn('API yanıtı beklenen formatta değil (data dizisi yok).');
      return [];
    }
  } catch (error) {
    console.error('getAppointmentsByPatientId Hatası:', error);
    return [];
  }
};

export const getDiagnosesByPatientId = async patientId => {
  const token = await AsyncStorage.getItem('userToken');
  try {
    const url = `${BASE_API_URL}/diagnoses`;
    console.log(`Tanılar çekiliyor (Hasta ID: ${patientId})...`);

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.error('Tanı isteği başarısız:', response.status);
      return [];
    }

    const responseJson = await response.json();

    if (responseJson.success && Array.isArray(responseJson.data)) {
      const allDiagnoses = responseJson.data;
      const filteredDiagnoses = allDiagnoses.filter(
        diagnosis => diagnosis.patient_id === patientId,
      );

      console.log(
        `Toplam ${allDiagnoses.length} tanıdan, bu hastaya ait ${filteredDiagnoses.length} tanı bulundu.`,
      );
      return filteredDiagnoses;
    } else {
      console.warn('Tanı API yanıtı beklenen formatta değil.');
      return [];
    }
  } catch (error) {
    console.error('getDiagnosesByPatientId Hatası:', error);
    return [];
  }
};

export const getPrescriptionsByPatientId = async patientId => {
  const token = await AsyncStorage.getItem('userToken');
  try {
    const url = `${BASE_API_URL}/prescriptions`;
    console.log(`ilaçlar çekiliyor (Hasta ID: ${patientId})...`);

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.error('ilaçlar isteği başarısız:', response.status);
      return [];
    }

    const responseJson = await response.json();

    if (responseJson.success && Array.isArray(responseJson.data)) {
      const allPrescription = responseJson.data;
      const filteredPrescription = allPrescription.filter(
        prescription => prescription.patient_id === patientId,
      );

      console.log(
        `Toplam ${allPrescription.length} ilaçtan, bu hastaya ait ${filteredPrescription.length} ilaç bulundu.`,
      );
      return filteredPrescription;
    } else {
      console.warn('İlaç API yanıtı beklenen formatta değil.');
      return [];
    }
  } catch (error) {
    console.error('getPrescriptionsByPatientId Hatası:', error);
    return [];
  }
};

export const getMedicalTestsByPatientId = async patientId => {
  const token = await AsyncStorage.getItem('userToken');
  try {
    const url = `${BASE_API_URL}/medical-tests`;
    console.log(`Tetkikler çekiliyor (Hasta ID: ${patientId})...`);

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.error('Tetkik isteği başarısız:', response.status);
      return [];
    }

    const responseJson = await response.json();

    if (responseJson.success && Array.isArray(responseJson.data)) {
      const allTests = responseJson.data;
      const filteredTests = allTests.filter(
        test => test.patient_id === patientId,
      );

      console.log(
        `Toplam ${allTests.length} tetkikten, bu hastaya ait ${filteredTests.length} tetkik bulundu.`,
      );
      return filteredTests;
    } else {
      console.warn('Tetkik API yanıtı beklenen formatta değil.');
      return [];
    }
  } catch (error) {
    console.error('getMedicalTestsByPatientId Hatası:', error);
    return [];
  }
};

export const getMedicalHistoryByPatientId = async patientId => {
  const token = await AsyncStorage.getItem('userToken');
  try {
    const url = `${BASE_API_URL}/medical-history`;
    console.log(`Tıbbi geçmiş çekiliyor (Hasta ID: ${patientId})...`);

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.error('Tıbbi geçmiş isteği başarısız:', response.status);
      return [];
    }

    const responseJson = await response.json();

    if (responseJson.success && Array.isArray(responseJson.data)) {
      const allHistory = responseJson.data;
      const filteredHistory = allHistory.filter(
        history => history.patient_id === patientId,
      );

      console.log(
        `Toplam ${allHistory.length} tıbbi geçmişten, bu hastaya ait ${filteredHistory.length} kayıt bulundu.`,
      );
      return filteredHistory;
    } else {
      console.warn('Tıbbi geçmiş API yanıtı beklenen formatta değil.');
      return [];
    }
  } catch (error) {
    console.error('getMedicalHistoryByPatientId Hatası:', error);
    return [];
  }
};

export const getSurgeryHistoryByPatientId = async patientId => {
  const token = await AsyncStorage.getItem('userToken');
  try {
    const url = `${BASE_API_URL}/surgery-history`;
    console.log(`Ameliyat geçmişi çekiliyor (Hasta ID: ${patientId})...`);

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.error('Ameliyat geçmişi isteği başarısız:', response.status);
      return [];
    }

    const responseJson = await response.json();

    if (responseJson.success && Array.isArray(responseJson.data)) {
      const allSurgeries = responseJson.data;
      const filteredSurgeries = allSurgeries.filter(
        surgery => surgery.patient_id === patientId,
      );

      console.log(
        `Toplam ${allSurgeries.length} ameliyat geçmişinden, bu hastaya ait ${filteredSurgeries.length} kayıt bulundu.`,
      );
      return filteredSurgeries;
    } else {
      console.warn('Ameliyat geçmişi API yanıtı beklenen formatta değil.');
      return [];
    }
  } catch (error) {
    console.error('getSurgeryHistoryByPatientId Hatası:', error);
    return [];
  }
};

export const getAllergiesByPatientId = async patientId => {
  const token = await AsyncStorage.getItem('userToken');
  try {
    const url = `${BASE_API_URL}/allergies`;
    console.log(`Alerjiler çekiliyor (Hasta ID: ${patientId})...`);

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      console.error('Alerji isteği başarısız:', response.status);
      return [];
    }

    const responseJson = await response.json();

    if (responseJson.success && Array.isArray(responseJson.data)) {
      const allAllergies = responseJson.data;
      const filteredAllergies = allAllergies.filter(
        allergy => allergy.patient_id === patientId,
      );

      console.log(
        `Toplam ${allAllergies.length} alerjiden, bu hastaya ait ${filteredAllergies.length} kayıt bulundu.`,
      );
      return filteredAllergies;
    } else {
      console.warn('Alerji API yanıtı beklenen formatta değil.');
      return [];
    }
  } catch (error) {
    console.error('getAllergiesByPatientId Hatası:', error);
    return [];
  }
};

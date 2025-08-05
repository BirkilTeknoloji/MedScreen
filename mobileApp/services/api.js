import { BASE_API_URL } from '@env';
import DeviceInfo from 'react-native-device-info';

export const addPatient = async (userId) => {
  try {
    const deviceId = DeviceInfo.getUniqueIdSync();
    const response = await fetch(`${BASE_API_URL}/devices/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ device_id: deviceId, user_id: userId }),
    });
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

export const getPatientData = async (id) => {
  try {
    const url = `${BASE_API_URL}/users/${id}/patientinfo`;
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

export const getPatientByDeviceId = async (deviceId) => {
  try {
    const url = `${BASE_API_URL}/users/device/${deviceId}/patientinfo`;
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });
    const data = await response.json();
    if (response.ok) {
      return data;
    } else {
      throw new Error(data.message || "Bir hata oluştu");
    }
  } catch (error) {
    console.error('Cihaza kayıtlı hasta verileri alınamadı:', error);
    throw error;
  }
};
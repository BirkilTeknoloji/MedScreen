export const addPatient = async (deviceId, userId) => {
  try {
    const response = await fetch('http://192.168.1.113:8080/api/v1/devices/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ device_id: deviceId, user_id: userId }),
    });
    if (!response.ok) {
      const errorData = await response.text();
      console.error('API response error:', errorData);
      throw new Error('Hasta kaydı başarısız');
    }

    const data = await response.json();
    console.log('Hasta kaydı başarılı:', data);
    return data;
  } catch (error) {
    console.error('Fetch ile hasta kaydı hatası:', error);
    throw error;
  }
};

export const getPatientData = async (id) => {
  try {
    const API_URL = "http://192.168.1.113:8080/api/v1/users/:id/patientinfo";
    const url = API_URL.replace(':id', id);
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });
    if (!response.ok) {
      console.log('Backend bağlantı hatası:', response.status);
      setError(`Veri alınamadı (Hata: ${response.status})`);
      return null;
    }
    const data = await response.json();
    console.log('Backend patient:', data);
    return data;
  } catch (error) {
    console.log('Backend gönderme hatası:', error.message);
    setError('Bir ağ hatası oluştu.');
    return null;
  }
}

export const getPatientByDeviceId = async (deviceId) => {
  try {
    const API_URL = "http://192.168.1.113:8080/api/v1/users/device/:deviceId/patientinfo";
    const url = API_URL.replace(':deviceId', deviceId);
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });
    const data = await response.json();
    console.log("device response:", data);
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
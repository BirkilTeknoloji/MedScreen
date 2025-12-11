import { BASE_API_URL } from '@env';
import DeviceInfo from 'react-native-device-info';
import AsyncStorage from '@react-native-async-storage/async-storage';

// --- Auth / Role helpers ---
// getUserInfo: returns parsed userInfo object from AsyncStorage or null
export const getUserInfo = async () => {
  try {
    const raw = await AsyncStorage.getItem('userInfo');
    if (!raw) return null;
    return JSON.parse(raw);
  } catch (e) {
    console.error('getUserInfo error:', e);
    return null;
  }
};

// getUserRoles: returns array of role strings (may be empty)
export const getUserRoles = async () => {
  try {
    const raw = await AsyncStorage.getItem('userRoles');
    if (!raw) return [];
    const parsed = JSON.parse(raw);
    return Array.isArray(parsed) ? parsed : [parsed];
  } catch (e) {
    console.error('getUserRoles error:', e);
    return [];
  }
};

// hasRole: checks if current user has one of the required roles (case-insensitive)
export const hasRole = async requiredRoles => {
  try {
    if (!requiredRoles) return false;
    const roles = await getUserRoles();
    const required = Array.isArray(requiredRoles)
      ? requiredRoles
      : [requiredRoles];

    // Debug logging for role check
    console.log('🔑 hasRole check - User roles:', roles);
    console.log('🔑 hasRole check - Required roles:', required);

    // Case-insensitive comparison
    const normalizedRoles = roles.map(r =>
      typeof r === 'string' ? r.toLowerCase() : '',
    );
    const normalizedRequired = required.map(r =>
      typeof r === 'string' ? r.toLowerCase() : '',
    );

    const result = normalizedRequired.some(r => normalizedRoles.includes(r));
    console.log('🔑 hasRole result:', result);

    return result;
  } catch (e) {
    console.error('hasRole error:', e);
    return false;
  }
};

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
    const userToken = await AsyncStorage.getItem('userToken');

    if (!userToken) {
      console.error('No user token for getPatientById');
      return null;
    }

    console.log('👤 Fetching patient by ID:', patientId);
    const url = `${BASE_API_URL}/patients`;

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${userToken}`,
      },
    });

    if (!response.ok) {
      console.error('Patient fetch failed:', response.status);
      return null;
    }

    const data = await response.json();
    const allPatients = data.data || data;

    if (Array.isArray(allPatients)) {
      const patient = allPatients.find(
        p => p.id === patientId || p.ID === patientId,
      );
      if (patient) {
        console.log('👤 Patient found:', patient.first_name, patient.last_name);
        return patient;
      } else {
        console.warn('Patient not found with ID:', patientId);
        return null;
      }
    }

    return null;
  } catch (error) {
    console.error('Error fetching patient by ID:', error);
    return null;
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

// Try new backendV2 QR parsing endpoint. Return parsed data or null.
// Updated parseQrCode for backendV2 - handles token-based QR validation
export const parseQrCode = async qrValue => {
  try {
    const userToken = await AsyncStorage.getItem('userToken');

    if (!userToken) {
      console.error('No user token available for QR validation');
      return null;
    }

    // Check if qrValue is a JSON (old format) or a token (new format)
    let isToken = false;
    try {
      const parsed = JSON.parse(qrValue);
      // If it parses as JSON and has id/field/itemId, it's old format
      if (parsed.id && parsed.field && parsed.itemId) {
        console.log('Old format QR detected (JSON with id/field/itemId)');
        return null; // Let the caller handle old format fallback
      }
    } catch {
      // Not JSON, so it's likely a token string
      isToken = true;
    }

    if (!isToken) {
      // Could be plain text that's not a token either
      return null;
    }

    // New format: token may be just a UUID string, a full URL, or a URL containing a token query param.
    // Try to extract a raw token string from common QR contents.
    let tokenStr = null;

    // If scanned value looks like JSON with fields used by the old format, bail out earlier
    // (handled above). Otherwise try parsing as URL to extract token param or path segment.
    try {
      // If qrValue is a URL, this will succeed
      const url = new URL(qrValue);

      // Check query params like ?token=...
      if (url.searchParams.has('token')) {
        tokenStr = url.searchParams.get('token');
      } else {
        // Look for path segments containing 'qr-tokens' and take the following segment as token
        const parts = url.pathname.split('/').filter(Boolean);
        const idx = parts.findIndex(p => p.toLowerCase() === 'qr-tokens');
        if (idx !== -1 && parts.length > idx + 1) {
          tokenStr = parts[idx + 1];
        }
      }
    } catch (e) {
      // Not a full URL, continue
    }

    // If still not found, check if the scanned value itself contains 'qr-tokens/<token>'
    if (!tokenStr) {
      const match = qrValue.match(/qr-tokens[\/:]([A-Za-z0-9\-_.]+)/i);
      if (match && match[1]) {
        tokenStr = decodeURIComponent(match[1]);
      }
    }

    // If still not found, assume the entire scanned value is the token
    if (!tokenStr) {
      tokenStr = qrValue.trim();
    }

    // Final sanity trim
    tokenStr = tokenStr.replace(/^\/+|\/+$/g, '');

    console.log('📱 RAW QR scanned value:', qrValue);
    console.log('📱 Extracted token string:', tokenStr);
    const validateUrl = `${BASE_API_URL}/qr-tokens/${encodeURIComponent(
      tokenStr,
    )}/validate`;
    console.log('📱 Validation URL:', validateUrl);

    const validateResponse = await fetch(validateUrl, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${userToken}`,
      },
    });

    if (!validateResponse.ok) {
      // Try to parse response body to surface structured error information
      let bodyText = '';
      let bodyJson = null;
      try {
        bodyText = await validateResponse.text();
        bodyJson = JSON.parse(bodyText);
      } catch (e) {
        /* ignore JSON parse errors */
      }

      // If token has already been used, return a structured result so UI can handle it
      const msg =
        (bodyJson && (bodyJson.message || bodyJson.error)) || bodyText || '';
      if (
        typeof msg === 'string' &&
        msg.toLowerCase().includes('token has already been used')
      ) {
        console.warn('Token already used:', tokenStr);
        return {
          type: 'token_used',
          token: tokenStr,
          tokenType:
            bodyJson && bodyJson.data && bodyJson.data.type
              ? bodyJson.data.type
              : null,
          data: bodyJson || { message: msg },
        };
      }

      console.error(
        'Token validation failed:',
        validateResponse.status,
        bodyText,
      );
      return null;
    }

    const tokenData = await validateResponse.json();
    console.log('QR token validated:', tokenData);

    // Extract token info from response
    // Expected: { success: true, data: { type, patient_id, device_id, expires_at, ... } }
    const qrTokenInfo = tokenData.data || tokenData;

    // If token type is patient_assignment and we have device MAC, try to assign
    if (qrTokenInfo.type === 'patient_assignment') {
      const deviceMac = await AsyncStorage.getItem('device_mac');

      if (deviceMac) {
        console.log('Attempting to assign patient to device:', deviceMac);
        const scanUrl = `${BASE_API_URL}/devices/${deviceMac}/scan-patient-qr`;

        try {
          const scanResponse = await fetch(scanUrl, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              Authorization: `Bearer ${userToken}`,
            },
            body: JSON.stringify({ token: qrValue }),
          });

          if (scanResponse.ok) {
            console.log('Patient successfully assigned to device');
            const scanData = await scanResponse.json();
            return {
              type: 'assignment_success',
              token: qrValue,
              tokenType: qrTokenInfo.type,
              data: scanData.data || scanData,
            };
          } else {
            console.warn('Device scan failed:', scanResponse.status);
            return {
              type: 'assignment_failed',
              token: qrValue,
              tokenType: qrTokenInfo.type,
              data: qrTokenInfo,
            };
          }
        } catch (scanError) {
          console.error('Device scan error:', scanError);
          return {
            type: 'assignment_failed',
            token: qrValue,
            tokenType: qrTokenInfo.type,
            data: qrTokenInfo,
          };
        }
      } else {
        console.log(
          'No device MAC stored, returning validated token info only',
        );
        return {
          type: 'token_validated',
          token: qrValue,
          tokenType: qrTokenInfo.type,
          data: qrTokenInfo,
        };
      }
    }

    // For other token types (e.g., prescription_info), just return the validated info
    return {
      type: 'token_validated',
      token: qrValue,
      tokenType: qrTokenInfo.type,
      data: qrTokenInfo,
    };
  } catch (err) {
    console.error('parseQrCode error:', err);
    return null;
  }
};

// Get prescriptions for a specific patient
export const getPrescriptionsByPatientId = async patientId => {
  try {
    const userToken = await AsyncStorage.getItem('userToken');

    if (!userToken) {
      console.error('No user token available');
      return null;
    }

    console.log('📋 Getting prescriptions for patient ID:', patientId);
    const url = `${BASE_API_URL}/prescriptions`;

    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${userToken}`,
      },
    });

    if (!response.ok) {
      console.error('Prescriptions fetch failed:', response.status);
      return null;
    }

    const data = await response.json();
    console.log('📋 All prescriptions response:', data);

    // Filter prescriptions for this patient
    const prescriptions = data.data || data;
    if (Array.isArray(prescriptions)) {
      const patientPrescriptions = prescriptions.filter(
        p => p.patient_id === patientId || p.patient?.id === patientId,
      );
      console.log(
        `📋 Found ${patientPrescriptions.length} prescriptions for patient ${patientId}`,
      );
      return patientPrescriptions;
    }

    return null;
  } catch (error) {
    console.error('Error fetching prescriptions:', error);
    return null;
  }
};

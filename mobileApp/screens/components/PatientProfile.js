import { View, Text } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from '../styles/PatientProfileStyle';
import UserDataCom from './UserDataCom';

export default function PatientProfile({
  userData,
  patientInfo,
  actionButtons,
}) {
  // Veri güvenliği için hasta ve hekim objelerini ayırıyoruz
  const patient = userData?.hasta || {};
  const doctor = userData?.hekim || {};

  const formatDate = dateString => {
    if (!dateString) return 'N/A';
    const date = new Date(dateString);
    return `${date.getDate().toString().padStart(2, '0')}.${(
      date.getMonth() + 1
    )
      .toString()
      .padStart(2, '0')}.${date.getFullYear()}`;
  };

  const calculateAge = birthDate => {
    if (!birthDate) return 'N/A';
    const birth = new Date(birthDate);
    const today = new Date();
    let age = today.getFullYear() - birth.getFullYear();
    if (
      today.getMonth() < birth.getMonth() ||
      (today.getMonth() === birth.getMonth() &&
        today.getDate() < birth.getDate())
    ) {
      age--;
    }
    return age;
  };

  const isFemale = patient.cinsiyet === 'K' || patient.cinsiyet === 'Kadın';
  const genderIcon = isFemale ? 'face-woman' : 'face-man';
  const genderColor = isFemale ? '#E91E63' : '#2196F3';

  return (
    <View style={styles.profil}>
      <View style={{ flexDirection: 'row', justifyContent: 'space-between' }}>
        <View style={{ flexDirection: 'row', gap: 24, flex: 1 }}>
          <View style={styles.avatarInfo}>
            <View
              style={[
                styles.profileImage,
                {
                  backgroundColor: genderColor + '20',
                  justifyContent: 'center',
                  alignItems: 'center',
                },
              ]}
            >
              <Icon name={genderIcon} size={50} color={genderColor} />
            </View>
          </View>

          <View style={{ flex: 1 }}>
            <View style={styles.container}>
              <Text style={styles.textName}>
                {patient.ad} {patient.soyadi}
              </Text>
              <Text style={{ color: '#666' }}>
                TC: {patient.tc_kimlik_numarasi || 'Belirtilmemiş'}
              </Text>
              <Text
                style={{ color: '#4A90E2', fontSize: 12, fontWeight: 'bold' }}
              >
                Protokol No: {userData.basvuru_protokol_numarasi}
              </Text>
            </View>

            <View style={styles.row}>
              <View>
                <Text style={styles.infoText}>📅 Doğum</Text>
                <Text style={styles.infoText2}>
                  {formatDate(patient.dogum_tarihi)}
                </Text>
              </View>
              <View>
                <Text style={styles.infoText}>Yaş</Text>
                <Text style={styles.infoText2}>
                  {calculateAge(patient.dogum_tarihi)}
                </Text>
              </View>
              <View>
                <Text style={styles.infoText}>Cinsiyet</Text>
                <Text style={styles.infoText2}>{patient.cinsiyet}</Text>
              </View>
              <View>
                <Text style={styles.infoText}>Kan Grubu</Text>
                <Text style={styles.infoText2}>
                  {patient.kan_grubu || 'N/A'}
                </Text>
              </View>
            </View>
          </View>
        </View>

        {/* DOKTOR BİLGİLERİ (API'DEN GELEN) */}
        <View style={{ justifyContent: 'center', marginBottom: 70 }}>
          {doctor.ad ? (
            <UserDataCom
              title="Sorumlu Doktor"
              name={`Dr. ${doctor.ad} ${doctor.soyadi}`}
              phone={doctor.personel_kodu || 'İletişim Yok'}
              color={'#1b8b05'}
              bgColor={'#e8f5e9'}
            />
          ) : (
            <Text style={{ color: '#999', fontStyle: 'italic' }}>
              Atanmış Doktor Yok
            </Text>
          )}
        </View>
      </View>

      <View style={styles.line}></View>

      <View style={styles.profilePerson}>
        <View>
          <Text style={styles.infoText}>Hasta Tipi</Text>
          <Text style={styles.infoText2}>
            {patient.hasta_tipi || 'Yatan Hasta'}
          </Text>
        </View>
        <View>
          <Text style={styles.infoText}>Birim / Servis</Text>
          <Text style={styles.infoText2}>
            {userData.birim_kodu || 'Genel Servis'}
          </Text>
        </View>
        <View>
          <Text style={styles.infoText}>Yatış Zamanı</Text>
          <Text style={styles.infoText2}>
            {formatDate(userData.hasta_kabul_zamani)}
          </Text>
        </View>
      </View>
    </View>
  );
}

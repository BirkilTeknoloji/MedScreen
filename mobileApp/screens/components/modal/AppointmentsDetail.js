import React from 'react';
import {
  View,
  Text,
  StyleSheet,
  Modal,
  TouchableOpacity,
  ScrollView,
} from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from '../styles/DetailModalStyle';
import InfoRow from '../InfoRow';
import UserCard from '../UserCard';

const AppointmentsDetail = ({ visible, appointment, onClose }) => {
  if (!visible || !appointment) return null;

  const formatDate = dateString => {
    if (!dateString) return 'Tarih belirtilmemiş';
    return new Date(dateString).toLocaleDateString('tr-TR');
  };

  const formatTime = dateString => {
    if (!dateString) return 'Saat belirtilmemiş';
    return new Date(dateString).toLocaleTimeString('tr-TR', {
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  return (
    <Modal
      animationType="slide"
      transparent={true}
      visible={visible}
      onRequestClose={onClose}
    >
      <View style={styles.overlay}>
        <View style={styles.modalContainer}>
          <View style={styles.header}>
            <View>
              <Text style={styles.title}>Randevu Detayı</Text>
              <View style={styles.statusBadge}>
                <Text style={styles.statusText}>
                  {appointment.status || 'Planlandı'}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>
          // AppointmentsDetail.js içindeki InfoRow ve UserCard kısımlarını
          güncelleyin:
          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Başvuru Bilgileri</Text>

            <View style={{ flexDirection: 'row', flexWrap: 'wrap' }}>
              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="Kabul Tarihi"
                // API: hasta_kabul_zamani
                value={formatDate(appointment.hasta_kabul_zamani)}
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="clock-time-four-outline"
                label="Kabul Saati"
                // API: hasta_kabul_zamani
                value={formatTime(appointment.hasta_kabul_zamani)}
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="identifier"
                label="Protokol No"
                // API: basvuru_protokol_numarasi
                value={appointment.basvuru_protokol_numarasi}
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="alert-circle-outline"
                label="Hayati Tehlike"
                // API: hayati_tehlike_durumu
                value={appointment.hayati_tehlike_durumu}
              />
            </View>

            {/* Hekim Bilgisi Bölümü */}
            {appointment.hekim && (
              <View>
                <View style={styles.divider} />
                <UserCard
                  icon="doctor"
                  name={`Dr. ${appointment.hekim.ad} ${appointment.hekim.soyadi}`}
                  // API: personel_gorev_kodu (Örn: HEKIM)
                  role={appointment.hekim.personel_gorev_kodu}
                />
                <Text
                  style={{ fontSize: 12, color: '#6B7280', marginLeft: 65 }}
                >
                  Branş Kodu: {appointment.hekim.medula_brans_kodu}
                </Text>
              </View>
            )}
          </ScrollView>
          <View style={styles.footer}>
            <TouchableOpacity style={styles.cancelButton} onPress={onClose}>
              <Text style={styles.cancelButtonText}>Kapat</Text>
            </TouchableOpacity>
          </View>
        </View>
      </View>
    </Modal>
  );
};
export default AppointmentsDetail;

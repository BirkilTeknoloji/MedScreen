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
import UserCard from '../UserCard';
import InfoRow from '../InfoRow';
const SurgeryHistoryDetail = ({ visible, surgeryHistory, onClose }) => {
  if (!visible || !surgeryHistory) return null;

  const formatDate = dateString => {
    if (!dateString) return 'Tarih belirtilmemiş';
    return new Date(dateString).toLocaleDateString('tr-TR');
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
              <Text style={styles.title}>Ameliyat Geçmişi Detayı</Text>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Ameliyat Bilgileri</Text>

            <View
              style={{
                flexDirection: 'row',
                flexWrap: 'wrap',
                justifyContent: 'space-between',
              }}
            >
              <InfoRow
                style={{ width: '48%' }}
                icon="medical-bag"
                label="Ameliyat Adı"
                value={
                  surgeryHistory.tibbi_bilgi_alt_turu_kodu || 'Belirtilmemiş'
                }
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="Ameliyat Tarihi"
                value={formatDate(surgeryHistory.kayit_zamani)}
              />

              {surgeryHistory.surgeon_name && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="account-tie"
                  label="Cerrah"
                  value={surgeryHistory.surgeon_name}
                />
              )}

              {surgeryHistory.hospital && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="hospital-building"
                  label="Hastane"
                  value={surgeryHistory.hospital}
                />
              )}

              {surgeryHistory.duration && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="clock-time-four-outline"
                  label="Süre"
                  value={`${surgeryHistory.duration} dakika`}
                />
              )}

              {surgeryHistory.anesthesia_type && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="needle"
                  label="Anestezi Türü"
                  value={surgeryHistory.anesthesia_type}
                />
              )}
            </View>

            {surgeryHistory.complications && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Komplikasyonlar</Text>
                <View
                  style={[
                    styles.notesContainer,
                    { borderLeftColor: '#2563EB' },
                  ]}
                >
                  <Icon
                    name="alert-circle"
                    size={20}
                    color="#2563EB"
                    style={styles.notesIcon}
                  />
                  <Text style={styles.notesText}>
                    {surgeryHistory.complications}
                  </Text>
                </View>
              </View>
            )}

            <View style={[styles.notesSection, { width: '100%' }]}>
              <View style={styles.divider} />
              <Text style={[styles.sectionTitle, { color: '#2563EB' }]}>
                Açıklama ve Reaksiyon
              </Text>

              <View
                style={[
                  styles.notesContainer,
                  {
                    borderLeftColor: '#2563EB',
                    backgroundColor: '#e4ecffff',
                    width: '100%',
                  },
                ]}
              >
                <Icon
                  name="alert-circle-outline"
                  size={20}
                  color="#2563EB"
                  style={styles.notesIcon}
                />
                <Text style={[styles.notesText, { color: '#2563EB' }]}>
                  {surgeryHistory.aciklama || 'Belirtilmemiş'}
                </Text>
              </View>
            </View>

            {surgeryHistory.added_by_doctor && (
              <View>
                <View style={styles.divider} />
                <UserCard
                  icon="doctor"
                  name={`Dr. ${surgeryHistory.added_by_doctor.first_name} ${surgeryHistory.added_by_doctor.last_name}`}
                  role={surgeryHistory.added_by_doctor.specialization}
                />
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

export default SurgeryHistoryDetail;

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
const MedicalHistoryDetail = ({ visible, medicalHistory, onClose }) => {
  if (!visible || !medicalHistory) return null;

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
              <Text style={styles.title}>Tıbbi Geçmişi Detayı</Text>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Tıbbi Geçmiş Bilgileri</Text>

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
                label="Tanı Adı"
                value={
                  medicalHistory.tibbi_bilgi_alt_turu_kodu || 'Belirtilmemiş'
                }
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="Tanı Tarihi"
                value={formatDate(medicalHistory.kayit_zamani)}
              />

              {medicalHistory.surgeon_name && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="account-tie"
                  label="Cerrah"
                  value={surgeryHistory.surgeon_name}
                />
              )}

              {medicalHistory.hospital && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="hospital-building"
                  label="Hastane"
                  value={surgeryHistory.hospital}
                />
              )}

              {medicalHistory.duration && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="clock-time-four-outline"
                  label="Süre"
                  value={`${surgeryHistory.duration} dakika`}
                />
              )}

              {medicalHistory.anesthesia_type && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="needle"
                  label="Anestezi Türü"
                  value={surgeryHistory.anesthesia_type}
                />
              )}
            </View>

            {medicalHistory.complications && (
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
                    {medicalHistory.complications}
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
                  {medicalHistory.aciklama || 'Belirtilmemiş'}
                </Text>
              </View>
            </View>

            {medicalHistory.added_by_doctor && (
              <View>
                <View style={styles.divider} />
                <UserCard
                  icon="doctor"
                  name={`Dr. ${medicalHistory.added_by_doctor.first_name} ${medicalHistory.added_by_doctor.last_name}`}
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

export default MedicalHistoryDetail;

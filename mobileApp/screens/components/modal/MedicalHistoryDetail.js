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
              <Text style={styles.title}>Tıbbi Geçmiş Detayı</Text>
              <View style={styles.statusBadge}>
                <Text style={styles.statusText}>
                  {medicalHistory.status || 'Kayıt'}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Geçmiş Bilgileri</Text>

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
                label="Durum/Hastalık"
                value={medicalHistory.condition_name || 'Belirtilmemiş'}
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="Tanı Tarihi"
                value={formatDate(medicalHistory.diagnosed_date)}
              />

              {medicalHistory.treatment && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="pill"
                  label="Tedavi"
                  value={medicalHistory.treatment}
                />
              )}

              {medicalHistory.severity && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="alert-circle-outline"
                  label="Şiddet"
                  value={medicalHistory.severity}
                />
              )}

              {medicalHistory.outcome && (
                <InfoRow
                  style={{ width: '100%' }}
                  icon="check-circle-outline"
                  label="Sonuç"
                  value={medicalHistory.outcome}
                />
              )}
            </View>

            {medicalHistory.family_history && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Aile Geçmişi</Text>
                <View style={styles.notesContainer}>
                  <Icon
                    name="account-group"
                    size={20}
                    color="#7C3AED"
                    style={styles.notesIcon}
                  />
                  <Text style={styles.notesText}>
                    {medicalHistory.family_history}
                  </Text>
                </View>
              </View>
            )}
            {medicalHistory.notes && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Notlar</Text>
                <View style={styles.notesContainer}>
                  <Icon
                    name="note-text"
                    size={20}
                    color="#2563EB"
                    style={styles.notesIcon}
                  />
                  <Text style={styles.notesText}>{medicalHistory.notes}</Text>
                </View>
              </View>
            )}

            {medicalHistory.added_by_doctor && (
              <View>
                <View style={styles.divider} />
                <UserCard
                  icon="doctor"
                  name={`Dr. ${medicalHistory.added_by_doctor.first_name} ${medicalHistory.added_by_doctor.last_name}`}
                  role={medicalHistory.added_by_doctor.specialization}
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

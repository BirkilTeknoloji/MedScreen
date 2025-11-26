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
import styles from './styles/DetailModalStyle';
const MedicalHistoryDetail = ({ visible, medicalHistory, onClose }) => {
  if (!visible || !medicalHistory) return null;

  const formatDate = (dateString) => {
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

            <InfoRow
              icon="medical-bag"
              label="Durum/Hastalık"
              value={medicalHistory.condition_name || 'Belirtilmemiş'}
            />

            <InfoRow
              icon="calendar-blank-outline"
              label="Tanı Tarihi"
              value={formatDate(medicalHistory.diagnosed_date)}
            />

            {medicalHistory.treatment && (
              <InfoRow
                icon="pill"
                label="Tedavi"
                value={medicalHistory.treatment}
              />
            )}

            {medicalHistory.severity && (
              <InfoRow
                icon="alert-circle-outline"
                label="Şiddet"
                value={medicalHistory.severity}
              />
            )}

            {medicalHistory.outcome && (
              <InfoRow
                icon="check-circle-outline"
                label="Sonuç"
                value={medicalHistory.outcome}
              />
            )}

            {medicalHistory.family_history && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Aile Geçmişi</Text>
                <View style={styles.notesContainer}>
                  <Icon name="account-group" size={20} color="#7C3AED" style={styles.notesIcon} />
                  <Text style={styles.notesText}>{medicalHistory.family_history}</Text>
                </View>
              </View>
            )}

            {medicalHistory.notes && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Notlar</Text>
                <View style={styles.notesContainer}>
                  <Icon name="note-text" size={20} color="#2563EB" style={styles.notesIcon} />
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

const InfoRow = ({ icon, label, value, isLast }) => (
  <View style={[styles.infoRow, isLast && { marginBottom: 0 }]}>
    <View style={styles.iconContainer}>
      <Icon name={icon} size={20} color="#2563EB" />
    </View>
    <View style={styles.infoTextContainer}>
      <Text style={styles.infoLabel}>{label}</Text>
      <Text style={styles.infoValue}>{value}</Text>
    </View>
  </View>
);

const UserCard = ({ icon, name, role }) => (
  <View style={styles.userCard}>
    <View style={styles.iconAvatar}>
      <Icon name={icon} size={24} color="#2563EB" />
    </View>
    <View style={styles.userInfo}>
      <Text style={styles.userName}>{name}</Text>
      <Text style={styles.userRole}>{role}</Text>
    </View>
  </View>
);

export default MedicalHistoryDetail;
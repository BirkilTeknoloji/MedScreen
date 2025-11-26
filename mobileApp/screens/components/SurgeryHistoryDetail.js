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
const SurgeryHistoryDetail = ({ visible, surgeryHistory, onClose }) => {
  if (!visible || !surgeryHistory) return null;

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
              <Text style={styles.title}>Ameliyat Geçmişi Detayı</Text>
              <View style={styles.statusBadge}>
                <Text style={styles.statusText}>
                  Tamamlandı
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Ameliyat Bilgileri</Text>

            <InfoRow
              icon="medical-bag"
              label="Ameliyat Adı"
              value={surgeryHistory.procedure_name || 'Belirtilmemiş'}
            />

            <InfoRow
              icon="calendar-blank-outline"
              label="Ameliyat Tarihi"
              value={formatDate(surgeryHistory.surgery_date)}
            />

            {surgeryHistory.surgeon_name && (
              <InfoRow
                icon="account-tie"
                label="Cerrah"
                value={surgeryHistory.surgeon_name}
              />
            )}

            {surgeryHistory.hospital && (
              <InfoRow
                icon="hospital-building"
                label="Hastane"
                value={surgeryHistory.hospital}
              />
            )}

            {surgeryHistory.duration && (
              <InfoRow
                icon="clock-time-four-outline"
                label="Süre"
                value={`${surgeryHistory.duration} dakika`}
              />
            )}

            {surgeryHistory.anesthesia_type && (
              <InfoRow
                icon="needle"
                label="Anestezi Türü"
                value={surgeryHistory.anesthesia_type}
              />
            )}

            {surgeryHistory.complications && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Komplikasyonlar</Text>
                <View style={[styles.notesContainer, { borderLeftColor: '#DC2626' }]}>
                  <Icon name="alert-circle" size={20} color="#DC2626" style={styles.notesIcon} />
                  <Text style={styles.notesText}>{surgeryHistory.complications}</Text>
                </View>
              </View>
            )}

            {surgeryHistory.notes && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Notlar</Text>
                <View style={styles.notesContainer}>
                  <Icon name="note-text" size={20} color="#2563EB" style={styles.notesIcon} />
                  <Text style={styles.notesText}>{surgeryHistory.notes}</Text>
                </View>
              </View>
            )}

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


export default SurgeryHistoryDetail;
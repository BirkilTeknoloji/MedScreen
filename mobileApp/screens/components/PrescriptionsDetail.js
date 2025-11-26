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

const PrescriptionsDetail = ({ visible, prescription, onClose }) => {
  if (!visible || !prescription) return null;

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
              <Text style={styles.title}>İlaç Detayı</Text>
              <View style={styles.statusBadge}>
                <Text style={styles.statusText}>
                  {prescription.status || 'Aktif'}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>İlaç Bilgileri</Text>

            <InfoRow
              icon="pill"
              label="İlaç Adı"
              value={prescription.medication_name || prescription.name || 'Belirtilmemiş'}
            />

            <InfoRow
              icon="calendar-blank-outline"
              label="Reçete Tarihi"
              value={formatDate(prescription.prescribed_date || prescription.date)}
            />

            {prescription.dosage && (
              <InfoRow
                icon="medical-bag"
                label="Doz"
                value={prescription.dosage}
              />
            )}

            {prescription.frequency && (
              <InfoRow
                icon="clock-time-four-outline"
                label="Kullanım Sıklığı"
                value={prescription.frequency}
              />
            )}

            {prescription.duration && (
              <InfoRow
                icon="calendar-range"
                label="Kullanım Süresi"
                value={prescription.duration}
              />
            )}

            {prescription.instructions && (
              <InfoRow
                icon="text-subject"
                label="Kullanım Talimatları"
                value={prescription.instructions}
              />
            )}

            {prescription.notes && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Notlar</Text>
                <View style={styles.notesContainer}>
                  <Icon name="note-text" size={20} color="#2563EB" style={styles.notesIcon} />
                  <Text style={styles.notesText}>{prescription.notes}</Text>
                </View>
              </View>
            )}

            {prescription.doctor && (
              <View>
                <View style={styles.divider} />
                <UserCard
                  icon="doctor"
                  name={`Dr. ${prescription.doctor.first_name} ${prescription.doctor.last_name}`}
                  role={prescription.doctor.specialization}
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





export default PrescriptionsDetail;
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

const styles = StyleSheet.create({
  overlay: {
    flex: 1,
    backgroundColor: 'rgba(0,0,0,0.5)',
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  modalContainer: {
    width: '85%',
    maxHeight: '75%',
    backgroundColor: '#fff',
    borderRadius: 16,
    overflow: 'hidden',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.25,
    shadowRadius: 3.84,
    elevation: 5,
  },
  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    padding: 20,
    borderBottomWidth: 1,
    borderBottomColor: '#F3F4F6',
  },
  title: {
    fontSize: 20,
    fontWeight: '700',
    color: '#111827',
    marginBottom: 8,
  },
  statusBadge: {
    paddingHorizontal: 12,
    paddingVertical: 4,
    borderRadius: 20,
    backgroundColor: '#DCFCE7',
    alignSelf: 'flex-start',
  },
  statusText: {
    fontSize: 12,
    fontWeight: '600',
    color: '#166534',
  },
  closeButton: {
    padding: 4,
  },
  scrollContent: {
    padding: 20,
  },
  sectionTitle: {
    fontSize: 16,
    fontWeight: '600',
    color: '#374151',
    marginBottom: 16,
  },
  infoRow: {
    flexDirection: 'row',
    marginBottom: 20,
    alignItems: 'flex-start',
  },
  iconContainer: {
    width: 32,
    alignItems: 'center',
    marginRight: 12,
    marginTop: 2,
  },
  infoTextContainer: {
    flex: 1,
  },
  infoLabel: {
    fontSize: 13,
    color: '#6B7280',
    marginBottom: 4,
  },
  infoValue: {
    fontSize: 15,
    color: '#1F2937',
    lineHeight: 22,
  },
  notesSection: {
    marginTop: 10,
  },
  notesContainer: {
    flexDirection: 'row',
    backgroundColor: '#F9FAFB',
    padding: 12,
    borderRadius: 8,
    borderLeftWidth: 4,
    borderLeftColor: '#2563EB',
  },
  notesIcon: {
    marginRight: 8,
    marginTop: 2,
  },
  notesText: {
    flex: 1,
    fontSize: 14,
    color: '#374151',
    lineHeight: 20,
  },
  divider: {
    height: 1,
    backgroundColor: '#E5E7EB',
    marginVertical: 20,
  },
  userCard: {
    flexDirection: 'row',
    borderWidth: 1,
    borderColor: '#E5E7EB',
    borderRadius: 12,
    padding: 12,
    alignItems: 'flex-start',
  },
  iconAvatar: {
    width: 48,
    height: 48,
    borderRadius: 24,
    backgroundColor: '#EBF4FF',
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: 12,
  },
  userInfo: {
    flex: 1,
  },
  userName: {
    fontSize: 16,
    fontWeight: '600',
    color: '#111827',
  },
  userRole: {
    fontSize: 13,
    color: '#6B7280',
  },
  footer: {
    padding: 16,
    borderTopWidth: 1,
    borderTopColor: '#F3F4F6',
    backgroundColor: '#F9FAFB',
  },
  cancelButton: {
    alignSelf: 'center',
    padding: 10,
  },
  cancelButtonText: {
    color: '#DC2626',
    fontWeight: '600',
    fontSize: 14,
  },
});

export default SurgeryHistoryDetail;
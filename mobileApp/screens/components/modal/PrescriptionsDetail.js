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

const PrescriptionsDetail = ({ visible, prescription, onClose }) => {
  if (!visible || !prescription) return null;

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

            <View
              style={{
                flexDirection: 'row',
                flexWrap: 'wrap',
                justifyContent: 'space-between',
              }}
            >
              <InfoRow
                style={{ width: '48%' }}
                icon="pill"
                label="İlaç Adı"
                value={
                  prescription.medication_name ||
                  prescription.name ||
                  'Belirtilmemiş'
                }
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="Reçete Tarihi"
                value={formatDate(
                  prescription.prescribed_date || prescription.date,
                )}
              />

              {prescription.dosage && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="medical-bag"
                  label="Doz"
                  value={prescription.dosage}
                />
              )}

              {prescription.frequency && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="clock-time-four-outline"
                  label="Kullanım Sıklığı"
                  value={prescription.frequency}
                />
              )}

              {prescription.duration && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="calendar-range"
                  label="Kullanım Süresi"
                  value={prescription.duration}
                />
              )}

              {prescription.instructions && (
                <InfoRow
                  style={{ width: '100%' }}
                  icon="text-subject"
                  label="Kullanım Talimatları"
                  value={prescription.instructions}
                />
              )}
            </View>

            {prescription.notes && (
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

export default PrescriptionsDetail;

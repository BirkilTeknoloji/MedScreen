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

const MedicalTestsDetail = ({ visible, medicalTest, onClose }) => {
  if (!visible || !medicalTest) return null;

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
              <Text style={styles.title}>Tetkik Detayı</Text>
              <View style={styles.statusBadge}>
                <Text style={styles.statusText}>
                  {medicalTest.status || 'Planlandı'}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Tetkik Bilgileri</Text>

            <View
              style={{
                flexDirection: 'row',
                flexWrap: 'wrap',
                justifyContent: 'space-between',
              }}
            >
              <InfoRow
                style={{ width: '48%' }}
                icon="test-tube"
                label="Tetkik Adı"
                value={medicalTest.test_name || 'Belirtilmemiş'}
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="medical-bag"
                label="Tetkik Türü"
                value={medicalTest.test_type || 'Belirtilmemiş'}
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="İstenme Tarihi"
                value={formatDate(medicalTest.ordered_date)}
              />

              {medicalTest.scheduled_date && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="calendar-clock"
                  label="Planlanan Tarih"
                  value={formatDate(medicalTest.scheduled_date)}
                />
              )}

              {medicalTest.completed_date && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="calendar-check"
                  label="Tamamlanma Tarihi"
                  value={formatDate(medicalTest.completed_date)}
                />
              )}

              {medicalTest.lab_name && (
                <InfoRow
                  style={{ width: '100%' }}
                  icon="hospital-building"
                  label="Laboratuvar"
                  value={medicalTest.lab_name}
                />
              )}
            </View>

            {medicalTest.results && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Sonuçlar</Text>
                <View style={styles.notesContainer}>
                  <Icon
                    name="clipboard-text"
                    size={20}
                    color="#059669"
                    style={styles.notesIcon}
                  />
                  <Text style={styles.notesText}>{medicalTest.results}</Text>
                </View>
              </View>
            )}

            {medicalTest.notes && (
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
                  <Text style={styles.notesText}>{medicalTest.notes}</Text>
                </View>
              </View>
            )}

            {medicalTest.ordered_by_doctor && (
              <View>
                <View style={styles.divider} />
                <UserCard
                  icon="doctor"
                  name={`Dr. ${medicalTest.ordered_by_doctor.first_name} ${medicalTest.ordered_by_doctor.last_name}`}
                  role={medicalTest.ordered_by_doctor.specialization}
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
export default MedicalTestsDetail;

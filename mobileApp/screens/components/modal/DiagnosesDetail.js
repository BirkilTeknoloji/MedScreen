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
const DiagnosesDetail = ({ visible, diagnosis, onClose }) => {
  if (!visible || !diagnosis) return null;

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
              <Text style={styles.title}>Tanı Detayı</Text>
              <View style={styles.statusBadge}>
                <Text style={styles.statusText}>
                  {diagnosis.status || 'Aktif'}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Tanı Bilgileri</Text>

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
                  diagnosis.diagnosis_name || diagnosis.name || 'Belirtilmemiş'
                }
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="Tanı Tarihi"
                value={formatDate(diagnosis.diagnosis_date || diagnosis.date)}
              />

              {diagnosis.icd_code && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="barcode"
                  label="ICD Kodu"
                  value={diagnosis.icd_code}
                />
              )}

              {diagnosis.severity && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="alert-circle-outline"
                  label="Şiddet"
                  value={diagnosis.severity}
                />
              )}

              {diagnosis.description && (
                <InfoRow
                  style={{ width: '100%' }}
                  icon="text-subject"
                  label="Açıklama"
                  value={diagnosis.description}
                />
              )}
            </View>

            {diagnosis.notes && (
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
                  <Text style={styles.notesText}>{diagnosis.notes}</Text>
                </View>
              </View>
            )}

            {diagnosis.doctor && (
              <View>
                <View style={styles.divider} />
                <UserCard
                  icon="doctor"
                  name={`Dr. ${diagnosis.doctor.first_name} ${diagnosis.doctor.last_name}`}
                  role={diagnosis.doctor.specialization}
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

export default DiagnosesDetail;

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
const DiagnosesDetail = ({ visible, diagnosis, onClose }) => {
  if (!visible || !diagnosis) return null;

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

            <InfoRow
              icon="medical-bag"
              label="Tanı Adı"
              value={diagnosis.diagnosis_name || diagnosis.name || 'Belirtilmemiş'}
            />

            <InfoRow
              icon="calendar-blank-outline"
              label="Tanı Tarihi"
              value={formatDate(diagnosis.diagnosis_date || diagnosis.date)}
            />

            {diagnosis.icd_code && (
              <InfoRow
                icon="barcode"
                label="ICD Kodu"
                value={diagnosis.icd_code}
              />
            )}

            {diagnosis.description && (
              <InfoRow
                icon="text-subject"
                label="Açıklama"
                value={diagnosis.description}
              />
            )}

            {diagnosis.severity && (
              <InfoRow
                icon="alert-circle-outline"
                label="Şiddet"
                value={diagnosis.severity}
              />
            )}

            {diagnosis.notes && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Notlar</Text>
                <View style={styles.notesContainer}>
                  <Icon name="note-text" size={20} color="#2563EB" style={styles.notesIcon} />
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




export default DiagnosesDetail;
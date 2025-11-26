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

const AppointmentsDetail = ({ visible, appointment, onClose }) => {
  if (!visible || !appointment) return null;

  const formatDate = (dateString) => {
    if (!dateString) return 'Tarih belirtilmemiş';
    return new Date(dateString).toLocaleDateString('tr-TR');
  };

  const formatTime = (dateString) => {
    if (!dateString) return 'Saat belirtilmemiş';
    return new Date(dateString).toLocaleTimeString('tr-TR', {hour: '2-digit', minute:'2-digit'});
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
              <Text style={styles.title}>Randevu Detayı</Text>
              <View style={styles.statusBadge}>
                <Text style={styles.statusText}>
                  {appointment.status || 'Planlandı'}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Randevu Bilgileri</Text>

            <InfoRow
              icon="calendar-blank-outline"
              label="Tarih"
              value={formatDate(appointment.appointment_date)}
            />

            <InfoRow
              icon="clock-time-four-outline"
              label="Saat"
              value={formatTime(appointment.appointment_date)}
            />

            {appointment.duration_minutes && (
              <InfoRow
                icon="timer-outline"
                label="Süre"
                value={`${appointment.duration_minutes} dakika`}
              />
            )}

            {appointment.appointment_type && (
              <InfoRow
                icon="medical-bag"
                label="Randevu Türü"
                value={appointment.appointment_type}
              />
            )}

            <InfoRow
              icon="text-subject"
              label="Sebep"
              value={appointment.reason || 'Belirtilmemiş'}
            />

            {appointment.notes && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Notlar</Text>
                <View style={styles.notesContainer}>
                  <Icon name="note-text" size={20} color="#2563EB" style={styles.notesIcon} />
                  <Text style={styles.notesText}>{appointment.notes}</Text>
                </View>
              </View>
            )}

            {appointment.doctor && (
              <View>
                <View style={styles.divider} />
                <UserCard
                  icon="doctor"
                  name={`Dr. ${appointment.doctor.first_name} ${appointment.doctor.last_name}`}
                  role={appointment.doctor.specialization}
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




export default AppointmentsDetail;
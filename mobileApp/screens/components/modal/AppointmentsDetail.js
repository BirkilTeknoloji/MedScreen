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
import InfoRow from '../InfoRow';
import UserCard from '../UserCard';

const AppointmentsDetail = ({ visible, appointment, onClose }) => {
  if (!visible || !appointment) return null;

  const formatDate = dateString => {
    if (!dateString) return 'Tarih belirtilmemiş';
    return new Date(dateString).toLocaleDateString('tr-TR');
  };

  const formatTime = dateString => {
    if (!dateString) return 'Saat belirtilmemiş';
    return new Date(dateString).toLocaleTimeString('tr-TR', {
      hour: '2-digit',
      minute: '2-digit',
    });
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

            <View
              style={{
                flexDirection: 'row',
                flexWrap: 'wrap',
              }}
            >
              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="Tarih"
                value={formatDate(appointment.appointment_date)}
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="clock-time-four-outline"
                label="Saat"
                value={formatTime(appointment.appointment_date)}
              />

              {appointment.duration_minutes && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="timer-outline"
                  label="Süre"
                  value={`${appointment.duration_minutes} dakika`}
                />
              )}

              {appointment.appointment_type && (
                <InfoRow
                  style={{ width: '48%' }}
                  icon="medical-bag"
                  label="Randevu Türü"
                  value={appointment.appointment_type}
                />
              )}

              <InfoRow
                style={{ width: '100%' }}
                icon="text-subject"
                label="Sebep"
                value={appointment.reason || 'Belirtilmemiş'}
              />
            </View>

            {appointment.notes && (
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
export default AppointmentsDetail;

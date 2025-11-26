import React from 'react';
import { View, Text } from 'react-native';
import styles from './styles/DetailModalStyle';

const AppointmentDetail = ({ appointment }) => {
  const formatDate = (dateString) => {
    if (!dateString) return 'Tarih belirtilmemiş';
    return new Date(dateString).toLocaleDateString('tr-TR');
  };

  return (
    <View>
      {appointment?.appointment_date && (
        <View style={styles.detailRow}>
          <Text style={styles.detailLabel}>Randevu Tarihi:</Text>
          <Text style={styles.detailValue}>{formatDate(appointment.appointment_date)}</Text>
        </View>
      )}
      {appointment?.doctor && (
        <View style={styles.detailRow}>
          <Text style={styles.detailLabel}>Doktor:</Text>
          <Text style={styles.detailValue}>{appointment.doctor.first_name} {appointment.doctor.last_name}</Text>
        </View>
      )}
      {appointment?.doctor?.specialization && (
        <View style={styles.detailRow}>
          <Text style={styles.detailLabel}>Bölüm:</Text>
          <Text style={styles.detailValue}>{appointment.doctor.specialization}</Text>
        </View>
      )}
      {appointment?.reason && (
        <View style={styles.detailRow}>
          <Text style={styles.detailLabel}>Şikayet:</Text>
          <Text style={styles.detailValue}>{appointment.reason}</Text>
        </View>
      )}
      {appointment?.appointment_type && (
        <View style={styles.detailRow}>
          <Text style={styles.detailLabel}>Randevu Tipi:</Text>
          <Text style={styles.detailValue}>{appointment.appointment_type}</Text>
        </View>
      )}
      {appointment?.duration_minutes && (
        <View style={styles.detailRow}>
          <Text style={styles.detailLabel}>Süre:</Text>
          <Text style={styles.detailValue}>{appointment.duration_minutes} dakika</Text>
        </View>
      )}
      {appointment?.status && (
        <View style={styles.detailRow}>
          <Text style={styles.detailLabel}>Durum:</Text>
          <Text style={styles.detailValue}>{appointment.status}</Text>
        </View>
      )}
    </View>
  );
};

export default AppointmentDetail;
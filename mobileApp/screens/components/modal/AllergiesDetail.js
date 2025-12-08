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

const AllergiesDetail = ({ visible, allergy, onClose }) => {
  if (!visible || !allergy) return null;

  const formatDate = dateString => {
    if (!dateString) return 'Tarih belirtilmemiş';
    return new Date(dateString).toLocaleDateString('tr-TR');
  };

  // Şiddete göre renk ve metin ayarları
  const getSeverityConfig = severity => {
    switch (severity) {
      case 'severe':
        return { color: '#DC2626', bg: '#FEE2E2', text: 'Şiddetli' }; // Kırmızı
      case 'moderate':
        return { color: '#D97706', bg: '#FEF3C7', text: 'Orta' }; // Turuncu
      case 'mild':
        return { color: '#059669', bg: '#D1FAE5', text: 'Hafif' }; // Yeşil
      default:
        return {
          color: '#4B5563',
          bg: '#F3F4F6',
          text: severity || 'Belirtilmemiş',
        }; // Gri
    }
  };

  const getAllergyTypeText = type => {
    switch (type) {
      case 'medication':
        return 'İlaç Alerjisi';
      case 'food':
        return 'Gıda Alerjisi';
      case 'environmental':
        return 'Çevresel Alerji';
      default:
        return type;
    }
  };

  const severityInfo = getSeverityConfig(allergy.severity);

  return (
    <Modal
      animationType="slide"
      transparent={true}
      visible={visible}
      onRequestClose={onClose}
    >
      <View style={styles.overlay}>
        <View style={styles.modalContainer}>
          {/* --- HEADER --- */}
          <View style={styles.header}>
            <View style={{ flex: 1 }}>
              <Text style={styles.title}>{allergy.allergen}</Text>
              <View
                style={[
                  styles.statusBadge,
                  { backgroundColor: severityInfo.bg },
                ]}
              >
                <Text
                  style={[styles.statusText, { color: severityInfo.color }]}
                >
                  {severityInfo.text}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Alerji Bilgileri</Text>

            <View
              style={{
                flexDirection: 'row',
                flexWrap: 'wrap',
                justifyContent: 'space-between',
              }}
            >
              <InfoRow
                style={{ width: '48%' }}
                icon="tag-outline"
                label="Alerji Türü"
                value={getAllergyTypeText(allergy.allergy_type)}
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="Tanı Tarihi"
                value={formatDate(allergy.diagnosed_date)}
              />
            </View>

            {/* --- REAKSİYON (Önemli - Kırmızı Kutu) --- */}
            {allergy.reaction && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={[styles.sectionTitle, { color: '#DC2626' }]}>
                  Reaksiyon
                </Text>
                <View
                  style={[
                    styles.notesContainer,
                    { borderLeftColor: '#DC2626', backgroundColor: '#FEF2F2' },
                  ]}
                >
                  <Icon
                    name="alert-circle-outline"
                    size={20}
                    color="#DC2626"
                    style={styles.notesIcon}
                  />
                  <Text style={[styles.notesText, { color: '#7F1D1D' }]}>
                    {allergy.reaction}
                  </Text>
                </View>
              </View>
            )}

            {/* --- NOTLAR (Mavi Kutu) --- */}
            {allergy.notes && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Notlar</Text>
                <View style={styles.notesContainer}>
                  <Icon
                    name="note-text-outline"
                    size={20}
                    color="#2563EB"
                    style={styles.notesIcon}
                  />
                  <Text style={styles.notesText}>{allergy.notes}</Text>
                </View>
              </View>
            )}

            {/* --- DOKTOR KARTI --- */}
            {allergy.added_by_doctor && (
              <View>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Ekleyen Doktor</Text>
                <UserCard
                  icon="doctor"
                  name={`Dr. ${allergy.added_by_doctor.first_name} ${allergy.added_by_doctor.last_name}`}
                  role="Doktor"
                />
              </View>
            )}
          </ScrollView>

          {/* --- FOOTER --- */}
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

export default AllergiesDetail;

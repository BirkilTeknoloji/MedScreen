import React from 'react';
import { View, Text, Modal, TouchableOpacity, ScrollView } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from '../styles/DetailModalStyle';
import UserCard from '../UserCard';
import InfoRow from '../InfoRow';

const DiagnosesDetail = ({ visible, diagnosis, onClose }) => {
  if (!visible || !diagnosis) return null;

  const formatDate = dateString => {
    if (!dateString) return 'Tarih belirtilmemiş';
    return new Date(dateString).toLocaleDateString('tr-TR', {
      day: '2-digit',
      month: 'long',
      year: 'numeric',
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
              <Text style={styles.title}>Tanı Detayı</Text>
              <View
                style={[
                  styles.statusBadge,
                  {
                    backgroundColor:
                      diagnosis.birincil_tani === 1 ? '#EBF5FF' : '#F3F4F6',
                  },
                ]}
              >
                <Text
                  style={[
                    styles.statusText,
                    {
                      color:
                        diagnosis.birincil_tani === 1 ? '#2563EB' : '#4B5563',
                    },
                  ]}
                >
                  {diagnosis.birincil_tani === 1
                    ? 'Birincil Tanı'
                    : 'Destekleyici Tanı'}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            <Text style={styles.sectionTitle}>Tıbbi Bilgiler</Text>

            <View
              style={{
                flexDirection: 'row',
                flexWrap: 'wrap',
                justifyContent: 'space-between',
              }}
            >
              <InfoRow
                style={{ width: '48%' }}
                icon="barcode"
                label="Tanı Kodu"
                value={diagnosis.tani_kodu || 'N/A'}
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="tag-outline"
                label="Tanı Türü"
                value={
                  diagnosis.tani_turu === 'ON_TANI' ? 'Ön Tanı' : 'Kesin Tanı'
                }
              />

              <InfoRow
                style={{ width: '100%' }}
                icon="calendar-clock"
                label="Tanı Konulma Zamanı"
                value={formatDate(diagnosis.tani_zamani)}
              />

              <InfoRow
                style={{ width: '100%' }}
                icon="file-document-outline"
                label="Başvuru Protokol No"
                value={
                  diagnosis.hasta_basvuru?.basvuru_protokol_numarasi || 'N/A'
                }
              />
            </View>

            {/* Hekim Bilgisi */}
            {diagnosis.hekim && (
              <View>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Tanıyı Koyan Hekim</Text>
                <UserCard
                  icon="doctor"
                  name={`Dr. ${diagnosis.hekim.ad} ${diagnosis.hekim.soyadi}`}
                  role={diagnosis.hekim.personel_gorev_kodu || 'Hekim'}
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

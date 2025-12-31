import React from 'react';
import { View, Text, Modal, TouchableOpacity, ScrollView } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from '../styles/DetailModalStyle';
import UserCard from '../UserCard';
import InfoRow from '../InfoRow';

const PrescriptionsDetail = ({ visible, prescription, onClose }) => {
  if (!visible || !prescription) return null;

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
              <Text style={styles.title}>Reçete Detayı</Text>
              <View style={styles.statusBadge}>
                <Text style={styles.statusText}>
                  {prescription.recete_turu_kodu || 'NORMAL'}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            {/* --- Reçete Genel Bilgileri --- */}
            <Text style={styles.sectionTitle}>Genel Bilgiler</Text>
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
                label="Reçete Kodu"
                value={prescription.recete_kodu}
              />
              <InfoRow
                style={{ width: '48%' }}
                icon="file-document-edit-outline"
                label="E-Reçete No"
                value={prescription.medula_e_recete_numarasi || 'N/A'}
              />
              <InfoRow
                style={{ width: '100%' }}
                icon="calendar-clock"
                label="Reçete Zamanı"
                value={formatDate(prescription.recete_zamani)}
              />
            </View>

            <View style={styles.divider} />

            {/* --- İlaç Listesi (Yeni Eklenen Kısım) --- */}
            <Text style={styles.sectionTitle}>
              İlaçlar ({prescription.ilaclar?.length || 0})
            </Text>
            {prescription.ilaclar &&
              prescription.ilaclar.map((ilac, index) => (
                <View
                  key={index}
                  style={{
                    backgroundColor: '#F9FAFB',
                    padding: 12,
                    borderRadius: 8,
                    marginBottom: 10,
                    borderLeftWidth: 4,
                    borderLeftColor: '#2563EB',
                  }}
                >
                  <Text
                    style={{
                      fontWeight: 'bold',
                      fontSize: 16,
                      color: '#111827',
                      marginBottom: 5,
                    }}
                  >
                    {ilac.ilac_adi}
                  </Text>
                  <View style={{ flexDirection: 'row', flexWrap: 'wrap' }}>
                    <Text
                      style={{
                        fontSize: 13,
                        color: '#4B5563',
                        marginRight: 15,
                      }}
                    >
                      💉 Doz: {ilac.ilac_kullanim_dozu} {ilac.doz_birim}
                    </Text>
                    <Text
                      style={{
                        fontSize: 13,
                        color: '#4B5563',
                        marginRight: 15,
                      }}
                    >
                      🕒 Periyot: {ilac.ilac_kullanim_periyodu}{' '}
                      {ilac.ilac_kullanim_periyodu_birimi}
                    </Text>
                    <Text style={{ fontSize: 13, color: '#4B5563' }}>
                      📦 Adet: {ilac.kutu_adeti}
                    </Text>
                  </View>
                  <Text
                    style={{ fontSize: 12, color: '#6B7280', marginTop: 4 }}
                  >
                    📌 Şekil: {ilac.ilac_kullanim_sekli} | Barkod: {ilac.barkod}
                  </Text>
                </View>
              ))}

            {/* --- Hekim Bilgisi --- */}
            {prescription.hekim && (
              <View>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Yazan Hekim</Text>
                <UserCard
                  icon="doctor"
                  name={`Dr. ${prescription.hekim.ad} ${prescription.hekim.soyadi}`}
                  role={prescription.hekim.personel_gorev_kodu || 'Hekim'}
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

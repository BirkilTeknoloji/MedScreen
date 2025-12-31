import React from 'react';
import { View, Text, Modal, TouchableOpacity, ScrollView } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from '../styles/DetailModalStyle';
import InfoRow from '../InfoRow';
import UserCard from '../UserCard';

const AllergiesDetail = ({ visible, allergy, onClose }) => {
  if (!visible || !allergy) return null;

  const formatDate = dateString => {
    if (!dateString) return 'Tarih belirtilmemiş';
    return new Date(dateString).toLocaleDateString('tr-TR', {
      day: '2-digit',
      month: 'long',
      year: 'numeric',
    });
  };

  // API'den gelen veriye göre şiddet tespiti (Varsayılan olarak 'moderate' aldık)
  const getSeverityConfig = () => {
    // Eğer API'den şiddet bilgisi gelmiyorsa alerji her zaman kritiktir
    return { color: '#DC2626', bg: '#FEE2E2' };
  };

  const severityInfo = getSeverityConfig();

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
              {/* API: tibbi_bilgi_alt_turu_kodu (Örn: PENISILIN) */}
              <Text style={styles.title}>
                {allergy.tibbi_bilgi_alt_turu_kodu || 'Bilinmeyen Alerjen'}
              </Text>
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
                label="Kategori"
                // API: tibbi_bilgi_turu_kodu (Örn: ALERJI)
                value={
                  allergy.tibbi_bilgi_turu_kodu === 'ALERJI'
                    ? 'Tıbbi Alerji'
                    : 'Diğer'
                }
              />

              <InfoRow
                style={{ width: '48%' }}
                icon="calendar-blank-outline"
                label="Kayıt Tarihi"
                // API: kayit_zamani
                value={formatDate(allergy.kayit_zamani)}
              />
            </View>

            {/* --- AÇIKLAMA VE REAKSİYON (Önemli - Kırmızı Kutu) --- */}
            {allergy.aciklama && (
              <View style={styles.notesSection}>
                <View style={styles.divider} />
                <Text style={[styles.sectionTitle, { color: '#DC2626' }]}>
                  Açıklama ve Reaksiyon
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
                    {allergy.aciklama}
                  </Text>
                </View>
              </View>
            )}

            {/* --- EKLEYEN HEKİM --- */}
            {allergy.hekim && (
              <View>
                <View style={styles.divider} />
                <Text style={styles.sectionTitle}>Ekleyen Hekim</Text>
                <UserCard
                  icon="doctor"
                  name={`Dr. ${allergy.hekim.ad} ${allergy.hekim.soyadi}`}
                  role={allergy.hekim.personel_gorev_kodu || 'Hekim'}
                />
              </View>
            )}

            {/* --- HASTA BİLGİSİ (Alt Bilgi) --- */}
            <View style={{ marginTop: 20, opacity: 0.6 }}>
              <Text style={{ fontSize: 12 }}>
                Bilgi Kodu: {allergy.hasta_tibbi_bilgi_kodu}
              </Text>
            </View>
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

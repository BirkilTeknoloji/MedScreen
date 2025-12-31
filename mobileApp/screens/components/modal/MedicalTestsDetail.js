import React from 'react';
import { View, Text, Modal, TouchableOpacity, ScrollView } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from '../styles/DetailModalStyle';
import InfoRow from '../InfoRow';

const MedicalTestsDetail = ({ visible, medicalTest, onClose }) => {
  if (!visible || !medicalTest) return null;

  const formatDate = dateString => {
    if (!dateString) return 'Belirtilmemiş';
    return new Date(dateString).toLocaleDateString('tr-TR', {
      day: '2-digit',
      month: 'long',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  // --- Değer Analizi Fonksiyonu ---
  const analyzeResult = () => {
    const valueStr = medicalTest.sonuc_degeri; // Örn: "102 mg/dL"
    const rangeStr = medicalTest.kritik_deger_araligi; // Örn: "70-110 mg/dL"

    if (!valueStr || !rangeStr)
      return {
        label: 'Analiz Edilemedi',
        color: '#6B7280',
        icon: 'help-circle',
      };

    // Sayısal değerleri ayıkla
    const value = parseFloat(valueStr.replace(',', '.'));
    const ranges = rangeStr
      .split('-')
      .map(r => parseFloat(r.replace(',', '.')));

    if (isNaN(value) || ranges.length < 2)
      return { label: 'Bilinmiyor', color: '#6B7280', icon: 'help-circle' };

    const [min, max] = ranges;

    if (value < min) {
      return {
        label: 'Değer Altında (Düşük)',
        color: '#EF4444',
        icon: 'arrow-down-bold',
      };
    } else if (value > max) {
      return {
        label: 'Değer Üstünde (Yüksek)',
        color: '#EF4444',
        icon: 'arrow-up-bold',
      };
    } else {
      return {
        label: 'Değer Aralığında (Normal)',
        color: '#059669',
        icon: 'check-circle',
      };
    }
  };

  const analysis = analyzeResult();

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
              <View
                style={[
                  styles.statusBadge,
                  { backgroundColor: analysis.color + '15' },
                ]}
              >
                <Text style={[styles.statusText, { color: analysis.color }]}>
                  {analysis.label}
                </Text>
              </View>
            </View>
            <TouchableOpacity onPress={onClose} style={styles.closeButton}>
              <Icon name="close" size={24} color="#6B7280" />
            </TouchableOpacity>
          </View>

          <ScrollView contentContainerStyle={styles.scrollContent}>
            {/* --- Sonuç Paneli (Küçültülmüş ve Renklendirilmiş) --- */}
            <View
              style={{
                backgroundColor: analysis.color + '10',
                padding: 15,
                borderRadius: 12,
                borderWidth: 1,
                borderColor: analysis.color + '30',
                alignItems: 'center',
                marginBottom: 20,
              }}
            >
              <Text
                style={{ fontSize: 14, color: '#4B5563', fontWeight: '600' }}
              >
                {medicalTest.tetkik_adi}
              </Text>
              <View
                style={{
                  flexDirection: 'row',
                  alignItems: 'center',
                  marginVertical: 5,
                }}
              >
                <Icon
                  name={analysis.icon}
                  size={20}
                  color={analysis.color}
                  style={{ marginRight: 8 }}
                />
                <Text
                  style={{
                    fontSize: 24,
                    fontWeight: 'bold',
                    color: analysis.color,
                  }}
                >
                  {medicalTest.sonuc_degeri}
                </Text>
              </View>
              <Text style={{ fontSize: 13, color: '#6B7280' }}>
                Referans: {medicalTest.kritik_deger_araligi}
              </Text>
            </View>

            <Text style={styles.sectionTitle}>Tetkik Detayları</Text>
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
                label="Kod"
                value={medicalTest.tetkik_sonuc_kodu}
              />
              <InfoRow
                style={{ width: '48%' }}
                icon="clock-outline"
                label="Onay"
                value={formatDate(medicalTest.onay_zamani)}
              />
              <InfoRow
                style={{ width: '100%' }}
                icon="file-document-outline"
                label="Protokol No"
                value={medicalTest.hasta_basvuru?.basvuru_protokol_numarasi}
              />
            </View>

            <View style={styles.divider} />

            <Text style={styles.sectionTitle}>Başvuru Durumu</Text>
            <View style={{ flexDirection: 'row', flexWrap: 'wrap' }}>
              <InfoRow
                style={{ width: '100%' }}
                icon="calendar-import"
                label="Yatış Zamanı"
                value={formatDate(
                  medicalTest.hasta_basvuru?.hasta_kabul_zamani,
                )}
              />
              <InfoRow
                style={{ width: '48%' }}
                icon="hospital-marker"
                label="Birim"
                value={medicalTest.hasta_basvuru?.basvuru_durumu}
              />
              <InfoRow
                style={{ width: '48%' }}
                icon="alert-decagram"
                label="Hayati Tehlike"
                value={medicalTest.hasta_basvuru?.hayati_tehlike_durumu}
              />
            </View>
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

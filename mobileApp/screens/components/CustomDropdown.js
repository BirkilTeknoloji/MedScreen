import React, { useState, useRef } from 'react';
import {
  Animated,
  Text,
  View,
  TouchableOpacity,
  ScrollView,
} from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons';
import TitleIcon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from './styles/CustomDropwdownStyle';
import AppointmentsDetail from './modal/AppointmentsDetail';
import DiagnosesDetail from './modal/DiagnosesDetail';
import PrescriptionsDetail from './modal/PrescriptionsDetail';
import MedicalTestsDetail from './modal/MedicalTestsDetail';
import MedicalHistoryDetail from './modal/MedicalHistoryDetail';
import SurgeryHistoryDetail from './modal/SurgeryHistoryDetail';
import AllergiesDetail from './modal/AllergiesDetail';

const CustomDropdown = ({ data, title, icon }) => {
  const [dropModal, setDropModal] = useState(false);
  const [detailModal, setDetailModal] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);
  const animation = useRef(new Animated.Value(0)).current;

  // --- Tarih Formatlama ---
  const formatDate = dateString => {
    if (!dateString) return 'Tarih belirtilmemiş';
    const date = new Date(dateString);
    if (isNaN(date.getTime())) return 'Geçersiz Tarih';

    return date.toLocaleDateString('tr-TR', {
      day: '2-digit',
      month: 'long',
      year: 'numeric',
    });
  };

  const formatDateTime = dateString => {
    if (!dateString) return 'Tarih belirtilmemiş';
    const date = new Date(dateString);
    if (isNaN(date.getTime())) return 'Geçersiz Tarih';

    return date.toLocaleDateString('tr-TR', {
      day: '2-digit',
      month: 'long',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  const toggleDropdown = () => {
    const toValue = dropModal ? 0 : 1;
    Animated.timing(animation, {
      toValue,
      duration: 300,
      useNativeDriver: false,
    }).start();
    setDropModal(!dropModal);
  };

  const maxHeight = animation.interpolate({
    inputRange: [0, 1],
    outputRange: [0, 400],
  });

  const openDetailModal = item => {
    setSelectedItem(item);
    setDetailModal(true);
  };

  const closeDetailModal = () => {
    setDetailModal(false);
    setSelectedItem(null);
  };

  // --- API Response'a Göre Görüntüleme Bilgileri ---
  const getItemDisplayInfo = item => {
    if (!item) {
      return { displayTitle: 'Bilgi Yok', displayDate: '', displayStatus: '' };
    }

    switch (title) {
      case 'Randevular':
        return {
          displayTitle: item.hekim
            ? `Dr. ${item.hekim.ad} ${item.hekim.soyadi}`
            : item.basvuru_protokol_numarasi || 'Başvuru',
          displayDate: formatDateTime(item.hasta_kabul_zamani),
          displayStatus: item.basvuru_durumu || 'Beklemede',
        };

      case 'Tanılar':
        return {
          displayTitle: `${item.tani_kodu || 'Tanı Kodu Yok'}`,
          displayDate: formatDateTime(item.tani_zamani),
          displayStatus: item.tani_turu === 'ON_TANI' ? 'Ön Tanı' : 'Kesin',
        };

      case 'Reçeteler':
        return {
          displayTitle: item.recete_kodu || 'Reçete',
          displayDate: formatDate(item.kayit_zamani),
          displayStatus: item.recete_turu_kodu === 'NORMAL' ? 'Normal' : 'Acil',
        };

      case 'Tetkikler':
        return {
          displayTitle: item.tetkik_adi || item.tetkik_kodu || 'Tetkik',
          displayDate: formatDateTime(item.sonuc_zamani || item.kayit_zamani),
          displayStatus: '',
        };

      case 'Tıbbi Geçmiş':
        return {
          displayTitle:
            item.tibbi_bilgi_alt_turu_kodu || item.turu || 'Tıbbi Bilgi',
          displayDate: formatDate(item.kayit_zamani),
          displayStatus: item.hasta_tibbi_bilgi_kodu,
        };

      case 'Ameliyat Geçmişi':
        return {
          displayTitle: item.tibbi_bilgi_alt_turu_kodu || 'Ameliyat Geçmişi',
          displayDate: formatDate(item.kayit_zamani),
          displayStatus: item.hasta_tibbi_bilgi_kodu,
        };

      case 'Alerjiler':
        return {
          displayTitle: item.tibbi_bilgi_alt_turu_kodu || 'Alerji',
          displayDate: formatDate(item.kayit_zamani),
          displayStatus: item.hasta_tibbi_bilgi_kodu,
        };

      default:
        return {
          displayTitle: 'Bilgi',
          displayDate: formatDate(item.kayit_zamani),
          displayStatus: 'Bilinmiyor',
        };
    }
  };

  return (
    <>
      {/* Dropdown Başlığı */}
      <TouchableOpacity activeOpacity={0.7} onPress={toggleDropdown}>
        <View style={styles.dropdownBtn}>
          <View style={{ flexDirection: 'row', alignItems: 'center', gap: 10 }}>
            {icon && <TitleIcon name={icon} size={24} color="#5e6977" />}
            <Text style={styles.dropwdownText}>{title}</Text>
          </View>
          <Icon
            name={dropModal ? 'keyboard-arrow-up' : 'keyboard-arrow-down'}
            size={30}
            color="#a8afb4"
          />
        </View>
      </TouchableOpacity>

      {/* Dropdown İçeriği */}
      {dropModal && (
        <Animated.View style={[styles.dropdownContent, { maxHeight }]}>
          <ScrollView
            nestedScrollEnabled={true}
            showsVerticalScrollIndicator={true}
          >
            {data && data.length > 0 ? (
              data.map((item, index) => {
                const info = getItemDisplayInfo(item);
                return (
                  <TouchableOpacity
                    key={`${title}-${index}`}
                    onPress={() => openDetailModal(item)}
                    style={styles.itemContainer}
                  >
                    <View style={styles.appointmentItem}>
                      <View style={{ flex: 1, justifyContent: 'space-around' }}>
                        <Text style={styles.appointmentTitle} numberOfLines={2}>
                          {info.displayTitle}
                        </Text>
                        <Text style={styles.appointmentDate}>
                          {info.displayDate}
                        </Text>
                      </View>
                      <View style={styles.appointmentStatus}>
                        <Text style={styles.statusText} numberOfLines={1}>
                          {info.displayStatus}
                        </Text>
                      </View>
                    </View>
                  </TouchableOpacity>
                );
              })
            ) : (
              <View style={{ padding: 20, alignItems: 'center' }}>
                <Text style={styles.noDataText}>{title} bulunamadı</Text>
              </View>
            )}
          </ScrollView>
        </Animated.View>
      )}

      {/* Dinamik Modal Yönetimi */}
      {detailModal && selectedItem && (
        <>
          {title === 'Randevular' && (
            <AppointmentsDetail
              visible={detailModal}
              appointment={selectedItem}
              onClose={closeDetailModal}
            />
          )}
          {title === 'Tanılar' && (
            <DiagnosesDetail
              visible={detailModal}
              diagnosis={selectedItem}
              onClose={closeDetailModal}
            />
          )}
          {(title === 'Reçeteler' || title === 'İlaçlar') && (
            <PrescriptionsDetail
              visible={detailModal}
              prescription={selectedItem}
              onClose={closeDetailModal}
            />
          )}
          {title === 'Tetkikler' && (
            <MedicalTestsDetail
              visible={detailModal}
              medicalTest={selectedItem}
              onClose={closeDetailModal}
            />
          )}
          {title === 'Tıbbi Geçmiş' && (
            <MedicalHistoryDetail
              visible={detailModal}
              medicalHistory={selectedItem}
              onClose={closeDetailModal}
            />
          )}
          {title === 'Ameliyat Geçmişi' && (
            <SurgeryHistoryDetail
              visible={detailModal}
              surgeryHistory={selectedItem}
              onClose={closeDetailModal}
            />
          )}
          {title === 'Alerjiler' && (
            <AllergiesDetail
              visible={detailModal}
              allergy={selectedItem}
              onClose={closeDetailModal}
            />
          )}
        </>
      )}
    </>
  );
};

export default CustomDropdown;

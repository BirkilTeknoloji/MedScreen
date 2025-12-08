import React, { useState, useRef } from 'react';
import { Animated, Text, View, TouchableOpacity } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons';
import TitleIcon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from './styles/CustomDropwdownStyle';
import DetailModal from './DetailModal';
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

  const formatDate = dateString => {
    if (!dateString) return 'Tarih belirtilmemiş';
    const months = [
      'Ocak',
      'Şubat',
      'Mart',
      'Nisan',
      'Mayıs',
      'Haziran',
      'Temmuz',
      'Ağustos',
      'Eylül',
      'Ekim',
      'Kasım',
      'Aralık',
    ];
    const date = new Date(dateString);
    return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
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
    outputRange: [0, 300],
  });

  const openDetailModal = item => {
    setSelectedItem(item);
    setDetailModal(true);
  };

  const closeDetailModal = () => {
    setDetailModal(false);
    setSelectedItem(null);
  };

  return (
    <>
      <TouchableOpacity activeOpacity={0.3} onPress={toggleDropdown}>
        <View style={styles.dropdownBtn}>
          <View style={{ flexDirection: 'row', alignItems: 'center', gap: 10 }}>
            {icon && <TitleIcon name={icon} size={24} color="#5e6977" />}
            <Text style={styles.dropwdownText}>{title}</Text>
          </View>
          <Icon
            name={dropModal ? 'keyboard-arrow-up' : 'keyboard-arrow-down'}
            size={30}
            color="#a8afb4ff"
          />
        </View>
      </TouchableOpacity>

      {dropModal && (
        <Animated.View style={[styles.dropdownContent, { maxHeight }]}>
          <View>
            {data && data.length > 0 ? (
              data.map((item, index) => (
                <TouchableOpacity
                  key={index}
                  onPress={() => openDetailModal(item)}
                >
                  <View style={styles.appointmentItem}>
                    <View style={{ justifyContent: 'space-around' }}>
                      <Text style={styles.appointmentTitle}>
                        {/* Tanılar ve İlaçlar için özel görünüm */}
                        {title === 'Tanılar'
                          ? item.diagnosis_name || item.name || 'Tanı'
                          : title === 'İlaçlar'
                          ? item.medication_name || item.name || 'İlaç'
                          : title === 'Tetkikler'
                          ? item.test_name || item.name || 'Tetkik'
                          : title === 'Tıbbi Geçmiş'
                          ? item.condition_name || item.name || 'Geçmiş'
                          : title === 'Ameliyat Geçmişi'
                          ? item.procedure_name || item.name || 'Ameliyat'
                          : title === 'Alerjiler'
                          ? item.allergen || item.name || 'Alerji'
                          : item.doctor?.specialization ||
                            item.reason ||
                            'Randevu'}
                      </Text>
                      <Text style={styles.appointmentDate}>
                        {title === 'Tanılar'
                          ? formatDate(item.diagnosis_date || item.date)
                          : title === 'İlaçlar'
                          ? formatDate(item.prescribed_date || item.date)
                          : title === 'Tetkikler'
                          ? formatDate(item.ordered_date || item.date)
                          : title === 'Tıbbi Geçmiş'
                          ? formatDate(item.diagnosed_date || item.date)
                          : title === 'Ameliyat Geçmişi'
                          ? formatDate(item.surgery_date || item.date)
                          : title === 'Alerjiler'
                          ? formatDate(item.diagnosed_date || item.date)
                          : formatDate(item.appointment_date || item.date)}
                      </Text>
                    </View>
                    <View style={styles.appointmentStatus}>
                      <Text style={styles.statusText}>
                        {item.status || 'Durum Yok'}
                      </Text>
                    </View>
                  </View>
                </TouchableOpacity>
              ))
            ) : (
              <Text style={styles.noDataText}>{title} bulunamadı</Text>
            )}
          </View>
        </Animated.View>
      )}

      {title === 'Randevular' ? (
        <AppointmentsDetail
          visible={detailModal}
          appointment={selectedItem}
          onClose={closeDetailModal}
        />
      ) : title === 'Tanılar' ? (
        <DiagnosesDetail
          visible={detailModal}
          diagnosis={selectedItem}
          onClose={closeDetailModal}
        />
      ) : title === 'İlaçlar' ? (
        <PrescriptionsDetail
          visible={detailModal}
          prescription={selectedItem}
          onClose={closeDetailModal}
        />
      ) : title === 'Tetkikler' ? (
        <MedicalTestsDetail
          visible={detailModal}
          medicalTest={selectedItem}
          onClose={closeDetailModal}
        />
      ) : title === 'Tıbbi Geçmiş' ? (
        <MedicalHistoryDetail
          visible={detailModal}
          medicalHistory={selectedItem}
          onClose={closeDetailModal}
        />
      ) : title === 'Ameliyat Geçmişi' ? (
        <SurgeryHistoryDetail
          visible={detailModal}
          surgeryHistory={selectedItem}
          onClose={closeDetailModal}
        />
      ) : title === 'Alerjiler' ? (
        <AllergiesDetail
          visible={detailModal}
          allergy={selectedItem}
          onClose={closeDetailModal}
        />
      ) : (
        <DetailModal
          visible={detailModal}
          selectedItem={selectedItem}
          onClose={closeDetailModal}
        />
      )}
    </>
  );
};

export default CustomDropdown;

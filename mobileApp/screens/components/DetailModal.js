import React from 'react';
import {
  Modal,
  TouchableOpacity,
  View,
  Text,
  ScrollView,
  Image,
} from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons';
import styles from './styles/DetailModalStyle';
import AppointmentDetail from './AppointmentDetail';

const DetailModal = ({ visible, selectedItem, onClose }) => {
  return (
    <Modal
      visible={visible}
      transparent={true}
      animationType="slide"
      onRequestClose={onClose}
    >
      <TouchableOpacity
        style={styles.modalOverlay}
        activeOpacity={1}
        onPress={onClose}
      >
        <View style={styles.modalContent}>
          <View style={styles.modalHeader}>
            <Text style={styles.modalTitle}>
              {selectedItem?.appointment_date
                ? selectedItem?.doctor?.specialization || 'Randevu Detayı'
                : selectedItem?.title ||
                  selectedItem?.type ||
                  selectedItem?.name ||
                  (typeof selectedItem === 'string' ? selectedItem : 'Detay')}
            </Text>
            <TouchableOpacity
              onPress={onClose}
              hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
              style={{ padding: 5 }}
            >
              <Icon name="close" size={24} color="#666" />
            </TouchableOpacity>
          </View>

          <ScrollView style={styles.modalBody}>
            {/* Randevu Detayları */}
            {selectedItem?.appointment_date && (
              <AppointmentDetail appointment={selectedItem} />
            )}
            {/* Eski detaylar - sadece randevu değilse göster */}
            {!selectedItem?.appointment_date && selectedItem?.date && (
              <View style={styles.detailRow}>
                <Text style={styles.detailLabel}>Tarih:</Text>
                <Text style={styles.detailValue}>{selectedItem.date}</Text>
              </View>
            )}
            {selectedItem?.result &&
              (Array.isArray(selectedItem.result) ? (
                selectedItem.result[0]?.imageUrl ? (
                  selectedItem.result.map((res, idx) => (
                    <Image
                      key={idx}
                      source={{ uri: res.imageUrl }}
                      style={styles.image}
                    />
                  ))
                ) : (
                  <View style={styles.table}>
                    <View style={styles.tableHeader}>
                      <Text style={styles.tableCellHeader}>Test</Text>
                      <Text style={styles.tableCellHeader}>Değer</Text>
                      <Text style={styles.tableCellHeader}>Birim</Text>
                      <Text style={styles.tableCellHeader}>Normal</Text>
                      <Text style={styles.tableCellHeader}>Durum</Text>
                    </View>
                    {selectedItem.result.map((res, idx) => (
                      <View key={idx} style={styles.tableRow}>
                        <Text style={styles.tableCell}>{res.name}</Text>
                        <Text style={styles.tableCell}>{res.value}</Text>
                        <Text style={styles.tableCell}>{res.unit}</Text>
                        <Text style={styles.tableCell}>{res.normalRange}</Text>
                        <Text
                          style={[
                            styles.tableCell,
                            res.status === 'normal' && styles.statusNormal,
                            res.status === 'high' && styles.statusHigh,
                            res.status === 'low' && styles.statusLow,
                          ]}
                        >
                          {res.status}
                        </Text>
                      </View>
                    ))}
                  </View>
                )
              ) : (
                selectedItem.result.imageUrl && (
                  <Image
                    source={{ uri: selectedItem.result.imageUrl }}
                    style={styles.image}
                  />
                )
              ))}
            {selectedItem?.pastIllnesses && (
              <View style={styles.listContainer}>
                <Text style={styles.listTitle}>Geçmiş Hastalıklar:</Text>
                {selectedItem.pastIllnesses.map((illness, idx) => (
                  <View key={idx} style={styles.listItem}>
                    <Text style={styles.listItemText}>• {illness}</Text>
                  </View>
                ))}
              </View>
            )}
            {selectedItem?.surgeries && (
              <View style={styles.listContainer}>
                <Text style={styles.listTitle}>Ameliyat Geçmişi:</Text>
                {selectedItem.surgeries.map((surgery, idx) => (
                  <View key={idx} style={styles.listItem}>
                    <Text style={styles.listItemText}>• {surgery}</Text>
                  </View>
                ))}
              </View>
            )}
            {typeof selectedItem === 'string' && (
              <View style={styles.listContainer}>
                <Text style={styles.listItemText}>{selectedItem}</Text>
              </View>
            )}
          </ScrollView>
        </View>
      </TouchableOpacity>
    </Modal>
  );
};

export default DetailModal;

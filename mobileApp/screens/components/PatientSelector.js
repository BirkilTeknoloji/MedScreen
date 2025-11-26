import React, { useState } from 'react';
import { View, Text, TouchableOpacity, Modal, FlatList } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons';
import styles from './styles/PatientSelectorStyle';

const PatientSelector = ({ selectedPatient, onPatientSelect, patientList }) => {
  const [modalVisible, setModalVisible] = useState(false);

  const mockPatients = [
    { id: 1, name: 'Mert Kurtoglu', tcNumber: '12345678901', age: 34 },
    { id: 2, name: 'Ayşe Yılmaz', tcNumber: '98765432109', age: 28 },
    { id: 3, name: 'Mehmet Demir', tcNumber: '11223344556', age: 45 },
    { id: 4, name: 'Fatma Özkan', tcNumber: '66778899001', age: 52 },
  ];

  const patients = patientList || mockPatients;

  const handlePatientSelect = (patient) => {
    onPatientSelect(patient);
    setModalVisible(false);
  };

  const renderPatientItem = ({ item }) => (
    <TouchableOpacity 
      style={styles.patientItem}
      onPress={() => handlePatientSelect(item)}
    >
      <View style={styles.patientInfo}>
        <Text style={styles.patientName}>{item.name}</Text>
        <Text style={styles.patientDetails}>TC: {item.tcNumber} • Yaş: {item.age}</Text>
      </View>
      <Icon name="chevron-right" size={24} color="#666" />
    </TouchableOpacity>
  );

  return (
    <>
      <TouchableOpacity 
        style={styles.selectorButton}
        onPress={() => setModalVisible(true)}
      >
        <View style={styles.selectorContent}>
          <Text style={styles.selectorLabel}>Seçili Hasta:</Text>
          <Text style={styles.selectedPatientName}>
            {selectedPatient ? selectedPatient.name : 'Hasta Seçiniz'}
          </Text>
        </View>
        <Icon name="expand-more" size={24} color="#666" />
      </TouchableOpacity>

      <Modal
        visible={modalVisible}
        transparent={true}
        animationType="slide"
        onRequestClose={() => setModalVisible(false)}
      >
        <View style={styles.modalOverlay}>
          <View style={styles.modalContent}>
            <View style={styles.modalHeader}>
              <Text style={styles.modalTitle}>Hasta Seçiniz</Text>
              <TouchableOpacity onPress={() => setModalVisible(false)}>
                <Icon name="close" size={24} color="#666" />
              </TouchableOpacity>
            </View>
            
            <FlatList
              data={patients}
              renderItem={renderPatientItem}
              keyExtractor={(item) => item.id.toString()}
              style={styles.patientList}
            />
          </View>
        </View>
      </Modal>
    </>
  );
};

export default PatientSelector;
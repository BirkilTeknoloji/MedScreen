// PatientScreen.js
import React, { useState } from 'react';
import { View, StyleSheet } from 'react-native';
import PatientProfile from './components/PatientProfile';
import TabBar from './components/TabBar';
import AppointmentsTestsTab from './components/AppointmentsTestsTab';
import MedicalHistoryTab from './components/MedicalHistoryTab';
import ActionButtons from './components/ActionButtons';
import { patientData } from './data/patientData';

export default function PatientScreen({ route, navigation }) {
  // const { userData } = route.params;

  const userData = patientData.fakeUserData;

  const [activeTab, setActiveTab] = useState('randevularTetkikler');

  const tabs = [
    { key: 'randevularTetkikler', label: 'Randevular & Tetkikler' },
    { key: 'saglikGecmisi', label: 'Sağlık Geçmişi' },
  ];

  return (
    <View style={styles.container}>
      <PatientProfile userData={userData} />
      
      <TabBar 
        tabs={tabs}
        activeTab={activeTab}
        onTabPress={setActiveTab}
      />

      {activeTab === 'randevularTetkikler' && (
        <AppointmentsTestsTab data={patientData} />
      )}

      {activeTab === 'saglikGecmisi' && (
        <MedicalHistoryTab data={patientData} />
      )}

      <ActionButtons navigation={navigation} />
    </View>
  );
}

const styles = StyleSheet.create({
  container: { 
    flex: 1, 
    padding: 20 
  },
});
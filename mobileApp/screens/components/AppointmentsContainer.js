import React from 'react';
import { Text, View } from 'react-native';
import CustomDropdown from './CustomDropdown';
import styles from './styles/AppointmentsContainerStyle';

const AppointmentsContainer = ({
  userData,
  appointments,
  diagnoses,
  prescriptions,
  medicalTests,
  medicalHistory,
  surgeryHistory,
  allergies,
}) => {
  return (
    <View
      style={{
        flexDirection: 'row',
        flex: 1,
        paddingHorizontal: 10,
        gap: 10,
      }}
    >
      <View
        style={{
          ...styles.appointmentContainer,
          flex: 1,
          width: undefined,
        }}
      >
        <Text style={styles.textTitle}>Randevular & Tetkikler</Text>
        <View style={styles.line}></View>

        <CustomDropdown data={appointments} title={'Randevular'} />
        <CustomDropdown data={diagnoses} title={'Tanılar'} />
        <CustomDropdown data={prescriptions} title={'İlaçlar'} />
        <CustomDropdown data={medicalTests} title={'Tetkikler'} />
      </View>
      <View
        style={{
          ...styles.appointmentContainer,
          flex: 1,
          width: undefined,
        }}
      >
        <Text style={styles.textTitle}>Sağlık Geçmişi</Text>
        <View style={styles.line}></View>

        <CustomDropdown data={medicalHistory} title={'Tıbbi Geçmiş'} />
        <CustomDropdown
          data={surgeryHistory}
          title={'Ameliyat Geçmişi'}
        />
        <CustomDropdown data={allergies} title={'Alerjiler'} />
        <CustomDropdown data={userData.Prescriptions} title={'Reçeteler'} />
      </View>
    </View>
  );
};

export default AppointmentsContainer;

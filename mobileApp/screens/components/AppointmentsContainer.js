import React from 'react';
import { Text, View } from 'react-native';
import CustomDropdown from './CustomDropdown';
import styles from './styles/AppointmentsContainerStyle';

const AppointmentsContainer = ({ userData }) => {
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

        <CustomDropdown data={userData.Appointments} title={'Randevular'} />
        <CustomDropdown data={userData.Diagnosis} title={'Tanılar'} />
        <CustomDropdown data={userData.Prescriptions} title={'İlaçlar'} />
        <CustomDropdown data={userData.Notes} title={'Notlar'} />
        <CustomDropdown data={userData.Tests} title={'Tetkikler'} />
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

        <CustomDropdown
          data={userData.MedicalHistory?.pastIllnesses}
          title={'Tıbbi Geçmiş'}
        />
        <CustomDropdown
          data={userData.SurgeryHistory?.surgeries}
          title={'Ameliyat Geçmişi'}
        />
        <CustomDropdown data={userData.Allergies} title={'Alerjiler'} />
        <CustomDropdown data={userData.Prescriptions} title={'Reçeteler'} />
      </View>
    </View>
  );
};

export default AppointmentsContainer;

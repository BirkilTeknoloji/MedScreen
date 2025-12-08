import React from 'react';
import { Text, View } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
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
        <View
          style={{
            flexDirection: 'row',
            alignItems: 'center',
            gap: 8,
            marginBottom: 10,
          }}
        >
          <Icon name="calendar-check" size={24} color="#4A90E2" />
          <Text style={{ ...styles.textTitle, marginBottom: 0 }}>
            Randevular & Tetkikler
          </Text>
        </View>
        <View style={styles.line}></View>

        <CustomDropdown
          data={appointments}
          title={'Randevular'}
          icon={'calendar-clock'}
        />
        <CustomDropdown
          data={diagnoses}
          title={'Tanılar'}
          icon={'stethoscope'}
        />
        <CustomDropdown data={prescriptions} title={'İlaçlar'} icon={'pill'} />
        <CustomDropdown
          data={medicalTests}
          title={'Tetkikler'}
          icon={'test-tube'}
        />
      </View>
      <View
        style={{
          ...styles.appointmentContainer,
          flex: 1,
          width: undefined,
        }}
      >
        <View
          style={{
            flexDirection: 'row',
            alignItems: 'center',
            gap: 8,
            marginBottom: 10,
          }}
        >
          <Icon name="history" size={24} color="#4A90E2" />
          <Text style={{ ...styles.textTitle, marginBottom: 0 }}>
            Sağlık Geçmişi
          </Text>
        </View>
        <View style={styles.line}></View>

        <CustomDropdown
          data={medicalHistory}
          title={'Tıbbi Geçmiş'}
          icon={'history'}
        />
        <CustomDropdown
          data={surgeryHistory}
          title={'Ameliyat Geçmişi'}
          icon={'hospital-box'}
        />
        <CustomDropdown
          data={allergies}
          title={'Alerjiler'}
          icon={'alert-circle-outline'}
        />
        <CustomDropdown
          data={userData.Prescriptions}
          title={'Reçeteler'}
          icon={'prescription'}
        />
      </View>
    </View>
  );
};

export default AppointmentsContainer;

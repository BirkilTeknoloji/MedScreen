import React from 'react';
import { Text, View } from 'react-native';
import cusStyle from '../styles/PatientProfileStyle';
const UserDataCom = ({ color, bgColor, title, name, phone }) => {
  return (
    <View>
      <Text
        style={[
          cusStyle.emergencyPerson,
          {
            color: color,
            backgroundColor: bgColor,
            fontWeight: '500',
          },
        ]}
      >
       {title}
      </Text>
      <View
        style={{
          alignContent: 'center',
          marginTop: 8,
          textAlign: 'center',
        }}
      >
        <View>
          <Text style={[cusStyle.infoText, { textAlign: 'right' }]}>
            Adı Soyadı
          </Text>
          <Text style={[cusStyle.infoText2, { textAlign: 'right' }]}>
            {name}
          </Text>
        </View>
        <View>
          <Text style={[cusStyle.infoText, { textAlign: 'right' }]}>
            Telefon Numarası
          </Text>
          <Text style={[cusStyle.infoText2, { textAlign: 'right' }]}>
            {phone}
          </Text>
        </View>
      </View>
    </View>
  );
};

export default UserDataCom;

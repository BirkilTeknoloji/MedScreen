import React from 'react';
import { View } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from './styles/DetailModalStyle';

const UserCard = ({ icon, name, role, details }) => (
  <View style={styles.userCard}>
    <View style={style.iconAvatar}>
      <Icon name={icon} size={24} color="#2563EB" />
    </View>
    <View style={styles.userInfo}>
      <Text style={styles.userName}>{name}</Text>
      <Text style={styles.userRole}>{role}</Text>
      {details && (
        <View style={styles.userDetailsContainer}>
          <View style={styles.dividerLight} />
          {details.map((det, index) => (
            <Text key={index} style={styles.userDetailText}>
              <Text style={{ fontWeight: '600' }}>{det.label} </Text>
              {det.value}
            </Text>
          ))}
        </View>
      )}
    </View>
  </View>
);

export default UserCard;

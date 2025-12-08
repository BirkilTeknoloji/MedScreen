import React from 'react';
import { View, Text } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from './styles/DetailModalStyle';

const UserCard = ({ icon, name, role }) => (
  <View style={styles.userCard}>
    <View style={styles.iconAvatar}>
      <Icon name={icon} size={24} color="#2563EB" />
    </View>
    <View style={styles.userInfo}>
      <Text style={styles.userName}>{name}</Text>
      <Text style={styles.userRole}>{role}</Text>
    </View>
  </View>
);
export default UserCard;

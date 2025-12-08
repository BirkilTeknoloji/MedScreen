import React from 'react';
import { View, Text } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';
import styles from './styles/DetailModalStyle';

const InfoRow = ({ icon, label, value, isLast, style }) => (
  <View style={[styles.infoRow, isLast && { marginBottom: 0 }, style]}>
    <View style={styles.iconContainer}>
      <Icon name={icon} size={20} color="#2563EB" />
    </View>
    <View style={styles.infoTextContainer}>
      <Text style={styles.infoLabel}>{label} </Text>
      <Text style={styles.infoValue}>{value}</Text>
    </View>
  </View>
);
export default InfoRow;

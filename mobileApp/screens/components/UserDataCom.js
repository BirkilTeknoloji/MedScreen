import { View, Text, StyleSheet } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialCommunityIcons';

const UserDataCom = ({ title, name, phone, color, bgColor }) => {
  return (
    <View style={styles.container}>
      <View style={[styles.badge, { backgroundColor: bgColor }]}>
        <Text style={[styles.badgeText, { color }]}>{title}</Text>
      </View>
      <View style={styles.infoRow}>
        <Icon name="account" size={16} color="#333" />
        <Text style={styles.infoText}>{name || 'Belirtilmemiş'}</Text>
      </View>
      <View style={styles.infoRow}>
        <Icon name="phone" size={16} color="#333" />
        <Text style={styles.infoText}>{phone || 'Belirtilmemiş'}</Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    minWidth: 180,
  },
  badge: {
    paddingHorizontal: 10,
    paddingVertical: 4,
    borderRadius: 12,
    alignSelf: 'flex-start',
    marginBottom: 6,
  },
  badgeText: {
    fontSize: 11,
    fontWeight: 'bold',
  },
  infoRow: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 6,
    marginBottom: 4,
  },
  infoText: {
    fontSize: 12,
    fontWeight: '500',
    color: '#333',
  },
});

export default UserDataCom;

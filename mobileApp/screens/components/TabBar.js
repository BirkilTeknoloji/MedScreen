import { View, Text, TouchableOpacity } from 'react-native';
import styles from './styles/TabBarStyle';

export default function TabBar({ tabs, activeTab, onTabPress }) {
  return (
    <View style={styles.tabBar}>
      {tabs.map(tab => (
        <TouchableOpacity
          key={tab.key}
          style={[styles.tabButton, activeTab === tab.key && styles.tabActive]}
          onPress={() => onTabPress(tab.key)}
        >
          <Text style={styles.tabText}>{tab.label}</Text>
        </TouchableOpacity> 
      ))}
    </View>
  );
}

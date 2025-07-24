// components/TabBar.js
import React from 'react';
import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';

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

const styles = StyleSheet.create({
  tabBar: { 
    flexDirection: 'row', 
    marginBottom: 10 
  },
  tabButton: {
    flex: 1,
    paddingVertical: 10,
    backgroundColor: '#eee',
    marginHorizontal: 2,
    alignItems: 'center',
    borderRadius: 4,
  },
  tabActive: {
    backgroundColor: '#87cefa',
  },
  tabText: { 
    fontWeight: 'bold',
    fontSize: 18, 
  },
});
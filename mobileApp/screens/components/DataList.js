// components/DataList.js
import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

export default function DataList({ data, emptyMessage = "Kayıt yok" }) {
  if (data.length === 0) {
    return <Text style={styles.emptyText}>{emptyMessage}</Text>;
  }

  return (
    <View>
      {data.map(item => (
        <View key={item.id} style={styles.listItem}>
          <Text style={styles.listTitle}>{item.title}</Text>
          <Text style={styles.listDate}>{item.date}</Text>
        </View>
      ))}
    </View>
  );
}

const styles = StyleSheet.create({
  listItem: {
    paddingVertical: 5,
    borderBottomWidth: 1,
    borderColor: '#eee',
  },
  listTitle: {
    fontSize: 16,
    fontWeight: '600',
  },
  listDate: {
    fontSize: 14,
    color: '#666',
  },
  emptyText: { 
    fontStyle: 'italic', 
    color: '#999', 
    paddingLeft: 10, 
    paddingVertical: 5 
  },
});
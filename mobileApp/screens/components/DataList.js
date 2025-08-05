import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import styles from './styles/DataListStyle';

export default function DataList({ data, onItemPress, emptyMessage = "Kayıt yok" }) {
  if (!data || data.length === 0) {
    return <Text style={styles.emptyText}>{emptyMessage}</Text>;
  }

  const handleItemPress = (item) => {
    if (onItemPress) {
      onItemPress(item);
    }
  };



  return (
    <View>
      {data.map(item => (
        <TouchableOpacity
          key={item.id}
          style={styles.listItem}
          onPress={() => handleItemPress(item)}
          activeOpacity={0.7}
        >
          <View style={{ flex: 1 }}>
            <Text style={styles.listTitle}>{item.title}</Text>
            <Text style={styles.listDate}>{item.date}</Text>
          </View>
          

        </TouchableOpacity>
      ))}
    </View>
  );
}
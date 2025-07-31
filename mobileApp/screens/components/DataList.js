import { View, Text } from 'react-native';
import styles from './styles/DataListStyle';

export default function DataList({ data, emptyMessage = "Kayıt yok" }) {
  if (!data || data.length === 0) {
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


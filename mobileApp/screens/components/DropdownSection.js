import { useState } from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import styles from './styles/DropdownSectionStyle';

export default function DropdownSection({ title, children }) {
  const [open, setOpen] = useState(false);

  return (
    <View style={styles.section}>
      <TouchableOpacity 
        onPress={() => setOpen(!open)} 
        style={styles.dropdownHeader}
      >
        <Text style={styles.sectionTitle}>{title}</Text>
        <Text style={styles.dropdownIcon}>{open ? '▲' : '▼'}</Text>
      </TouchableOpacity>
      {open && (
        <View style={styles.dropdownContent}>
          {children}
        </View>
      )}
    </View>
  );
}

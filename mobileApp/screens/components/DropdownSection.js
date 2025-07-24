// components/DropdownSection.js
import React, { useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';

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

const styles = StyleSheet.create({
  section: { 
    marginBottom: 15, 
    borderBottomWidth: 1, 
    borderColor: '#ddd' 
  },
  dropdownHeader: { 
    flexDirection: 'row', 
    justifyContent: 'space-between', 
    paddingVertical: 10, 
    paddingHorizontal: 5 
  },
  sectionTitle: { 
    fontSize: 20, 
    fontWeight: 'bold' 
  },
  dropdownIcon: { 
    fontSize: 18 
  },
  dropdownContent: { 
    paddingLeft: 10, 
    paddingBottom: 10 
  },
});
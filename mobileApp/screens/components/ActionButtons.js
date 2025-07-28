// components/ActionButtons.js
import React from 'react';
import { View, Button, StyleSheet } from 'react-native';
import { useNavigation } from '@react-navigation/native';

export default function ActionButtons() {
  const navigation = useNavigation();

  return (
    <View style={styles.buttonGroup}>
      <Button 
        title="📷 QR Kod Okut" 
        onPress={() => navigation.navigate('QrScannerScreen')} 
      />
      <Button 
        title="↩️ Çıkış Yap" 
        onPress={() => navigation.navigate('HomeScreen')} 
      />
    </View>
  );
}

const styles = StyleSheet.create({
  buttonGroup: { 
    marginTop: 20, 
    flexDirection: 'row', 
    justifyContent: 'space-between' 
  },
});
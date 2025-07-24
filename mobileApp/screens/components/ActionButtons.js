// components/ActionButtons.js
import React from 'react';
import { View, Button, StyleSheet } from 'react-native';

export default function ActionButtons({ navigation }) {
  return (
    <View style={styles.buttonGroup}>
      <Button 
        title="📷 QR Kodumu Göster" 
        onPress={() => {
          // QR kod gösterme işlevi buraya eklenecek
          console.log('QR Kod göster');
        }} 
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
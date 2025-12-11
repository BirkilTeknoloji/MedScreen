import React, { useEffect, useState } from 'react';
import { View, Button, Text, TouchableOpacity } from 'react-native';
import { useNavigation } from '@react-navigation/native';
import styles from './styles/ActionButtonsStyle';
import { hasRole } from '../../services/api';

export default function ActionButtons() {
  const navigation = useNavigation();
  const [allowed, setAllowed] = useState(false);
  const [checked, setChecked] = useState(false);

  useEffect(() => {
    let mounted = true;
    (async () => {
      try {
        const ok = await hasRole(['doctor', 'admin']);
        if (mounted) setAllowed(!!ok);
      } catch (e) {
        console.error('Role check failed:', e);
      } finally {
        if (mounted) setChecked(true);
      }
    })();
    return () => {
      mounted = false;
    };
  }, []);

  const handleNavigate = screen => navigation.navigate(screen);

  // While role is being checked, show nothing (or a placeholder)
  if (!checked) return null;

  // If user is not doctor, show access denied UI
  if (!allowed) {
    return (
      <View style={[styles.buttonGroup, { alignItems: 'center', padding: 16 }]}>
        <Text
          style={{ color: '#b00020', fontWeight: 'bold', marginBottom: 12 }}
        >
          Yetkiniz yok
        </Text>
        <Text style={{ color: '#444', textAlign: 'center', marginBottom: 12 }}>
          Bu işlemi gerçekleştirmek için doktor rolüne sahip olmanız gerekir.
        </Text>
        <TouchableOpacity
          onPress={() => navigation.navigate('Home')}
          style={{
            backgroundColor: '#1976d2',
            paddingHorizontal: 16,
            paddingVertical: 10,
            borderRadius: 8,
          }}
        >
          <Text style={{ color: '#fff', fontWeight: 'bold' }}>
            Ana Sayfaya Dön
          </Text>
        </TouchableOpacity>
      </View>
    );
  }

  return (
    <View style={styles.buttonGroup}>
      <TouchableOpacity
        style={styles.button}
        onPress={() => handleNavigate('QrScannerScreen')}
      >
        <Text style={styles.buttonText}>📷 QR Kod Okut</Text>
      </TouchableOpacity>
      <TouchableOpacity
        style={styles.button}
        onPress={() => handleNavigate('Home')}
      >
        <Text style={styles.buttonText}>↩️ Çıkış Yap</Text>
      </TouchableOpacity>
    </View>
  );
}

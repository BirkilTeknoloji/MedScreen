import React, { useEffect } from 'react';
import { SafeAreaView, LogBox, Platform } from 'react-native';
import Orientation from 'react-native-orientation-locker';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { RootSiblingParent } from 'react-native-root-siblings';
import NfcManager from 'react-native-nfc-manager'; // NFC Manager Eklendi

import HomeScreen from './screens/HomeScreen';
import PatientScreen from './screens/PatientScreen';
import QrScannerScreen from './screens/QrScannerScreen';
import AddPatientScreen from './screens/AddPatientScreen';

const Stack = createNativeStackNavigator();

export default function App() {
  useEffect(() => {
    // 1. Ekran Yönlendirmesi
    Orientation.lockToLandscapeLeft();

    // 2. NFC Manager Başlatma (Hata çözümü burada)
    const initNfc = async () => {
      try {
        // Android'de activity tam hazır olmadan başlarsa hata verebilir
        // Küçük bir gecikme activity'nin hazır olmasını sağlar
        if (Platform.OS === 'android') {
          setTimeout(async () => {
            await NfcManager.start();
            console.log('✅ NFC Manager başarıyla başlatıldı');
          }, 1000);
        } else {
          await NfcManager.start();
        }
      } catch (ex) {
        console.warn('NFC Donanım hatası veya desteklenmiyor:', ex);
      }
    };

    initNfc();

    return () => {
      Orientation.unlockAllOrientations();
    };
  }, []);

  LogBox.ignoreLogs([
    'Text strings must be rendered within a <Text> component',
    'Non-serializable values were found in the navigation state' // Navigasyon uyarıları için
  ]);

  return (
    <RootSiblingParent>
      <SafeAreaView style={{ flex: 1 }}>
        <NavigationContainer>
          <Stack.Navigator 
            initialRouteName="Home" 
            screenOptions={{ 
              headerShown: false,
              animation: 'fade' // Tablet geçişleri için daha akıcı bir animasyon
            }}
          >
            <Stack.Screen name="Home" component={HomeScreen} />
            <Stack.Screen name="PatientScreen" component={PatientScreen} />
            <Stack.Screen name="AddPatientScreen" component={AddPatientScreen} />
            <Stack.Screen name="QrScannerScreen" component={QrScannerScreen} />
          </Stack.Navigator>
        </NavigationContainer>
      </SafeAreaView>
    </RootSiblingParent>
  );
}
import React, { useEffect } from 'react';
import { SafeAreaView } from 'react-native';
import Orientation from 'react-native-orientation-locker';

import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { RootSiblingParent } from 'react-native-root-siblings';

import HomeScreen from './screens/HomeScreen';
import PatientScreen from './screens/PatientScreen';
import QrScannerScreen from './screens/QrScannerScreen';
// import PersonnelScreen from './screens/PersonnelScreen';

const Stack = createNativeStackNavigator();

export default function App() {
  useEffect(() => {
    Orientation.lockToLandscapeLeft();
    return () => {
      Orientation.unlockAllOrientations();
    };
  }, []);

  return (
    <RootSiblingParent> {/* Toast mesajları için gerekli */}
      <SafeAreaView style={{ flex: 1 }}>
        <NavigationContainer>
          <Stack.Navigator initialRouteName="Home" screenOptions={{ headerShown: false }}>
            <Stack.Screen name="Home" component={HomeScreen} />
            <Stack.Screen name="PatientScreen" component={PatientScreen} />
            {/* <Stack.Screen name="PersonnelScreen" component={PersonnelScreen} /> */}
            <Stack.Screen name="QrScannerScreen" component={QrScannerScreen} />
          </Stack.Navigator>
        </NavigationContainer>
      </SafeAreaView>
    </RootSiblingParent>
  );
}

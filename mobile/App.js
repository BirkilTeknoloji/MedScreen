import React, { useEffect } from 'react';
import { SafeAreaView } from 'react-native';
import * as ScreenOrientation from 'expo-screen-orientation';
import HomeScreen from './screens/HomeScreen';

export default function App() {
  useEffect(() => {
    ScreenOrientation.lockAsync(ScreenOrientation.OrientationLock.LANDSCAPE_RIGHT);
  }, []);

  return (
    <SafeAreaView style={{ flex: 1 }}>
      <HomeScreen />
    </SafeAreaView>
  );
}

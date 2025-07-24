import React, { useState } from 'react';
import { View, Text, StyleSheet, Button, TouchableOpacity, FlatList } from 'react-native';


export default function PatientScreen({ hasta, onLogout }) {
  const [aktifSekme, setAktifSekme] = useState('randevular');

  const sekmeler = ['randevular', 'receteler', 'tetkikler'];

  const veriler = {
    randevular: [
      { id: '1', tarih: '24 Temmuz 2025', saat: '10:30', bolum: 'Kardiyoloji' },
      { id: '2', tarih: '10 Ağustos 2025', saat: '11:00', bolum: 'Nöroloji' },
    ],
    receteler: [
      { id: '1', ilac: 'Parol', tarih: '22 Temmuz 2025' },
    ],
    tetkikler: [
      { id: '1', isim: 'Kan Testi', durum: 'Bekleniyor' },
    ],
  };

  return (
    <View style={styles.container}>
      <View style={styles.profil}>
        <Text style={styles.isim}>{hasta.ad} {hasta.soyad}</Text>
        <Text>📅 Doğum: {hasta.dogumTarihi}</Text>
        <Text>🏥 Poliklinik: {hasta.poliklinik}</Text>
      </View>

      <View style={styles.tabBar}>
        {sekmeler.map(sekme => (
          <TouchableOpacity
            key={sekme}
            style={[
              styles.tabButton,
              aktifSekme === sekme && styles.tabAktif
            ]}
            onPress={() => setAktifSekme(sekme)}
          >
            <Text style={styles.tabText}>{sekme.toUpperCase()}</Text>
          </TouchableOpacity>
        ))}
      </View>

      <FlatList
        data={veriler[aktifSekme]}
        keyExtractor={item => item.id}
        renderItem={({ item }) => (
          <View style={styles.listeEleman}>
            {aktifSekme === 'randevular' && (
              <Text>{item.tarih} - {item.saat} - {item.bolum}</Text>
            )}
            {aktifSekme === 'receteler' && (
              <Text>{item.ilac} - {item.tarih}</Text>
            )}
            {aktifSekme === 'tetkikler' && (
              <Text>{item.isim} - {item.durum}</Text>
            )}
          </View>
        )}
      />

      <View style={styles.buttonGroup}>
        <Button title="📷 QR Kodumu Göster" onPress={() => {}} />
        <Button title="↩️ Çıkış Yap" onPress={onLogout} />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, padding: 20 },
  profil: { marginBottom: 20 },
  isim: { fontSize: 24, fontWeight: 'bold' },
  tabBar: { flexDirection: 'row', marginBottom: 10 },
  tabButton: {
    flex: 1,
    padding: 10,
    backgroundColor: '#eee',
    marginHorizontal: 2,
    alignItems: 'center',
  },
  tabAktif: {
    backgroundColor: '#87cefa',
  },
  tabText: { fontWeight: 'bold' },
  listeEleman: {
    padding: 10,
    borderBottomWidth: 1,
    borderColor: '#ccc',
  },
  buttonGroup: { marginTop: 20 },
});

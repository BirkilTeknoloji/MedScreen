// components/PatientProfile.js
import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

export default function PatientProfile({ userData }) {
  const formatDate = (dateString) => {
    if (!dateString) return "01.01.1990";
    const date = new Date(dateString);
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const year = date.getFullYear();
    return `${day}.${month}.${year}`;
  };

  return (
    <View style={styles.profil}>
      <Text style={styles.isim}>{userData.Name}</Text>
      
      <View style={styles.infoContainer}>
        <View style={styles.infoGrid}>
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>🆔 TC No:</Text>
            <Text style={styles.infoValue}>{userData.TCNumber}</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>📅 Doğum:</Text>
            <Text style={styles.infoValue}>{formatDate(userData.BirthDate)}</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>👤 Cinsiyet:</Text>
            <Text style={styles.infoValue}>{userData.Gender}</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>🩸 Kan Grubu:</Text>
            <Text style={styles.infoValue}>{userData.BloodType}</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>📏 Boy/Kilo:</Text>
            <Text style={styles.infoValue}>{userData.Height}cm / {userData.Weight}kg</Text>
          </View>
          
          <View style={styles.infoItem}>
            <Text style={styles.infoLabel}>📞 Telefon:</Text>
            <Text style={styles.infoValue}>{userData.Phone}</Text>
          </View>
        </View>
        
        {/* Adres ayrı bir satırda tam genişlikte */}
        <View style={styles.addressRow}>
          <Text style={styles.infoLabel}>📍 Adres:</Text>
          <Text style={styles.infoValue}>{userData.Address}</Text>
        </View>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  profil: { 
    marginBottom: 20,
    backgroundColor: '#f8f9fa',
    borderRadius: 10,
    padding: 15,
    borderWidth: 1,
    borderColor: '#e9ecef'
  },
  isim: { 
    fontSize: 36, 
    fontWeight: 'bold', 
    marginBottom: 20,
    textAlign: 'center',
    color: '#2c3e50'
  },
  infoContainer: {
    gap: 10,
  },
  infoGrid: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'space-between',
  },
  infoItem: {
    width: '32%',
    marginBottom: 10,
    flexDirection: 'row',
    alignItems: 'center',
  },
  infoLabel: {
    fontSize: 18,
    fontWeight: '600',
    color: '#495057',
    marginRight: 5,
    flexShrink: 0,
  },
  infoValue: {
    fontSize: 18,
    color: '#343a40',
    fontWeight: '500',
    flex: 1,
  },
  addressRow: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingTop: 5,
    borderTopWidth: 1,
    borderTopColor: '#e9ecef',
  },
});
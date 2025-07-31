import { StyleSheet } from 'react-native';

export default StyleSheet.create({
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
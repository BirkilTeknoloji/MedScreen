import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  buttonGroup: {
    flexDirection: 'column',
    gap: 8,
    alignItems: 'flex-end',
  },
  button: {
    backgroundColor: '#1976d2',
    paddingHorizontal: 16,
    paddingVertical: 10,
    borderRadius: 8,
    minWidth: 140,
    alignItems: 'center',
  },
  logoutButton: {
    backgroundColor: '#757575',
  },
  buttonText: {
    color: '#fff',
    fontWeight: 'bold',
    fontSize: 14,
  },
});

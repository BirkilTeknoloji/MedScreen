import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
    backgroundColor: '#e5edf5ff',
  },
  loadingText: {
    textAlign: 'center',
    fontSize: 16,
    marginTop: 50,
  },
  errorText: {
    color: 'red',
    textAlign: 'center',
    fontSize: 16,
    marginTop: 50,
  },
  noDataText: {
    textAlign: 'center',
    fontSize: 16,
    marginTop: 50,
  },

  contentRow: {
    flex: 1,
    flexDirection: 'row',
    gap: 10,
    marginTop: 20,
  },
});

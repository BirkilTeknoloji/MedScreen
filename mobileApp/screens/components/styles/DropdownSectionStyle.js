import { StyleSheet } from 'react-native';

export default StyleSheet.create({
  section: { 
    marginBottom: 15, 
    borderBottomWidth: 1, 
    borderColor: '#ddd' 
  },
  dropdownHeader: { 
    flexDirection: 'row', 
    justifyContent: 'space-between', 
    paddingVertical: 10, 
    paddingHorizontal: 5 
  },
  sectionTitle: { 
    fontSize: 20, 
    fontWeight: 'bold' 
  },
  dropdownIcon: { 
    fontSize: 18 
  },
  dropdownContent: { 
    paddingLeft: 10, 
    paddingBottom: 10 
  },
});
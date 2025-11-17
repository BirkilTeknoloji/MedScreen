import React, { useState, useRef } from 'react';
import { Animated, Text, View, TouchableOpacity } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons';
import styles from './styles/CustomDropwdownStyle';
import DetailModal from './DetailModal';

const CustomDropdown = ({ data, title }) => {
  const [dropModal, setDropModal] = useState(false);
  const [detailModal, setDetailModal] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);
  const animation = useRef(new Animated.Value(0)).current;
  console.log('CustomDropdown data:', data);
  const toggleDropdown = () => {
    const toValue = dropModal ? 0 : 1;

    Animated.timing(animation, {
      toValue,
      duration: 300,
      useNativeDriver: false,
    }).start();

    setDropModal(!dropModal);
  };

  // ✅ maxHeight interpolation eklendi
  const maxHeight = animation.interpolate({
    inputRange: [0, 1],
    outputRange: [0, 300],
  });

  const openDetailModal = item => {
    setSelectedItem(item);
    setDetailModal(true);
  };

  const closeDetailModal = () => {
    setDetailModal(false);
    setSelectedItem(null);
  };

  return (
    <>
      <TouchableOpacity activeOpacity={0.3} onPress={toggleDropdown}>
        <View style={styles.dropdownBtn}>
          <Text style={styles.dropwdownText}>{title}</Text>
          <Icon
            name={dropModal ? 'keyboard-arrow-up' : 'keyboard-arrow-down'}
            size={30}
            color="#a8afb4ff"
          />
        </View>
      </TouchableOpacity>

      {dropModal && (
        <Animated.View
          style={[
            styles.dropdownContent,
            {
              maxHeight,
            },
          ]}
        >
          <View>
            {data && data.length > 0 ? (
              data.map((item, index) => (
                <TouchableOpacity
                  key={index}
                  onPress={() => openDetailModal(item)}
                >
                  <View style={styles.appointmentItem}>
                    <View style={{ justifyContent: 'space-around' }}>
                      <Text style={styles.appointmentTitle}>
                        {item.title || item.type || item.name || item}
                      </Text>
                      <Text style={styles.appointmentDate}>{item.date}</Text>
                    </View>
                    <View style={styles.appointmentStatus}>
                      <Text style={styles.statusText}>Gerçekleşti</Text>
                    </View>
                  </View>
                </TouchableOpacity>
              ))
            ) : (
              <Text style={styles.noDataText}>{title} bulunamadı</Text>
            )}
          </View>
        </Animated.View>
      )}

      <DetailModal 
        visible={detailModal}
        selectedItem={selectedItem}
        onClose={closeDetailModal}
      />
    </>
  );
};

export default CustomDropdown;

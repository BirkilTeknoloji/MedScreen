import React, { useState } from 'react';
import { ScrollView, View, Text, Modal, TouchableOpacity, Pressable, Image } from 'react-native';
import DropdownSection from './DropdownSection';
import DataList from './DataList';

export default function AppointmentsTestsTab({ data }) {
  const [selectedItem, setSelectedItem] = useState(null);
  const [modalVisible, setModalVisible] = useState(false);
  const [modalImageUrl, setModalImageUrl] = useState(null);
  const [imageModalVisible, setImageModalVisible] = useState(false);

  const handleItemPress = (item) => {
    setSelectedItem(item);
    setModalVisible(true);
  };

  const closeModal = () => {
    setModalVisible(false);
    setSelectedItem(null);
  };

  const openImageModal = (imageUrl) => {
    setModalImageUrl(imageUrl);
    setImageModalVisible(true);
  };

  const closeImageModal = () => {
    setImageModalVisible(false);
    setModalImageUrl(null);
  };

  const capitalize = (str) => str.charAt(0).toUpperCase() + str.slice(1);

  const renderItemDetails = (item) => {
    if (!item) return null;

    return (
      <View style={{ marginTop: 16 }}>
        <Text style={{ fontSize: 16, fontWeight: 'bold', color: '#333', marginBottom: 8 }}>
          Detaylar:
        </Text>

        {Array.isArray(item.result) ? (
          item.result.map((resultItem, index) => (
            <View
              key={index}
              style={{
                marginBottom: 12,
                padding: 12,
                backgroundColor: '#f9f9f9',
                borderRadius: 8,
                borderWidth: 1,
                borderColor: '#ddd',
              }}
            >
              {Object.entries(resultItem).map(([key, value]) => {
                if (key.toLowerCase() === 'imageurl' && typeof value === 'string') {
                  return (
                    <Pressable
                      key={key}
                      onPress={() => openImageModal(value)}
                    >
                      <Image
                        source={{ uri: value }}
                        style={{ width: 150, height: 150, marginBottom: 8, borderRadius: 8 }}
                        resizeMode="contain"
                      />
                    </Pressable>
                  );
                }
                return (
                  <Text key={key} style={{ fontSize: 14, marginBottom: 4, color: '#444' }}>
                    <Text style={{ fontWeight: 'bold' }}>{capitalize(key)}: </Text>
                    {String(value)}
                  </Text>
                );
              })}
            </View>
          ))
        ) : item.result ? (
          <View
            style={{
              marginBottom: 12,
              padding: 12,
              backgroundColor: '#f9f9f9',
              borderRadius: 8,
              borderWidth: 1,
              borderColor: '#ddd',
            }}
          >
            {Object.entries(item.result).map(([key, value]) => {
              if (key.toLowerCase() === 'imageurl' && typeof value === 'string') {
                return (
                  <Pressable
                    key={key}
                    onPress={() => openImageModal(value)}
                  >
                    <Image
                      source={{ uri: value }}
                      style={{ width: 150, height: 150, marginBottom: 8, borderRadius: 8 }}
                      resizeMode="contain"
                    />
                  </Pressable>
                );
              }
              return (
                <Text key={key} style={{ fontSize: 14, marginBottom: 4, color: '#444' }}>
                  <Text style={{ fontWeight: 'bold' }}>{capitalize(key)}: </Text>
                  {String(value)}
                </Text>
              );
            })}
          </View>
        ) : (
          // Eğer result yoksa, item'ın diğer özelliklerini göster
          <View
            style={{
              marginBottom: 12,
              padding: 12,
              backgroundColor: '#f9f9f9',
              borderRadius: 8,
              borderWidth: 1,
              borderColor: '#ddd',
            }}
          >
            {Object.entries(item)
              .filter(([key]) => key !== 'id' && key !== 'title' && key !== 'date')
              .map(([key, value]) => {
                if (key.toLowerCase() === 'imageurl' && typeof value === 'string') {
                  return (
                    <Pressable
                      key={key}
                      onPress={() => openImageModal(value)}
                    >
                      <Image
                        source={{ uri: value }}
                        style={{ width: 150, height: 150, marginBottom: 8, borderRadius: 8 }}
                        resizeMode="contain"
                      />
                    </Pressable>
                  );
                }
                return (
                  <Text key={key} style={{ fontSize: 14, marginBottom: 4, color: '#444' }}>
                    <Text style={{ fontWeight: 'bold' }}>{capitalize(key)}: </Text>
                    {String(value)}
                  </Text>
                );
              })}
          </View>
        )}
      </View>
    );
  };

  return (
    <>
      <ScrollView style={{ flex: 1 }}>
        <DropdownSection title="Randevular">
          <DataList data={data.Appointments} onItemPress={handleItemPress} />
        </DropdownSection>

        <DropdownSection title="Tanılar">
          <DataList data={data.Diagnosis} onItemPress={handleItemPress} />
        </DropdownSection>

        <DropdownSection title="Reçeteler">
          <DataList data={data.Prescriptions} onItemPress={handleItemPress} />
        </DropdownSection>

        <DropdownSection title="Notler">
          <DataList data={data.Notes} onItemPress={handleItemPress} />
        </DropdownSection>

        <DropdownSection title="Tetkikler">
          <DataList data={data.Tests} onItemPress={handleItemPress} />
        </DropdownSection>
      </ScrollView>

      {/* Detay modal */}
      <Modal
        visible={modalVisible}
        transparent={true}
        animationType="slide"
        onRequestClose={closeModal}
      >
        <View style={{
          flex: 1,
          backgroundColor: 'rgba(0,0,0,0.5)',
          justifyContent: 'center',
          alignItems: 'center',
          padding: 20,
        }}>
          <View style={{
            backgroundColor: 'white',
            borderRadius: 12,
            padding: 20,
            width: '95%',
            maxHeight: '90%',
          }}>
            <ScrollView showsVerticalScrollIndicator={false}>
              {selectedItem && (
                <>
                  <Text style={{
                    fontSize: 18,
                    fontWeight: 'bold',
                    color: '#333',
                    marginBottom: 16,
                    textAlign: 'center',
                  }}>
                    {selectedItem.title}
                  </Text>
                  
                  <Text style={{
                    fontSize: 14,
                    color: '#666',
                    marginBottom: 8,
                    textAlign: 'center',
                  }}>
                    Tarih: {selectedItem.date}
                  </Text>

                  {renderItemDetails(selectedItem)}
                </>
              )}
            </ScrollView>

            <TouchableOpacity
              style={{
                backgroundColor: '#007AFF',
                padding: 12,
                borderRadius: 8,
                alignItems: 'center',
                marginTop: 16,
              }}
              onPress={closeModal}
            >
              <Text style={{ color: 'white', fontSize: 16, fontWeight: 'bold' }}>
                Kapat
              </Text>
            </TouchableOpacity>
          </View>
        </View>
      </Modal>

      {/* Resim modal */}
      <Modal
        visible={imageModalVisible}
        transparent={true}
        animationType="fade"
        onRequestClose={closeImageModal}
      >
        <Pressable
          style={{
            flex: 1,
            backgroundColor: 'rgba(0,0,0,0.8)',
            justifyContent: 'center',
            alignItems: 'center',
          }}
          onPress={closeImageModal}
        >
          <Image
            source={{ uri: modalImageUrl }}
            style={{ width: '90%', height: '90%', borderRadius: 12 }}
            resizeMode="contain"
          />
        </Pressable>
      </Modal>
    </>
  );
}
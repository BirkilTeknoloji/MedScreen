// components/MedicalHistoryTab.js
import React from 'react';
import { ScrollView } from 'react-native';
import DropdownSection from './DropdownSection';
import DataList from './DataList';

export default function MedicalHistoryTab({ data }) {
  return (
    <ScrollView style={{ flex: 1 }}>
      <DropdownSection title="Tıbbi Geçmiş">
        <DataList data={data.medicalhistory} />
      </DropdownSection>

      <DropdownSection title="Ameliyat Geçmişi">
        <DataList data={data.surgeryhistory} />
      </DropdownSection>

      <DropdownSection title="Alerjiler">
        <DataList data={data.allergies} />
      </DropdownSection>
    </ScrollView>
  );
}
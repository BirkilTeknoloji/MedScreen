// components/AppointmentsTestsTab.js
import React from 'react';
import { ScrollView } from 'react-native';
import DropdownSection from './DropdownSection';
import DataList from './DataList';

export default function AppointmentsTestsTab({ data }) {
  return (
    <ScrollView style={{ flex: 1 }}>
      <DropdownSection title="Randevular">
        <DataList data={data.appointments} />
      </DropdownSection>

      <DropdownSection title="Tanılar">
        <DataList data={data.diagnosis} />
      </DropdownSection>

      <DropdownSection title="Reçeteler">
        <DataList data={data.prescriptions} />
      </DropdownSection>

      <DropdownSection title="Notlar">
        <DataList data={data.notes} />
      </DropdownSection>

      <DropdownSection title="Tetkikler">
        <DataList data={data.tests} />
      </DropdownSection>
    </ScrollView>
  );
}
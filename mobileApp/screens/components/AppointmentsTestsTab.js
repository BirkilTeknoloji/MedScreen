import React from 'react';
import { ScrollView } from 'react-native';
import DropdownSection from './DropdownSection';
import DataList from './DataList';

export default function AppointmentsTestsTab({ data }) {
  return (
    <ScrollView style={{ flex: 1 }}>
      <DropdownSection title="Randevular">
        <DataList data={data.Appointments} />
      </DropdownSection>

      <DropdownSection title="Tanılar">
        <DataList data={data.Diagnosis} />
      </DropdownSection>

      <DropdownSection title="Reçeteler">
        <DataList data={data.Prescriptions} />
      </DropdownSection>

      <DropdownSection title="Notlar">
        <DataList data={data.Notes} />
      </DropdownSection>

      <DropdownSection title="Tetkikler">
        <DataList data={data.Tests} />
      </DropdownSection>
    </ScrollView>
  );
}
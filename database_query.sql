--MedScreen projesi veritabanı

-- Encoding ayarla
SET client_encoding = 'UTF8';

-- ============================================
-- TABLO OLUŞTURMA
-- ============================================

-- Users tablosu
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL,
    specialization VARCHAR(100) NULL,
    license_number VARCHAR(50) NULL,
    phone VARCHAR(20),
    is_active BOOLEAN DEFAULT TRUE,
    nfc_card_id VARCHAR(100) UNIQUE NULL
);

CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_nfc_card_id ON users(nfc_card_id);

-- NFC Cards tablosu
CREATE TABLE nfc_cards (
    id SERIAL PRIMARY KEY,
    card_uid VARCHAR(100) UNIQUE NOT NULL,
    assigned_user_id INTEGER NULL,
    is_active BOOLEAN DEFAULT TRUE,
    issued_at TIMESTAMP DEFAULT NOW(),
    last_used_at TIMESTAMP NULL,
    created_by_user_id INTEGER NOT NULL
);

CREATE INDEX idx_nfc_cards_card_uid ON nfc_cards(card_uid);
CREATE INDEX idx_nfc_cards_assigned_user_id ON nfc_cards(assigned_user_id);
CREATE INDEX idx_nfc_cards_is_active ON nfc_cards(is_active);

-- Patients tablosu
CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    user_id INTEGER NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    tc_number VARCHAR(11) UNIQUE NOT NULL,
    birth_date DATE NOT NULL,
    gender VARCHAR(20) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(255) NULL,
    address TEXT,
    emergency_contact_name VARCHAR(100),
    emergency_contact_phone VARCHAR(20),
    blood_type VARCHAR(5) NULL,
    height DECIMAL(5,2) NULL,
    weight DECIMAL(5,2) NULL,
    primary_doctor_id INTEGER NOT NULL
);

CREATE INDEX idx_patients_tc_number ON patients(tc_number);
CREATE INDEX idx_patients_last_name ON patients(last_name);
CREATE INDEX idx_patients_primary_doctor_id ON patients(primary_doctor_id);

-- Appointments tablosu
CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    patient_id INTEGER NOT NULL,
    doctor_id INTEGER NOT NULL,
    appointment_date TIMESTAMP NOT NULL,
    duration_minutes INTEGER DEFAULT 30,
    appointment_type VARCHAR(50) NOT NULL,
    status VARCHAR(30) NOT NULL,
    reason TEXT,
    notes TEXT,
    created_by_user_id INTEGER NOT NULL
);

CREATE INDEX idx_appointments_patient_id ON appointments(patient_id);
CREATE INDEX idx_appointments_doctor_id ON appointments(doctor_id);
CREATE INDEX idx_appointments_appointment_date ON appointments(appointment_date);
CREATE INDEX idx_appointments_status ON appointments(status);

-- Diagnoses tablosu
CREATE TABLE diagnoses (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    patient_id INTEGER NOT NULL,
    appointment_id INTEGER NULL,
    doctor_id INTEGER NOT NULL,
    diagnosis_date DATE NOT NULL,
    icd_code VARCHAR(10) NULL,
    diagnosis_name VARCHAR(255) NOT NULL,
    description TEXT,
    severity VARCHAR(20) NULL,
    status VARCHAR(30) DEFAULT 'active'
);

CREATE INDEX idx_diagnoses_patient_id ON diagnoses(patient_id);
CREATE INDEX idx_diagnoses_appointment_id ON diagnoses(appointment_id);
CREATE INDEX idx_diagnoses_doctor_id ON diagnoses(doctor_id);
CREATE INDEX idx_diagnoses_diagnosis_date ON diagnoses(diagnosis_date);
CREATE INDEX idx_diagnoses_icd_code ON diagnoses(icd_code);

-- Prescriptions tablosu
CREATE TABLE prescriptions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    patient_id INTEGER NOT NULL,
    appointment_id INTEGER NULL,
    doctor_id INTEGER NOT NULL,
    prescribed_date DATE NOT NULL,
    medication_name VARCHAR(255) NOT NULL,
    dosage VARCHAR(100) NOT NULL,
    frequency VARCHAR(100) NOT NULL,
    duration VARCHAR(100) NOT NULL,
    quantity INTEGER NOT NULL,
    refills_allowed INTEGER DEFAULT 0,
    instructions TEXT,
    status VARCHAR(30) DEFAULT 'active'
);

CREATE INDEX idx_prescriptions_patient_id ON prescriptions(patient_id);
CREATE INDEX idx_prescriptions_appointment_id ON prescriptions(appointment_id);
CREATE INDEX idx_prescriptions_doctor_id ON prescriptions(doctor_id);
CREATE INDEX idx_prescriptions_prescribed_date ON prescriptions(prescribed_date);

-- Medical Tests tablosu
CREATE TABLE medical_tests (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    patient_id INTEGER NOT NULL,
    appointment_id INTEGER NULL,
    ordered_by_doctor_id INTEGER NOT NULL,
    test_type VARCHAR(100) NOT NULL,
    test_name VARCHAR(255) NOT NULL,
    ordered_date DATE NOT NULL,
    scheduled_date DATE NULL,
    completed_date DATE NULL,
    results TEXT NULL,
    result_file_path VARCHAR(500) NULL,
    status VARCHAR(30) DEFAULT 'ordered',
    lab_name VARCHAR(255) NULL,
    notes TEXT
);

CREATE INDEX idx_medical_tests_patient_id ON medical_tests(patient_id);
CREATE INDEX idx_medical_tests_appointment_id ON medical_tests(appointment_id);
CREATE INDEX idx_medical_tests_ordered_by_doctor_id ON medical_tests(ordered_by_doctor_id);
CREATE INDEX idx_medical_tests_test_type ON medical_tests(test_type);
CREATE INDEX idx_medical_tests_status ON medical_tests(status);

-- Medical History tablosu
CREATE TABLE medical_history (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    patient_id INTEGER NOT NULL,
    condition_name VARCHAR(255) NOT NULL,
    diagnosed_date DATE NULL,
    resolved_date DATE NULL,
    status VARCHAR(30) DEFAULT 'active',
    notes TEXT,
    added_by_doctor_id INTEGER NOT NULL
);

CREATE INDEX idx_medical_history_patient_id ON medical_history(patient_id);
CREATE INDEX idx_medical_history_status ON medical_history(status);

-- Surgery History tablosu
CREATE TABLE surgery_history (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    patient_id INTEGER NOT NULL,
    procedure_name VARCHAR(255) NOT NULL,
    surgery_date DATE NOT NULL,
    surgeon_name VARCHAR(255) NULL,
    complications TEXT NULL,
    notes TEXT,
    added_by_doctor_id INTEGER NOT NULL
);

CREATE INDEX idx_surgery_history_patient_id ON surgery_history(patient_id);
CREATE INDEX idx_surgery_history_surgery_date ON surgery_history(surgery_date);

-- Allergies tablosu
CREATE TABLE allergies (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    patient_id INTEGER NOT NULL,
    allergen VARCHAR(255) NOT NULL,
    allergy_type VARCHAR(50) NOT NULL,
    reaction TEXT NOT NULL,
    severity VARCHAR(20) NOT NULL,
    diagnosed_date DATE NULL,
    notes TEXT,
    added_by_doctor_id INTEGER NOT NULL
);

CREATE INDEX idx_allergies_patient_id ON allergies(patient_id);
CREATE INDEX idx_allergies_severity ON allergies(severity);

-- Vital Signs tablosu
CREATE TABLE vital_signs (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    patient_id INTEGER NOT NULL,
    appointment_id INTEGER NULL,
    recorded_by_user_id INTEGER NOT NULL,
    recorded_at TIMESTAMP DEFAULT NOW(),
    blood_pressure_systolic INTEGER NULL,
    blood_pressure_diastolic INTEGER NULL,
    heart_rate INTEGER NULL,
    temperature DECIMAL(4,2) NULL,
    respiratory_rate INTEGER NULL,
    oxygen_saturation INTEGER NULL,
    height DECIMAL(5,2) NULL,
    weight DECIMAL(5,2) NULL,
    bmi DECIMAL(4,2) NULL,
    notes TEXT
);

CREATE INDEX idx_vital_signs_patient_id ON vital_signs(patient_id);
CREATE INDEX idx_vital_signs_appointment_id ON vital_signs(appointment_id);
CREATE INDEX idx_vital_signs_recorded_at ON vital_signs(recorded_at);

-- ============================================
-- FOREIGN KEY CONSTRAINTS
-- ============================================

ALTER TABLE nfc_cards
    ADD CONSTRAINT fk_nfc_cards_assigned_user 
    FOREIGN KEY (assigned_user_id) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE nfc_cards
    ADD CONSTRAINT fk_nfc_cards_created_by 
    FOREIGN KEY (created_by_user_id) REFERENCES users(id);

ALTER TABLE patients
    ADD CONSTRAINT fk_patients_primary_doctor 
    FOREIGN KEY (primary_doctor_id) REFERENCES users(id);

ALTER TABLE patients
    ADD CONSTRAINT fk_patients_user 
    FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE appointments
    ADD CONSTRAINT fk_appointments_patient 
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE;

ALTER TABLE appointments
    ADD CONSTRAINT fk_appointments_doctor 
    FOREIGN KEY (doctor_id) REFERENCES users(id);

ALTER TABLE appointments
    ADD CONSTRAINT fk_appointments_created_by 
    FOREIGN KEY (created_by_user_id) REFERENCES users(id);

ALTER TABLE diagnoses
    ADD CONSTRAINT fk_diagnoses_patient 
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE;

ALTER TABLE diagnoses
    ADD CONSTRAINT fk_diagnoses_appointment 
    FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON DELETE SET NULL;

ALTER TABLE diagnoses
    ADD CONSTRAINT fk_diagnoses_doctor 
    FOREIGN KEY (doctor_id) REFERENCES users(id);

ALTER TABLE prescriptions
    ADD CONSTRAINT fk_prescriptions_patient 
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE;

ALTER TABLE prescriptions
    ADD CONSTRAINT fk_prescriptions_appointment 
    FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON DELETE SET NULL;

ALTER TABLE prescriptions
    ADD CONSTRAINT fk_prescriptions_doctor 
    FOREIGN KEY (doctor_id) REFERENCES users(id);

ALTER TABLE medical_tests
    ADD CONSTRAINT fk_medical_tests_patient 
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE;

ALTER TABLE medical_tests
    ADD CONSTRAINT fk_medical_tests_appointment 
    FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON DELETE SET NULL;

ALTER TABLE medical_tests
    ADD CONSTRAINT fk_medical_tests_doctor 
    FOREIGN KEY (ordered_by_doctor_id) REFERENCES users(id);

ALTER TABLE medical_history
    ADD CONSTRAINT fk_medical_history_patient 
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE;

ALTER TABLE medical_history
    ADD CONSTRAINT fk_medical_history_doctor 
    FOREIGN KEY (added_by_doctor_id) REFERENCES users(id);

ALTER TABLE surgery_history
    ADD CONSTRAINT fk_surgery_history_patient 
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE;

ALTER TABLE surgery_history
    ADD CONSTRAINT fk_surgery_history_doctor 
    FOREIGN KEY (added_by_doctor_id) REFERENCES users(id);

ALTER TABLE allergies
    ADD CONSTRAINT fk_allergies_patient 
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE;

ALTER TABLE allergies
    ADD CONSTRAINT fk_allergies_doctor 
    FOREIGN KEY (added_by_doctor_id) REFERENCES users(id);

ALTER TABLE vital_signs
    ADD CONSTRAINT fk_vital_signs_patient 
    FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE;

ALTER TABLE vital_signs
    ADD CONSTRAINT fk_vital_signs_appointment 
    FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON DELETE SET NULL;

ALTER TABLE vital_signs
    ADD CONSTRAINT fk_vital_signs_recorded_by 
    FOREIGN KEY (recorded_by_user_id) REFERENCES users(id);

-- ============================================
-- ORNEK VERILER (Turkce karaktersiz)
-- ============================================

-- Users
INSERT INTO users (first_name, last_name, role, specialization, license_number, phone, is_active, nfc_card_id) VALUES
('Mehmet', 'Yilmaz', 'doctor', 'Kardiyoloji', 'DR123456', '05321234567', true, 'NFC001'),
('Ayse', 'Demir', 'doctor', 'Dahiliye', 'DR123457', '05321234568', true, 'NFC002'),
('Fatma', 'Kara', 'nurse', NULL, 'NS123458', '05321234569', true, 'NFC003'),
('Ali', 'Celik', 'doctor', 'Ortopedi', 'DR123459', '05321234570', true, NULL),
('Zeynep', 'Sahin', 'receptionist', NULL, NULL, '05321234571', true, 'NFC004'),
('Hasan', 'Ozturk', 'admin', NULL, NULL, '05321234572', true, 'NFC005');

-- NFC Cards
INSERT INTO nfc_cards (card_uid, assigned_user_id, is_active, created_by_user_id, last_used_at) VALUES
('NFC-CARD-001-UID', 1, true, 6, NOW() - INTERVAL '2 hours'),
('NFC-CARD-002-UID', 2, true, 6, NOW() - INTERVAL '1 day'),
('NFC-CARD-003-UID', 3, true, 6, NOW() - INTERVAL '3 hours'),
('NFC-CARD-004-UID', 5, true, 6, NOW() - INTERVAL '30 minutes'),
('NFC-CARD-005-UID', 6, true, 6, NOW() - INTERVAL '5 hours');

-- Patients
INSERT INTO patients (first_name, last_name, tc_number, birth_date, gender, phone, email, address, emergency_contact_name, emergency_contact_phone, blood_type, height, weight, primary_doctor_id) VALUES
('Ahmet', 'Yildiz', '12345678901', '1985-03-15', 'male', '05331234567', 'ahmet.yildiz@email.com', 'Cankaya Mahallesi, No:45, Ankara', 'Elif Yildiz', '05331234568', 'A+', 175.00, 80.50, 1),
('Emine', 'Aksoy', '23456789012', '1990-07-22', 'female', '05341234567', 'emine.aksoy@email.com', 'Kizilay Caddesi, No:12/3, Ankara', 'Mustafa Aksoy', '05341234568', 'O+', 165.00, 62.00, 2),
('Can', 'Oz', '34567890123', '2000-11-08', 'male', '05351234567', 'can.oz@email.com', 'Bahcelievler Sokak, No:78, Ankara', 'Ayse Oz', '05351234568', 'B+', 180.00, 75.00, 1),
('Selin', 'Aydin', '45678901234', '1978-05-30', 'female', '05361234567', 'selin.aydin@email.com', 'Ataturk Bulvari, No:156, Ankara', 'Kemal Aydin', '05361234568', 'AB+', 168.00, 70.00, 4),
('Burak', 'Koc', '56789012345', '1995-09-18', 'male', '05371234567', NULL, 'Yenimahalle, No:23, Ankara', 'Deniz Koc', '05371234568', 'O-', 172.00, 68.00, 2);

-- Appointments
INSERT INTO appointments (patient_id, doctor_id, appointment_date, duration_minutes, appointment_type, status, reason, notes, created_by_user_id) VALUES
(1, 1, '2024-11-20 10:00:00', 30, 'consultation', 'scheduled', 'Gogus agrisi sikayeti', NULL, 5),
(2, 2, '2024-11-18 14:30:00', 45, 'checkup', 'completed', 'Yillik kontrol', 'Normal muayene, kan tahlili onerildi', 5),
(3, 1, '2024-11-19 11:00:00', 30, 'follow-up', 'confirmed', 'Tansiyon kontrolu', NULL, 5),
(4, 4, '2024-11-21 09:00:00', 60, 'consultation', 'scheduled', 'Diz agrisi', NULL, 5),
(5, 2, '2024-11-17 15:00:00', 30, 'consultation', 'completed', 'Bas agrisi', 'Migran teshisi kondu', 5),
(1, 1, '2024-11-15 10:30:00', 30, 'consultation', 'completed', 'Kalp carpintisi', 'EKG normal, stres kaynakli olabilir', 5);

-- Diagnoses
INSERT INTO diagnoses (patient_id, appointment_id, doctor_id, diagnosis_date, icd_code, diagnosis_name, description, severity, status) VALUES
(1, 6, 1, '2024-11-15', 'I20', 'Angina Pectoris', 'Stres kaynakli gogus agrisi, koroner arter hastaligi suphesi', 'moderate', 'under_observation'),
(5, 5, 2, '2024-11-17', 'G43', 'Migren', 'Aura ile birlikte migren ataklari', 'moderate', 'active'),
(2, 2, 2, '2024-11-18', 'E78', 'Hiperkolesterolemi', 'Yuksek kolesterol', 'mild', 'active');

-- Prescriptions
INSERT INTO prescriptions (patient_id, appointment_id, doctor_id, prescribed_date, medication_name, dosage, frequency, duration, quantity, refills_allowed, instructions, status) VALUES
(1, 6, 1, '2024-11-15', 'Aspirin', '100mg', 'Gunde bir kez', '30 gun', 30, 2, 'Yemekten sonra aliniz', 'active'),
(5, 5, 2, '2024-11-17', 'Sumatriptan', '50mg', 'Gerektiginde (atak sirasinda)', '30 gun', 10, 0, 'Migren basladiginda aliniz', 'active'),
(2, 2, 2, '2024-11-18', 'Atorvastatin', '20mg', 'Gunde bir kez (aksam)', '90 gun', 90, 3, 'Aksam yemegindensonra aliniz', 'active'),
(3, 3, 1, '2024-11-19', 'Ramipril', '5mg', 'Gunde bir kez (sabah)', '30 gun', 30, 5, 'Sabah ac karina aliniz', 'active');

-- Medical Tests
INSERT INTO medical_tests (patient_id, appointment_id, ordered_by_doctor_id, test_type, test_name, ordered_date, scheduled_date, completed_date, results, status, lab_name, notes) VALUES
(1, 6, 1, 'blood_test', 'Lipid Profili', '2024-11-15', '2024-11-16', '2024-11-16', 'Total kolesterol: 220 mg/dL, LDL: 150 mg/dL', 'completed', 'Ankara Merkez Laboratuvar', 'Aclik kan sekeri de istendi'),
(2, 2, 2, 'blood_test', 'Tam Kan Sayimi', '2024-11-18', '2024-11-19', NULL, NULL, 'scheduled', 'Ankara Merkez Laboratuvar', NULL),
(4, 4, 4, 'x-ray', 'Diz Grafisi (AP/Lateral)', '2024-11-21', '2024-11-22', NULL, NULL, 'ordered', 'Radyoloji Merkezi', 'Her iki diz'),
(1, 6, 1, 'ecg', 'EKG', '2024-11-15', '2024-11-15', '2024-11-15', 'Sinus ritmi, normal', 'completed', NULL, 'Muayenehanede yapildi');

-- Medical History
INSERT INTO medical_history (patient_id, condition_name, diagnosed_date, resolved_date, status, notes, added_by_doctor_id) VALUES
(1, 'Hipertansiyon', '2020-05-10', NULL, 'chronic', 'Ilac tedavisi ile kontrol altinda', 1),
(2, 'Astim', '2015-03-20', NULL, 'active', 'Inhaler kullaniyor', 2),
(4, 'Diyabet Tip 2', '2018-07-15', NULL, 'chronic', 'Metformin tedavisi', 2),
(5, 'Alerji', '2019-11-05', NULL, 'active', 'Polen alerjisi', 2);

-- Surgery History
INSERT INTO surgery_history (patient_id, procedure_name, surgery_date, surgeon_name, complications, notes, added_by_doctor_id) VALUES
(4, 'Apandektomi', '2010-06-15', 'Dr. Mehmet Aslan', NULL, 'Komplikasyonsuz', 4),
(1, 'Safra Kesesi Ameliyati', '2015-08-20', 'Dr. Ayse Yilmaz', NULL, 'Laparoskopik, basarili', 1),
(3, 'Burun Egriligi Operasyonu', '2021-03-10', 'Dr. Can Demir', 'Hafif kanama', 'Iyilesme sureci normal', 4);

-- Allergies
INSERT INTO allergies (patient_id, allergen, allergy_type, reaction, severity, diagnosed_date, notes, added_by_doctor_id) VALUES
(1, 'Penisilin', 'medication', 'Dokuntu ve kasiinti', 'moderate', '2010-05-15', 'Beta-laktam antibiyotiklerden kacinilmali', 1),
(2, 'Findik', 'food', 'Agiz sismesi, nefes darligi', 'severe', '2015-06-20', 'EpiPen tasiyor', 2),
(5, 'Polen', 'environmental', 'Hapsirma, burun akintisi, goz kasintisi', 'mild', '2019-11-05', 'Bahar aylarinda siddetleniyor', 2),
(4, 'Ibuprofen', 'medication', 'Mide agrisi ve bulanti', 'mild', '2020-02-10', 'NSAID grubu ilaclarla dikkatli olunmali', 4);

-- Vital Signs
INSERT INTO vital_signs (patient_id, appointment_id, recorded_by_user_id, recorded_at, blood_pressure_systolic, blood_pressure_diastolic, heart_rate, temperature, respiratory_rate, oxygen_saturation, height, weight, bmi, notes) VALUES
(1, 6, 3, '2024-11-15 10:15:00', 135, 85, 78, 36.6, 16, 98, 175.00, 80.50, 26.29, 'Hafif yuksek tansiyon'),
(2, 2, 3, '2024-11-18 14:30:00', 120, 75, 72, 36.5, 14, 99, 165.00, 62.00, 22.77, 'Normal'),
(3, 3, 3, '2024-11-19 11:00:00', 140, 90, 80, 36.7, 15, 97, 180.00, 75.00, 23.15, 'Takip gerekli'),
(5, 5, 3, '2024-11-17 15:00:00', 110, 70, 68, 36.4, 14, 99, 172.00, 68.00, 23.00, 'Normal vital bulgular');

-- ============================================
-- SONUC KONTROLU
-- ============================================
SELECT 
    'users' as tablo, COUNT(*) as kayit_sayisi FROM users
UNION ALL
SELECT 'nfc_cards', COUNT(*) FROM nfc_cards
UNION ALL
SELECT 'patients', COUNT(*) FROM patients
UNION ALL
SELECT 'appointments', COUNT(*) FROM appointments
UNION ALL
SELECT 'diagnoses', COUNT(*) FROM diagnoses
UNION ALL
SELECT 'prescriptions', COUNT(*) FROM prescriptions
UNION ALL
SELECT 'medical_tests', COUNT(*) FROM medical_tests
UNION ALL
SELECT 'medical_history', COUNT(*) FROM medical_history
UNION ALL
SELECT 'surgery_history', COUNT(*) FROM surgery_history
UNION ALL
SELECT 'allergies', COUNT(*) FROM allergies
UNION ALL
SELECT 'vital_signs', COUNT(*) FROM vital_signs;
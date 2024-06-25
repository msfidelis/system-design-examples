-- Seed para Medicos
INSERT INTO Medicos (id, nome, especialidade, crm) VALUES
(1, 'Dr. João Silva', 'Cardiologia', 'CRM12345'),
(2, 'Dra. Maria Fernandes', 'Dermatologia', 'CRM23456'),
(3, 'Dr. Carlos Sousa', 'Neurologia', 'CRM34567'),
(4, 'Dra. Helena Costa', 'Pediatria', 'CRM45678'),
(5, 'Dr. Pedro Lima', 'Ortopedia', 'CRM56789');

-- Seed para Pacientes
INSERT INTO Pacientes (id, nome, data_nascimento, endereco) VALUES
(1, 'Maria Oliveira', '1985-07-10', 'Rua das Flores, 123'),
(2, 'João Santos', '1970-01-20', 'Avenida Brasil, 456'),
(3, 'Ana Costa', '1992-03-15', 'Rua das Palmeiras, 789'),
(4,'Pedro Almeida', '1980-08-22', 'Rua das Acácias, 101'),
(5, 'Julia Ribeiro', '1975-11-30', 'Avenida Central, 202');

-- Seed para Medicamentos
INSERT INTO Medicamentos (nome, descricao) VALUES
('Aspirina', 'Analgésico e anti-inflamatório'),
('Paracetamol', 'Analgésico'),
('Amoxicilina', 'Antibiótico'),
('Ibuprofeno', 'Analgésico e anti-inflamatório'),
('Simvastatina', 'Reduz colesterol');

-- Seed para Prescricoes
INSERT INTO Prescricoes (id_medico, id_paciente, data_prescricao) VALUES
(1, 1, '2023-05-20 14:30:00'),
(2, 2, '2023-06-10 09:00:00'),
(3, 3, '2023-07-05 11:15:00'),
(4, 4, '2023-08-01 10:00:00'),
(5, 5, '2023-09-15 16:30:00'),
(1, 3, '2023-10-20 08:45:00'),
(2, 4, '2023-11-25 13:20:00'),
(3, 5, '2023-12-05 09:10:00'),
(4, 1, '2024-01-12 11:30:00'),
(5, 2, '2024-02-22 15:00:00');

-- Seed para Prescricao_Medicamentos
INSERT INTO Prescricao_Medicamentos (id_prescricao, id_medicamento, horario, dosagem) VALUES
(1, 1, '08:00', '100mg'),
(1, 2, '20:00', '500mg'),
(2, 3, '12:00', '250mg'),
(3, 1, '07:00', '100mg'),
(3, 3, '19:00', '250mg'),
(4, 4, '08:00', '200mg'),
(4, 5, '21:00', '10mg'),
(5, 1, '06:00', '100mg'),
(5, 2, '18:00', '500mg'),
(6, 3, '09:00', '250mg'),
(6, 4, '19:00', '200mg'),
(7, 1, '07:30', '100mg'),
(7, 2, '22:00', '500mg'),
(8, 5, '10:00', '10mg'),
(8, 3, '14:00', '250mg'),
(9, 2, '09:00', '500mg'),
(9, 4, '17:00', '200mg'),
(10, 1, '08:00', '100mg'),
(10, 5, '20:00', '10mg'),
(1, 1, '08:00', '100mg'),
(2, 2, '20:00', '500mg'),
(3, 3, '12:00', '250mg'),
(4, 1, '07:00', '100mg'),
(5, 3, '19:00', '250mg'),
(6, 4, '08:00', '200mg'),
(7, 5, '21:00', '10mg')

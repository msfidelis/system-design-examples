CREATE TABLE IF NOT EXISTS Medicos (
    id SERIAL primary key ,
    nome VARCHAR(255) NOT NULL,
    especialidade VARCHAR(255) NOT NULL,
    crm VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Pacientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    data_nascimento DATE NOT NULL,
    endereco VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Medicamentos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao TEXT
);

CREATE TABLE IF NOT EXISTS Prescricoes (
    id SERIAL PRIMARY KEY,
    id_medico INT NOT NULL,
    id_paciente INT NOT NULL,
    data_prescricao TIMESTAMP NOT NULL,
    FOREIGN KEY (id_medico) REFERENCES Medicos(id),
    FOREIGN KEY (id_paciente) REFERENCES Pacientes(id)
);

CREATE TABLE IF NOT EXISTS Prescricao_Medicamentos (
	id SERIAL PRIMARY KEY,
    id_prescricao INT NOT NULL,
    id_medicamento INT NOT NULL,
    horario VARCHAR(50) NOT NULL,
    dosagem VARCHAR(50) NOT NULL,
    FOREIGN KEY (id_prescricao) REFERENCES Prescricoes(id),
    FOREIGN KEY (id_medicamento) REFERENCES Medicamentos(id)
);